package api

import (
	"github.com/IanZC0der/kubecenter/apps/pods"
	_ "github.com/IanZC0der/kubecenter/docs"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/response"
	"github.com/gin-gonic/gin"
)

type PodsApiHandler struct {
	svc pods.Service
}

var _ ioc.GinApiHandler = &PodsApiHandler{}

func init() {
	ioc.DefaultApiHandlerContainer().Register(&PodsApiHandler{})
}

func (p *PodsApiHandler) Init() error {
	p.svc = ioc.DefaultControllerContainer().Get(pods.AppName).(pods.Service)
	return nil
}

func (p *PodsApiHandler) Name() string {
	return pods.AppName
}
func (p *PodsApiHandler) Registry(router gin.IRouter) {
	v1 := router.Group("pods")
	v1.GET("", p.GetPods)
	v1.GET("/namespacelist", p.GetNamespaceList)
}

// @Summary      get the pods list
// @Description  get the pods list and name spaces
// @Tags         pods
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /pods [get]
func (p *PodsApiHandler) GetPods(c *gin.Context) {
	podsList, err := p.svc.GetPods(c.Request.Context())

	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "get the pods list", podsList)
}

// @Summary      get the namespace list
// @Description  get the namespace list
// @Tags         pods
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /pods/namespacelist [get]
func (p *PodsApiHandler) GetNamespaceList(c *gin.Context) {
	namespaceList, err := p.svc.GetNamespaceList(c.Request.Context())

	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, "get the namespace list", namespaceList)
}
