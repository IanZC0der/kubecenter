package persistentvolume

import (
	"github.com/IanZC0der/kubecenter/util"
	corev1 "k8s.io/api/core/v1"
)

type Base struct {
	Name          string                               `json:"name"`
	Labels        []*util.ListItem                     `json:"labels"`
	Capacity      int32                                `json:"capacity"`
	AccessModes   []corev1.PersistentVolumeAccessMode  `json:"accessModes"`
	ReClaimPolicy corev1.PersistentVolumeReclaimPolicy `json:"reClaimPolicy"`
}

type PersistentVolumeRes struct {
	*Base
	Status corev1.PersistentVolumePhase `json:"status"`
	Claim  string                       `json:"claim"`
	Age    int64                        `json:"age"`
	Reason string                       `json:"reason"`
}

func NewPersistentVolumeRes() *PersistentVolumeRes {
	return &PersistentVolumeRes{
		Base: &Base{
			Labels:      make([]*util.ListItem, 0),
			AccessModes: make([]corev1.PersistentVolumeAccessMode, 0),
		},
	}
}

type PersistentVolumeReq struct {
	Name          string                               `json:"name"`
	Labels        []*util.ListItem                     `json:"labels"`
	Capacity      int32                                `json:"capacity"`
	AccessModes   []corev1.PersistentVolumeAccessMode  `json:"accessModes"`
	ReClaimPolicy corev1.PersistentVolumeReclaimPolicy `json:"reClaimPolicy"`
	VolumeSource  *VolumeSource                        `json:"volumeSource"`
}

func NewPersistentVolumeReq() *PersistentVolumeReq {
	return &PersistentVolumeReq{

		Labels:      make([]*util.ListItem, 0),
		AccessModes: make([]corev1.PersistentVolumeAccessMode, 0),

		VolumeSource: &VolumeSource{
			NfsVolumeSource: &NfsVolumeSource{},
		},
	}
}

type VolumeSource struct {
	Type            string           `json:"type"`
	NfsVolumeSource *NfsVolumeSource `json:"nfsVolumeSource"`
}

type NfsVolumeSource struct {
	NfsPath     string `json:"nfsPath"`
	NfsServer   string `json:"nfsServer"`
	NfsReadOnly bool   `json:"nfsReadOnly"`
}

type PersistentVolumeClaim struct {
	Name             string                              `json:"name"`
	Namespace        string                              `json:"namespace"`
	Labels           []*util.ListItem                    `json:"labels"`
	AccessModes      []corev1.PersistentVolumeAccessMode `json:"accessModes"`
	Capacity         int32                               `json:"capacity"`
	Selector         []*util.ListItem                    `json:"selector"`
	StorageClassName string                              `json:"storageClassName"`
}

func NewCreatePersistentVolumeClaimReq() *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		Labels:   make([]*util.ListItem, 0),
		Selector: make([]*util.ListItem, 0),
	}
}
