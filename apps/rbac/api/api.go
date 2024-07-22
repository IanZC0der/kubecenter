package api

import (
	"github.com/IanZC0der/kubecenter/apps/rbac"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/response"
	"github.com/gin-gonic/gin"
)

type RBACApiHandler struct {
	svc rbac.Service
}

var _ ioc.GinApiHandler = &RBACApiHandler{}

func init() {
	ioc.DefaultApiHandlerContainer().Register(&RBACApiHandler{})
}

func (r *RBACApiHandler) Init() error {
	r.svc = ioc.DefaultControllerContainer().Get(rbac.AppName).(rbac.Service)
	return nil
}

func (r *RBACApiHandler) Name() string {
	return rbac.AppName
}

func (r *RBACApiHandler) Registry(router gin.IRouter) {
	v1 := router.Group("/rbac")
	v1.GET("/serviceaccount", r.GetServiceAccountList)
	v1.POST("/serviceaccount", r.CreateServiceAccount)
	v1.DELETE("/serviceaccount", r.DeleteServiceAccount)
}

func (r *RBACApiHandler) GetServiceAccountList(c *gin.Context) {
	namespace := c.Query("namespace")
	keyword := c.Query("keyword")
	list, err := r.svc.GetServiceAccountList(c.Request.Context(), namespace, keyword)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, "get service account list success", list)
}

func (r *RBACApiHandler) CreateServiceAccount(c *gin.Context) {
	req := rbac.NewCreateServiceAccountRequest()
	if err := c.ShouldBind(req); err != nil {
		response.Failed(c, err)
		return
	}
	result, err := r.svc.CreateServiceAccount(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "create service account success", result)
}

func (r *RBACApiHandler) DeleteServiceAccount(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Query("name")
	err := r.svc.DeleteServiceAccount(c.Request.Context(), namespace, name)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "delete service account success", name)
}
