package healthcheck

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/windmeup/goreleaser/internal/testctx"
)

func TestDependencies(t *testing.T) {
	ctx := testctx.New()
	require.Equal(t, []string{"git", "go"}, system{}.Dependencies(ctx))
}

func TestStringer(t *testing.T) {
	require.NotEmpty(t, system{}.String())
}
