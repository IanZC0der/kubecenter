package api

import (
	"github.com/IanZC0der/kubecenter/apps/pods"
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
	v1 := router.Group("v1").Group("pods")
	v1.GET("/", p.GetPods)
}

func (p *PodsApiHandler) GetPods(c *gin.Context) {
	podsList, err := p.svc.GetPods(c.Request.Context())

	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "get the pods list", podsList)
}
