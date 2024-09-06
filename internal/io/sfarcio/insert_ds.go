package sfarcio

import (
	"log/slog"

	"github.com/sfborg/from-dwca/internal/ent/schema"
)

func (s *sfarcio) InsertDataSource(data *schema.DataSource) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
	INSERT OR IGNORE INTO data_source (
			id, gn_id, title, title_short, version, revision_date,
		  doi, citation, authors, description, website_url, data_url,
		  record_count, updated_at
		  )
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		data.ID, data.GnID, data.Title, data.TitleShort, data.Version,
		data.RevisionDate, data.DOI, data.Citation, data.Authors,
		data.Description, data.WebsiteURL, data.DataURL,
		data.RecordCount, data.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		slog.Error("Error inserting data_source data", "error", err)
		return err
	}

	return tx.Commit()
}
