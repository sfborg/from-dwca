package fdwca

import (
	dwca "github.com/gnames/dwca/pkg"
	dwcacfg "github.com/gnames/dwca/pkg/config"
	"github.com/gnames/gnparser"
	"github.com/sfborg/from-coldp/pkg/ent/sfgarc"
	"github.com/sfborg/from-dwca/pkg/config"
)

type fdwca struct {
	cfg     config.Config
	s       sfgarc.Archive
	d       dwca.Archive
	gnpPool chan gnparser.GNparser
}

func New(cfg config.Config, sfgarc sfgarc.Archive) FromDwCA {
	res := &fdwca{cfg: cfg}

	poolSize := cfg.JobsNum
	gnpPool := make(chan gnparser.GNparser, poolSize)
	for i := 0; i < poolSize; i++ {
		cfgGNP := gnparser.NewConfig()
		gnpPool <- gnparser.New(cfgGNP)
	}
	res.gnpPool = gnpPool
	res.s = sfgarc

	return res
}

func (fd *fdwca) GetDwCA(fileDwCA string) (dwca.Archive, error) {
	opts := []dwcacfg.Option{
		dwcacfg.OptJobsNum(fd.cfg.JobsNum),
		dwcacfg.OptWrongFieldsNum(fd.cfg.BadRow),
	}

	dwcaCfg := dwcacfg.New(opts...)
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
	fd.d = arc
	err := fd.importCore()
	if err != nil {
		return err
	}

	err = fd.importExtensions(arc)
	if err != nil {
		return err
	}

	err = fd.importEML(arc.EML())
	if err != nil {
		return err
	}

	return nil
}

func (f *fdwca) ExportSFGA(outputPath string) error {
	err := f.s.Export(outputPath)
	if err != nil {
		return err
	}

	return nil
}
