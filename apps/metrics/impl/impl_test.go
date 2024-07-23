package impl_test

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/metrics"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

var (
	metricsSvc metrics.Service
	ctx        = context.Background()
)

func init() {
	metricsSvc = ioc.DefaultControllerContainer().Get(metrics.AppName).(metrics.Service)
	kubeconfig := "../../../.kube/config"

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	global.KubeConfigSet = clientset

}

func TestMetricsServiceImpl_GetClusterInfo(t *testing.T) {
	res := metricsSvc.GetClusterInfo(ctx)
	t.Log(res)
}
