package persistence

import "context"

type (
	NodeStore interface {
		Closeable
		LeaseNodeList(ctx context.Context)
	}
)
