package pods

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
)

// type podsUtil struct {
// }
func GetPodListItemFromPod(pod *corev1.Pod) *PodListItem {
	var total, ready, restart int32

	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.Ready {
			ready++
		}
		restart += containerStatus.RestartCount
		total++
	}
	var podStatus string

	if pod.Status.Phase != "Running" {
		podStatus = "Error"
	} else {
		podStatus = "Running"
	}

	return &PodListItem{
		Name:     pod.Name,
		Ready:    fmt.Sprintf("%d/%d", ready, total),
		Status:   podStatus,
		Restarts: restart,
		Age:      pod.CreationTimestamp.Unix(),
		IP:       pod.Status.PodIP,
		Node:     pod.Spec.NodeName,
	}

}
