package static

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/windmeup/goreleaser/v2/pkg/config"
)

func TestGoExampleConfig(t *testing.T) {
	cfg, err := config.LoadReader(bytes.NewReader(GoExampleConfig))
	require.NoError(t, err)
	require.NotEmpty(t, GoExampleConfig)
	require.Equal(t, 2, cfg.Version)
}

func TestZigExampleConfig(t *testing.T) {
	cfg, err := config.LoadReader(bytes.NewReader(ZigExampleConfig))
	require.NoError(t, err)
	require.NotEmpty(t, ZigExampleConfig)
	require.Equal(t, 2, cfg.Version)
	require.Equal(t, "zig", cfg.Builds[0].Builder)
}
