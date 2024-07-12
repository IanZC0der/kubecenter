package nodes

import corev1 "k8s.io/api/core/v1"

func GetNodeInfoFromK8SNode(k8sNode *corev1.Node) *Node {
	newNode := NewNode()
	newNode.Name = k8sNode.Name

	newNode.OsImage = k8sNode.Status.NodeInfo.OSImage
	newNode.Version = k8sNode.Status.NodeInfo.KubeletVersion
	newNode.KernelVersion = k8sNode.Status.NodeInfo.KernelVersion
	newNode.ContainerRuntime = k8sNode.Status.NodeInfo.ContainerRuntimeVersion
	newNode.Status = getNodeStatusFromK8SNode(k8sNode)
	newNode.Age = k8sNode.CreationTimestamp.Unix()
	newNode.InternalIp = getNodeIPFromK8SNode(k8sNode, corev1.NodeInternalIP)
	newNode.ExternalIp = getNodeIPFromK8SNode(k8sNode, corev1.NodeExternalIP)

	return newNode
}

func GetNodeDetailFromK8SNode(k8sNode *corev1.Node) *Node {
	newNode := GetNodeInfoFromK8SNode(k8sNode)
	for _, item := range k8sNode.Spec.Taints {
		newNode.Taints = append(newNode.Taints, &item)
	}

	for k, v := range k8sNode.Labels {
		newNode.Labels = append(newNode.Labels, &ListItem{
			Key:   k,
			Value: v,
		})
	}
	return newNode
}
func getNodeStatusFromK8SNode(k8sNode *corev1.Node) string {
	s := "NotReady"
	for _, condition := range k8sNode.Status.Conditions {
		if condition.Type == corev1.NodeReady && condition.Status == corev1.ConditionTrue {
			s = "Ready"
			break
		}
	}
	return s
}

func getNodeIPFromK8SNode(k8sNode *corev1.Node, addressType corev1.NodeAddressType) string {
	for _, item := range k8sNode.Status.Addresses {
		if item.Type == addressType {
			return item.Address
		}
	}
	return "<none>"
}
