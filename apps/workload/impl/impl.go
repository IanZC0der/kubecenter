package impl

import (
	"context"
	"fmt"
	"github.com/IanZC0der/kubecenter/apps/persistentvolume"
	"github.com/IanZC0der/kubecenter/apps/pods"
	"github.com/IanZC0der/kubecenter/apps/workload"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/util"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
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
	pvcList := make([]*persistentvolume.PersistentVolumeClaim, 0)

	// get pvc
	for _, item := range statefulsetK8S.Spec.VolumeClaimTemplates {
		scName := ""
		if item.Spec.StorageClassName != nil {
			scName = *item.Spec.StorageClassName
		}
		pvcList = append(pvcList, &persistentvolume.PersistentVolumeClaim{
			Name:             item.Name,
			AccessModes:      item.Spec.AccessModes,
			Capacity:         int32(item.Spec.Resources.Requests.Storage().Value() / (1024 * 1024)),
			StorageClassName: scName,
		})
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
func (w *WorkloadServiceImpl) CreateOrUpdateStatefulSet(ctx context.Context, req *workload.StatefulSet) (*workload.StatefulSetResponse, string, error) {

	pvcTemplates := make([]corev1.PersistentVolumeClaim, len(req.VolumeClaimTemplates))

	for i, item := range req.VolumeClaimTemplates {
		pvcTemplates[i] = corev1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Name:   item.Name,
				Labels: util.MapConverter(item.Labels),
			},
			Spec: corev1.PersistentVolumeClaimSpec{
				AccessModes: item.AccessModes,
				Resources: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceStorage: resource.MustParse(strconv.Itoa(int(item.Capacity)) + "Mi"),
					},
				},
			},
		}
	}
	k8sPod := pods.CreatePodFromPodRequest(req.Template)
	statefulserK8S := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
			Labels:    util.MapConverter(req.Labels),
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas:    &req.Replicas,
			ServiceName: req.ServiceName,
			Selector: &metav1.LabelSelector{
				MatchLabels: util.MapConverter(req.Selector),
			},
			VolumeClaimTemplates: pvcTemplates,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: k8sPod.ObjectMeta,
				Spec:       k8sPod.Spec,
			},
		},
	}
	existing, err := global.KubeConfigSet.AppsV1().StatefulSets(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
	msg := ""
	var res *appsv1.StatefulSet
	if err != nil {
		//create
		res, err = global.KubeConfigSet.AppsV1().StatefulSets(req.Namespace).Create(ctx, statefulserK8S, metav1.CreateOptions{})
		if err != nil {
			msg = fmt.Sprintf("Create statefulset err: %s", err.Error())
			return nil, msg, err
		} else {
			msg = fmt.Sprintf("Create statefulset success: %s", res.Name)
		}
	} else {
		existing.Spec = statefulserK8S.Spec
		res, err = global.KubeConfigSet.AppsV1().StatefulSets(req.Namespace).Update(ctx, existing, metav1.UpdateOptions{})
		if err != nil {
			msg = fmt.Sprintf("Update statefulset err: %s", err.Error())
			return nil, msg, err
		} else {
			msg = fmt.Sprintf("Update statefulset success: %s", res.Name)
		}
	}
	workloadResp := &workload.StatefulSetResponse{
		Name:      res.Name,
		Namespace: res.Namespace,
		Replicas:  res.Status.Replicas,
		Ready:     res.Status.ReadyReplicas,
		Age:       res.CreationTimestamp.Unix(),
	}
	return workloadResp, msg, nil
}
func (w *WorkloadServiceImpl) DeleteStatefulSet(ctx context.Context, namespace, name string) error {
	return global.KubeConfigSet.AppsV1().StatefulSets(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}
