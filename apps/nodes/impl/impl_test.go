package impl_test

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/nodes"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	corev1 "k8s.io/api/core/v1"
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

/*
"labels": [

	{
	    "key": "beta.kubernetes.io/arch",
	    "value": "amd64"
	},
	{
	    "key": "beta.kubernetes.io/os",
	    "value": "linux"
	},
	{
	    "key": "kubernetes.io/arch",
	    "value": "amd64"
	},
	{
	    "key": "kubernetes.io/hostname",
	    "value": "ubuntu-s-2vcpu-4gb-sfo3-02"
	},
	{
	    "key": "kubernetes.io/os",
	    "value": "linux"
	}

],
*/
func TestNodesServiceImpl_UpdateLabel(t *testing.T) {
	req := nodes.NewUpdateLabelRequest()
	req.Name = "ubuntu-s-2vcpu-4gb-sfo3-02"
	req.Labels = append(req.Labels, &nodes.ListItem{
		Key:   "test2",
		Value: "app2",
	})
	req.Labels = append(req.Labels, &nodes.ListItem{
		Key:   "test",
		Value: "app",
	})
	err := nodesSvc.UpdateLabel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	n, err := nodesSvc.GetNodeDetail(ctx, req.Name)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(n)

}

func TestNodesServiceImpl_UpdateTaints(t *testing.T) {
	req := nodes.NewUpdateTaintRequest()
	req.Name = "ubuntu-s-2vcpu-4gb-sfo3-02"
	req.Taints = append(req.Taints, &corev1.Taint{
		Key:    "test",
		Value:  "app",
		Effect: corev1.TaintEffectNoSchedule,
		//TimeAdded: &v1.Time{
		//	time.Now(),
		//},
	})

	err := nodesSvc.UpdateTaints(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	n, err := nodesSvc.GetNodeDetail(ctx, req.Name)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(n)

}
