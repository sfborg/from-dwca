package sysio

import (
	"os"

	"github.com/gnames/gnsys"
	"github.com/sfborg/from-dwca/internal/ent/sys"
	"github.com/sfborg/from-dwca/pkg/config"
)

type sysio struct {
	cfg config.Config
}

func New(cfg config.Config) sys.Sys {
	return &sysio{cfg: cfg}
}

func (s *sysio) Init() error {
	err := s.cleanup()
	if err != nil {
		return err
	}
	gnsys.MakeDir(s.cfg.DBPath)
	return nil
}

func (s *sysio) Close() error {
	return s.cleanup()
}

func (s *sysio) cleanup() error {
	return os.RemoveAll(s.cfg.RootPath)
}
