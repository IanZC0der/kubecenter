package impl

import (
	"context"
	"fmt"
	"github.com/IanZC0der/kubecenter/apps/persistentvolume"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"strings"
)

const (
	NFS_VOLUMESOURCE_TYPE = "nfs"
)

func init() {
	ioc.DefaultControllerContainer().Register(&PersistentVolumeServiceImpl{})
}

var _ persistentvolume.Service = &PersistentVolumeServiceImpl{}

type PersistentVolumeServiceImpl struct{}

func (p *PersistentVolumeServiceImpl) Init() error {
	return nil
}

func (p *PersistentVolumeServiceImpl) Name() string {
	return persistentvolume.AppName
}

func (p *PersistentVolumeServiceImpl) GetPVList(ctx context.Context, keyword string) ([]*corev1.PersistentVolume, error) {
	pvList, err := global.KubeConfigSet.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	res := make([]*corev1.PersistentVolume, 0)

	for _, item := range pvList.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		res = append(res, &item)
	}
	return res, nil
}

func (p *PersistentVolumeServiceImpl) DeletePV(ctx context.Context, name string) error {
	err := global.KubeConfigSet.CoreV1().PersistentVolumes().Delete(ctx, name, metav1.DeleteOptions{})
	return err
}

func (p *PersistentVolumeServiceImpl) CreatePV(ctx context.Context, req *persistentvolume.PersistentVolumeReq) (*corev1.PersistentVolume, error) {
	var volumeSource corev1.PersistentVolumeSource
	//nfs source only
	if req.VolumeSource.Type != NFS_VOLUMESOURCE_TYPE {
		return nil, fmt.Errorf("invalid volume type: %s", req.VolumeSource.Type)
	} else {
		volumeSource.NFS = &corev1.NFSVolumeSource{
			Server:   req.VolumeSource.NfsVolumeSource.NfsServer,
			Path:     req.VolumeSource.NfsVolumeSource.NfsPath,
			ReadOnly: req.VolumeSource.NfsVolumeSource.NfsReadOnly,
		}
	}

	newReq := &corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:   req.Name,
			Labels: util.MapConverter(req.Labels),
		},
		Spec: corev1.PersistentVolumeSpec{
			Capacity: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceStorage: resource.MustParse(strconv.Itoa(int(req.Capacity)) + "Mi"),
			},
			AccessModes:                   req.AccessModes,
			PersistentVolumeSource:        volumeSource,
			PersistentVolumeReclaimPolicy: req.ReClaimPolicy,
		},
	}

	res, err := global.KubeConfigSet.CoreV1().PersistentVolumes().Create(ctx, newReq, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PersistentVolumeServiceImpl) DeletePVC(ctx context.Context, namespace, name string) error {
	err := global.KubeConfigSet.CoreV1().PersistentVolumeClaims(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	return err
}

func (p *PersistentVolumeServiceImpl) GetPVCList(ctx context.Context, namespace, name string) ([]*corev1.PersistentVolumeClaim, error) {
	pvcList, err := global.KubeConfigSet.CoreV1().PersistentVolumeClaims(namespace).List(ctx, metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	res := make([]*corev1.PersistentVolumeClaim, 0)
	for _, item := range pvcList.Items {
		if !strings.Contains(item.Name, name) {
			continue
		}
		res = append(res, &item)
	}
	return res, nil
}

func (p *PersistentVolumeServiceImpl) CreatePVC(ctx context.Context, req *persistentvolume.PersistentVolumeClaim) (*corev1.PersistentVolumeClaim, error) {
	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
			Labels:    util.MapConverter(req.Labels),
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: util.MapConverter(req.Labels),
			},
			AccessModes: req.AccessModes,
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse(strconv.Itoa(int(req.Capacity)) + "Mi"),
				},
			},
			StorageClassName: &req.StorageClassName,
		},
	}
	if pvc.Spec.StorageClassName != nil {
		pvc.Spec.Selector = nil
	}
	res, err := global.KubeConfigSet.CoreV1().PersistentVolumeClaims(req.Namespace).Create(ctx, pvc, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
