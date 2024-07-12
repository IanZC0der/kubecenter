package impl

import (
	"context"
	"fmt"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	//"k8s.io/kube-openapi/pkg/util"
	k8serror "k8s.io/apimachinery/pkg/api/errors"
	"strings"

	"github.com/IanZC0der/kubecenter/apps/pods"
)

func init() {
	ioc.DefaultControllerContainer().Register(&PodsServerImpl{})
}

var _ pods.Service = &PodsServerImpl{}

type PodsServerImpl struct {
}

func (s *PodsServerImpl) GetPods(ctx context.Context) (*pods.Pods, error) {
	podsList := pods.NewPodsList()
	c := context.TODO()
	pods, err := global.KubeConfigSet.CoreV1().Pods("").List(c, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, pod := range pods.Items {
		podsList.Items[pod.Namespace] = append(podsList.Items[pod.Namespace], pod.Name)
	}

	return podsList, nil

}

func (s *PodsServerImpl) GetNamespaceList(ctx context.Context) (*pods.NamespaceList, error) {
	namespaceList := pods.NewNamespaceList()
	c := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Namespaces().List(c, metav1.ListOptions{})

	if err != nil {
		return nil, err
	}
	for _, namespace := range list.Items {
		namespaceList.Items = append(namespaceList.Items, &pods.Namespace{
			Name:              namespace.Name,
			CreationTimestamp: namespace.CreationTimestamp.Unix(),
			Status:            string(namespace.Status.Phase),
		})
	}
	return namespaceList, nil
}

func (s *PodsServerImpl) GetPodsListWithinNode(ctx context.Context, keyword string, nodeName string) (*pods.PodsList, error) {
	podsList := pods.NewPodsItemsList()

	c := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Pods("").List(c, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range list.Items {
		if nodeName != "" && item.Spec.NodeName != nodeName {
			continue
		}
		if strings.Contains(item.Name, keyword) {
			podsList.Items = append(podsList.Items, pods.GetPodListItemFromPod(&item))
		}
	}
	return podsList, nil
}

func (s *PodsServerImpl) GetPodsListUnderNamespaceWithKeyword(ctx context.Context, namespace string, keyword string) (*pods.PodsList, error) {
	podsList := pods.NewPodsItemsList()

	c := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Pods(namespace).List(c, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range list.Items {
		if strings.Contains(item.Name, keyword) {
			podsList.Items = append(podsList.Items, pods.GetPodListItemFromPod(&item))
		}
	}
	return podsList, nil
}

func (s *PodsServerImpl) GetPodDetail(ctx context.Context, namespace string, name string) (*pods.Pod, error) {
	c := context.TODO()
	k8sPod, err := global.KubeConfigSet.CoreV1().Pods(namespace).Get(c, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return pods.GetPodInfoFromPod(k8sPod), err
}

func (s *PodsServerImpl) CreatePod(ctx context.Context, pod *pods.Pod) (*corev1.Pod, error) {
	c := context.TODO()
	k8sPod := pods.CreatePodFromPodRequest(pod)
	createdPod, err := global.KubeConfigSet.CoreV1().Pods(k8sPod.Namespace).Create(c, k8sPod, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return createdPod, nil
}

// update logic: delete + create
// 1. query the pod, if exists, initiate a dry run of creating the pod to validate the params of the request
// 2. select the pod and delete it
// 3. create a watched to listen for the event deleted, creating starts only after the pod is deleted
func (s *PodsServerImpl) UpdatePod(ctx context.Context, pod *pods.Pod) (*corev1.Pod, string, error) {
	c := context.TODO()
	k8sPod := pods.CreatePodFromPodRequest(pod)
	//look up the pod first, return if it doesn't exist
	k8sHandler := global.KubeConfigSet.CoreV1().Pods(k8sPod.Namespace)
	k8sPodQueried, err := k8sHandler.Get(c, k8sPod.Name, metav1.GetOptions{})
	if err == nil {
		//validate the params of k8s pod by initiating a dry run
		podCopy := *k8sPod
		podCopy.Name = k8sPod.Name + "-validate"
		_, err := k8sHandler.Create(c, &podCopy, metav1.CreateOptions{
			DryRun: []string{metav1.DryRunAll},
		})
		if err != nil {
			msg := fmt.Sprintf("data validation error, failed to update pod [%s] under namespace [%v], detail: %s", k8sPod.Name, k8sPod.Namespace, err)
			return nil, msg, err
		}

		//select the pod to be updated based on the lable, example: app=test
		var labelSelector []string
		for k, v := range k8sPodQueried.Labels {
			labelSelector = append(labelSelector, fmt.Sprintf("%s=%s", k, v))
		}
		//create a watcher listening for delete event, initiating the creating only after the pod is deleted
		watcher, err := k8sHandler.Watch(c, metav1.ListOptions{
			LabelSelector: strings.Join(labelSelector, ","),
		})
		if err != nil {
			msg := fmt.Sprintf("Failed to update pod [%s] under namespace [%v], detail: %s", k8sPod.Name, k8sPod.Namespace, err)
			return nil, msg, err
		}

		background := metav1.DeletePropagationBackground
		var gracePeriodSeconds int64 = 0
		err = k8sHandler.Delete(c, k8sPod.Name, metav1.DeleteOptions{
			GracePeriodSeconds: &gracePeriodSeconds,
			PropagationPolicy:  &background,
		})
		if err != nil {
			msg := fmt.Sprintf("deletion error, failed to update pod [%s] under namespace [%v], detail: %s", k8sPod.Name, k8sPod.Namespace, err)
			return nil, msg, err
		}
		for event := range watcher.ResultChan() {

			//
			if _, err := k8sHandler.Get(c, k8sPod.Name, metav1.GetOptions{}); k8serror.IsNotFound(err) {
				if createdPod, err := s.CreatePod(ctx, pod); err != nil {
					msg := fmt.Sprintf("error when creating the pod, failed to update pod [%s] under namespace [%v], detail: %s", k8sPod.Name, k8sPod.Namespace, err)
					return nil, msg, err
				} else {
					msg := fmt.Sprintf("successfully update pod [%s] under namespace [%v]", k8sPod.Name, k8sPod.Namespace)
					return createdPod, msg, nil
				}
			}
			k8sPodFromWatcher := event.Object.(*corev1.Pod)

			switch event.Type {
			case watch.Deleted:
				if k8sPodFromWatcher.Name != k8sPod.Name {
					continue
				}
				if createdPod, err := s.CreatePod(ctx, pod); err != nil {
					msg := fmt.Sprintf("Failed to update pod [%s] under namespace [%v], detail: %s", k8sPod.Name, k8sPod.Namespace, err)
					return nil, msg, err
				} else {
					msg := fmt.Sprintf("successfully update pod [%s] under namespace [%v]", k8sPod.Name, k8sPod.Namespace)
					return createdPod, msg, nil
				}

			}
		}
		return k8sPod, "", nil
	}
	msg := fmt.Sprintf("pod doesn't exist, please create the pod first, failed to update pod [%s] under namespace [%v], detail: %s", k8sPod.Name, k8sPod.Namespace, err)
	return k8sPod, msg, err
}

func (s *PodsServerImpl) DeletePod(ctx context.Context, namespace string, name string) (*corev1.Pod, string, error) {
	c := context.TODO()
	k8sPod, err := global.KubeConfigSet.CoreV1().Pods(namespace).Get(c, name, metav1.GetOptions{})
	if err != nil {
		msg := fmt.Sprintf("pod doesn't exist, failed to delete pod [%s] under namespace [%v], detail: %s", name, namespace, err)
		return nil, msg, err
	}
	background := metav1.DeletePropagationBackground
	var gracePeriodSeconds int64 = 0
	err = global.KubeConfigSet.CoreV1().Pods(namespace).Delete(c, name, metav1.DeleteOptions{
		GracePeriodSeconds: &gracePeriodSeconds,
		PropagationPolicy:  &background,
	})

	if err != nil {
		msg := fmt.Sprintf("failed to delete pod [%s] under namespace [%v]", name, namespace)
		return k8sPod, msg, err
	}
	msg := fmt.Sprintf("successfully deleted pod [%s] under namespace [%v]", name, namespace)

	return k8sPod, msg, nil
}
func (s *PodsServerImpl) Init() error {
	return nil
}

func (s *PodsServerImpl) Name() string {
	return pods.AppName
}
