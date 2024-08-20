package storio_test

import (
	"testing"

	"github.com/sfborg/from-dwca/internal/io/storio"
	"github.com/sfborg/from-dwca/internal/io/sysio"
	"github.com/sfborg/from-dwca/pkg/config"
	"github.com/sfborg/sflib/io/sfgaio"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)
	cfg := config.New()
	err := sysio.New(cfg).Init()
	assert.Nil(err)

	sfga := sfgaio.New(cfg.GitRepo, cfg.TempRepoPath)

	st := storio.New(cfg, sfga)
	err = st.Init()
	assert.Nil(err)

	err = st.Close()
	assert.Nil(err)
}
