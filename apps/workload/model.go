package workload

import (
	"github.com/IanZC0der/kubecenter/apps/persistentvolume"
	"github.com/IanZC0der/kubecenter/apps/pods"
	"github.com/IanZC0der/kubecenter/util"
)

type StatefulSetBaseInfo struct {
	Name                 string           `json:"name"`
	Namespace            string           `json:"namespace"`
	Replicas             int32            `json:"replicas"`
	Labels               []*util.ListItem `json:"labels"`
	Selector             []*util.ListItem `json:"selector"`
	ServiceName          string           `json:"serviceName"`
	VolumeClaimTemplates []*persistentvolume.PersistentVolumeClaim
}

type StatefulSet struct {
	*StatefulSetBaseInfo
	Template *pods.Pod
}

func NewStatefulSet() *StatefulSet {
	return &StatefulSet{
		Template: pods.NewPod(),
		StatefulSetBaseInfo: &StatefulSetBaseInfo{
			Labels:               make([]*util.ListItem, 0),
			Selector:             make([]*util.ListItem, 0),
			VolumeClaimTemplates: make([]*persistentvolume.PersistentVolumeClaim, 0),
		},
	}
}

type StatefulSetResponse struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace"`
	Ready     int32  `json:"ready"`
	Replicas  int32  `json:"replicas"`
	Age       int64  `json:"age"`
}

func NewStatefulSetResponse() *StatefulSetResponse {
	return &StatefulSetResponse{}
}
