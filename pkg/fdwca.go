package fdwca

import (
	dwca "github.com/gnames/dwca/pkg"
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

func (f *fdwca) GetDwCA(fileDwCA string) error {
	return nil
}

func (f *fdwca) ExportData() error {
	return nil
}

func (f *fdwca) DumpData() error {
	return nil
}
