// Package blob provides the pipe implementation that uploads files to "blob" providers, such as s3, gcs and azure.
package blob

import (
	"errors"

	"github.com/windmeup/goreleaser/v2/internal/pipe"
	"github.com/windmeup/goreleaser/v2/internal/semerrgroup"
	"github.com/windmeup/goreleaser/v2/internal/tmpl"
	"github.com/windmeup/goreleaser/v2/pkg/context"
)

// Pipe for blobs.
type Pipe struct{}

// String returns the description of the pipe.
func (Pipe) String() string                 { return "blobs" }
func (Pipe) Skip(ctx *context.Context) bool { return len(ctx.Config.Blobs) == 0 }

// Default sets the pipe defaults.
func (Pipe) Default(ctx *context.Context) error {
	for i := range ctx.Config.Blobs {
		blob := &ctx.Config.Blobs[i]
		if blob.Bucket == "" || blob.Provider == "" {
			return errors.New("bucket or provider cannot be empty")
		}
		if blob.Directory == "" {
			blob.Directory = "{{ .ProjectName }}/{{ .Tag }}"
		}

		switch blob.ContentDisposition {
		case "":
			blob.ContentDisposition = "attachment;filename={{.Filename}}"
		case "-":
			blob.ContentDisposition = ""
		}
	}
	return nil
}

// Publish to specified blob bucket url.
func (Pipe) Publish(ctx *context.Context) error {
	g := semerrgroup.New(ctx.Parallelism)
	skips := pipe.SkipMemento{}
	for _, conf := range ctx.Config.Blobs {
		g.Go(func() error {
			b, err := tmpl.New(ctx).Bool(conf.Disable)
			if err != nil {
				return err
			}
			if b {
				skips.Remember(pipe.Skip("configuration is disabled"))
				return nil
			}
			return doUpload(ctx, conf)
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return skips.Evaluate()
}
