package metrics

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

type MetricsItem struct {
	Name  string `json:"title"`
	Label string `json:"label"`
	Value string `json:"value"`
	// in RGB, easier for the frontend to present the data
	Color string `json:"color"`
}

type NodeMetrics struct {
	Metadata  metav1.ObjectMeta   `json:"metadata"`
	Timestamp time.Time           `json:"timestamp"`
	Window    string              `json:"window"`
	Usage     corev1.ResourceList `json:"usage"`
}

type NodeMetricsList struct {
	Kind       string          `json:"kind"`
	ApiVersion string          `json:"apiVersion"`
	Metadata   metav1.ListMeta `json:"metadata"`
	Items      []*NodeMetrics  `json:"items"`
}

func NewNodeMetricsList() *NodeMetricsList {
	return &NodeMetricsList{
		Items: make([]*NodeMetrics, 0),
	}
}
