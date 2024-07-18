package impl_test

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/workload"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

var (
	workloadSvc workload.Service
	ctx         = context.Background()
)

func init() {
	workloadSvc = ioc.DefaultControllerContainer().Get(workload.AppName).(workload.Service)
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

func TestWorkloadServiceImpl_GetStatefulSetList(t *testing.T) {
	res, err := workloadSvc.GetStatefulSetList(ctx, "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)

}

func TestWorkloadServiceImpl_GetDeploymentList(t *testing.T) {
	res, err := workloadSvc.GetStatefulSetDetail(ctx, "", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
