package storio

import (
	"log/slog"

	"github.com/sfborg/from-dwca/internal/ent/vern"
)

func (s *storio) InsertVernData(data []*vern.Data) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
	INSERT OR IGNORE INTO vernaculars (
			dwc_taxon_id, vernacular_name_id, dwc_vernacular_name,
		dcterms_language, lang_code, lang_eng_name, dwc_locality, dwc_country_code
		  )
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, v := range data {
		_, err = stmt.Exec(
			v.TaxonID, v.VernacularID, v.VernacularName,
			v.Language, v.LangCode, v.LangInEnglish, v.Locality, v.CountryCode,
		)
		if err != nil {
			tx.Rollback()
			slog.Error("Error inserting core data", "error", err)
			return err
		}
	}

	return tx.Commit()
}
