package apps

import (
	_ "github.com/IanZC0der/kubecenter/apps/configmap/api"
	_ "github.com/IanZC0der/kubecenter/apps/configmap/impl"
	_ "github.com/IanZC0der/kubecenter/apps/nodes/api"
	_ "github.com/IanZC0der/kubecenter/apps/nodes/impl"
	_ "github.com/IanZC0der/kubecenter/apps/pods/api"
	_ "github.com/IanZC0der/kubecenter/apps/pods/impl"
	_ "github.com/IanZC0der/kubecenter/apps/secret/api"
	_ "github.com/IanZC0der/kubecenter/apps/secret/impl"
)
