package repoio

import (
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sfborg/from-dwca/internal/ent/repo"
	"github.com/sfborg/from-dwca/pkg/config"
)

type repoio struct {
	cfg config.Config
}

func New(cfg config.Config) repo.Repo {
	return &repoio{cfg: cfg}
}

func (r *repoio) FetchSchema() ([]byte, error) {
	var err error
	err = os.RemoveAll(r.cfg.RepoPath)
	if err != nil {
		slog.Error("Cannot remove path", "path", r.cfg.RepoPath, "err", err)
		return nil, err
	}

	err = r.cloneRepo()
	if err != nil {
		return nil, err
	}

	schemaPath := filepath.Join(r.cfg.RepoPath, "schema.sql")
	res, err := os.ReadFile(schemaPath)
	if err != nil {
		slog.Error("Cannot read schema", "path", schemaPath, "err", err)
		return nil, err
	}

	return res, nil
}

func (r *repoio) cloneRepo() error {
	var err error
	var currentDir string
	cmd := exec.Command("git", "clone", r.cfg.RepoURL, r.cfg.RepoPath)
	err = cmd.Run()
	if err != nil {
		slog.Error("Cannot clone repo", "repo", r.cfg.RepoURL, "err", err)
		return err
	}

	currentDir, err = os.Getwd()
	if err != nil {
		slog.Error("Cannot get current dir", "err", err)
		return err
	}
	defer os.Chdir(currentDir)

	err = os.Chdir(r.cfg.RepoPath)
	if err != nil {
		slog.Error("Cannot change dir", "path", r.cfg.RepoPath, "err", err)
		return err
	}

	if r.cfg.RepoTag == "" {
		return nil
	}

	cmd = exec.Command("git", "checkout", r.cfg.RepoTag)
	err = cmd.Run()
	if err != nil {
		slog.Error("Cannot checkout tag", "tag", r.cfg.RepoTag, "err", err)
		return err
	}

	return nil
}
