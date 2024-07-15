package impl_test

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/secret"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

var (
	secretSvc secret.Service
	ctx       = context.Background()
)

func init() {
	secretSvc = ioc.DefaultControllerContainer().Get(secret.AppName).(secret.Service)
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

func TestSecretServiceImpl_GetSecretsList(t *testing.T) {
	l, err := secretSvc.GetSecretsList(ctx, "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(l)
}
