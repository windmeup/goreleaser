package testlib

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/windmeup/goreleaser/internal/tmpl"
)

// RequireTemplateError requires thqt an error happens and that it is a template error.
func RequireTemplateError(tb testing.TB, err error) {
	tb.Helper()

	require.Error(tb, err)
	require.ErrorAs(tb, err, &tmpl.Error{})
}
