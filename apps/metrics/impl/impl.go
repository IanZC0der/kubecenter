package impl

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/metrics"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

func init() {
	ioc.DefaultControllerContainer().Register(&MetricsServiceImpl{})
}

func (m *MetricsServiceImpl) Init() error {
	return nil
}

func (m *MetricsServiceImpl) Name() string {
	return metrics.AppName
}

var _ metrics.Service = &MetricsServiceImpl{}

type MetricsServiceImpl struct{}

func (m *MetricsServiceImpl) GetClusterInfo(ctx context.Context) []*metrics.MetricsItem {
	result := make([]*metrics.MetricsItem, 0)

	// get  the cluster creation time by finding the master node, the master created the earliest should be the creation time of the cluster
	list, err := global.KubeConfigSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err == nil {
		var creationTime int64 = 0
		for _, item := range list.Items {
			if _, ok := item.Labels["node-role.kubernetes.io/control-plane"]; ok {
				if creationTime == 0 || (creationTime > 0 && item.CreationTimestamp.Unix() < creationTime) {
					creationTime = item.CreationTimestamp.Unix()
				}
			}
		}
		formarttedTime := util.FormatTime(creationTime)
		result = append(result, &metrics.MetricsItem{
			Name: "Cluster Creation Time",
			//Label: "Creation Time",
			Value: formarttedTime,
		})

	}

	// get number of nodes

	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Name:  "Nodes",
			Value: strconv.Itoa(len(list.Items)),
		})
	}

	// add a color value in each item so  that the frontend can use the value to present the data
	for _, item := range result {
		item.Color = util.GenerateRGB(item.Name)
	}
	return result
}
