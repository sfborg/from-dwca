package config

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/sfborg/sflib/ent/sfga"
)

var (
	// repoURL is the URL to the SFGA schema repository.
	repoURL = "https://github.com/sfborg/sfga"

	// tag of the sfga repo to get correct schema version.
	repoTag = "v1.2.6"

	// schemaHash is the sha256 sum of the correponding schema version.
	schemaHash = "f863ecadb42"

	// jobsNum is the default number of concurrent jobs to run.
	jobsNum = 5
)

// Config is a configuration object for the Darwin Core Archive (DwCA)
// data processing.
type Config struct {
	// GitRepo contains data for sfga schema Git repository.
	sfga.GitRepo

	// TempRepoPath is a temporary location to schema files downloaded from GitHub.
	TempRepoPath string

	// RootPath is the root path for all temporary files.
	RootPath string

	// DBPath is the path SFGA database.
	DBPath string

	// DumpPath is the path to store the resulting sqlite file with data.
	DumpPath string

	// JobsNum is the number of concurrent jobs to run.
	JobsNum int

	// BatchSize is the number of records to insert in one transaction.
	BatchSize int

	// InMemory is a flag to use in-memory sqlite database.
	InMemory bool

	// WithSqlOutput is a flag to output SQL dump instead of SQLite binary.
	WithSqlOutput bool
}

// Option is a function type that allows to standardize how options to
// the configuration are organized.
type Option func(*Config)

// OptRootPath sets the root path for all temporary files.
func OptRootPath(s string) Option {
	return func(c *Config) {
		c.RootPath = s
	}
}

// OptSchemaPath sets the path to store the sqlite schema file.
func OptSchemaPath(s string) Option {
	return func(c *Config) {
		c.DBPath = s
	}
}

// OptDumpPath sets the path to store resulting sqlite file with data imported
// from DwCA file.
func OptDumpPath(s string) Option {
	return func(c *Config) {
		c.DumpPath = s
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

// OptWithSqlOutput sets output as text-only SQL file.
func OptWithSqlOutput(b bool) Option {
	return func(c *Config) {
		c.WithSqlOutput = b
	}
}

// OptInMemory sets the flag to use in-memory sqlite database.
func OptInMemory(b bool) Option {
	return func(c *Config) {
		c.InMemory = b
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

	schemaRepo := filepath.Join(tmpDir, "sfborg_sfda")

	res := Config{
		GitRepo: sfga.GitRepo{
			URL:          repoURL,
			Tag:          repoTag,
			ShaSchemaSQL: schemaHash,
		},
		TempRepoPath: schemaRepo,
		RootPath:     path,
		JobsNum:      jobsNum,
		BatchSize:    50_000,
	}

	for _, opt := range opts {
		opt(&res)
	}

	res.DBPath = filepath.Join(res.RootPath, "from", "dwca", "db")
	res.DumpPath = filepath.Join(res.RootPath, "from", "dwca", "dump")
	return res
}
