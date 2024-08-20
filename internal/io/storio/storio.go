package storio

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sfborg/from-dwca/internal/ent/stor"
	"github.com/sfborg/from-dwca/pkg/config"
	"github.com/sfborg/sflib/ent/sfga"
	_ "modernc.org/sqlite"
)

type storio struct {
	cfg  config.Config
	sfga sfga.SFGA
	db   *sql.DB
}

func New(cfg config.Config, sfga sfga.SFGA) stor.Storage {
	return &storio{cfg: cfg, sfga: sfga}
}

func (s *storio) Init() error {
	schema, err := s.sfga.FetchSchema()
	if err != nil {
		slog.Error("Cannot fetch schema", "error", err)
		return err
	}

	schFile := filepath.Join(s.cfg.DBPath, "schema.sql")
	err = os.WriteFile(schFile, schema, 0644)
	if err != nil {
		slog.Error("Cannot write schema file", "error", err)
		return err
	}

	dbFile := filepath.Join(s.cfg.DBPath, "sfga.db")

	read := fmt.Sprintf(".read %s", schFile)
	fmt.Println()

	cmd := exec.Command("sqlite3", dbFile, read)
	err = cmd.Run()
	if err != nil {
		slog.Error("Cannot import database using schema", "error", err)
		return err
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		slog.Error("Cannot open database", "error", err)
		return err
	}

	// Enable in-memory temporary tables
	_, err = db.Exec("PRAGMA temp_store = MEMORY")
	if err != nil {
		slog.Error("Cannot enable in-memory temporary tables", "error", err)
		return err
	}

	// Enable Write-Ahead Logging. Allow many reads and one write concurrently,
	// usually boosts write performance.
	_, err = db.Exec("PRAGMA journal_mode = WAL")
	if err != nil {
		slog.Error("Cannot enable WAL journal mode", "error", err)
		return err
	}

	s.db = db

	return nil
}

func (s *storio) Close() error {
	if s.db == nil {
		return nil
	}
	return s.db.Close()
}
