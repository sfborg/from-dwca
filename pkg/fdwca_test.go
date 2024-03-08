package fdwca_test

import (
	"path/filepath"
	"testing"

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
	assert.NotNil(arc)
}
