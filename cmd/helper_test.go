package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/windmeup/goreleaser/internal/skips"
	"github.com/windmeup/goreleaser/pkg/context"
)

func requireAll(tb testing.TB, ctx *context.Context, keys ...skips.Key) {
	tb.Helper()
	for _, key := range keys {
		require.True(tb, ctx.Skips[string(key)], "expected %q to be true, but was false", key)
	}
}
