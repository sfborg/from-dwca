package stor

import (
	"github.com/sfborg/from-dwca/internal/ent/schema"
)

// Storage provides connection to the SQLite databse and gives
// methods to insert data according to the SFGA schema.
type Storage interface {
	Init() error
	InsertCoreData(data []*schema.Core) error
	InsertVernData(data []*schema.Vern) error
	InsertDataSource(data *schema.DataSource) error
	Exists() bool
	DumpSFGA(outPath string) error
	Close() error
}
