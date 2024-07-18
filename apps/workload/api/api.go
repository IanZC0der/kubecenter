package api

import (
	"github.com/IanZC0der/kubecenter/apps/workload"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/response"
	"github.com/gin-gonic/gin"
)

type WorkloadApiHandler struct {
	svc workload.Service
}

var _ ioc.GinApiHandler = &WorkloadApiHandler{}

func init() {
	ioc.DefaultApiHandlerContainer().Register(&WorkloadApiHandler{})
}

func (w *WorkloadApiHandler) Init() error {
	w.svc = ioc.DefaultControllerContainer().Get(workload.AppName).(workload.Service)
	return nil
}

func (w *WorkloadApiHandler) Name() string {
	return workload.AppName
}
func (w *WorkloadApiHandler) Registry(router gin.IRouter) {
	v1 := router.Group("/workload")
	v1.GET("")
	v1.GET("/detail")
	v1.POST("")
	v1.DELETE("")
}

func (w *WorkloadApiHandler) GetStatefulSetList(c *gin.Context) {
	namespace := c.Query("namespace")
	keyword := c.Query("keyword")

	res, err := w.svc.GetStatefulSetList(c.Request.Context(), namespace, keyword)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "get statefulSet list success", res)
}

func (w *WorkloadApiHandler) GetStatefulSetDetail(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Query("name")

	res, err := w.svc.GetStatefulSetDetail(c.Request.Context(), namespace, name)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "get statefulSet detail success", res)
}

func (w *WorkloadApiHandler) CreateOrUpdateStatefulSet(c *gin.Context) {
	req := workload.NewStatefulSet()
	if err := c.ShouldBind(req); err != nil {
		response.Failed(c, err)
		return
	}

	res, msg, err := w.svc.CreateOrUpdateStatefulSet(c.Request.Context(), req)
	if err != nil {
		response.FailedWithMsg(c, msg, err)
		return
	}
	response.Success(c, "create statefulSet success", res)
}

func (w *WorkloadApiHandler) DeleteStatefulSet(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Query("name")

	err := w.svc.DeleteStatefulSet(c.Request.Context(), namespace, name)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "delete statefulSet success", nil)
}
