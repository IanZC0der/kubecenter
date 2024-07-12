package api

import (
	"github.com/IanZC0der/kubecenter/apps/configmap"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/response"
	"github.com/gin-gonic/gin"
)

type ConfigmapApiHandler struct {
	svc configmap.Service
}

var _ ioc.GinApiHandler = &ConfigmapApiHandler{}

func init() {
	ioc.DefaultApiHandlerContainer().Register(&ConfigmapApiHandler{})
}

func (h *ConfigmapApiHandler) Init() error {
	h.svc = ioc.DefaultControllerContainer().Get(configmap.AppName).(configmap.Service)
	return nil
}

func (h *ConfigmapApiHandler) Name() string {
	return configmap.AppName
}

func (h *ConfigmapApiHandler) Registry(router gin.IRouter) {
	v1 := router.Group("configmap")
	v1.GET("", h.GetConfigMaps)
	v1.GET("/detail", h.GetConfigMapDetail)
	v1.POST("", h.UpdateConfigMap)
}

// @Summary      get configmap list
// @Description	 get configmap list
// @Tags         configmap
// @Accept       json
// @Produce      json
// @Param namespace query string false "Retrieve the configmap list based on the namespace, not required"
// @Param keyword query string false "Retrieve the configmap list based on the keyword, not required"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /configmap [get]
func (h *ConfigmapApiHandler) GetConfigMaps(c *gin.Context) {
	namespace := c.Query("namespace")
	keyword := c.Query("keyword")
	maps, err := h.svc.GetConfigMapList(c.Request.Context(), namespace, keyword)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "get the config maps", maps)
}

// @Summary      get configmap detail
// @Description	 get configmap detail
// @Tags         configmap
// @Accept       json
// @Produce      json
// @Param namespace query string true "Retrieve the configmap detail based on the namespace"
// @Param name query string true "Retrieve the configmap detail based on the name"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /configmap/detail [get]
func (h *ConfigmapApiHandler) GetConfigMapDetail(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Query("name")
	cMap, err := h.svc.GetConfigMapDetail(c.Request.Context(), namespace, name)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "get the config map detail", cMap)
}

// @Summary      create/update configmap
// @Description.markdown updateconfigmap
// @Tags         configmap
// @Accept       json
// @Produce      json
// @Param configmap body object true "the configs of the config map"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /configmap [post]
func (h *ConfigmapApiHandler) UpdateConfigMap(c *gin.Context) {
	cMap := configmap.NewConfigMap()
	if err := c.BindJSON(cMap); err != nil {
		response.Failed(c, err)
		return
	}
	k8sConfigMap, msg, err := h.svc.CreateOrUpdateConfigMap(c.Request.Context(), cMap)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, msg, k8sConfigMap)
}
