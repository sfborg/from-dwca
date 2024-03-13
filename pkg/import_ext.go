package fdwca

import (
	"context"
	"path/filepath"
	"strings"

	dwca "github.com/gnames/dwca/pkg"
	"github.com/gnames/dwca/pkg/ent/meta"
	"github.com/gnames/gnlib"
	"github.com/gnames/gnuuid"
	"github.com/sfborg/from-dwca/internal/ent/vern"
	"golang.org/x/sync/errgroup"
)

func (f *fdwca) importExtensions(arc dwca.Archive) error {
	for i, ext := range arc.Meta().Extensions {
		if ext == nil {
			continue
		}

		rowType := filepath.Base(ext.RowType)
		rowType = strings.ToLower(rowType)
		if !strings.Contains(rowType, "vernacular") {
			continue
		}
		f.importVernacular(i, ext)
	}
	return nil
}

func (fd *fdwca) importVernacular(idx int, ext *meta.Extension) error {
	chIn := make(chan []string)
	chOut := make(chan []*vern.Data)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		defer close(chOut)
		return fd.vernWorker(ctx, ext, chIn, chOut)
	})

	g.Go(func() error {
		return fd.writeVernData(ctx, chOut)
	})

	err := fd.arc.ExtensionStream(ctx, idx, chIn)
	if err != nil {
		return err
	}

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func (fd *fdwca) vernWorker(
	ctx context.Context,
	ext *meta.Extension,
	chIn <-chan []string,
	chOut chan<- []*vern.Data,
) error {
	fieldsMap := fieldsMap(ext.Fields)
	coreID := ext.CoreID.Idx

	batch := make([]*vern.Data, 0, fd.cfg.BatchSize)
	for v := range chIn {
		vrn := fd.processVernRow(v, coreID, fieldsMap)
		if vrn == nil {
			continue
		}
		if len(batch) == fd.cfg.BatchSize {
			chOut <- batch
			batch = make([]*vern.Data, 0, fd.cfg.BatchSize)
		}
		batch = append(batch, vrn)

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}
	chOut <- batch
	return nil
}

func (fd *fdwca) processVernRow(
	row []string,
	coreID int,
	fieldsMap map[string]int,
) *vern.Data {
	var res vern.Data
	res.TaxonID = row[coreID]

	res.VernacularName = fieldVal(row, fieldsMap, "vernacularname")
	if res.VernacularName == "" {
		return nil
	}
	res.VernacularID = gnuuid.New(res.VernacularName).String()

	res.Language = fieldVal(row, fieldsMap, "language")
	res.LangCode = gnlib.LangCode(res.Language)
	res.LangInEnglish = gnlib.LangName(res.LangCode)
	res.Locality = fieldVal(row, fieldsMap, "locality")
	res.CountryCode = fieldVal(row, fieldsMap, "countrycode")
	return &res
}

func (fd *fdwca) writeVernData(
	ctx context.Context,
	chOut <-chan []*vern.Data,
) error {
	var err error
	for cd := range chOut {
		// write to db
		err = fd.stor.InsertVernData(cd)
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
