package api

import (
	"github.com/IanZC0der/kubecenter/apps/k8sservice"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/response"
	"github.com/gin-gonic/gin"
)

type K8SServiceApiHandler struct {
	svc k8sservice.Service
}

var _ ioc.GinApiHandler = &K8SServiceApiHandler{}

func init() {
	ioc.DefaultApiHandlerContainer().Register(&K8SServiceApiHandler{})
}

func (k *K8SServiceApiHandler) Init() error {
	k.svc = ioc.DefaultControllerContainer().Get(k8sservice.AppName).(k8sservice.Service)
	return nil
}

func (k *K8SServiceApiHandler) Name() string {
	return k8sservice.AppName
}

func (k *K8SServiceApiHandler) Registry(router gin.IRouter) {
	v1 := router.Group("k8sservice")
	v1.GET("", k.GetSvcList)
	v1.GET("/detail", k.GetSvcDetail)
	v1.POST("", k.CreateOrUpdateSvc)
	v1.DELETE("", k.DeleteSvc)
}

// @Summary      get service list
// @Description	 get service list
// @Tags         service
// @Accept       json
// @Produce      json
// @Param namespace query string false "Retrieve the service list based on the namespace, not required"
// @Param keyword query string false "Retrieve the service list based on the keyword, not required"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /k8sservice [get]
func (k *K8SServiceApiHandler) GetSvcList(ctx *gin.Context) {
	namespace := ctx.Query("namespace")
	keyword := ctx.Query("keyword")

	res, err := k.svc.GetSvcList(ctx.Request.Context(), namespace, keyword)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, "get service list success", res)
}

// @Summary      get service detail
// @Description	 get service detail
// @Tags         service
// @Accept       json
// @Produce      json
// @Param namespace query string true "Retrieve the service info based on the namespace, required"
// @Param name query string true "Retrieve the service info based on the name, required"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /k8sservice/detail [get]
func (k *K8SServiceApiHandler) GetSvcDetail(ctx *gin.Context) {
	namespace := ctx.Query("namespace")
	name := ctx.Query("name")
	res, err := k.svc.GetSvcDetail(ctx.Request.Context(), namespace, name)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, "get service detail success", res)
}

// @Summary      delete service
// @Description	 delete service
// @Tags         service
// @Accept       json
// @Produce      json
// @Param namespace query string true "Delete the service based on the namespace, required"
// @Param name query string true "Delete the service based on the name, required"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /k8sservice [delete]
func (k *K8SServiceApiHandler) DeleteSvc(ctx *gin.Context) {
	namespace := ctx.Query("namespace")
	name := ctx.Query("name")
	err := k.svc.DeleteSvc(ctx.Request.Context(), namespace, name)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, "delete service success", struct{}{})
}

// @Summary     create/update service
// @Description.markdown createsvc
// @Tags         service
// @Accept       json
// @Produce      json
// @Param Service body object true "The configs of the service to be created/updated"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /k8sservice [post]
func (k *K8SServiceApiHandler) CreateOrUpdateSvc(ctx *gin.Context) {
	req := k8sservice.NewK8SService()
	if err := ctx.ShouldBind(req); err != nil {
		response.Failed(ctx, err)
		return
	}
	res, msg, err := k.svc.CreateOrUpdateSvc(ctx.Request.Context(), req)

	if err != nil {
		response.FailedWithMsg(ctx, msg, err)
		return
	}
	response.Success(ctx, msg, res)
}
