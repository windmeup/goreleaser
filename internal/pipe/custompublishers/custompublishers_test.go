package custompublishers

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/windmeup/goreleaser/internal/testctx"
	"github.com/windmeup/goreleaser/pkg/config"
)

func TestDescription(t *testing.T) {
	require.NotEmpty(t, Pipe{}.String())
}

func TestSkip(t *testing.T) {
	t.Run("skip", func(t *testing.T) {
		require.True(t, Pipe{}.Skip(testctx.New()))
	})

	t.Run("skip on skip-publish", func(t *testing.T) {
		ctx := testctx.NewWithCfg(config.Project{
			Publishers: []config.Publisher{
				{},
			},
		}, testctx.SkipPublish)
		require.True(t, Pipe{}.Skip(ctx))
	})

	t.Run("dont skip", func(t *testing.T) {
		ctx := testctx.NewWithCfg(config.Project{
			Publishers: []config.Publisher{
				{},
			},
		})
		require.False(t, Pipe{}.Skip(ctx))
	})
}

func TestPublish(t *testing.T) {
	require.NoError(t, Pipe{}.Publish(testctx.NewWithCfg(config.Project{
		Publishers: []config.Publisher{
			{
				Cmd: "echo",
			},
		},
	})))
}
