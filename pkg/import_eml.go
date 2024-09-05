package fdwca

import (
	"strings"

	"github.com/gnames/dwca/pkg/ent/eml"
	"github.com/sfborg/from-dwca/internal/ent/ds"
)

func (fd *fdwca) importEML(data *eml.EML, recNum int) error {
	var doi string
	if strings.Contains(data.Dataset.ID, "doi.") {
		doi = data.Dataset.ID
	}

	ds := ds.DataSource{
		ID:          data.Dataset.ID,
		Title:       data.Dataset.Title,
		DOI:         doi,
		Description: data.Dataset.Abstract.Para,
		RecordCount: recNum,
	}

	err := fd.stor.InsertDataSource(&ds)
	if err != nil {
		return err
	}

	return nil
}
