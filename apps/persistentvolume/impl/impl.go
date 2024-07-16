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
	if req.Type != NFS_VOLUMESOURCE_TYPE {
		return nil, fmt.Errorf("invalid volume type: %s", req.Type)
	} else {
		volumeSource.NFS = &corev1.NFSVolumeSource{
			Server:   req.NfsServer,
			Path:     req.NfsPath,
			ReadOnly: req.NfsReadOnly,
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
