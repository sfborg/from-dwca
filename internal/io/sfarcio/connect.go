package sfarcio

import (
	"log/slog"
	"os"
	"path/filepath"
)

func (s *sfarcio) Connect() error {
	err := s.setSchema()
	if err != nil {
		slog.Error("Cannot get SFGA schema", "error", err)
	}

	// get Connector to SQLite db
	s.db, err = s.sfdb.Connect()
	if err != nil {
		slog.Error("Cannot connect to SQLite database", "error", err)
	}

	return nil
}

func (s *sfarcio) setSchema() error {
	defer s.sch.Clean()

	schema, err := s.sch.Fetch()
	if err != nil {
		slog.Error("Cannot fetch schema", "error", err)
		return err
	}

	schFile := filepath.Join(s.cfg.CacheDbDir, "schema.sql")
	err = os.WriteFile(schFile, schema, 0644)
	if err != nil {
		slog.Error("Cannot write schema file", "error", err)
		return err
	}
	return nil
}