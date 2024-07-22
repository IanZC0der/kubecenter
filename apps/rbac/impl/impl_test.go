package impl_test

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/rbac"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

var (
	rbacSvc rbac.Service
	ctx     = context.Background()
)

func init() {
	rbacSvc = ioc.DefaultControllerContainer().Get(rbac.AppName).(rbac.Service)
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

func TestRBACServiceImpl_GetServiceAccountList(t *testing.T) {
	namespace := ""
	keyword := ""
	res, err := rbacSvc.GetServiceAccountList(ctx, namespace, keyword)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
