package fdwca

import (
	"strings"

	"github.com/gnames/coldp/ent/coldp"
	"github.com/gnames/dwca/pkg/ent/eml"
)

func (fd *fdwca) importEML(data *eml.EML) error {
	var doi string
	if strings.Contains(data.Dataset.ID, "doi.") {
		doi = data.Dataset.ID
	}

	meta := coldp.Meta{
		Title:       data.Dataset.Title,
		DOI:         doi,
		Description: data.Dataset.Abstract.Para,
	}

	err := fd.s.InsertMeta(&meta)
	if err != nil {
		return err
	}

	return nil
}
