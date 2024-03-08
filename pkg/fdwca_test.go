package fdwca_test

import (
	"path/filepath"
	"testing"

	"github.com/sfborg/from-dwca/internal/io/storio"
	"github.com/sfborg/from-dwca/internal/io/sysio"
	fdwca "github.com/sfborg/from-dwca/pkg"
	"github.com/sfborg/from-dwca/pkg/config"
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
	path := filepath.Join("testdata", "dwca", "gymnodiniales.tar.gz")
	cfg := config.New()

	err = sysio.New(cfg).Init()
	assert.Nil(err)

	stor := storio.New(cfg)
	err = stor.Init()
	assert.Nil(err)

	fd := fdwca.New(cfg, stor)
	arc, err := fd.GetDwCA(path)
	assert.Nil(err)
	assert.NotNil(arc.Meta())
	err = fd.ImportDwCA(arc)
	assert.Nil(err)
}
