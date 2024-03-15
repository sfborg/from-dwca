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
	dumpFile := filepath.Join(s.cfg.DBPath, "sfga.sql")
	dbFile := filepath.Join(s.cfg.DBPath, "sfga.db")

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

	slog.Info("Creating zip file", "file", outPath)
	err = s.Zip(dumpFile, outPath)
	if err != nil {
		return err
	}

	slog.Info("Zip file created", "file", outPath)

	return nil
}

// 	dumpFile := filepath.Join(s.cfg.DBPath, "sfga.sql")
// 	dbFile := filepath.Join(s.cfg.DBPath, "sfga.db")
//
// 	dump := fmt.Sprintf("sqlite3 %s .dump > %s", dbFile, dumpFile)
// 	cmd := exec.Command("sqlite3", dbFile, ".dump")
// 	err := cmd.Run()
// 	if err != nil {
// 		return err
// 	}
//
// 	dumpWriter, err := cmd.StdoutPipe()
// 	if err != nil {
// 		return fmt.Errorf("error creating dump pipe: %w", err)
// 	}
//
// 	slog.Info("Creating zip file", "file", outPath)
// 	err = s.Zip(dumpFile, outPath)
// 	if err != nil {
// 		return err
// 	}
//
// 	slog.Info("Zip file created", "file", outPath)
// 	return nil
// }

func (s *storio) Zip(inputPath, outputPath string) error {
	zipFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	fileInfo, err := f.Stat()
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

	_, err = io.Copy(writer, f)
	if err != nil {
		return err
	}
	return nil
}
