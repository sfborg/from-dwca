package fdwca_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gnames/gnsys"
	"github.com/sfborg/from-dwca/internal/io/storio"
	"github.com/sfborg/from-dwca/internal/io/sysio"
	fdwca "github.com/sfborg/from-dwca/pkg"
	"github.com/sfborg/from-dwca/pkg/config"
	"github.com/sfborg/sflib/io/sfgaio"
	"github.com/stretchr/testify/assert"
)

func TestGetDwca(t *testing.T) {
	assert := assert.New(t)
	path := filepath.Join("testdata", "dwca", "gymnodiniales.tar.gz")
	cfg := config.New()
	fd := fdwca.New(cfg, nil)
	arc, err := fd.GetDwCA(path)
	assert.Nil(err)
	assert.NotNil(arc.Meta())
}

func TestImportDwCA(t *testing.T) {
	assert := assert.New(t)
	var err error
	path := filepath.Join("testdata", "dwca", "aos-birds.tar.gz")
	cfg := config.New()

	err = sysio.New(cfg).Init()
	assert.Nil(err)

	sfga := sfgaio.New(cfg.GitRepo, cfg.TempRepoPath)

	stor := storio.New(cfg, sfga)
	err = stor.Init()
	assert.Nil(err)

	fd := fdwca.New(cfg, stor)
	arc, err := fd.GetDwCA(path)
	assert.Nil(err)
	assert.NotNil(arc.Meta())
	err = fd.ImportDwCA(arc)
	assert.Nil(err)
}

func TestOutSFGA(t *testing.T) {
	assert := assert.New(t)
	var err error
	var exists bool

	path := filepath.Join("testdata", "dwca", "gymnodiniales.tar.gz")
	cfg := config.New()

	err = sysio.New(cfg).Init()
	assert.Nil(err)

	sfga := sfgaio.New(cfg.GitRepo, cfg.TempRepoPath)

	stor := storio.New(cfg, sfga)
	err = stor.Init()
	assert.Nil(err)

	fd := fdwca.New(cfg, stor)

	err = fd.OutSFGA("test")
	assert.NotNil(err)

	arc, err := fd.GetDwCA(path)
	assert.Nil(err)
	assert.NotNil(arc.Meta())
	err = fd.ImportDwCA(arc)
	assert.Nil(err)

	outPath := filepath.Join(os.TempDir(), "sfga.zip")
	exists, err = gnsys.FileExists(outPath)
	assert.Nil(err)
	assert.False(exists)

	err = fd.OutSFGA(outPath)
	assert.Nil(err)

	exists, err = gnsys.FileExists(outPath + ".sqlite")
	assert.Nil(err)
	assert.True(exists)
	os.Remove(outPath)
}
