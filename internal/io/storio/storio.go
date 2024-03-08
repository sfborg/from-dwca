package storio

import (
	"database/sql"

	"github.com/sfborg/from-dwca/internal/ent/stor"
	"github.com/sfborg/from-dwca/pkg/config"
)

type storio struct {
	db *sql.DB
}

func New(cfg config.Config) (stor.Storage, error) {
	return &storio{}, nil
}
