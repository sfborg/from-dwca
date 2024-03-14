package fdwca

import (
	"github.com/gnames/dwca/pkg/ent/eml"
	"github.com/sfborg/from-dwca/internal/ent/ds"
)

func (fd *fdwca) importEML(data *eml.EML, recNum int) error {
	ds := ds.DataSource{
		ID:          data.Dataset.ID,
		Title:       data.Dataset.Title,
		Description: data.Dataset.Abstract.Para,
		RecordCount: recNum,
	}

	err := fd.stor.InsertDataSource(&ds)
	if err != nil {
		return err
	}

	return nil
}
