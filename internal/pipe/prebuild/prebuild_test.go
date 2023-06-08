package prebuild

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/windmeup/goreleaser/pkg/config"
	"github.com/windmeup/goreleaser/pkg/context"
)

func TestRun(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		ctx := context.New(config.Project{
			Env:    []string{"FOO=bar"},
			Builds: []config.Build{{Main: "{{ .Env.FOO }}"}},
		})
		require.NoError(t, Pipe{}.Run(ctx))
		require.Equal(t, "bar", ctx.Config.Builds[0].Main)
	})

	t.Run("empty", func(t *testing.T) {
		ctx := context.New(config.Project{
			Env:    []string{"FOO="},
			Builds: []config.Build{{Main: "{{ .Env.FOO }}"}},
		})
		require.NoError(t, Pipe{}.Run(ctx))
		require.Equal(t, ".", ctx.Config.Builds[0].Main)
	})

	t.Run("bad", func(t *testing.T) {
		ctx := context.New(config.Project{
			Builds: []config.Build{{Main: "{{ .Env.FOO }}"}},
		})
		require.EqualError(t, Pipe{}.Run(ctx), `template: tmpl:1:7: executing "tmpl" at <.Env.FOO>: map has no entry for key "FOO"`)
	})
}

func TestString(t *testing.T) {
	require.NotEmpty(t, Pipe{}.String())
}
