package fdwca

import dwca "github.com/gnames/dwca/pkg"

// FromDwCA provies methods to convert DarwinCore Archive to
// Species File Group Archive.
type FromDwCA interface {
	// GetDwCA reads a DarwinCore Archive from a file, and returns
	// a normalized dwca.Archive object.
	GetDwCA(fileDwCA string) (dwca.Archive, error)

	// ImportDwCA converts a dwca.Archive to a Species File Group Archive
	// database.
	ImportDwCA(arc dwca.Archive) error

	// OutSFGA writes a Species File Group Archive to a file.
	OutSFGA(outputPath string) error
}
