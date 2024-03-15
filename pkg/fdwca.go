package fdwca

import (
	"errors"
	"path/filepath"

	dwca "github.com/gnames/dwca/pkg"
	dwcacfg "github.com/gnames/dwca/pkg/config"
	"github.com/gnames/gnparser"
	"github.com/sfborg/from-dwca/internal/ent/stor"
	"github.com/sfborg/from-dwca/pkg/config"
)

type fdwca struct {
	cfg     config.Config
	stor    stor.Storage
	arc     dwca.Archive
	gnpPool chan gnparser.GNparser
}

func New(cfg config.Config, stor stor.Storage) FromDwCA {
	res := &fdwca{cfg: cfg}

	poolSize := cfg.JobsNum
	gnpPool := make(chan gnparser.GNparser, poolSize)
	for i := 0; i < poolSize; i++ {
		cfgGNP := gnparser.NewConfig()
		gnpPool <- gnparser.New(cfgGNP)
	}
	res.gnpPool = gnpPool
	res.stor = stor

	return res
}

func (fd *fdwca) GetDwCA(fileDwCA string) (dwca.Archive, error) {
	dwcaCfg := dwcacfg.New()
	arc, err := dwca.Factory(fileDwCA, dwcaCfg)
	if err != nil {
		return nil, err
	}
	err = arc.Load(arc.Config().ExtractPath)
	if err != nil {
		return nil, err
	}

	err = arc.Normalize()
	if err != nil {
		return nil, err
	}

	arc, err = dwca.FactoryOutput(dwcaCfg)
	if err != nil {
		return nil, err
	}

	err = arc.Load(arc.Config().OutputPath)
	if err != nil {
		return nil, err
	}

	return arc, nil
}

func (fd *fdwca) ImportDwCA(arc dwca.Archive) error {
	fd.arc = arc
	num, err := fd.importCore()
	if err != nil {
		return err
	}

	err = fd.importExtensions(arc)
	if err != nil {
		return err
	}

	err = fd.importEML(arc.EML(), num)
	if err != nil {
		return err
	}

	return nil
}

func (f *fdwca) OutSFGA(path string) error {
	var err error
	if !f.checkSFGA() {
		return errors.New("SFGA not found")
	}

	ext := filepath.Ext(path)
	if ext != ".zip" {
		path = path + ".zip"
	}

	err = f.stor.DumpSFGA(path)
	if err != nil {
		return err
	}

	return nil
}
