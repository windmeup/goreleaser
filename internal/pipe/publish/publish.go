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
	"github.com/windmeup/goreleaser/internal/pipe/winget"
	"github.com/windmeup/goreleaser/internal/skips"
	"github.com/windmeup/goreleaser/pkg/context"
)

// Publisher should be implemented by pipes that want to publish artifacts.
type Publisher interface {
	fmt.Stringer

	// Default sets the configuration defaults
	Publish(ctx *context.Context) error
}

// New publish pipeline.
func New() Pipe {
	return Pipe{
		pipeline: []Publisher{
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
			winget.Pipe{},
			brew.Pipe{},
			aur.Pipe{},
			krew.Pipe{},
			scoop.Pipe{},
			chocolatey.Pipe{},
			milestone.Pipe{},
		},
	}
}

// Pipe that publishes artifacts.
type Pipe struct {
	pipeline []Publisher
}

func (Pipe) String() string                 { return "publishing" }
func (Pipe) Skip(ctx *context.Context) bool { return skips.Any(ctx, skips.Publish) }

func (p Pipe) Run(ctx *context.Context) error {
	memo := errhandler.Memo{}
	for _, publisher := range p.pipeline {
		if err := skip.Maybe(
			publisher,
			logging.PadLog(
				publisher.String(),
				errhandler.Handle(publisher.Publish),
			),
		)(ctx); err != nil {
			if ig, ok := publisher.(Continuable); ok && ig.ContinueOnError() && !ctx.FailFast {
				memo.Memorize(fmt.Errorf("%s: %w", publisher.String(), err))
				continue
			}
			return fmt.Errorf("%s: failed to publish artifacts: %w", publisher.String(), err)
		}
	}
	return memo.Error()
}

type Continuable interface {
	ContinueOnError() bool
}
