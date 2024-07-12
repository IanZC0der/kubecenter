package configmap

type ListItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ConfigMap struct {
	Name      string      `json:"name"`
	Namespace string      `json:"namespace"`
	Labels    []*ListItem `json:"labels"`
	Data      []*ListItem `json:"data"`
}

type ConfigMapResponse struct {
	Name      string      `json:"name"`
	Namespace string      `json:"namespace"`
	Labels    []*ListItem `json:"labels"`
	Data      []*ListItem `json:"data"`
	DataNum   int         `json:"dataNum"`
	Age       int64       `json:"age"`
}

func NewConfigMap() *ConfigMap {
	return &ConfigMap{
		Labels: make([]*ListItem, 0),
		Data:   make([]*ListItem, 0),
	}
}
