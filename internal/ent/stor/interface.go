package stor

import (
	"github.com/sfborg/from-dwca/internal/ent/core"
	"github.com/sfborg/from-dwca/internal/ent/vern"
)

type Storage interface {
	Init() error
	InsertCoreData(data []*core.Data) error
	InsertVernData(data []*vern.Data) error
	Close() error
}
