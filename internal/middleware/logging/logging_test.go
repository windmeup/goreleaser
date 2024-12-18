package logging

import (
	"testing"

	"github.com/caarlos0/log"
	"github.com/stretchr/testify/require"
	"github.com/windmeup/goreleaser/v2/pkg/context"
)

func TestLogging(t *testing.T) {
	require.NoError(t, Log("foo", func(_ *context.Context) error {
		return nil
	})(nil))

	require.NoError(t, PadLog("foo", func(_ *context.Context) error {
		log.Info("a")
		return nil
	})(nil))
}
