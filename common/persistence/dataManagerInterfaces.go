package persistence

import (
	"context"
	"time"
)

type (

	// Closeable is an interface for any entity that supports a close operation to release resources
	Closeable interface {
		Close()
	}

	NodeManager interface {
		Closeable
		GetName() string
		CreateNode(ctx context.Context, request *CreateNodeRequest) (*CreateNodeResponse, error)
		DeleteNode(ctx context.Context, request *DeleteNodeRequest) (*DeleteNodeResponse, error)
		GetNode(ctx context.Context, request *GetNodeRequest) (*GetNodesResponse, error)
	}

	NodeInfo struct {
		ListID      string
		NodeID      int64
		NextNodeID  int64
		PrevNodeId  int64
		CreatedTime time.Time
	}

	CreateNodeRequest struct {
	}
	CreateNodeResponse struct {
	}
	DeleteNodeRequest struct {
	}
	DeleteNodeResponse struct {
	}
	GetNodeRequest struct {
	}

	GetNodesResponse struct {
		Nodes []*NodeInfo
	}
)
