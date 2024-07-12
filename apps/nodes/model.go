package nodes

import (
	"encoding/json"
	corev1 "k8s.io/api/core/v1"
)

type ListItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (lI *ListItem) String() string {
	jsonlI, _ := json.Marshal(lI)
	return string(jsonlI)

}

type Node struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	Age        int64  `json:"age"`
	InternalIp string `json:"internalIp"`
	ExternalIp string `json:"externalIp"`

	Version          string          `json:"version"`
	OsImage          string          `json:"osImage"`
	KernelVersion    string          `json:"kernelVersion"`
	ContainerRuntime string          `json:"containerRuntime"`
	Labels           []*ListItem     `json:"labels"`
	Taints           []*corev1.Taint `json:"taints"`
}

func NewNode() *Node {
	return &Node{
		Labels: make([]*ListItem, 0),
		Taints: make([]*corev1.Taint, 0),
	}
}

type UpdateLabelRequest struct {
	Name   string      `json:"name"`
	Labels []*ListItem `json:"labels"`
}

func NewUpdateLabelRequest() *UpdateLabelRequest {
	return &UpdateLabelRequest{
		Labels: make([]*ListItem, 0),
	}
}

type UpdateTaintRequest struct {
	Name   string          `json:"name"`
	Taints []*corev1.Taint `json:"taints"`
}

func NewUpdateTaintRequest() *UpdateTaintRequest {
	return &UpdateTaintRequest{
		Taints: make([]*corev1.Taint, 0),
	}
}
