package impl_test

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/persistentvolume"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

var (
	pvSvc persistentvolume.Service
	ctx   = context.Background()
)

func init() {
	pvSvc = ioc.DefaultControllerContainer().Get(persistentvolume.AppName).(persistentvolume.Service)
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

func TestPersistentVolumeServiceImpl_GetPVList(t *testing.T) {
	l, err := pvSvc.GetPVList(ctx, "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(l)
}
