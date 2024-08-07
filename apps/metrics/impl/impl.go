package impl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IanZC0der/kubecenter/apps/metrics"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/util"
	promapi "github.com/prometheus/client_golang/api"
	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"time"
)

func init() {
	ioc.DefaultControllerContainer().Register(&MetricsServiceImpl{})
}

func (m *MetricsServiceImpl) Init() error {
	return nil
}

func (m *MetricsServiceImpl) Name() string {
	return metrics.AppName
}

var _ metrics.Service = &MetricsServiceImpl{}

type MetricsServiceImpl struct{}

func (m *MetricsServiceImpl) GetClusterInfo(ctx context.Context) []*metrics.MetricsItem {
	result := make([]*metrics.MetricsItem, 0)

	// get  the cluster creation time by finding the master node, the master created the earliest should be the creation time of the cluster
	list, err := global.KubeConfigSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err == nil {
		var creationTime int64 = 0
		for _, item := range list.Items {
			if _, ok := item.Labels["node-role.kubernetes.io/control-plane"]; ok {
				if creationTime == 0 || (creationTime > 0 && item.CreationTimestamp.Unix() < creationTime) {
					creationTime = item.CreationTimestamp.Unix()
				}
			}
		}
		formarttedTime := util.FormatTime(creationTime)
		result = append(result, &metrics.MetricsItem{
			Name: "Cluster Creation Time",
			//Label: "Creation Time",
			Value: formarttedTime,
		})

	}

	// get number of nodes

	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Name:  "Nodes",
			Value: strconv.Itoa(len(list.Items)),
		})
	}

	// add a color value in each item so  that the frontend can use the value to present the data
	for _, item := range result {
		item.Color = util.GenerateRGB(item.Name)
	}
	return result
}

func (m *MetricsServiceImpl) GetResourceInfo(ctx context.Context) []*metrics.MetricsItem {
	result := make([]*metrics.MetricsItem, 0)

	// get namespace info
	ns, err := global.KubeConfigSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			// values should be the number of namespaces
			Value: strconv.Itoa(len(ns.Items)),
			Name:  "Namespaces",
		})
	}

	// get pods info
	podsList, err := global.KubeConfigSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})

	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(podsList.Items)),
			Name:  "Pods",
		})
	}

	// get configMap info

	cmlist, err := global.KubeConfigSet.CoreV1().ConfigMaps("").List(ctx, metav1.ListOptions{})

	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(cmlist.Items)),
			Name:  "ConfigMaps",
		})
	}

	// get sercret info

	scList, err := global.KubeConfigSet.CoreV1().Secrets("").List(ctx, metav1.ListOptions{})
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(scList.Items)),
			Name:  "Secrets",
		})
	}

	// get persisten volumes info

	pvList, err := global.KubeConfigSet.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(pvList.Items)),
			Name:  "PersistentVolumes",
		})
	}

	pvcList, err := global.KubeConfigSet.CoreV1().PersistentVolumeClaims("").List(ctx, metav1.ListOptions{})
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(pvcList.Items)),
			Name:  "PersistentVolumeClaims",
		})
	}

	// get services info
	svcList, err := global.KubeConfigSet.CoreV1().Services("").List(ctx, metav1.ListOptions{})
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(svcList.Items)),
			Name:  "Services",
		})
	}

	// get ingress

	ingrs, err := global.KubeConfigSet.NetworkingV1().Ingresses("").List(ctx, metav1.ListOptions{})

	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(ingrs.Items)),
			Name:  "Ingresses",
		})
	}

	// deployment info

	dplmt, err := global.KubeConfigSet.AppsV1().Deployments("").List(ctx, metav1.ListOptions{})
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(dplmt.Items)),
			Name:  "Deployments",
		})
	}

	// Daemonsets

	dmsList, err := global.KubeConfigSet.AppsV1().DaemonSets("").List(ctx, metav1.ListOptions{})
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(dmsList.Items)),
			Name:  "DaemonSets",
		})
	}

	// StatefulSets
	stSets, err := global.KubeConfigSet.AppsV1().StatefulSets("").List(ctx, metav1.ListOptions{})
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(stSets.Items)),
			Name:  "StatefulSets",
		})
	}

	// Jobs

	jobsList, err := global.KubeConfigSet.BatchV1().Jobs("").List(ctx, metav1.ListOptions{})

	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(jobsList.Items)),
			Name:  "Jobs",
		})
	}

	//cronJobs

	cronJobsList, err := global.KubeConfigSet.BatchV1().CronJobs("").List(ctx, metav1.ListOptions{})
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(cronJobsList.Items)),
			Name:  "CronJobs",
		})
	}

	//service accounts
	svcActs, err := global.KubeConfigSet.CoreV1().ServiceAccounts("").List(ctx, metav1.ListOptions{})
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(svcActs.Items)),
			Name:  "ServiceAccounts",
		})
	}

	// cluster roles
	rls, err := global.KubeConfigSet.RbacV1().ClusterRoles().List(ctx, metav1.ListOptions{})
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(rls.Items)),
			Name:  "ClusterRoles",
		})
	}

	//role binding
	rb, err := global.KubeConfigSet.RbacV1().RoleBindings("").List(ctx, metav1.ListOptions{})
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(rb.Items)),
			Name:  "RoleBindings",
		})
	}
	// cluster role bindings

	clsrb, err := global.KubeConfigSet.RbacV1().ClusterRoleBindings().List(ctx, metav1.ListOptions{})
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Value: strconv.Itoa(len(clsrb.Items)),
			Name:  "ClusterRoleBindings",
		})
	}

	//get rgb

	for _, item := range result {
		item.Color = util.GenerateRGB(item.Value)
	}
	return result
}

func (m *MetricsServiceImpl) GetClusterUsageInfo(ctx context.Context) []*metrics.MetricsItem {
	result := make([]*metrics.MetricsItem, 0)
	//url := "/apis/metrics.k8s.io/v1beta1/nodes"
	raw, err := global.KubeConfigSet.RESTClient().Get().AbsPath(global.CONF.System.MetricsServerUrl).DoRaw(ctx)
	if err != nil {
		return result
	}
	nodeMetricsList := metrics.NewNodeMetricsList()

	err = json.Unmarshal(raw, &nodeMetricsList)
	if err != nil {
		return result
	}

	nodeList, err := global.KubeConfigSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return result
	}
	if len(nodeList.Items) != len(nodeMetricsList.Items) {
		return result
	}
	var cpuUsage, cpuTotal, memUsage, memTotal int64

	podsList, err := global.KubeConfigSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return result
	}
	var podUsage, podsTotal int64 = int64(len(podsList.Items)), 0

	for i, item := range nodeList.Items {
		cpuUsage += nodeMetricsList.Items[i].Usage.Cpu().MilliValue() // cpu usaage value should be millivalue
		memUsage += nodeMetricsList.Items[i].Usage.Memory().Value()
		cpuTotal += item.Status.Capacity.Cpu().MilliValue()
		memTotal += item.Status.Capacity.Memory().Value()
		podsTotal += item.Status.Capacity.Pods().Value()
	}

	// get usage percentage

	podUsagePercentage := fmt.Sprintf("%.2f", float64(podUsage)/float64(podsTotal)*100)
	result = append(result, &metrics.MetricsItem{
		Value: podUsagePercentage,
		Name:  "Pods Usage",
	})

	memUsagePercentage := fmt.Sprintf("%.2f", float64(memUsage)/float64(memTotal)*100)
	result = append(result, &metrics.MetricsItem{
		Value: memUsagePercentage,
		Label: util.CLUSTER_MEMORY,
		Name:  "Memory Usage",
	})

	cpuUsagePercentage := fmt.Sprintf("%.2f", float64(cpuUsage)/float64(cpuTotal)*100)
	result = append(result, &metrics.MetricsItem{
		Value: cpuUsagePercentage,
		Label: util.CLUSTER_CPU,
		Name:  "CPU Usage",
	})
	return result
}

func (m *MetricsServiceImpl) GetClusterUsageTrends(ctx context.Context) []*metrics.MetricsItem {
	result := make([]*metrics.MetricsItem, 0)
	data, err := m.getMetricsFromPrometheus(util.CLUSTER_CPU)
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Name:  "Cluster CPU Usage",
			Value: data,
		})
	}

	data, err = m.getMetricsFromPrometheus(util.CLUSTER_MEMORY)
	if err == nil {
		result = append(result, &metrics.MetricsItem{
			Name:  "Cluster Memory Usage",
			Value: data,
		})
	}

	return result
}

func (m *MetricsServiceImpl) getMetricsFromPrometheus(metricsName string) (string, error) {
	result := make(map[string][]string)
	addr := fmt.Sprintf("%s://%s:%v", global.CONF.System.Prometheus.Pscheme, global.CONF.System.Prometheus.Phost, global.CONF.System.Prometheus.Pport)
	client, err := promapi.NewClient(promapi.Config{
		Address: addr,
	})

	if err != nil {
		return "", err
	}
	promConn := promv1.NewAPI(client)
	end := time.Now()
	start := end.Add(-time.Hour * 24)
	r := promv1.Range{
		Start: start,
		End:   end,
		Step:  5 * time.Minute,
	}

	queryRange, _, err := promConn.QueryRange(context.TODO(), metricsName, r)
	if err != nil {
		return "", err
	}

	matrix := queryRange.(model.Matrix)
	if len(matrix) == 0 {
		return "", errors.New("prometheus returns no data")
	}
	// x is timestamp, y is the value (usage)
	x, y := make([]string, 0), make([]string, 0)

	for _, value := range matrix[0].Values {
		format := value.Timestamp.Time().Format("15:04")
		x = append(x, format)
		y = append(y, value.Value.String())

	}
	result["x"] = x
	result["y"] = y

	raw, _ := json.Marshal(result)

	return string(raw), nil
}
