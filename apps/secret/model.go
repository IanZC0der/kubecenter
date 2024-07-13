package secret

import (
	"github.com/IanZC0der/kubecenter/util"
	corev1 "k8s.io/api/core/v1"
)

type Secret struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Type      corev1.SecretType `json:"type"`
	Labels    []*util.ListItem  `json:"labels"`
	Data      []*util.ListItem  `json:"data"`
}

func NewSecret() *Secret {
	return &Secret{
		Labels: make([]*util.ListItem, 0),
		Data:   make([]*util.ListItem, 0),
	}
}

type SecretResponse struct {
	*Secret
	DataNum int   `json:"dataNum"`
	Age     int64 `json:"age"`
}

func NewSecretResponse() *SecretResponse {
	return &SecretResponse{
		Secret: NewSecret(),
	}
}
