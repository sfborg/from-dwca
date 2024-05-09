package fdwca

import (
	"context"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/gnames/dwca/pkg/ent/meta"
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/ent/parsed"
	"github.com/gnames/gnuuid"
	"github.com/sfborg/from-dwca/internal/ent/core"
	"golang.org/x/sync/errgroup"
)

func (fd *fdwca) importCore() (int, error) {
	chIn := make(chan []string)
	chOut := make(chan []*core.Data)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	var wg sync.WaitGroup

	for i := 0; i < fd.cfg.JobsNum; i++ {
		wg.Add(1)
		g.Go(func() error {
			defer wg.Done()
			return fd.coreParserWorker(ctx, chIn, chOut)
		})
	}

	g.Go(func() error {
		return fd.writeCoreData(ctx, chOut)
	})

	// close chOut when all workers are done
	go func() {
		wg.Wait()
		close(chOut)
	}()

	num, err := fd.arc.CoreStream(ctx, chIn)
	if err != nil {
		return 0, err
	}

	if err := g.Wait(); err != nil {
		return 0, err
	}
	return num, nil
}

func (f *fdwca) writeCoreData(
	ctx context.Context,
	chOut chan []*core.Data,
) error {
	var err error
	for cd := range chOut {
		// write to db
		err = f.stor.InsertCoreData(cd)
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

func (fd *fdwca) coreParserWorker(
	ctx context.Context,
	chIn chan []string,
	chOut chan []*core.Data,
) error {
	p := <-fd.gnpPool
	defer func() {
		fd.gnpPool <- p
	}()

	fieldsMap := fieldsMap(fd.arc.Meta().Core.Fields)
	coreID := fd.arc.Meta().Core.ID.Idx

	batch := make([]*core.Data, 0, fd.cfg.BatchSize)
	for v := range chIn {
		row := fd.processCoreRow(v, p, coreID, fieldsMap)
		if len(batch) == fd.cfg.BatchSize {
			chOut <- batch
			batch = make([]*core.Data, 0, fd.cfg.BatchSize)
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
	p gnparser.GNparser,
	idIdx int,
	fieldsMap map[string]int,
) *core.Data {
	res := core.Data{RecordID: row[idIdx]}
	name := fieldVal(row, fieldsMap, "scientificnamestring")
	parsed := p.ParseName(name)
	addParsedData(&res, parsed)
	res.RecordID = row[idIdx]
	res.Classification = fieldVal(row, fieldsMap, "breadcrumbnames")
	res.ClassificationIDs = fieldVal(row, fieldsMap, "breadcrumbids")
	res.ClassificationRanks = fieldVal(row, fieldsMap, "breadcrumbranks")
	res.AcceptedNameUsageID = fieldVal(row, fieldsMap, "acceptednameusageid")
	res.NomeclaturalCode = fieldVal(row, fieldsMap, "nomenclaturalcode")
	res.Rank = fieldVal(row, fieldsMap, "taxonrank")

	return &res
}

func addParsedData(cd *core.Data, parsed parsed.Parsed) {
	cd.NameID = parsed.VerbatimID
	cd.Name = parsed.Verbatim
	if !parsed.Parsed {
		cd.Virus = parsed.Virus
		return
	}
	cd.Canonical = parsed.Canonical.Simple
	cd.CanonicalID = gnuuid.New(parsed.Canonical.Simple).String()
	cd.CanonicalFull = parsed.Canonical.Full
	cd.CanonicalFullID = gnuuid.New(parsed.Canonical.Full).String()
	cd.CanonicalStem = parsed.Canonical.Stemmed
	cd.CanonicalStemID = gnuuid.New(parsed.Canonical.Stemmed).String()
	if parsed.Bacteria != nil {
		cd.Bacteria = parsed.Bacteria.Value == 1
	}
	cd.Surrogate = parsed.Surrogate != nil
	cd.ParseQuality = parsed.ParseQuality
	cd.Authorship, cd.Year = authYear(parsed)
}

func authYear(p parsed.Parsed) (string, int) {
	var year int
	var auth string
	if p.Authorship == nil {
		return "", 0
	}

	auth = p.Authorship.Verbatim
	if p.Authorship.Year == "" {
		return auth, year
	}

	yr := strings.Trim(p.Authorship.Year, "()")
	yrInt, err := strconv.Atoi(yr[0:4])
	if err != nil {
		return auth, year
	}
	return auth, yrInt
}
