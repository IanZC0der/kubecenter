package pods

import "encoding/json"

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

func (n *Namespace) String() string {
	jsonNs, _ := json.Marshal(n)
	return string(jsonNs)

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

type ListItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (lI *ListItem) String() string {
	jsonlI, _ := json.Marshal(lI)
	return string(jsonlI)

}

type Base struct {
	Name      string      `json:"name"`
	Labels    []*ListItem `json:"labels"`
	Namespace string      `json:"namespace"`
	//always || never || on failure
	RestartPolicy string `json:"restartPolicy"`
}

func NewBase() *Base {
	return &Base{
		Labels: make([]*ListItem, 0),
	}
}

type Volume struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type DnsConfig struct {
	Nameservers []string `json:"nameservers"`
}

func NewDnsConfig() *DnsConfig {
	return &DnsConfig{
		Nameservers: make([]string, 0),
	}
}

type NetWorking struct {
	HostNetwork bool        `json:"hostNetwork"`
	HostName    string      `json:"hostName"`
	DnsPolicy   string      `json:"dnsPolicy"`
	DnsConfig   *DnsConfig  `json:"dnsConfig"`
	HostAliases []*ListItem `json:"hostAliases"`
}

func NewNetWorking() *NetWorking {
	return &NetWorking{
		HostAliases: make([]*ListItem, 0),
		DnsConfig:   NewDnsConfig(),
	}
}

type Resources struct {
	Enable     bool  `json:"enable"`
	MemRequest int32 `json:"memRequest"`
	MemLimit   int32 `json:"memLimit"`
	CpuRequest int32 `json:"cpuRequest"`
	CpuLimit   int32 `json:"cpuLimit"`
}

func NewResources() *Resources {
	return &Resources{}
}

type VolumeMount struct {
	MountName string `json:"mountName"`
	MountPath string `json:"mountPath"`
	ReadOnly  bool   `json:"readOnly"`
}

type ProbeHttpGet struct {
	// http/https
	Scheme string `json:"scheme"`
	//internal request: ""
	Host        string      `json:"host"`
	Path        string      `json:"path"`
	Port        int32       `json:"port"`
	HttpHeaders []*ListItem `json:"httpHeaders"`
}

func NewProbeHttpGet() *ProbeHttpGet {
	return &ProbeHttpGet{
		HttpHeaders: make([]*ListItem, 0),
	}
}

type ProbeCommand struct {
	Command []string `json:"command"`
}

func NewProbeCommand() *ProbeCommand {
	return &ProbeCommand{
		Command: make([]string, 0),
	}
}

type ProbeTcpSocket struct {
	Host string `json:"host"`
	// probe port
	Port int32 `json:"port"`
}

type ProbeTime struct {
	//probing starts after the delay
	InitialDelaySeconds int32 `json:"initialDelaySeconds"`
	//interval
	PeriodSeconds int32 `json:"periodSeconds"`
	//probe failed after the timeout
	TimeoutSeconds int32 `json:"timeoutSeconds"`
	//success after passing the threshold
	SuccessThreshold int32 `json:"successThreshold"`
	//failed after pass the threshold
	FailureThreshold int32 `json:"failureThreshold"`
}

type ContainerProbe struct {
	// whether the probed is enabled
	Enable bool `json:"enable"`
	//probe type
	Type      string          `json:"type"`
	HttpGet   *ProbeHttpGet   `json:"httpGet"`
	Exec      *ProbeCommand   `json:"exec"`
	TcpSocket *ProbeTcpSocket `json:"tcpSocket"`
	*ProbeTime
}

func NewContainerProbe() *ContainerProbe {
	return &ContainerProbe{
		HttpGet:   NewProbeHttpGet(),
		Exec:      NewProbeCommand(),
		TcpSocket: &ProbeTcpSocket{},
		ProbeTime: &ProbeTime{},
	}
}

type ContainerPort struct {
	Name          string `json:"name"`
	ContainerPort int32  `json:"containerPort"`
	HostPort      int32  `json:"hostPort"`
}

type Container struct {
	Name            string           `json:"name"`
	Image           string           `json:"image"`
	ImagePullPolicy string           `json:"imagePullPolicy"`
	Tty             bool             `json:"tty"`
	Ports           []*ContainerPort `json:"ports"`
	WorkingDir      string           `json:"workingDir"`
	Command         []string         `json:"command"`
	Args            []string         `json:"args"`
	Envs            []*ListItem      `json:"envs"`
	Privileged      bool             `json:"privileged"`
	Resources       *Resources       `json:"resources"`
	VolumeMounts    []*VolumeMount   `json:"volumeMounts"`
	StartupProbe    *ContainerProbe  `json:"startupProbe"`
	LivenessProbe   *ContainerProbe  `json:"livenessProbe"`
	ReadinessProbe  *ContainerProbe  `json:"readinessProbe"`
}

func NewContainer() *Container {
	return &Container{
		Ports:          make([]*ContainerPort, 0),
		Command:        make([]string, 0),
		Args:           make([]string, 0),
		Envs:           make([]*ListItem, 0),
		Privileged:     false,
		Resources:      NewResources(),
		VolumeMounts:   make([]*VolumeMount, 0),
		StartupProbe:   nil,
		LivenessProbe:  nil,
		ReadinessProbe: nil,
	}
}

type Pod struct {
	Base    *Base     `json:"base"`
	Volumes []*Volume `json:"volumes"`
	//网络相关
	NetWorking *NetWorking `json:"netWorking"`
	///init containers
	InitContainers []*Container `json:"initContainers"`
	//containers
	Containers []*Container `json:"containers"`
}

func NewPod() *Pod {
	return &Pod{
		Base:           NewBase(),
		Volumes:        make([]*Volume, 0),
		NetWorking:     NewNetWorking(),
		InitContainers: make([]*Container, 0),
		Containers:     make([]*Container, 0),
	}
}

type PodListItem struct {
	Name     string `json:"name"`
	Ready    string `json:"ready"`
	Status   string `json:"status"`
	Restarts int32  `json:"restarts"`
	Age      int64  `json:"age"`
	IP       string `json:"IP"`
	Node     string `json:"node"`
}

func (pI *PodListItem) String() string {
	jsonpI, _ := json.Marshal(pI)
	return string(jsonpI)

}

type PodsList struct {
	Items []*PodListItem `json:"items"`
}

func NewPodsItemsList() *PodsList {
	return &PodsList{
		Items: make([]*PodListItem, 0),
	}
}
