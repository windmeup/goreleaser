// Package publish contains the publishing pipe.
package publish

import (
	"fmt"

	"github.com/windmeup/goreleaser/internal/middleware/errhandler"
	"github.com/windmeup/goreleaser/internal/middleware/logging"
	"github.com/windmeup/goreleaser/internal/middleware/skip"
	"github.com/windmeup/goreleaser/internal/pipe/artifactory"
	"github.com/windmeup/goreleaser/internal/pipe/aur"
	"github.com/windmeup/goreleaser/internal/pipe/blob"
	"github.com/windmeup/goreleaser/internal/pipe/brew"
	"github.com/windmeup/goreleaser/internal/pipe/chocolatey"
	"github.com/windmeup/goreleaser/internal/pipe/custompublishers"
	"github.com/windmeup/goreleaser/internal/pipe/docker"
	"github.com/windmeup/goreleaser/internal/pipe/ko"
	"github.com/windmeup/goreleaser/internal/pipe/krew"
	"github.com/windmeup/goreleaser/internal/pipe/milestone"
	"github.com/windmeup/goreleaser/internal/pipe/nix"
	"github.com/windmeup/goreleaser/internal/pipe/release"
	"github.com/windmeup/goreleaser/internal/pipe/scoop"
	"github.com/windmeup/goreleaser/internal/pipe/sign"
	"github.com/windmeup/goreleaser/internal/pipe/snapcraft"
	"github.com/windmeup/goreleaser/internal/pipe/upload"
	"github.com/windmeup/goreleaser/pkg/context"
)

// Publisher should be implemented by pipes that want to publish artifacts.
type Publisher interface {
	fmt.Stringer

	// Default sets the configuration defaults
	Publish(ctx *context.Context) error
}

// nolint: gochecknoglobals
var publishers = []Publisher{
	blob.Pipe{},
	upload.Pipe{},
	artifactory.Pipe{},
	custompublishers.Pipe{},
	docker.Pipe{},
	docker.ManifestPipe{},
	ko.Pipe{},
	sign.DockerPipe{},
	snapcraft.Pipe{},
	// This should be one of the last steps
	release.Pipe{},
	// brew et al use the release URL, so, they should be last
	nix.NewPublish(),
	brew.Pipe{},
	aur.Pipe{},
	krew.Pipe{},
	scoop.Pipe{},
	chocolatey.Pipe{},
	milestone.Pipe{},
}

// Pipe that publishes artifacts.
type Pipe struct{}

func (Pipe) String() string                 { return "publishing" }
func (Pipe) Skip(ctx *context.Context) bool { return ctx.SkipPublish }

func (Pipe) Run(ctx *context.Context) error {
	for _, publisher := range publishers {
		if err := skip.Maybe(
			publisher,
			logging.PadLog(
				publisher.String(),
				errhandler.Handle(publisher.Publish),
			),
		)(ctx); err != nil {
			return fmt.Errorf("%s: failed to publish artifacts: %w", publisher.String(), err)
		}
	}
	return nil
}
