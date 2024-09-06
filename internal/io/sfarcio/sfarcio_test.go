package sfarcio_test

import (
	"testing"

	"github.com/sfborg/from-dwca/internal/io/sfarcio"
	"github.com/sfborg/from-dwca/internal/io/sysio"
	"github.com/sfborg/from-dwca/pkg/config"
	"github.com/sfborg/sflib/io/dbio"
	"github.com/sfborg/sflib/io/schemaio"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)
	cfg := config.New()
	err := sysio.New(cfg).Init()
	assert.Nil(err)

	schema := schemaio.New(cfg.GitRepo, cfg.TempRepoDir)
	sfdb := dbio.New(cfg.CacheSfgaDir)

	st := sfarcio.New(cfg, schema, sfdb)
	err = st.Connect()
	assert.Nil(err)

	err = st.Close()
	assert.Nil(err)
}
