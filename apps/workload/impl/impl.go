package impl

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/persistentvolume"
	"github.com/IanZC0der/kubecenter/apps/pods"
	"github.com/IanZC0der/kubecenter/apps/workload"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type WorkloadServiceImpl struct{}

var _ workload.Service = &WorkloadServiceImpl{}

func init() {
	ioc.DefaultControllerContainer().Register(&WorkloadServiceImpl{})
}

func (w *WorkloadServiceImpl) Init() error {
	return nil
}

func (w *WorkloadServiceImpl) Name() string {
	return workload.AppName
}

func (w *WorkloadServiceImpl) GetStatefulSetDetail(ctx context.Context, namespace, name string) (*workload.StatefulSet, error) {
	res := workload.NewStatefulSet()

	statefulsetK8S, err := global.KubeConfigSet.AppsV1().StatefulSets(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	pvcList := make([]*persistentvolume.PersistentVolumeClaim, len(statefulsetK8S.Spec.VolumeClaimTemplates))

	// get pvc
	for i, item := range statefulsetK8S.Spec.VolumeClaimTemplates {
		pvcList[i] = &persistentvolume.PersistentVolumeClaim{
			Name:             item.Name,
			AccessModes:      item.Spec.AccessModes,
			Capacity:         int32(item.Spec.Resources.Requests.Storage().Value() / (1024 * 1024)),
			StorageClassName: *item.Spec.StorageClassName,
		}
	}

	podInStatefulset := pods.GetPodInfoFromPod(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: statefulsetK8S.Spec.Template.Labels,
		},
		Spec: statefulsetK8S.Spec.Template.Spec,
	})

	res.Name = statefulsetK8S.Name
	res.Namespace = statefulsetK8S.Namespace
	res.Replicas = *statefulsetK8S.Spec.Replicas
	res.Labels = util.ListConverter(statefulsetK8S.Labels)
	res.Selector = util.ListConverter(statefulsetK8S.Spec.Selector.MatchLabels)
	res.ServiceName = statefulsetK8S.Spec.ServiceName
	res.VolumeClaimTemplates = pvcList
	res.Template = podInStatefulset

	return res, nil
}
func (w *WorkloadServiceImpl) GetStatefulSetList(ctx context.Context, namespace, keyword string) ([]*workload.StatefulSetResponse, error) {
	res := make([]*workload.StatefulSetResponse, 0)

	list, err := global.KubeConfigSet.AppsV1().StatefulSets(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range list.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		newResponse := workload.NewStatefulSetResponse()
		newResponse.Name = item.Name
		newResponse.Namespace = item.Namespace
		newResponse.Replicas = item.Status.Replicas
		newResponse.Ready = item.Status.ReadyReplicas
		newResponse.Age = item.CreationTimestamp.Unix()
		res = append(res, newResponse)
	}

	return res, nil
}
func (w *WorkloadServiceImpl) CreateOrUpdateStatefulSet(ctx context.Context, req *workload.StatefulSet) (*workload.StatefulSetResponse, error) {
	return nil, nil
}
func (w *WorkloadServiceImpl) DeleteStatefulSet(ctx context.Context, namespace, name string) error {
	return nil
}
