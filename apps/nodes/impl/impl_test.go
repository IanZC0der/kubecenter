package impl_test

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/nodes"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

var (
	nodesSvc nodes.Service
	ctx      = context.Background()
)

func init() {
	nodesSvc = ioc.DefaultControllerContainer().Get(nodes.AppName).(nodes.Service)
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

func TestNodesServiceImpl_GetNodeList(t *testing.T) {
	nl, err := nodesSvc.GetNodeList(ctx, "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(nl)
}

func TestNodesServiceImpl_GetNodeDetail(t *testing.T) {
	n, err := nodesSvc.GetNodeDetail(ctx, "ubuntu-s-2vcpu-4gb-sfo3-01")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(n)
}
