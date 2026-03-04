package lifecycle

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ErrorReporter is called when a subroutine returns an error.
type ErrorReporter interface {
	Report(ctx context.Context, err error, info ErrorInfo)
}

// ErrorInfo provides context about the subroutine that failed.
type ErrorInfo struct {
	Subroutine string
	Object     client.Object
	Action     string // "process", "finalize", "initialize", "terminate"
}
