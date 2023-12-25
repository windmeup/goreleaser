package static

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/windmeup/goreleaser/pkg/config"
)

func TestExampleConfig(t *testing.T) {
	_, err := config.LoadReader(bytes.NewReader(ExampleConfig))
	require.NoError(t, err)
	require.NotEmpty(t, ExampleConfig)
}
