package pods

import (
	"context"
	corev1 "k8s.io/api/core/v1"
)

const (
	AppName = "pods"
)

type Service interface {
	GetPods(ctx context.Context) (*Pods, error)
	GetNamespaceList(ctx context.Context) (*NamespaceList, error)
	GetPodsListUnderNamespaceWithKeyword(ctx context.Context, namespace string, keyword string) (*PodsList, error)
	GetPodsListWithinNode(ctx context.Context, keyword string, nodeName string) (*PodsList, error)
	GetPodDetail(ctx context.Context, namespace string, name string) (*Pod, error)
	CreatePod(ctx context.Context, pod *Pod) (*corev1.Pod, error)
	UpdatePod(ctx context.Context, pod *Pod) (*corev1.Pod, string, error)
	DeletePod(ctx context.Context, namespace string, name string) (*corev1.Pod, string, error)
}
