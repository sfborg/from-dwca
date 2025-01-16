package sfarc

import "github.com/gnames/coldp/ent/coldp"

// Archive is an SFGArchive factory. It imports data from a
type Archive interface {
	Exists() bool
	Connect() error
	Close() error

	InsertNameUsage(data []*coldp.NameUsage) error
	InsertVernacular(data []*coldp.Vernacular) error
	InsertMeta(data *coldp.Meta) error

	Export(outPath string) error
}
