package pods

type Pods struct {
	Items map[string][]string `json:"items"`
}

type Namespace struct {
	Name              string `json:"name"`
	CreationTimestamp int64  `json:"creationTimestamp"`
	Status            string `json:"status"`
}

type NamespaceList struct {
	Items []*Namespace `json:"items"`
}

func NewPodsList() *Pods {
	return &Pods{
		Items: make(map[string][]string),
	}
}

func NewNamespaceList() *NamespaceList {
	return &NamespaceList{
		Items: make([]*Namespace, 0),
	}
}
