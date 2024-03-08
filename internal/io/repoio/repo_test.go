package repoio_test

import (
	"testing"

	"github.com/sfborg/from-dwca/internal/io/repoio"
	"github.com/sfborg/from-dwca/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestFetchSchema(t *testing.T) {
	assert := assert.New(t)
	cfg := config.New()
	r := repoio.New(cfg)
	schema, err := r.FetchSchema()
	assert.Nil(err)
	assert.True(len(schema) > 200)
	assert.Contains(string(schema), "CREATE TABLE")
}
