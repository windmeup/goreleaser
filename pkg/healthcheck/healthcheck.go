// Package healthcheck checks for missing binaries that the user needs to
// install.
package healthcheck

import (
	"fmt"

	"github.com/windmeup/goreleaser/internal/pipe/chocolatey"
	"github.com/windmeup/goreleaser/internal/pipe/docker"
	"github.com/windmeup/goreleaser/internal/pipe/sbom"
	"github.com/windmeup/goreleaser/internal/pipe/sign"
	"github.com/windmeup/goreleaser/internal/pipe/snapcraft"
	"github.com/windmeup/goreleaser/pkg/context"
)

// Healthchecker should be implemented by pipes that want checks.
type Healthchecker interface {
	fmt.Stringer

	// Dependencies return the binaries of the dependencies needed.
	Dependencies(ctx *context.Context) []string
}

// Healthcheckers is the list of healthchekers.
// nolint: gochecknoglobals
var Healthcheckers = []Healthchecker{
	system{},
	snapcraft.Pipe{},
	sign.Pipe{},
	sign.DockerPipe{},
	sbom.Pipe{},
	docker.Pipe{},
	docker.ManifestPipe{},
	chocolatey.Pipe{},
}

type system struct{}

func (system) String() string                           { return "system" }
func (system) Dependencies(_ *context.Context) []string { return []string{"git", "go"} }
