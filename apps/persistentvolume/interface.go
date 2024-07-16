package persistentvolume

import (
	"context"
	corev1 "k8s.io/api/core/v1"
)

const (
	AppName = "persistentvolume"
)

type Service interface {
	GetPVList(context.Context, string) ([]*corev1.PersistentVolume, error)
	DeletePV(context.Context, string) error
	CreatePV(context.Context, *PersistentVolumeReq) (*corev1.PersistentVolume, error)
}
