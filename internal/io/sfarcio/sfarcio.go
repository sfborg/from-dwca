package sfarcio

import (
	"database/sql"
	"log/slog"

	"github.com/sfborg/from-dwca/internal/ent/sfarc"
	"github.com/sfborg/from-dwca/pkg/config"
	"github.com/sfborg/sflib/ent/sfga"
)

type sfarcio struct {
	cfg  config.Config
	sch  sfga.Schema
	sfdb sfga.DB
	db   *sql.DB
}

// New creates an instance of SFGArchive store
func New(cfg config.Config, sch sfga.Schema, sfdb sfga.DB) sfarc.Archive {
	return &sfarcio{cfg: cfg, sch: sch, sfdb: sfdb}
}

func (s *sfarcio) Exists() bool {
	if s.db == nil {
		return false
	}

	q := "SELECT dwc_taxon_id FROM core LIMIT 5"

	var id string
	err := s.db.QueryRow(q).Scan(&id)
	if err != nil {
		slog.Error("Cannot get data from core", "error", err)
		return false
	}
	if id == "" {
		slog.Error("No dwc_taxon_id in core")
		return false
	}

	return true
}

func (s *sfarcio) Close() error {
	if s.db == nil {
		return nil
	}
	return s.sfdb.Close()
}
