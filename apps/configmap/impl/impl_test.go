package impl_test

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/configmap"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

var (
	configmapSvc configmap.Service
	ctx          = context.Background()
)

func init() {
	configmapSvc = ioc.DefaultControllerContainer().Get(configmap.AppName).(configmap.Service)
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

func TestConfigmapServiceImpl_GetConfigMapList(t *testing.T) {
	l, err := configmapSvc.GetConfigMapList(ctx, "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(l)
}

func TestConfigmapServiceImpl_GetConfigMap(t *testing.T) {
	m, err := configmapSvc.GetConfigMapDetail(ctx, "kube-public", "kube-root-ca.crt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(m)
}
