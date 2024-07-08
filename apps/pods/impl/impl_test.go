package impl

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/pods"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

var (
	podsSvc pods.Service
	ctx     = context.Background()
)

func init() {
	podsSvc = ioc.DefaultControllerContainer().Get(pods.AppName).(pods.Service)
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
func TestPodsServerImpl_GetPods(t *testing.T) {
	pl := pods.NewPodsList()

	pl, err := podsSvc.GetPods(ctx)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(pl)
}

func TestPodsServerImpl_GetNamespaceList(t *testing.T) {
	nl := pods.NewNamespaceList()
	nl, err := podsSvc.GetNamespaceList(ctx)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(nl)
}

func TestPodsServerImpl_GetPodsListUnderNamespaceWithKeyword(t *testing.T) {
	podsList := pods.NewPodsItemsList()
	ns := "kube-system"
	podsList, err := podsSvc.GetPodsListUnderNamespaceWithKeyword(ctx, ns, "")

	if err != nil {
		t.Fatal(err)
	}
	t.Log(podsList)
}
