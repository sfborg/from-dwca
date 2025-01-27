package fdwca

import (
	"context"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gnames/coldp/ent/coldp"
	"github.com/gnames/dwca/pkg/ent/meta"
	"golang.org/x/sync/errgroup"
)

func (fd *fdwca) importCore() error {
	chIn := make(chan []string)
	chOut := make(chan []coldp.NameUsage)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	var wg sync.WaitGroup
	wg.Add(1)

	g.Go(func() error {
		defer wg.Done()
		return fd.coreWorker(ctx, chIn, chOut)
	})

	g.Go(func() error {
		return fd.writeCoreData(ctx, chOut)
	})

	// close chOut when all workers are done
	go func() {
		wg.Wait()
		close(chOut)
	}()

	_, err := fd.d.CoreStream(ctx, chIn)
	if err != nil {
		return err
	}

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func (fd *fdwca) coreWorker(
	ctx context.Context,
	chIn chan []string,
	chOut chan []coldp.NameUsage,
) error {
	fieldsMap := fieldsMap(fd.d.Meta().Core.Fields)
	coreID := fd.d.Meta().Core.ID.Idx

	batch := make([]coldp.NameUsage, 0, fd.cfg.BatchSize)
	for v := range chIn {
		row := fd.processCoreRow(v, coreID, fieldsMap)
		if len(batch) == fd.cfg.BatchSize {
			chOut <- batch
			batch = make([]coldp.NameUsage, 0, fd.cfg.BatchSize)
		}
		batch = append(batch, row)

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}
	chOut <- batch
	return nil
}

func (f *fdwca) writeCoreData(
	ctx context.Context,
	chOut chan []coldp.NameUsage,
) error {
	var err error
	for cd := range chOut {
		// write to db
		err = f.s.InsertNameUsages(cd)
		if err != nil {
			for range chOut {
			}
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}
	return nil
}

func fieldsMap(fields []meta.Field) map[string]int {
	fieldsMap := make(map[string]int)
	for _, v := range fields {
		term := filepath.Base(v.Term)
		term = strings.ToLower(term)
		fieldsMap[term] = v.Idx
	}
	return fieldsMap
}

func fieldVal(row []string, fielsMap map[string]int, name string) string {
	if idx, ok := fielsMap[name]; ok {
		if idx >= len(row) {
			return ""
		}
		return row[idx]
	}
	return ""
}

func (fd *fdwca) processCoreRow(
	row []string,
	idIdx int,
	fieldsMap map[string]int,
) coldp.NameUsage {
	res := coldp.NameUsage{ID: row[idIdx]}
	res.ScientificNameString = fieldVal(row, fieldsMap, "scientificnamestring")
	res.ScientificName = fieldVal(row, fieldsMap, "scientificname")
	res.Authorship = fieldVal(row, fieldsMap, "scientificnameauthorship")
	res.LocalID = fieldVal(row, fieldsMap, "localid")
	res.GlobalID = fieldVal(row, fieldsMap, "globalid")
	res.ParentID = fieldVal(row, fieldsMap, "parentnameusageid")
	code := fieldVal(row, fieldsMap, "nomenclaturalcode")
	res.Code = coldp.NewNomCode(code)
	rank := fieldVal(row, fieldsMap, "taxonrank")
	res.Rank = coldp.NewRank(rank)
	acceptedNameUsageID := fieldVal(row, fieldsMap, "acceptednameusageid")
	ts := fieldVal(row, fieldsMap, "taxonomicstatus")
	res.TaxonomicStatus = coldp.NewTaxonomicStatus(ts)

	if acceptedNameUsageID != "" && res.ID != acceptedNameUsageID {
		// for NameUsage's synonyms accepted ID goes to parent, and
		// synonymy is expressed by taxonomic status
		res.ParentID = acceptedNameUsageID
		if res.TaxonomicStatus == coldp.UnknownTaxSt {
			res.TaxonomicStatus = coldp.SynonymTS
		}
	} else {
		if res.TaxonomicStatus == coldp.UnknownTaxSt {
			res.TaxonomicStatus = coldp.AcceptedTS
		}
	}

	res.Kingdom = fieldVal(row, fieldsMap, "kingdom")
	res.Phylum = fieldVal(row, fieldsMap, "phylum")
	domain := fieldVal(row, fieldsMap, "domain")
	if domain != "" {
		res.Kingdom = domain
	}
	res.Class = fieldVal(row, fieldsMap, "class")
	res.Order = fieldVal(row, fieldsMap, "order")
	res.Family = fieldVal(row, fieldsMap, "family")
	res.Genus = fieldVal(row, fieldsMap, "genus")
	return res
}
