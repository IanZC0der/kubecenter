package metrics

import (
	"context"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/util"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

var (
	metricsSvc Service
)

type KubecenterCollector struct {
	clusterCpu    prometheus.Gauge
	clusterMemory prometheus.Gauge
}

func (k KubecenterCollector) Describe(descs chan<- *prometheus.Desc) {
	k.clusterCpu.Describe(descs)
	k.clusterMemory.Describe(descs)
}

func (k KubecenterCollector) Collect(metrics chan<- prometheus.Metric) {
	metricsSvc = ioc.DefaultControllerContainer().Get(AppName).(Service)
	usageInfo := metricsSvc.GetClusterUsageInfo(context.Background())
	for _, item := range usageInfo {
		switch item.Label {
		case util.CLUSTER_CPU:
			newValue, _ := strconv.ParseFloat(item.Value, 64)
			k.clusterCpu.Set(newValue)
			k.clusterCpu.Collect(metrics)
		case util.CLUSTER_MEMORY:
			newValue, _ := strconv.ParseFloat(item.Value, 64)
			k.clusterMemory.Set(newValue)
			k.clusterMemory.Collect(metrics)
		}

	}
}

func NewKubecenterCollector() *KubecenterCollector {
	return &KubecenterCollector{
		clusterCpu: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: util.CLUSTER_CPU,
				Help: "collector cluster cpu info",
			},
		),
		clusterMemory: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: util.CLUSTER_MEMORY,
				Help: "collector cluster memory info",
			},
		),
	}

}

func init() {
	//metricsSvc = ioc.DefaultControllerContainer().Get(AppName).(Service)
	prometheus.MustRegister(NewKubecenterCollector())
}
