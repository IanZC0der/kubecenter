package impl_test

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/k8sservice"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

var (
	k8ssvcService k8sservice.Service
	ctx           = context.Background()
)

func init() {
	k8ssvcService = ioc.DefaultControllerContainer().Get(k8sservice.AppName).(k8sservice.Service)
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

func TestK8sServiceImpl_GetSvcList(t *testing.T) {
	l, err := k8ssvcService.GetSvcList(ctx, "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(l)
}

func TestK8sServiceImpl_GetSvcDetail(t *testing.T) {
	namespace := "kube-system"
	name := "kube-dns"
	res, err := k8ssvcService.GetSvcDetail(ctx, namespace, name)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
