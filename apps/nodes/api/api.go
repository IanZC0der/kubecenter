package api

import (
	"github.com/IanZC0der/kubecenter/apps/nodes"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/response"
	"github.com/gin-gonic/gin"
)

type NodesApiHandler struct {
	svc nodes.Service
}

var _ ioc.GinApiHandler = &NodesApiHandler{}

func init() {
	ioc.DefaultApiHandlerContainer().Register(&NodesApiHandler{})
}

func (n *NodesApiHandler) Init() error {
	n.svc = ioc.DefaultControllerContainer().Get(nodes.AppName).(nodes.Service)
	return nil
}

func (n *NodesApiHandler) Name() string {
	return nodes.AppName
}

func (n *NodesApiHandler) Registry(router gin.IRouter) {
	v1 := router.Group("nodes")
	v1.GET("", n.GetNodesList)
	v1.GET("/detail", n.GetNodeDetail)
	v1.PATCH("/updatelabels", n.UpdateLabels)
	v1.PATCH("/updatetaints", n.UpdateTaints)
}

// @Summary      get nodes list
// @Description	 get nodes list
// @Tags         nodes
// @Accept       json
// @Produce      json
// @Param keyword query string false "Retrieve the nodes list based on the keyword, not required"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /nodes [get]
func (n *NodesApiHandler) GetNodesList(c *gin.Context) {
	keyword := c.Query("keyword")
	nodesList, err := n.svc.GetNodeList(c.Request.Context(), keyword)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "ok", nodesList)
}

// @Summary      get the detail of a node
// @Description	 get the detail of a node based on its name
// @Tags         nodes
// @Accept       json
// @Produce      json
// @Param name query string true "name of the node"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /nodes/detail [get]
func (n *NodesApiHandler) GetNodeDetail(c *gin.Context) {
	name := c.Query("name")
	node, err := n.svc.GetNodeDetail(c.Request.Context(), name)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "ok", node)

}

// @Summary      update the labels of a node
// @Description.markdown updatelabel
// @Tags         nodes
// @Accept       json
// @Produce      json
// @Param labels body object true "new labels config"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /nodes/updatelabels [patch]
func (n *NodesApiHandler) UpdateLabels(c *gin.Context) {
	req := nodes.NewUpdateLabelRequest()
	if err := c.ShouldBindJSON(req); err != nil {
		response.Failed(c, err)
		return
	}
	err := n.svc.UpdateLabel(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
	}
	response.Success(c, "update label success", nil)
}

// @Summary      update the taints of a node
// @Description.markdown updatetaint
// @Tags         nodes
// @Accept       json
// @Produce      json
// @Param taints body object true "new taints config"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /nodes/updatetaints [patch]
func (n *NodesApiHandler) UpdateTaints(c *gin.Context) {
	req := nodes.NewUpdateTaintRequest()
	if err := c.ShouldBindJSON(req); err != nil {
		response.Failed(c, err)
		return
	}
	err := n.svc.UpdateTaints(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
	}
	response.Success(c, "update taint success", nil)
}
