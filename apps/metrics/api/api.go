package api

import (
	"github.com/IanZC0der/go-myblog/response"
	"github.com/IanZC0der/kubecenter/apps/metrics"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsApiHandler struct {
	svc metrics.Service
}

var _ ioc.GinApiHandler = &MetricsApiHandler{}

func init() {
	ioc.DefaultApiHandlerContainer().Register(&MetricsApiHandler{})
}

func (h *MetricsApiHandler) Init() error {
	h.svc = ioc.DefaultControllerContainer().Get(metrics.AppName).(metrics.Service)
	return nil
}

func (h *MetricsApiHandler) Name() string {
	return metrics.AppName
}

func (h *MetricsApiHandler) Registry(router gin.IRouter) {
	v1 := router.Group("/metrics")
	v1.GET("/dashboard", h.GetDashBoardData)
	v1.GET("/prometheus", h.GetPrometheus)
}

// @Summary      get the metrics of the cluster
// @Description	 get the metrics of the cluster, the result will contain cluster info, resources info in the cluster and cluster memory/cpu usage
// @Tags         metrics
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /metrics/dashboard [get]
func (h *MetricsApiHandler) GetDashBoardData(ctx *gin.Context) {

	res := make(map[string][]*metrics.MetricsItem)
	res["cluster"] = h.svc.GetClusterInfo(ctx.Request.Context())
	res["resource"] = h.svc.GetResourceInfo(ctx.Request.Context())
	res["usage"] = h.svc.GetClusterUsageInfo(ctx.Request.Context())
	res["trends"] = h.svc.GetClusterUsageTrends(ctx.Request.Context())
	response.Success(ctx, res)
}

// @Summary      the api is for Prometheus to collect the data, visit: http://23.251.33.63:30090 to see data collected to prometheus
// @Description	 result will include the data exported to prometheus
// @Tags         metrics
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /metrics/prometheus [get]
func (h *MetricsApiHandler) GetPrometheus(ctx *gin.Context) {
	handler := promhttp.Handler()
	handler.ServeHTTP(ctx.Writer, ctx.Request)
}
