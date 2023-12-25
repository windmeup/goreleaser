package testlib

import (
	"testing"

	"github.com/windmeup/goreleaser/internal/pipe"
)

func TestAssertSkipped(t *testing.T) {
	AssertSkipped(t, pipe.Skip("skip"))
}
