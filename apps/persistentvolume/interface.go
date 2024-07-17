package persistentvolume

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
)

const (
	AppName = "persistentvolume"
)

type Service interface {
	GetPVList(context.Context, string) ([]*corev1.PersistentVolume, error)
	DeletePV(context.Context, string) error
	CreatePV(context.Context, *PersistentVolumeReq) (*corev1.PersistentVolume, error)
	GetPVCList(context.Context, string, string) ([]*corev1.PersistentVolumeClaim, error)
	DeletePVC(context.Context, string, string) error
	CreatePVC(context.Context, *PersistentVolumeClaim) (*corev1.PersistentVolumeClaim, error)
	GetSCList(context.Context, string) ([]*storagev1.StorageClass, error)
	DeleteSC(context.Context, string) error
	CreateSC(context.Context, *StorageClass) (*storagev1.StorageClass, error)
}
