// Package pipeline provides generic errors for pipes to use.
package pipeline

import (
	"fmt"

	"github.com/windmeup/goreleaser/internal/pipe/announce"
	"github.com/windmeup/goreleaser/internal/pipe/archive"
	"github.com/windmeup/goreleaser/internal/pipe/aur"
	"github.com/windmeup/goreleaser/internal/pipe/before"
	"github.com/windmeup/goreleaser/internal/pipe/brew"
	"github.com/windmeup/goreleaser/internal/pipe/build"
	"github.com/windmeup/goreleaser/internal/pipe/changelog"
	"github.com/windmeup/goreleaser/internal/pipe/checksums"
	"github.com/windmeup/goreleaser/internal/pipe/chocolatey"
	"github.com/windmeup/goreleaser/internal/pipe/defaults"
	"github.com/windmeup/goreleaser/internal/pipe/dist"
	"github.com/windmeup/goreleaser/internal/pipe/docker"
	"github.com/windmeup/goreleaser/internal/pipe/effectiveconfig"
	"github.com/windmeup/goreleaser/internal/pipe/env"
	"github.com/windmeup/goreleaser/internal/pipe/git"
	"github.com/windmeup/goreleaser/internal/pipe/gomod"
	"github.com/windmeup/goreleaser/internal/pipe/krew"
	"github.com/windmeup/goreleaser/internal/pipe/metadata"
	"github.com/windmeup/goreleaser/internal/pipe/nfpm"
	"github.com/windmeup/goreleaser/internal/pipe/nix"
	"github.com/windmeup/goreleaser/internal/pipe/prebuild"
	"github.com/windmeup/goreleaser/internal/pipe/publish"
	"github.com/windmeup/goreleaser/internal/pipe/reportsizes"
	"github.com/windmeup/goreleaser/internal/pipe/sbom"
	"github.com/windmeup/goreleaser/internal/pipe/scoop"
	"github.com/windmeup/goreleaser/internal/pipe/semver"
	"github.com/windmeup/goreleaser/internal/pipe/sign"
	"github.com/windmeup/goreleaser/internal/pipe/snapcraft"
	"github.com/windmeup/goreleaser/internal/pipe/snapshot"
	"github.com/windmeup/goreleaser/internal/pipe/sourcearchive"
	"github.com/windmeup/goreleaser/internal/pipe/universalbinary"
	"github.com/windmeup/goreleaser/internal/pipe/upx"
	"github.com/windmeup/goreleaser/internal/pipe/winget"
	"github.com/windmeup/goreleaser/pkg/context"
)

// Piper defines a pipe, which can be part of a pipeline (a series of pipes).
type Piper interface {
	fmt.Stringer

	// Run the pipe
	Run(ctx *context.Context) error
}

// BuildPipeline contains all build-related pipe implementations in order.
// nolint:gochecknoglobals
var BuildPipeline = []Piper{
	// load and validate environment variables
	env.Pipe{},
	// get and validate git repo state
	git.Pipe{},
	// parse current tag to a semver
	semver.Pipe{},
	// load default configs
	defaults.Pipe{},
	// snapshot version handling
	snapshot.Pipe{},
	// run global hooks before build
	before.Pipe{},
	// ensure ./dist is clean
	dist.Pipe{},
	// setup gomod-related stuff
	gomod.Pipe{},
	// run prebuild stuff
	prebuild.Pipe{},
	// proxy gomod if needed
	gomod.CheckGoModPipe{},
	// proxy gomod if needed
	gomod.ProxyPipe{},
	// writes the actual config (with defaults et al set) to dist
	effectiveconfig.Pipe{},
	// build
	build.Pipe{},
	// universal binary handling
	universalbinary.Pipe{},
	// upx
	upx.Pipe{},
}

// BuildCmdPipeline is the pipeline run by goreleaser build.
// nolint:gochecknoglobals
var BuildCmdPipeline = append(
	BuildPipeline,
	reportsizes.Pipe{},
	metadata.Pipe{},
)

// Pipeline contains all pipe implementations in order.
// nolint: gochecknoglobals
var Pipeline = append(
	BuildPipeline,
	// builds the release changelog
	changelog.Pipe{},
	// archive in tar.gz, zip or binary (which does no archiving at all)
	archive.Pipe{},
	// archive the source code using git-archive
	sourcearchive.Pipe{},
	// archive via fpm (deb, rpm) using "native" go impl
	nfpm.Pipe{},
	// archive via snapcraft (snap)
	snapcraft.Pipe{},
	// create SBOMs of artifacts
	sbom.Pipe{},
	// checksums of the files
	checksums.Pipe{},
	// sign artifacts
	sign.Pipe{},
	// create arch linux aur pkgbuild
	aur.Pipe{},
	// create nixpkgs
	nix.NewBuild(),
	// winget installers
	winget.Pipe{},
	// create brew tap
	brew.Pipe{},
	// krew plugins
	krew.Pipe{},
	// create scoop buckets
	scoop.Pipe{},
	// create chocolatey pkg and publish
	chocolatey.Pipe{},
	// reports artifacts sizes to the log and to artifacts.json
	reportsizes.Pipe{},
	// create and push docker images
	docker.Pipe{},
	// publishes artifacts
	publish.New(),
	// creates a metadata.json and an artifacts.json files in the dist folder
	metadata.Pipe{},
	// announce releases
	announce.Pipe{},
)
