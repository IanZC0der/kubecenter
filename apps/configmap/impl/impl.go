package impl

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/configmap"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func init() {
	ioc.DefaultControllerContainer().Register(&ConfigmapServiceImpl{})
}

var _ configmap.Service = &ConfigmapServiceImpl{}

type ConfigmapServiceImpl struct{}

func (s *ConfigmapServiceImpl) Init() error {
	return nil
}

func (s *ConfigmapServiceImpl) Name() string {
	return configmap.AppName
}

func (s *ConfigmapServiceImpl) GetConfigMapList(ctx context.Context, namespace string, keyword string) ([]*configmap.ConfigMapResponse, error) {
	l, err := global.KubeConfigSet.CoreV1().ConfigMaps(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	configs := make([]*configmap.ConfigMapResponse, 0)

	for _, item := range l.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		configs = append(configs, configmap.K8SConfigmapToConfigmap(&item))
	}
	return configs, nil
}
func (s *ConfigmapServiceImpl) GetConfigMapDetail(ctx context.Context, namespace string, name string) (*configmap.ConfigMapResponse, error) {
	item, err := global.KubeConfigSet.CoreV1().ConfigMaps(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return configmap.K8SConfigmapToConfigmap(item), nil
}
