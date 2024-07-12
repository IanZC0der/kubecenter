package nodes

import "context"

const (
	AppName = "nodes"
)

type Service interface {
	GetNodeList(context.Context, string) ([]*Node, error)
	GetNodeDetail(context.Context, string) (*Node, error)
	UpdateLabel(context.Context, *UpdateLabelRequest) error
	UpdateTaints(ctx context.Context, request *UpdateTaintRequest) error
}
