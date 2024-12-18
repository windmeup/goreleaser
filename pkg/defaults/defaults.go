// Package defaults make the list of Defaulter implementations available
// so projects extending GoReleaser are able to use it, namely, GoDownloader.
package defaults

import (
	"fmt"

	"github.com/windmeup/goreleaser/v2/internal/pipe/archive"
	"github.com/windmeup/goreleaser/v2/internal/pipe/artifactory"
	"github.com/windmeup/goreleaser/v2/internal/pipe/aur"
	"github.com/windmeup/goreleaser/v2/internal/pipe/aursources"
	"github.com/windmeup/goreleaser/v2/internal/pipe/blob"
	"github.com/windmeup/goreleaser/v2/internal/pipe/bluesky"
	"github.com/windmeup/goreleaser/v2/internal/pipe/brew"
	"github.com/windmeup/goreleaser/v2/internal/pipe/build"
	"github.com/windmeup/goreleaser/v2/internal/pipe/changelog"
	"github.com/windmeup/goreleaser/v2/internal/pipe/checksums"
	"github.com/windmeup/goreleaser/v2/internal/pipe/chocolatey"
	"github.com/windmeup/goreleaser/v2/internal/pipe/discord"
	"github.com/windmeup/goreleaser/v2/internal/pipe/dist"
	"github.com/windmeup/goreleaser/v2/internal/pipe/docker"
	"github.com/windmeup/goreleaser/v2/internal/pipe/gomod"
	"github.com/windmeup/goreleaser/v2/internal/pipe/ko"
	"github.com/windmeup/goreleaser/v2/internal/pipe/krew"
	"github.com/windmeup/goreleaser/v2/internal/pipe/linkedin"
	"github.com/windmeup/goreleaser/v2/internal/pipe/mastodon"
	"github.com/windmeup/goreleaser/v2/internal/pipe/mattermost"
	"github.com/windmeup/goreleaser/v2/internal/pipe/milestone"
	"github.com/windmeup/goreleaser/v2/internal/pipe/nfpm"
	"github.com/windmeup/goreleaser/v2/internal/pipe/nix"
	"github.com/windmeup/goreleaser/v2/internal/pipe/notary"
	"github.com/windmeup/goreleaser/v2/internal/pipe/opencollective"
	"github.com/windmeup/goreleaser/v2/internal/pipe/project"
	"github.com/windmeup/goreleaser/v2/internal/pipe/reddit"
	"github.com/windmeup/goreleaser/v2/internal/pipe/release"
	"github.com/windmeup/goreleaser/v2/internal/pipe/sbom"
	"github.com/windmeup/goreleaser/v2/internal/pipe/scoop"
	"github.com/windmeup/goreleaser/v2/internal/pipe/sign"
	"github.com/windmeup/goreleaser/v2/internal/pipe/slack"
	"github.com/windmeup/goreleaser/v2/internal/pipe/smtp"
	"github.com/windmeup/goreleaser/v2/internal/pipe/snapcraft"
	"github.com/windmeup/goreleaser/v2/internal/pipe/snapshot"
	"github.com/windmeup/goreleaser/v2/internal/pipe/sourcearchive"
	"github.com/windmeup/goreleaser/v2/internal/pipe/teams"
	"github.com/windmeup/goreleaser/v2/internal/pipe/telegram"
	"github.com/windmeup/goreleaser/v2/internal/pipe/twitter"
	"github.com/windmeup/goreleaser/v2/internal/pipe/universalbinary"
	"github.com/windmeup/goreleaser/v2/internal/pipe/upload"
	"github.com/windmeup/goreleaser/v2/internal/pipe/upx"
	"github.com/windmeup/goreleaser/v2/internal/pipe/webhook"
	"github.com/windmeup/goreleaser/v2/internal/pipe/winget"
	"github.com/windmeup/goreleaser/v2/pkg/context"
)

// Defaulter can be implemented by a Piper to set default values for its
// configuration.
type Defaulter interface {
	fmt.Stringer

	// Default sets the configuration defaults
	Default(ctx *context.Context) error
}

// Defaulters is the list of defaulters.
//
//nolint:gochecknoglobals
var Defaulters = []Defaulter{
	dist.Pipe{},
	snapshot.Pipe{},
	release.Pipe{},
	project.Pipe{},
	changelog.Pipe{},
	gomod.Pipe{},
	build.Pipe{},
	universalbinary.Pipe{},
	sign.BinaryPipe{},
	notary.MacOS{},
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
	aursources.Pipe{},
	nix.Pipe{},
	winget.Pipe{},
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
	bluesky.Pipe{},
}
