package impl

import (
	"context"
	"encoding/json"
	"github.com/IanZC0der/kubecenter/apps/nodes"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"strings"
)

func init() {
	ioc.DefaultControllerContainer().Register(&NodesServiceImpl{})
}

var _ nodes.Service = &NodesServiceImpl{}

type NodesServiceImpl struct{}

func (s *NodesServiceImpl) Init() error {
	return nil
}

func (s *NodesServiceImpl) Name() string {
	return nodes.AppName
}

func (s *NodesServiceImpl) GetNodeList(ctx context.Context, keyword string) ([]*nodes.Node, error) {
	c := context.TODO()

	list, err := global.KubeConfigSet.CoreV1().Nodes().List(c, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	nodesList := make([]*nodes.Node, 0)

	for _, node := range list.Items {
		if strings.Contains(node.Name, keyword) {
			nodesList = append(nodesList, nodes.GetNodeInfoFromK8SNode(&node))
		}
	}
	return nodesList, nil
}

func (s *NodesServiceImpl) GetNodeDetail(ctx context.Context, name string) (*nodes.Node, error) {
	c := context.TODO()
	k8sNode, err := global.KubeConfigSet.CoreV1().Nodes().Get(c, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	res := nodes.GetNodeDetailFromK8SNode(k8sNode)
	return res, nil
}

func (s *NodesServiceImpl) UpdateLabel(ctx context.Context, req *nodes.UpdateLabelRequest) error {
	labelsMap := make(map[string]string, 0)
	for _, label := range req.Labels {
		labelsMap[label.Key] = label.Value
	}
	labelsMap["$patch"] = "replace"
	patchData := map[string]any{
		"metadata": map[string]any{
			"labels": labelsMap,
		},
	}
	patchDataBytes, _ := json.Marshal(&patchData)
	_, err := global.KubeConfigSet.CoreV1().Nodes().Patch(
		ctx,
		req.Name,
		types.StrategicMergePatchType,
		patchDataBytes,
		metav1.PatchOptions{},
	)
	return err
}

func (s *NodesServiceImpl) UpdateTaints(ctx context.Context, req *nodes.UpdateTaintRequest) error {
	patchData := map[string]any{
		"spec": map[string]any{
			"taints": req.Taints,
		},
	}
	patchDataBytes, _ := json.Marshal(&patchData)
	_, err := global.KubeConfigSet.CoreV1().Nodes().Patch(
		context.TODO(),
		req.Name,
		types.StrategicMergePatchType,
		patchDataBytes,
		metav1.PatchOptions{},
	)
	return err
}
