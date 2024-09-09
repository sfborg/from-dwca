package sfarcio

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/sfborg/from-dwca/internal/ent/schema"
)

func (s *sfarcio) InsertCore(data []*schema.Core) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
	INSERT OR IGNORE INTO core
			(dwc_taxon_id, local_id, global_id, dwc_scientific_name_id,
		   dwc_scientific_name, dwc_scientific_name_authorship,
		   dwc_name_published_in_year,

		   cardinality, canonical_id, canonical,  canonical_full_id,
		   canonical_full,

		   canonical_stem_id, canonical_stem, dwc_accepted_name_usage_id,
		   dwc_higher_classification, higher_classification_ids, 
			 
		   higher_classification_ranks, dwc_taxon_rank, is_virus, is_bacteria,
		   is_surrogate,
		   
		   dwc_nomenclatural_code, parse_quality)
		VALUES
			(?, ?, ?, ?, ?, ?, ?,  ?, ?, ?, ?, ?,  ?, ?, ?, ?, ?,  ?, ?, ?, ?, ?,  ?, ?)`,
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, v := range data {
		_, err = stmt.Exec(
			v.RecordID, v.LocalID, v.GlobalID, v.NameID, v.Name, v.Authorship,
			v.Year,

			v.Cardinality, v.CanonicalID, v.Canonical, v.CanonicalFullID,
			v.CanonicalFull,

			v.CanonicalStemID, v.CanonicalStem, v.AcceptedNameUsageID, v.Classification,
			v.ClassificationIDs,

			v.ClassificationRanks, v.Rank, boolToInt(v.Virus), boolToInt(v.Bacteria),
			boolToInt(v.Surrogate),

			v.NomeclaturalCode, v.ParseQuality,
		)
		if err != nil {
			tx.Rollback()
			slog.Error("Error inserting core data", "error", err)
			return err
		}
	}

	return tx.Commit()
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func FinishTx(tx *sql.Tx, err error) {
	if err != nil {
		// Rollback automatically if any errors occur within the transaction
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("Error rolling back transaction:", rollbackErr)
		}
	} else {
		// Commit the transaction if successful
		if err := tx.Commit(); err != nil {
			fmt.Println("Error committing transaction:", err)
		}
	}
}
