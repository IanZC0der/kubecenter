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
	v1.GET("/statefulset", w.GetStatefulSetList)
	v1.GET("/statefulset/detail", w.GetStatefulSetDetail)
	v1.POST("/statefulset", w.CreateOrUpdateStatefulSet)
	v1.DELETE("/statefulset", w.DeleteStatefulSet)
}

// @Summary      get statefulSet list
// @Description	 get statefulSet list
// @Tags         workload
// @Accept       json
// @Produce      json
// @Param namespace query string false "Retrieve the statefulSet list based on the namespace, not required"
// @Param keyword query string false "Retrieve the statefulSet list based on the keyword, not required"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /workload/statefulset [get]
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

// @Summary      get statefulset detail
// @Description	 get statefulset detail
// @Tags         workload
// @Accept       json
// @Produce      json
// @Param namespace query string true "Retrieve the statefulset detail based on the namespace"
// @Param name query string true "Retrieve the statefulset detail based on the name"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /workload/statefulset/detail [get]
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

// @Summary      create/update statefulSet
// @Description.markdown createstatefulset
// @Tags         workload
// @Accept       json
// @Produce      json
// @Param statefulSet body object true "the configs of the statefulSet"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /workload/statefulset [post]
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
	response.Success(c, msg, res)
}

// @Summary      delete a statefulSet
// @Description	 delete a statefulSet
// @Tags         workload
// @Accept       json
// @Produce      json
// @Param namespace query string true "the namespace of the statefulSet to be deleted"
// @Param name query string true "the name of the statefulSet to be deleted"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /workload/statefulset [delete]
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
