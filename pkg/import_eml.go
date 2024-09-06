package fdwca

import (
	"strings"

	"github.com/gnames/dwca/pkg/ent/eml"
	"github.com/sfborg/from-dwca/internal/ent/schema"
)

func (fd *fdwca) importEML(data *eml.EML, recNum int) error {
	var doi string
	if strings.Contains(data.Dataset.ID, "doi.") {
		doi = data.Dataset.ID
	}

	ds := schema.DataSource{
		ID:          data.Dataset.ID,
		Title:       data.Dataset.Title,
		DOI:         doi,
		Description: data.Dataset.Abstract.Para,
		RecordCount: recNum,
	}

	err := fd.s.InsertDataSource(&ds)
	if err != nil {
		return err
	}

	return nil
}
