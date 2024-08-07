package apps

import (
	_ "github.com/IanZC0der/kubecenter/apps/configmap/api"
	_ "github.com/IanZC0der/kubecenter/apps/configmap/impl"
	_ "github.com/IanZC0der/kubecenter/apps/k8sservice/api"
	_ "github.com/IanZC0der/kubecenter/apps/k8sservice/impl"
	_ "github.com/IanZC0der/kubecenter/apps/metrics"
	_ "github.com/IanZC0der/kubecenter/apps/metrics/api"
	_ "github.com/IanZC0der/kubecenter/apps/metrics/impl"
	_ "github.com/IanZC0der/kubecenter/apps/nodes/api"
	_ "github.com/IanZC0der/kubecenter/apps/nodes/impl"
	_ "github.com/IanZC0der/kubecenter/apps/persistentvolume/api"
	_ "github.com/IanZC0der/kubecenter/apps/persistentvolume/impl"
	_ "github.com/IanZC0der/kubecenter/apps/pods/api"
	_ "github.com/IanZC0der/kubecenter/apps/pods/impl"
	_ "github.com/IanZC0der/kubecenter/apps/rbac/api"
	_ "github.com/IanZC0der/kubecenter/apps/rbac/impl"
	_ "github.com/IanZC0der/kubecenter/apps/secret/api"
	_ "github.com/IanZC0der/kubecenter/apps/secret/impl"
	_ "github.com/IanZC0der/kubecenter/apps/workload/api"
	_ "github.com/IanZC0der/kubecenter/apps/workload/impl"
)
