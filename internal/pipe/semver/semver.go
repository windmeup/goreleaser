package semver

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/windmeup/goreleaser/pkg/context"
)

// Pipe is a global hook pipe.
type Pipe struct{}

// String is the name of this pipe.
func (Pipe) String() string {
	return "parsing tag"
}

// Run executes the hooks.
func (Pipe) Run(ctx *context.Context) error {
	version, err := monorepo(ctx)
	if err != nil {
		return err
	}
	sv, err := semver.NewVersion(version)
	if err != nil {
		return fmt.Errorf("failed to parse tag '%s' as semver: %w", ctx.Git.CurrentTag, err)
	}
	ctx.Semver = context.Semver{
		Major:      sv.Major(),
		Minor:      sv.Minor(),
		Patch:      sv.Patch(),
		Prerelease: sv.Prerelease(),
	}
	return nil
}

func monorepo(ctx *context.Context) (string, error) {
	currentTag := ctx.Git.CurrentTag
	if prefix := ctx.Config.Monorepo.TagPrefix; !strings.HasPrefix(currentTag, prefix) {
		return "", fmt.Errorf("failed to parse monorepo tag '%s': must starts with '%s'", currentTag, prefix)
	}
	parts := strings.Split(currentTag, "/")
	return parts[len(parts)-1], nil
}
