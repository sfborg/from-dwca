package stor

import (
	"github.com/sfborg/from-dwca/internal/ent/core"
	"github.com/sfborg/from-dwca/internal/ent/ds"
	"github.com/sfborg/from-dwca/internal/ent/vern"
)

// Storage provides connection to the SQLite databse and gives
// methods to insert data according to the SFGA schema.
type Storage interface {
	Init() error
	InsertCoreData(data []*core.Data) error
	InsertVernData(data []*vern.Data) error
	InsertDataSource(data *ds.DataSource) error
	Exists() bool
	DumpSFGA(outPath string) error
	Close() error
}
