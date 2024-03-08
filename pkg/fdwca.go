package fdwca

import (
	dwca "github.com/gnames/dwca/pkg"
	dwcacfg "github.com/gnames/dwca/pkg/config"
	"github.com/sfborg/from-dwca/internal/ent/stor"
	"github.com/sfborg/from-dwca/pkg/config"
)

type fdwca struct {
	cfg  config.Config
	stor stor.Storage
	dc   dwca.Archive
}

func New(cfg config.Config, stor stor.Storage) FromDwCA {
	return &fdwca{cfg: cfg}
}

func (f *fdwca) GetDwCA(fileDwCA string) (dwca.Archive, error) {
	dwcaCfg := dwcacfg.New()
	arc, err := dwca.Factory(fileDwCA, dwcaCfg)
	if err != nil {
		return nil, err
	}
	arc.Load(arc.Config().ExtractPath)
	arc.Normalize()

	arc, err = dwca.FactoryOutput(dwcaCfg)
	if err != nil {
		return nil, err
	}

	return arc, nil
}

func (f *fdwca) ExportData(arc dwca.Archive) error {
	return nil
}

func (f *fdwca) DumpData() error {
	return nil
}
