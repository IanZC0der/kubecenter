package global

import (
	"github.com/IanZC0der/kubecenter/conf"
	"k8s.io/client-go/kubernetes"
)

var (
	CONF          = Init()
	KubeConfigSet *kubernetes.Clientset
)

func C() *conf.Config {

	return CONF
}

func Init() *conf.Config {
	c := &conf.Config{
		App: &conf.App{},
		System: &conf.System{
			Prometheus: &conf.Prometheus{},
		},
	}
	return c
}

func APP() *conf.App {
	return CONF.App
}
