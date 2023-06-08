// Package defaults make the list of Defaulter implementations available
// so projects extending GoReleaser are able to use it, namely, GoDownloader.
package defaults

import (
	"fmt"

	"github.com/windmeup/goreleaser/internal/pipe/archive"
	"github.com/windmeup/goreleaser/internal/pipe/artifactory"
	"github.com/windmeup/goreleaser/internal/pipe/aur"
	"github.com/windmeup/goreleaser/internal/pipe/blob"
	"github.com/windmeup/goreleaser/internal/pipe/brew"
	"github.com/windmeup/goreleaser/internal/pipe/build"
	"github.com/windmeup/goreleaser/internal/pipe/checksums"
	"github.com/windmeup/goreleaser/internal/pipe/chocolatey"
	"github.com/windmeup/goreleaser/internal/pipe/discord"
	"github.com/windmeup/goreleaser/internal/pipe/docker"
	"github.com/windmeup/goreleaser/internal/pipe/gomod"
	"github.com/windmeup/goreleaser/internal/pipe/ko"
	"github.com/windmeup/goreleaser/internal/pipe/krew"
	"github.com/windmeup/goreleaser/internal/pipe/linkedin"
	"github.com/windmeup/goreleaser/internal/pipe/mastodon"
	"github.com/windmeup/goreleaser/internal/pipe/mattermost"
	"github.com/windmeup/goreleaser/internal/pipe/milestone"
	"github.com/windmeup/goreleaser/internal/pipe/nfpm"
	"github.com/windmeup/goreleaser/internal/pipe/nix"
	"github.com/windmeup/goreleaser/internal/pipe/opencollective"
	"github.com/windmeup/goreleaser/internal/pipe/project"
	"github.com/windmeup/goreleaser/internal/pipe/reddit"
	"github.com/windmeup/goreleaser/internal/pipe/release"
	"github.com/windmeup/goreleaser/internal/pipe/sbom"
	"github.com/windmeup/goreleaser/internal/pipe/scoop"
	"github.com/windmeup/goreleaser/internal/pipe/sign"
	"github.com/windmeup/goreleaser/internal/pipe/slack"
	"github.com/windmeup/goreleaser/internal/pipe/smtp"
	"github.com/windmeup/goreleaser/internal/pipe/snapcraft"
	"github.com/windmeup/goreleaser/internal/pipe/snapshot"
	"github.com/windmeup/goreleaser/internal/pipe/sourcearchive"
	"github.com/windmeup/goreleaser/internal/pipe/teams"
	"github.com/windmeup/goreleaser/internal/pipe/telegram"
	"github.com/windmeup/goreleaser/internal/pipe/twitter"
	"github.com/windmeup/goreleaser/internal/pipe/universalbinary"
	"github.com/windmeup/goreleaser/internal/pipe/upload"
	"github.com/windmeup/goreleaser/internal/pipe/upx"
	"github.com/windmeup/goreleaser/internal/pipe/webhook"
	"github.com/windmeup/goreleaser/pkg/context"
)

// Defaulter can be implemented by a Piper to set default values for its
// configuration.
type Defaulter interface {
	fmt.Stringer

	// Default sets the configuration defaults
	Default(ctx *context.Context) error
}

// Defaulters is the list of defaulters.
// nolint: gochecknoglobals
var Defaulters = []Defaulter{
	snapshot.Pipe{},
	release.Pipe{},
	project.Pipe{},
	gomod.Pipe{},
	build.Pipe{},
	universalbinary.Pipe{},
	upx.Pipe{},
	sourcearchive.Pipe{},
	archive.Pipe{},
	nfpm.Pipe{},
	snapcraft.Pipe{},
	checksums.Pipe{},
	sign.Pipe{},
	sign.DockerPipe{},
	sbom.Pipe{},
	docker.Pipe{},
	docker.ManifestPipe{},
	artifactory.Pipe{},
	blob.Pipe{},
	upload.Pipe{},
	aur.Pipe{},
	nix.Pipe{},
	brew.Pipe{},
	krew.Pipe{},
	ko.Pipe{},
	scoop.Pipe{},
	discord.Pipe{},
	reddit.Pipe{},
	slack.Pipe{},
	teams.Pipe{},
	twitter.Pipe{},
	smtp.Pipe{},
	mastodon.Pipe{},
	mattermost.Pipe{},
	milestone.Pipe{},
	linkedin.Pipe{},
	telegram.Pipe{},
	webhook.Pipe{},
	chocolatey.Pipe{},
	opencollective.Pipe{},
}
