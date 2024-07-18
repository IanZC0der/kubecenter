package workload

import "context"

const (
	AppName = "workload"
)

type Service interface {
	GetStatefulSetDetail(ctx context.Context, namespace, name string) (*StatefulSet, error)
	GetStatefulSetList(ctx context.Context, namespace, keyword string) ([]*StatefulSetResponse, error)
	CreateOrUpdateStatefulSet(ctx context.Context, req *StatefulSet) (*StatefulSetResponse, string, error)
	DeleteStatefulSet(ctx context.Context, namespace, name string) error
}
