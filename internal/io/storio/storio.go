package storio

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sfborg/from-dwca/internal/ent/stor"
	"github.com/sfborg/from-dwca/internal/io/repoio"
	"github.com/sfborg/from-dwca/pkg/config"
	_ "modernc.org/sqlite"
)

type storio struct {
	cfg config.Config
	db  *sql.DB
}

func New(cfg config.Config) stor.Storage {
	return &storio{cfg: cfg}
}

func (s *storio) Init() error {
	r := repoio.New(s.cfg)
	schema, err := r.FetchSchema()
	if err != nil {
		return err
	}

	schFile := filepath.Join(s.cfg.DBPath, "schema.sql")
	err = os.WriteFile(schFile, schema, 0644)
	if err != nil {
		slog.Error("Cannot write schema file", "error", err)
		return err
	}

	dbFile := filepath.Join(s.cfg.DBPath, "sfga.db")

	// read := fmt.Sprintf("\".read %s\"", schFile)
	read := fmt.Sprintf(".read %s", schFile)
	fmt.Println()

	cmd := exec.Command("sqlite3", dbFile, read)
	fmt.Printf("CMD %s\n\n", cmd)
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

	s.db = db

	return nil
}

func (s *storio) Close() error {
	if s.db == nil {
		return nil
	}
	return s.db.Close()
}
