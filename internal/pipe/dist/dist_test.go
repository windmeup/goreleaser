package dist

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/windmeup/goreleaser/internal/testctx"
	"github.com/windmeup/goreleaser/pkg/config"
)

func TestDistDoesNotExist(t *testing.T) {
	folder := t.TempDir()
	dist := filepath.Join(folder, "dist")
	require.NoError(t, Pipe{}.Run(testctx.NewWithCfg(config.Project{Dist: dist})))
}

func TestPopulatedDistExists(t *testing.T) {
	folder := t.TempDir()
	dist := filepath.Join(folder, "dist")
	require.NoError(t, os.Mkdir(dist, 0o755))
	f, err := os.Create(filepath.Join(dist, "mybin"))
	require.NoError(t, err)
	require.NoError(t, f.Close())
	ctx := testctx.NewWithCfg(config.Project{Dist: dist})
	require.Error(t, Pipe{}.Run(ctx))
	ctx.Clean = true
	require.NoError(t, Pipe{}.Run(ctx))
	_, err = os.Stat(dist)
	require.False(t, os.IsExist(err))
}

func TestEmptyDistExists(t *testing.T) {
	folder := t.TempDir()
	dist := filepath.Join(folder, "dist")
	require.NoError(t, os.Mkdir(dist, 0o755))
	ctx := testctx.NewWithCfg(config.Project{Dist: dist})
	require.NoError(t, Pipe{}.Run(ctx))
	_, err := os.Stat(dist)
	require.False(t, os.IsNotExist(err))
}

func TestDescription(t *testing.T) {
	require.NotEmpty(t, Pipe{}.String())
}
