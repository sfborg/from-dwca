package sfarcio

import (
	"errors"
	"log/slog"
	"path/filepath"
	"strings"
)

func (s *sfarcio) Export(outPath string) error {
	// check if core table has data
	if !s.Exists() {
		return errors.New("cannot find SFGA archive")
	}

	outPath = trimExtentions(outPath)

	// Determine the desired file extension based on configuration
	ext := ".sql"
	if s.cfg.WithBinOutput {
		ext = ".sqlite"
	}
	outPath += ext

	// Perform the export
	err := s.sfdb.Export(outPath, s.cfg.WithBinOutput, s.cfg.WithZipOutput)
	if err != nil {
		return err
	}

	return nil
}

func trimExtentions(outPath string) string {
	hasExt := false
	ext := filepath.Ext(outPath)
	var trimmed string
	if ext == ".zip" {
		hasExt = true
		outPath = strings.TrimSuffix(outPath, ext)
		trimmed += ext
		ext = filepath.Ext(outPath)
	}
	if ext == ".sql" || ext == ".sqlite" {
		hasExt = true
		outPath = strings.TrimSuffix(outPath, ext)
		trimmed = ext + trimmed
	}
	if hasExt {
		slog.Warn("Trimmed extentions from output File", "ext", trimmed)
	}
	return outPath
}
