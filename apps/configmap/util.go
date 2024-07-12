package configmap

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ConfigmapToK8SConfigmap(configmap *ConfigMap) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configmap.Name,
			Namespace: configmap.Namespace,
			Labels:    mapConverter(configmap.Labels),
		},
		Data: mapConverter(configmap.Data),
	}
}

func mapConverter(data []*ListItem) map[string]string {
	dataMap := make(map[string]string)
	for _, data := range data {
		dataMap[data.Key] = data.Value
	}
	return dataMap
}

func listConverter(data map[string]string) []*ListItem {
	l := make([]*ListItem, 0)
	for k, v := range data {
		l = append(l, &ListItem{Key: k, Value: v})
	}
	return l
}

func K8SConfigmapToConfigmap(configmap *corev1.ConfigMap) *ConfigMapResponse {
	return &ConfigMapResponse{
		Name:      configmap.Name,
		Namespace: configmap.Namespace,
		DataNum:   len(configmap.Data),
		Age:       configmap.CreationTimestamp.Unix(),
		Labels:    listConverter(configmap.Labels),
		Data:      listConverter(configmap.Data),
	}
}
