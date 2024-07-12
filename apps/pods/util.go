package pods

import (
	"errors"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
	"strings"
)

// type podsUtil struct {
// }

const (
	IMAGE_PULL_POLICY_IFNOTPRESENT = "IfNotPresent"
	RESTART_POLICY_ALWAYS          = "Always"
	VOLUME_TYPE_EMPTYDIR           = "emptyDir"
	PROBE_HTTP                     = "http"
	PROBE_TCP                      = "tcp"
	PROBE_EXEC                     = "exec"
)

var volumeMap = make(map[string]string)

func GetPodListItemFromPod(pod *corev1.Pod) *PodListItem {
	var total, ready, restart int32

	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.Ready {
			ready++
		}
		restart += containerStatus.RestartCount
		total++
	}
	var podStatus string

	if pod.Status.Phase != "Running" {
		podStatus = "Error"
	} else {
		podStatus = "Running"
	}

	return &PodListItem{
		Name:     pod.Name,
		Ready:    fmt.Sprintf("%d/%d", ready, total),
		Status:   podStatus,
		Restarts: restart,
		Age:      pod.CreationTimestamp.Unix(),
		IP:       pod.Status.PodIP,
		Node:     pod.Spec.NodeName,
	}

}

func GetBaseFromPod(pod *corev1.Pod) *Base {
	newBase := NewBase()
	newBase.Name = pod.Name
	newBase.Namespace = pod.Namespace

	for k, v := range pod.Labels {
		newBase.Labels = append(newBase.Labels, &ListItem{
			Key:   k,
			Value: v,
		})
	}
	newBase.RestartPolicy = string(pod.Spec.RestartPolicy)
	return newBase
}

func GetTolerationsFromPod(pod *corev1.Pod) []*corev1.Toleration {
	tls := make([]*corev1.Toleration, 0)
	for _, tl := range pod.Spec.Tolerations {
		tls = append(tls, &tl)
	}
	return tls
}

func GetPodInfoFromPod(pod *corev1.Pod) *Pod {
	newPod := NewPod()
	newPod.Base = GetBaseFromPod(pod)
	newPod.Tolerations = GetTolerationsFromPod(pod)
	newPod.NetWorking = GetNetworkingFromPod(pod)
	for _, volume := range pod.Spec.Volumes {
		if volume.EmptyDir == nil {
			continue
		}
		if volumeMap == nil {
			volumeMap = make(map[string]string)
		}
		volumeMap[volume.Name] = ""
		newPod.Volumes = append(newPod.Volumes, &Volume{
			Type: VOLUME_TYPE_EMPTYDIR,
			Name: volume.Name,
		})
	}
	//initcontainers
	for _, ctner := range pod.Spec.Containers {
		newPod.Containers = append(newPod.Containers, GetContainerFromPod(&ctner))
	}

	for _, ctner := range pod.Spec.InitContainers {
		newPod.InitContainers = append(newPod.InitContainers, GetContainerFromPod(&ctner))
	}
	return newPod
}

func GetNetworkingFromPod(pod *corev1.Pod) *NetWorking {
	newNetWorking := NewNetWorking()
	newNetWorking.HostName = pod.Spec.Hostname
	newNetWorking.HostNetwork = pod.Spec.HostNetwork
	newNetWorking.DnsPolicy = string(pod.Spec.DNSPolicy)
	if pod.Spec.DNSConfig != nil {
		if len(pod.Spec.DNSConfig.Nameservers) > 0 {
			for _, ns := range pod.Spec.DNSConfig.Nameservers {
				newNetWorking.DnsConfig.Nameservers = append(newNetWorking.DnsConfig.Nameservers, ns)
			}
		}
	}

	for _, a := range pod.Spec.HostAliases {
		newNetWorking.HostAliases = append(newNetWorking.HostAliases, &ListItem{
			Key:   a.IP,
			Value: strings.Join(a.Hostnames, "."),
		})
	}
	return newNetWorking

}

func GetContainerFromPod(ctner *corev1.Container) *Container {
	newContainer := NewContainer()
	newContainer.Name = ctner.Name
	newContainer.Image = ctner.Image
	newContainer.ImagePullPolicy = string(ctner.ImagePullPolicy)
	newContainer.Tty = ctner.TTY
	newContainer.WorkingDir = ctner.WorkingDir
	newContainer.Command = ctner.Command
	newContainer.Args = ctner.Args

	for _, env := range ctner.Env {
		newContainer.Envs = append(newContainer.Envs, &ListItem{
			Key:   env.Name,
			Value: env.Value,
		})
	}

	for _, port := range ctner.Ports {
		newContainer.Ports = append(newContainer.Ports, &ContainerPort{
			Name:          port.Name,
			ContainerPort: port.ContainerPort,
			HostPort:      port.HostPort,
		})
	}

	if ctner.SecurityContext != nil {
		newContainer.Privileged = *ctner.SecurityContext.Privileged
	}

	newContainer.Resources = GetResourcesFromContainer(ctner)

	for _, vm := range ctner.VolumeMounts {
		if _, ok := volumeMap[vm.Name]; ok {
			newContainer.VolumeMounts = append(newContainer.VolumeMounts, &VolumeMount{
				MountPath: vm.MountPath,
				MountName: vm.Name,
				ReadOnly:  vm.ReadOnly,
			})
		}
	}
	//probes
	newContainer.StartupProbe = GetProbeFromContainerProbe(ctner.ReadinessProbe)
	newContainer.LivenessProbe = GetProbeFromContainerProbe(ctner.LivenessProbe)
	newContainer.ReadinessProbe = GetProbeFromContainerProbe(ctner.ReadinessProbe)
	return newContainer

}

func GetResourcesFromContainer(ctner *corev1.Container) *Resources {
	newResources := NewResources()
	newResources.Enable = false
	requests := ctner.Resources.Requests
	limits := ctner.Resources.Limits
	if requests != nil {
		newResources.Enable = true
		newResources.CpuRequest = int32(requests.Cpu().MilliValue())
		//MiB
		newResources.MemRequest = int32(requests.Memory().Value() / (1024 * 1024))
	}
	if limits != nil {
		newResources.Enable = true
		newResources.CpuLimit = int32(limits.Cpu().MilliValue())
		newResources.MemLimit = int32(limits.Memory().Value() / (1024 * 1024))
	}
	return newResources
}

func GetProbeFromContainerProbe(prb *corev1.Probe) *ContainerProbe {
	newContainerProb := NewContainerProbe()
	newContainerProb.Enable = false
	if prb != nil {
		newContainerProb.Enable = true
		if prb.Exec != nil {
			newContainerProb.Type = PROBE_EXEC
			newContainerProb.Exec.Command = prb.Exec.Command
		} else if prb.HTTPGet != nil {
			newContainerProb.Type = PROBE_HTTP
			httpGet := prb.HTTPGet
			for _, headerK8s := range httpGet.HTTPHeaders {
				newContainerProb.HttpGet.HttpHeaders = append(newContainerProb.HttpGet.HttpHeaders, &ListItem{
					Key:   headerK8s.Name,
					Value: headerK8s.Value,
				})
			}
			newContainerProb.HttpGet.Host = httpGet.Host
			newContainerProb.HttpGet.Port = httpGet.Port.IntVal
			newContainerProb.HttpGet.Scheme = string(httpGet.Scheme)
			newContainerProb.HttpGet.Path = httpGet.Path
		} else if prb.TCPSocket != nil {
			newContainerProb.Type = PROBE_TCP
			newContainerProb.TcpSocket.Port = prb.TCPSocket.Port.IntVal
			newContainerProb.TcpSocket.Host = prb.TCPSocket.Host
		} else {
			newContainerProb.Type = PROBE_HTTP
			return newContainerProb
		}
		newContainerProb.InitialDelaySeconds = prb.InitialDelaySeconds
		newContainerProb.PeriodSeconds = prb.PeriodSeconds
		newContainerProb.TimeoutSeconds = prb.TimeoutSeconds
		newContainerProb.SuccessThreshold = prb.SuccessThreshold
		newContainerProb.FailureThreshold = prb.FailureThreshold
	}
	return newContainerProb
}

// required fields:
// pod.Base.Name
// pod.Base.Containers should not be nil
//if base restartlicy is nil, use default restart policy

// image name/image in the containers should be well-defined
// if pull policy is nil, use default
func PodCreateValidate(pod *Pod) error {
	if pod.Base.Name == "" {
		return errors.New("pod name is empty")
	}

	if pod.Base.RestartPolicy == "" {
		pod.Base.RestartPolicy = RESTART_POLICY_ALWAYS
	}

	if len(pod.Containers) == 0 {
		return errors.New("Pod containers are not defined")
	}

	if len(pod.InitContainers) > 0 {
		for i, container := range pod.InitContainers {
			if container.Name == "" {
				return errors.New("pod init container name is empty")
			}

			if container.Image == "" {
				return errors.New("pod init container image is not defined")
			}

			if container.ImagePullPolicy == "" {
				pod.InitContainers[i].ImagePullPolicy = IMAGE_PULL_POLICY_IFNOTPRESENT
			}
		}
	}

	if len(pod.Containers) > 0 {
		for i, container := range pod.Containers {
			if container.Name == "" {
				return errors.New("pod container name is empty")
			}

			if container.Image == "" {
				return errors.New("pod container image is not defined")
			}

			if container.ImagePullPolicy == "" {
				pod.Containers[i].ImagePullPolicy = IMAGE_PULL_POLICY_IFNOTPRESENT
			}
		}
	}
	return nil
}

func CreatePodFromPodRequest(pod *Pod) *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pod.Base.Name,
			Namespace: pod.Base.Namespace,
			Labels:    GetK8SPodBaseLabels(pod),
		},

		Spec: corev1.PodSpec{
			Tolerations:    GetK8sTolerations(pod.Tolerations),
			InitContainers: GetK8SContainers(pod.InitContainers),
			Containers:     GetK8SContainers(pod.Containers),
			Volumes:        GetK8SPodVolumes(pod.Volumes),
			DNSConfig: &corev1.PodDNSConfig{
				Nameservers: pod.NetWorking.DnsConfig.Nameservers,
			},
			DNSPolicy:     corev1.DNSPolicy(pod.NetWorking.DnsPolicy),
			Hostname:      pod.NetWorking.HostName,
			HostAliases:   GetK8SHostAliases(pod.NetWorking.HostAliases),
			RestartPolicy: corev1.RestartPolicy(pod.Base.RestartPolicy),
		},
	}
}

func GetK8sTolerations(tolerations []*corev1.Toleration) []corev1.Toleration {
	k8sTolerations := make([]corev1.Toleration, 0)
	for _, toleration := range tolerations {
		k8sTolerations = append(k8sTolerations, *toleration)
	}
	return k8sTolerations
}

func GetK8SPodBaseLabels(pod *Pod) map[string]string {
	k8sLabels := make(map[string]string)

	for _, label := range pod.Base.Labels {
		k8sLabels[label.Key] = label.Value
	}
	return k8sLabels
}

func GetK8SContainers(ctners []*Container) []corev1.Container {
	k8sContainers := make([]corev1.Container, 0)

	for _, ctner := range ctners {
		k8sContainers = append(k8sContainers, GetK8SConinter(ctner))
	}
	return k8sContainers
}

func GetK8SConinter(ctner *Container) corev1.Container {
	return corev1.Container{
		Name:            ctner.Name,
		Image:           ctner.Image,
		ImagePullPolicy: corev1.PullPolicy(ctner.ImagePullPolicy),
		TTY:             ctner.Tty,
		Command:         ctner.Command,
		Args:            ctner.Args,
		WorkingDir:      ctner.WorkingDir,
		SecurityContext: &corev1.SecurityContext{
			Privileged: &ctner.Privileged,
		},
		Ports:          GetK8SContainerPorts(ctner.Ports),
		Env:            GetK8SEnv(ctner.Envs),
		VolumeMounts:   GetK8SVolumeMounts(ctner.VolumeMounts),
		StartupProbe:   GetK8SContainerProbe(ctner.StartupProbe),
		ReadinessProbe: GetK8SContainerProbe(ctner.ReadinessProbe),
		LivenessProbe:  GetK8SContainerProbe(ctner.LivenessProbe),
		Resources:      GetK8SResources(ctner.Resources),
	}
}

func GetK8SContainerPorts(ports []*ContainerPort) []corev1.ContainerPort {
	k8sContainterPorts := make([]corev1.ContainerPort, 0)

	for _, item := range ports {
		k8sContainterPorts = append(k8sContainterPorts, corev1.ContainerPort{
			Name:          item.Name,
			HostPort:      item.HostPort,
			ContainerPort: item.ContainerPort,
		})
	}
	return k8sContainterPorts
}

func GetK8SEnv(env []*ListItem) []corev1.EnvVar {
	envs := make([]corev1.EnvVar, 0)
	for _, item := range env {
		envs = append(envs, corev1.EnvVar{
			Name:  item.Key,
			Value: item.Value,
		})
	}
	return envs
}

func GetK8SVolumeMounts(vmts []*VolumeMount) []corev1.VolumeMount {
	k8sVMounts := make([]corev1.VolumeMount, 0)
	for _, item := range vmts {
		k8sVMounts = append(k8sVMounts, corev1.VolumeMount{
			Name:      item.MountName,
			MountPath: item.MountPath,
			ReadOnly:  item.ReadOnly,
		})
	}
	return k8sVMounts
}

func GetK8SContainerProbe(prb *ContainerProbe) *corev1.Probe {
	if prb == nil {
		return nil
	}
	if !prb.Enable {
		return nil
	}
	k8sProbe := &corev1.Probe{
		InitialDelaySeconds: prb.InitialDelaySeconds,
		PeriodSeconds:       prb.PeriodSeconds,
		TimeoutSeconds:      prb.TimeoutSeconds,
		SuccessThreshold:    prb.SuccessThreshold,
		FailureThreshold:    prb.FailureThreshold,
	}

	switch prb.Type {
	case PROBE_HTTP:
		httpGet := prb.HttpGet
		k8sHttpHeaders := make([]corev1.HTTPHeader, 0)
		for _, header := range httpGet.HttpHeaders {
			k8sHttpHeaders = append(k8sHttpHeaders, corev1.HTTPHeader{
				Name:  header.Key,
				Value: header.Value,
			})
		}
		k8sProbe.HTTPGet = &corev1.HTTPGetAction{
			Scheme:      corev1.URIScheme(httpGet.Scheme),
			Host:        httpGet.Host,
			Port:        intstr.FromInt(int(httpGet.Port)),
			Path:        httpGet.Path,
			HTTPHeaders: k8sHttpHeaders,
		}
	case PROBE_TCP:
		tcpskt := prb.TcpSocket
		k8sProbe.TCPSocket = &corev1.TCPSocketAction{
			Host: tcpskt.Host,
			Port: intstr.FromInt(int(tcpskt.Port)),
		}
	case PROBE_EXEC:
		exec := prb.Exec
		k8sProbe.Exec = &corev1.ExecAction{
			Command: exec.Command,
		}
	}
	return k8sProbe
}

func GetK8SResources(rscs *Resources) corev1.ResourceRequirements {
	var k8sPodRscs corev1.ResourceRequirements
	if rscs == nil || !rscs.Enable {
		return k8sPodRscs
	}
	k8sPodRscs.Requests = corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse(strconv.Itoa(int(rscs.CpuRequest)) + "m"),
		corev1.ResourceMemory: resource.MustParse(strconv.Itoa(int(rscs.MemRequest)) + "Mi"),
	}

	k8sPodRscs.Limits = corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse(strconv.Itoa(int(rscs.CpuLimit)) + "m"),
		corev1.ResourceMemory: resource.MustParse(strconv.Itoa(int(rscs.MemLimit)) + "Mi"),
	}
	return k8sPodRscs

}

func GetK8SPodVolumes(vlms []*Volume) []corev1.Volume {
	podk8svlms := make([]corev1.Volume, 0)
	for _, item := range vlms {
		if item.Type != VOLUME_TYPE_EMPTYDIR {
			continue
		}
		source := corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{},
		}
		podk8svlms = append(podk8svlms, corev1.Volume{
			VolumeSource: source,
			Name:         item.Name,
		})
	}
	return podk8svlms
}

func GetK8SHostAliases(podAls []*ListItem) []corev1.HostAlias {
	podk8sAliases := make([]corev1.HostAlias, 0)
	for _, item := range podAls {
		podk8sAliases = append(podk8sAliases, corev1.HostAlias{
			IP:        item.Key,
			Hostnames: strings.Split(item.Value, ","),
		})
	}
	return podk8sAliases
}
