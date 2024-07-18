package k8sservice

import (
	"github.com/IanZC0der/kubecenter/util"
	corev1 "k8s.io/api/core/v1"
)

type K8SService struct {
	Name      string             `json:"name"`
	Namespace string             `json:"namespace"`
	Labels    []*util.ListItem   `json:"labels"`
	Type      corev1.ServiceType `json:"type"`
	Selector  []*util.ListItem   `json:"selector"`
	Ports     []*ServicePort     `json:"ports"`
}

func NewK8SService() *K8SService {
	return &K8SService{
		Labels:   make([]*util.ListItem, 0),
		Selector: make([]*util.ListItem, 0),
		Ports:    make([]*ServicePort, 0),
	}
}

type ServicePort struct {
	Name       string `json:"name"`
	Port       int32  `json:"port"`
	TargetPort int32  `json:"targetPort"`
	NodePort   int32  `json:"nodePort"`
}
