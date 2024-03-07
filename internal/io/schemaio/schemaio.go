package schemaio

import (
	"github.com/sfborg/from-dwca/internal/ent/schema"
	"github.com/sfborg/from-dwca/pkg/config"
)

type sch struct {
	cfg config.Config
}

func New(cfg config.Config) schema.SchemaSQL {
	return &sch{cfg: cfg}
}

func (s *sch) GetSchema(tag string) error {
	return nil
}
