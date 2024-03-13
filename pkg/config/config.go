package config

import (
	"log/slog"
	"os"
	"path/filepath"
)

var (
	// repoURL is the URL to the SFGA schema repository.
	repoURL = "https://github.com/sfborg/sfga"

	// tag of the sfga repo to get correct schema version.
	repoTag = "v1.2.1"

	// jobsNum is the default number of concurrent jobs to run.
	jobsNum = 5
)

// Config is a configuration object for the Darwin Core Archive (DwCA)
// data processing.
type Config struct {
	// RepoURL is the URL to the SFGA schema repository.
	RepoURL string

	// RepoPath is a temporary location to schema files downloaded from GitHub.
	RepoPath string

	// RepoTag is a tag of the SFGA repository to use.
	RepoTag string

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
	path = filepath.Join(path, "sfborg")

	tmp := os.TempDir()
	schemaRepo := filepath.Join(tmp, "sfborg_sfda")

	res := Config{
		RepoURL:   repoURL,
		RepoTag:   repoTag,
		RepoPath:  schemaRepo,
		RootPath:  path,
		JobsNum:   jobsNum,
		BatchSize: 50_000,
	}

	for _, opt := range opts {
		opt(&res)
	}

	res.DBPath = filepath.Join(res.RootPath, "from", "dwca", "db")
	res.DumpPath = filepath.Join(res.RootPath, "from", "dwca", "dump")
	return res
}
