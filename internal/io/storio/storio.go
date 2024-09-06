package storio

import (
	"github.com/sfborg/from-dwca/internal/ent/schema"
	"github.com/sfborg/from-dwca/internal/ent/stor"
	"github.com/sfborg/from-dwca/pkg/config"
	"github.com/sfborg/sflib/ent/sfga"
)

type storio struct {
}

func New(cfg config.Config, sch sfga.Schema) stor.Storage {
	res := storio{}
	return &res
}

func (s *storio) Init() error {
	return nil
}

func (s *storio) InsertCoreData(data []*schema.Core) error {
	return nil
}
func (s *storio) InsertVernData(data []*schema.Vern) error {
	return nil
}
func (s *storio) InsertDataSource(data *schema.DataSource) error {
	return nil
}

func (s *storio) Exists() bool {
	return false
}

func (s *storio) DumpSFGA(outPath string) error {
	return nil
}

func (s *storio) Close() error {
	return nil
}
