package secret

import (
	"github.com/IanZC0der/kubecenter/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetSecretResponseFromK8SSecret(k8sSecret *corev1.Secret) *SecretResponse {
	res := NewSecretResponse()
	res.Name = k8sSecret.Name
	res.Namespace = k8sSecret.Namespace
	res.Type = k8sSecret.Type
	res.Age = k8sSecret.CreationTimestamp.Unix()
	res.DataNum = len(k8sSecret.Data)
	res.Labels = util.ListConverter(k8sSecret.Labels)
	res.Data = util.ListConverterFromByte(k8sSecret.Data)
	return res
}

func GetK8SSecretFromSecretReq(sc *Secret) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sc.Name,
			Namespace: sc.Namespace,
			Labels:    util.MapConverter(sc.Labels),
		},
		Type:       sc.Type,
		StringData: util.MapConverter(sc.Data),
	}
}
