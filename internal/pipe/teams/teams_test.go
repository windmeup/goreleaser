package teams

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/windmeup/goreleaser/internal/testctx"
	"github.com/windmeup/goreleaser/internal/testlib"
	"github.com/windmeup/goreleaser/pkg/config"
)

func TestStringer(t *testing.T) {
	require.Equal(t, Pipe{}.String(), "teams")
}

func TestDefault(t *testing.T) {
	ctx := testctx.New()
	require.NoError(t, Pipe{}.Default(ctx))
	require.Equal(t, ctx.Config.Announce.Teams.MessageTemplate, defaultMessageTemplate)
}

func TestAnnounceInvalidTemplate(t *testing.T) {
	ctx := testctx.NewWithCfg(config.Project{
		Announce: config.Announce{
			Teams: config.Teams{
				Enabled:         true,
				MessageTemplate: "{{ .Foo }",
			},
		},
	})
	testlib.RequireTemplateError(t, Pipe{}.Announce(ctx))
}

func TestAnnounceMissingEnv(t *testing.T) {
	ctx := testctx.NewWithCfg(config.Project{
		Announce: config.Announce{
			Teams: config.Teams{
				Enabled: true,
			},
		},
	})
	require.NoError(t, Pipe{}.Default(ctx))
	require.EqualError(t, Pipe{}.Announce(ctx), `teams: env: environment variable "TEAMS_WEBHOOK" should not be empty`)
}

func TestSkip(t *testing.T) {
	t.Run("skip", func(t *testing.T) {
		require.True(t, Pipe{}.Skip(testctx.New()))
	})

	t.Run("dont skip", func(t *testing.T) {
		ctx := testctx.NewWithCfg(config.Project{
			Announce: config.Announce{
				Teams: config.Teams{
					Enabled: true,
				},
			},
		})
		require.False(t, Pipe{}.Skip(ctx))
	})
}
