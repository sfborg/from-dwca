package config

import (
	"log/slog"
	"os"
	"path/filepath"
)

var (
	// tag of sfgma schema used to create sqlite database.
	schemaTag = "v1.0.1"

	// jobsNum is the default number of concurrent jobs to run.
	jobsNum = 5
)

// Config is a configuration object for the Darwin Core Archive (DwCA)
// data processing.
type Config struct {
	// RootPath is the root path for all temporary files.
	RootPath string

	// SchemaRepo is a temporary location to schema files downloaded from GitHub.
	SchemaRepo string

	// SchemaPath is the path to store the schema file.
	SchemaPath string

	// DumpPath is the path to store the resulting sqlite file with data.
	DumpPath string

	// JobsNum is the number of concurrent jobs to run.
	JobsNum int

	// InMemory is a flag to use in-memory sqlite database.
	InMemory bool
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
		c.SchemaPath = s
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

// OptInMemory sets the flag to use in-memory sqlite database.
func OptInMemory(b bool) Option {
	return func(c *Config) {
		c.InMemory = b
	}
}

// New creates a new Config object with default values, and allows to
// override them with options.
func New(opts ...Option) Config {
	path, err := os.UserCacheDir()
	if err != nil {
		path = os.TempDir()
	}

	tmp := os.TempDir()
	schemaRepo := filepath.Join(tmp, "sfgma")

	path = filepath.Join(path, "sfborg", "from", "dwca")
	res := Config{
		RootPath:   path,
		JobsNum:    jobsNum,
		SchemaRepo: schemaRepo,
	}

	for _, opt := range opts {
		opt(&res)
	}

	res.SchemaPath = filepath.Join(res.RootPath, "db")
	res.DumpPath = filepath.Join(res.RootPath, "dump")
	return res
}
