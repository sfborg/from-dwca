package sfarc

import (
	"github.com/sfborg/from-dwca/internal/ent/schema"
)

// Archive is an SFGArchive factory. It imports data from a
// \\
type Archive interface {
	Exists() bool
	Connect() error
	Close() error

	InsertCore(data []*schema.Core) error
	InsertVern(data []*schema.Vern) error
	InsertDataSource(data *schema.DataSource) error

	Export(outPath string) error
}
