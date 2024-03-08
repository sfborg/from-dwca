package storio_test

import (
	"testing"

	"github.com/sfborg/from-dwca/internal/io/storio"
	"github.com/sfborg/from-dwca/internal/io/sysio"
	"github.com/sfborg/from-dwca/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)
	cfg := config.New()
	err := sysio.New(cfg).Init()
	assert.Nil(err)

	st := storio.New(cfg)
	err = st.Init()
	assert.Nil(err)

	err = st.Close()
	assert.Nil(err)
}
