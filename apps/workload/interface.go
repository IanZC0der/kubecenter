package workload

import "context"

const (
	AppName = "workload"
)

type Workload interface {
	GetStatefulSetDetail(ctx context.Context, namespace, name string) (*StatefulSet, error)
	GetStatefulSetList(ctx context.Context, namespace, keyword string) ([]*StatefulSetResponse, error)
	CreateOrUpdateStatefulSet(ctx context.Context, req *StatefulSet) (*StatefulSetResponse, error)
	DeleteStatefulSet(ctx context.Context, namespace, name string) error
}
