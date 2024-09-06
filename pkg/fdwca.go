package fdwca

import (
	dwca "github.com/gnames/dwca/pkg"
	dwcacfg "github.com/gnames/dwca/pkg/config"
	"github.com/gnames/gnparser"
	"github.com/sfborg/from-dwca/internal/ent/sfarc"
	"github.com/sfborg/from-dwca/pkg/config"
)

type fdwca struct {
	cfg     config.Config
	s       sfarc.Archive
	d       dwca.Archive
	gnpPool chan gnparser.GNparser
}

func New(cfg config.Config, sfarc sfarc.Archive) FromDwCA {
	res := &fdwca{cfg: cfg}

	poolSize := cfg.JobsNum
	gnpPool := make(chan gnparser.GNparser, poolSize)
	for i := 0; i < poolSize; i++ {
		cfgGNP := gnparser.NewConfig()
		gnpPool <- gnparser.New(cfgGNP)
	}
	res.gnpPool = gnpPool
	res.s = sfarc

	return res
}

func (fd *fdwca) GetDwCA(fileDwCA string) (dwca.Archive, error) {
	opts := []dwcacfg.Option{dwcacfg.OptJobsNum(fd.cfg.JobsNum)}
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

func (f *fdwca) ExportSFGA(outputPath string) error {
	err := f.s.Export(outputPath)
	if err != nil {
		return err
	}

	return nil
}

func (fd *fdwca) checkSFGA() bool {
	return fd.s.Exists()
}
