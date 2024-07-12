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
	v1.GET("/:namespace", p.GetPodsListUnderNamespace)
	v1.GET("/nodepods/:nodename", p.GetPodsListWithNode)
	v1.GET("/poddetail/:namespace", p.GetPodDetail)
	v1.POST("", p.CreatePods)
	v1.PUT("", p.UpdatePod)
	v1.DELETE("", p.DeletePod)
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

// @Summary      get the pods list under a namespace
// @Description	 get the pods list given a namespace, with optional keyword
// @Tags         pods
// @Accept       json
// @Produce      json
// @Param namespace path string true "Namespace"
// @Param keyword query string false "Filter pods by keyword"
// @Param nodename query string false "Filter pods by nodename"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /pods/{namespace} [get]
func (p *PodsApiHandler) GetPodsListUnderNamespace(c *gin.Context) {
	namespace := c.Param("namespace")
	keyword := c.Query("keyword")
	//nodeName := c.Query("nodename")

	podsList, err := p.svc.GetPodsListUnderNamespaceWithKeyword(c.Request.Context(), namespace, keyword)

	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "get the pods list", podsList)
}

// @Summary      get the pods list within a node
// @Description	 get the pods list within a node, with optional keyword
// @Tags         pods
// @Accept       json
// @Produce      json
// @Param nodename path string true "Name of the node"
// @Param keyword query string false "Filter pods by keyword"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /pods/nodepods/{nodename} [get]
func (p *PodsApiHandler) GetPodsListWithNode(c *gin.Context) {
	nodeName := c.Param("nodename")
	keyword := c.Query("keyword")

	podsList, err := p.svc.GetPodsListWithinNode(c.Request.Context(), keyword, nodeName)

	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "get the pods list", podsList)
}

// @Summary      get a pod detail
// @Description	 get a pod detail under a namespace
// @Tags         pods
// @Accept       json
// @Produce      json
// @Param namespace path string true "Namespace"
// @Param keyword query string true "The name of the pod"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /pods/poddetail/{namespace} [get]
func (p *PodsApiHandler) GetPodDetail(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Query("name")
	podDetail, err := p.svc.GetPodDetail(c.Request.Context(), namespace, name)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "get the pod detail success", podDetail)
}

// @Summary create a Kubernetes pod
//
// @Description
// @Description.markdown podexample
// @Tags pods
// @Accept json
// @Produce json
//
// @Param pod body object true "Pod Configuration"
//
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure      500  {object}  response.Response
// @Router /pods [post]
func (p *PodsApiHandler) CreatePods(c *gin.Context) {
	newPod := pods.NewPod()
	if err := c.ShouldBind(newPod); err != nil {
		response.Failed(c, err)
		return
	}

	err := pods.PodCreateValidate(newPod)
	if err != nil {
		response.Failed(c, err)
		return
	}
	//k8sPod := pods.CreatePodFromPodRequest(newPod)
	k8sPod, err := p.svc.CreatePod(c.Request.Context(), newPod)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, "create pod success", k8sPod)
}

// @Summary update a Kubernetes pod
//
// @Description.markdown updatepod
// @Tags pods
// @Accept json
// @Produce json
//
// @Param pod body object true "Pod Configuration"
//
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure      500  {object}  response.Response
// @Router /pods [put]
func (p *PodsApiHandler) UpdatePod(c *gin.Context) {
	newPod := pods.NewPod()
	if err := c.ShouldBind(newPod); err != nil {
		response.Failed(c, err)
		return
	}

	err := pods.PodCreateValidate(newPod)
	if err != nil {
		response.Failed(c, err)
		return
	}

	k8sPod, msg, err := p.svc.UpdatePod(c.Request.Context(), newPod)
	if err != nil {
		response.FailedWithMsg(c, msg, err)
		return
	}

	response.Success(c, msg, k8sPod)

}

// @Summary      delete a pod
// @Description	 delete a pod based on namespace and name
// @Tags         pods
// @Accept       json
// @Produce      json
// @Param namespace query string true "Namespace"
// @Param name query string true "The name of the pod"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /pods [delete]
func (p *PodsApiHandler) DeletePod(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Query("name")
	k8sPod, msg, err := p.svc.DeletePod(c.Request.Context(), namespace, name)
	if err != nil {
		response.FailedWithMsg(c, msg, err)
		return
	}
	response.Success(c, "delete pod success", k8sPod)
}
