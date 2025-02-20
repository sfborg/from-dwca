package config

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gnames/gnfmt"
)

var jobsNum = 5

// Config is a configuration object for the Darwin Core Archive (DwCA)
// data processing.
type Config struct {
	// TempRepoDir is a temporary location to schema files downloaded from GitHub.
	TempRepoDir string

	// CacheDir is the root path for all cached files.
	CacheDir string

	// CacheSfgaDir is the path SFGA database.
	CacheSfgaDir string

	// // CacheSfgaDir is the path to store the resulting sqlite file with data.
	// CacheSfgaDir string

	// JobsNum is the number of concurrent jobs to run.
	JobsNum int

	// BatchSize is the number of records to insert in one transaction.
	BatchSize int

	// BadRow dets decision what to do if a row has more/less fields
	// than it should.
	BadRow gnfmt.BadRow

	// WithBinOutput is a flag to output binary SQLite database instead of
	// SQL dump.
	WithBinOutput bool

	// WithZipOutput is a flag to return zipped SFGAarchive outpu.
	WithZipOutput bool
}

// Option is a function type that allows to standardize how options to
// the configuration are organized.
type Option func(*Config)

// OptCacheDir sets the root path for all temporary files.
func OptCacheDir(s string) Option {
	return func(c *Config) {
		c.CacheDir = s
	}
}

// OptCacheSfgaDir sets the path to store resulting sqlite file with data imported
// from DwCA file.
func OptCacheSfgaDir(s string) Option {
	return func(c *Config) {
		c.CacheSfgaDir = s
	}
}

// OptJobsNum sets the number of concurrent jobs to run.
func OptJobsNum(n int) Option {
	return func(c *Config) {
		if n < 1 || n > 100 {
			slog.Warn(
				"Unsupported number of jobs (supported: 1-100). Using default value",
				"bad-input", n, "default", jobsNum,
			)
			n = jobsNum
		}
		c.JobsNum = n
	}
}

func OptWrongFieldsNum(br gnfmt.BadRow) Option {
	return func(c *Config) {
		c.BadRow = br
	}
}

// OptWithBinOutput sets output as binary SQLite file.
func OptWithBinOutput(b bool) Option {
	return func(c *Config) {
		c.WithBinOutput = b
	}
}

// OptWithZipOutput sets output as binary SQLite file.
func OptWithZipOutput(b bool) Option {
	return func(c *Config) {
		c.WithZipOutput = b
	}
}

// New creates a new Config object with default values, and allows to
// override them with options.
func New(opts ...Option) Config {
	tmpDir := os.TempDir()
	path, err := os.UserCacheDir()
	if err != nil {
		path = tmpDir
	}
	path = filepath.Join(path, "sfborg")

	schemaRepo := filepath.Join(tmpDir, "sfborg", "sfga")

	res := Config{
		TempRepoDir: schemaRepo,
		CacheDir:    path,
		JobsNum:     jobsNum,
		BatchSize:   50_000,
		BadRow:      gnfmt.ErrorBadRow,
	}

	for _, opt := range opts {
		opt(&res)
	}

	res.CacheSfgaDir = filepath.Join(res.CacheDir, "from", "dwca", "sfga")
	return res
}
