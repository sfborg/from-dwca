package stor

import "github.com/sfborg/from-dwca/internal/ent/core"

type Storage interface {
	Init() error
	InsertCoreData(data []*core.Data) error
	Close() error
}
