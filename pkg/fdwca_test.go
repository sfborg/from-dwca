package fdwca_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnsys"
	"github.com/sfborg/from-coldp/pkg/io/sfgarcio"
	"github.com/sfborg/from-dwca/internal/io/sysio"
	fdwca "github.com/sfborg/from-dwca/pkg"
	"github.com/sfborg/from-dwca/pkg/config"
	"github.com/sfborg/sflib/io/dbio"
	"github.com/sfborg/sflib/io/schemaio"
	"github.com/stretchr/testify/assert"
)

func TestGetDwca(t *testing.T) {
	assert := assert.New(t)
	path := filepath.Join("testdata", "dwca", "gymnodiniales.tar.gz")
	cfg := config.New(config.OptWrongFieldsNum(gnfmt.SkipBadRow))
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
	coldpCfg := cfg.ToColdpConfig()

	err = sysio.New(cfg).Init()
	assert.Nil(err)

	schema := schemaio.New(cfg.GitRepo, cfg.TempRepoDir)
	db := dbio.New(cfg.CacheSfgaDir)

	sfgarc := sfgarcio.New(coldpCfg, schema, db)
	err = sfgarc.Connect()
	assert.Nil(err)

	fd := fdwca.New(cfg, sfgarc)
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
	cfg := config.New(config.OptWrongFieldsNum(gnfmt.SkipBadRow))
	coldpCfg := cfg.ToColdpConfig()

	err = sysio.New(cfg).Init()
	assert.Nil(err)

	schema := schemaio.New(cfg.GitRepo, cfg.TempRepoDir)
	sfdb := dbio.New(cfg.CacheSfgaDir)

	sfgarc := sfgarcio.New(coldpCfg, schema, sfdb)
	err = sfgarc.Connect()
	assert.Nil(err)

	fd := fdwca.New(cfg, sfgarc)

	arc, err := fd.GetDwCA(path)
	assert.Nil(err)
	assert.NotNil(arc.Meta())
	err = fd.ImportDwCA(arc)
	assert.Nil(err)

	outPath := filepath.Join(os.TempDir(), "sfga")
	exists, err = gnsys.FileExists(outPath)
	assert.Nil(err)
	assert.False(exists)

	err = fd.ExportSFGA(outPath)
	assert.Nil(err)

	exists, err = gnsys.FileExists(outPath + ".sql")
	assert.Nil(err)
	assert.True(exists)
	os.Remove(outPath)
}
