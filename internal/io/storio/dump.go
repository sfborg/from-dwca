package storio

import (
	"archive/zip"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
)

func (s *storio) Exists() bool {
	if s.db == nil {
		return false
	}

	var count int
	err := s.db.QueryRow("SELECT count(*) from core").Scan(&count)
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}

	return true
}

func (s *storio) DumpSFGA(outPath string) error {
	var err error
	dumpFile := outPath
	dbFile := filepath.Join(s.cfg.DBPath, "sfga.db")

	if s.cfg.WithSqlOutput {
		err = dumpSQL(dbFile, dumpFile)
	} else {
		err = dumpBinary(dbFile, dumpFile)
	}

	if err != nil {
		return err
	}
	return nil
}

func dumpBinary(dbFile, dumpFile string) error {
	f, err := os.Open(dbFile)
	if err != nil {
		return err
	}
	defer f.Close()

	w, err := os.Create(dumpFile)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = io.Copy(w, f)
	if err != nil {
		return err
	}
	slog.Info("SQLite database file is created", "file", dumpFile)
	return nil
}

func dumpSQL(dbFile, dumpFile string) error {
	cmd := exec.Command("sqlite3", dbFile, ".dump")
	dumpWriter, err := os.Create(dumpFile)
	if err != nil {
		return err
	}
	defer dumpWriter.Close() // Ensure file gets closed

	cmd.Stdout = dumpWriter // Set command's output to the file

	if err = cmd.Start(); err != nil {
		return err
	}

	if err = cmd.Wait(); err != nil {
		return err
	}

	slog.Info("SQLite SQL file is created", "file", dumpFile)

	return nil
}

func (s *storio) Zip(inputPath, outputPath string) error {
	zipFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	w, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer w.Close()

	fileInfo, err := w.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, w)
	if err != nil {
		return err
	}
	return nil
}
