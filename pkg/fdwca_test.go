package fdwca_test

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	"github.com/gnames/gnfmt"
	"github.com/gnames/gnsys"
	"github.com/sfborg/from-dwca/internal/io/sysio"
	fdwca "github.com/sfborg/from-dwca/pkg"
	"github.com/sfborg/from-dwca/pkg/config"
	"github.com/sfborg/sflib/io/sfgaio"
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

	err = sysio.New(cfg).Init()
	assert.Nil(err)

	sfga := sfgaio.New()
	sfga.Create(cfg.CacheSfgaDir)
	_, err = sfga.Connect()
	assert.Nil(err)

	fd := fdwca.New(cfg, sfga)
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

	err = sysio.New(cfg).Init()
	assert.Nil(err)

	sfga := sfgaio.New()
	err = sfga.Create(cfg.CacheSfgaDir)
	_, err = sfga.Connect()
	assert.Nil(err)

	fd := fdwca.New(cfg, sfga)

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

// TestOutSimpleSFGA checks if ZIP files are created and if the
// taxonomid status is set to unknown.
func TestOutSimpleSFGA(t *testing.T) {
	assert := assert.New(t)
	var err error
	var exists bool

	path := filepath.Join("testdata", "dwca", "eol.tar.gz")
	cfg := config.New(config.OptWithZipOutput(true))

	err = sysio.New(cfg).Init()
	assert.Nil(err)

	sfga := sfgaio.New()
	_, err = sfga.Connect()
	err = sfga.Create(cfg.CacheSfgaDir)
	assert.Nil(err)

	fd := fdwca.New(cfg, sfga)

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

	suffixes := []string{".sql", ".sqlite", ".sql.zip", ".sqlite.zip"}
	for _, v := range suffixes {
		exists, err = gnsys.FileExists(outPath + v)
		assert.Nil(err)
		assert.True(exists)
	}

	db := Connect(assert, outPath+".sqlite")

	var count int
	err = db.QueryRow(
		`SELECT count(*) FROM taxon WHERE status_id != ''`,
	).Scan(&count)
	assert.Nil(err)
	assert.Equal(0, count)

	err = db.QueryRow(
		`SELECT count(*) FROM synonym`,
	).Scan(&count)
	assert.Nil(err)
	assert.Equal(0, count)

	os.Remove(outPath)
}

func Connect(assert *assert.Assertions, path string) *sql.DB {
	db, err := sql.Open("sqlite", path)
	assert.Nil(err)
	return db
}
