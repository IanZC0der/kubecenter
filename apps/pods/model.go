package pods

type Pods struct {
	Items map[string][]string `json:"items"`
}

func NewPodsList() *Pods {
	return &Pods{
		Items: make(map[string][]string),
	}
}
