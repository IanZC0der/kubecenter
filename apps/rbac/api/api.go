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

	v1.GET("/role", r.GetRoleList)
	v1.GET("/role/detail", r.GetRoleDetail)
	v1.POST("/role", r.CreateRole)
	v1.DELETE("/role", r.DeleteRole)
}

// @Summary      get service account list
// @Description	 get service account list
// @Tags         rbac
// @Accept       json
// @Produce      json
// @Param namespace query string false "Retrieve the service account list based on the namespace, not required"
// @Param keyword query string false "Retrieve the service account list based on the keyword, not required"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /rbac/serviceaccount [get]
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

// @Summary      create service account
// @Description.markdown createsa
// @Tags         rbac
// @Accept       json
// @Produce      json
// @Param ServiceAccount body object true "the configs of the service account"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /rbac/serviceaccount [post]
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

// @Summary      delete a service account
// @Description	 delete a service account
// @Tags         rbac
// @Accept       json
// @Produce      json
// @Param namespace query string true "the namespace of the service account to be deleted"
// @Param name query string true "the name of the service account to be deleted"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /rbac/serviceaccount [delete]
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

// @Summary      get role list
// @Description	 get role list
// @Tags         rbac
// @Accept       json
// @Produce      json
// @Param namespace query string false "Retrieve the role list based on the namespace, not required"
// @Param keyword query string false "Retrieve the role list based on the keyword, not required"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /rbac/role [get]
func (r *RBACApiHandler) GetRoleList(c *gin.Context) {
	namespace := c.Query("namespace")
	keyword := c.Query("keyword")
	res, err := r.svc.GetRoleList(c.Request.Context(), namespace, keyword)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "get role list success", res)
}

// @Summary      get role detail
// @Description	 get role detail
// @Tags         rbac
// @Accept       json
// @Produce      json
// @Param namespace query string false "Retrieve the role detail based on the namespace"
// @Param name query string true "Retrieve the role detail based on the name"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /rbac/role/detail [get]
func (r *RBACApiHandler) GetRoleDetail(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Query("name")
	res, err := r.svc.GetRoleDetail(c.Request.Context(), namespace, name)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "get role detail success", res)
}

// @Summary      delete a role
// @Description	 delete a role
// @Tags         rbac
// @Accept       json
// @Produce      json
// @Param namespace query string false "the namespace of the role to be deleted"
// @Param name query string true "the name of the role to be deleted"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /rbac/role [delete]
func (r *RBACApiHandler) DeleteRole(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Query("name")
	err := r.svc.DeleteRole(c.Request.Context(), namespace, name)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "delete role success", name)
}

// @Summary      create role
// @Description.markdown createrole
// @Tags         rbac
// @Accept       json
// @Produce      json
// @Param ServiceAccount body object true "the configs of the role"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /rbac/role [post]
func (r *RBACApiHandler) CreateRole(c *gin.Context) {
	req := rbac.NewRoleRequest()
	if err := c.ShouldBind(req); err != nil {
		response.Failed(c, err)
		return
	}
	result, err := r.svc.CreateRole(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "create role success", result)
}
