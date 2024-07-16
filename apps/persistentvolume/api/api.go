package api

import (
	"github.com/IanZC0der/kubecenter/apps/persistentvolume"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/response"
	"github.com/gin-gonic/gin"
)

type PersistentVolumeApiHandler struct {
	svc persistentvolume.Service
}

var _ ioc.GinApiHandler = &PersistentVolumeApiHandler{}

func init() {
	ioc.DefaultApiHandlerContainer().Register(&PersistentVolumeApiHandler{})
}

func (p *PersistentVolumeApiHandler) Init() error {
	p.svc = ioc.DefaultControllerContainer().Get(persistentvolume.AppName).(persistentvolume.Service)
	return nil
}

func (p *PersistentVolumeApiHandler) Name() string {
	return persistentvolume.AppName
}
func (p *PersistentVolumeApiHandler) Registry(router gin.IRouter) {
	v1 := router.Group("pv")
	v1.GET("", p.GetPVList)
	v1.POST("", p.CreatePV)
	v1.DELETE("", p.DeletePV)
	v1.GET("/pvc", p.GetPVCList)
	v1.POST("/pvc", p.CreatePVC)
	v1.DELETE("/pvc", p.DeletePVC)
}

// @Summary      get persistent volume list
// @Description	 get persistent volume list
// @Tags         persistent volume
// @Accept       json
// @Produce      json
// @Param keyword query string false "Retrieve the volume list based on the keyword, not required"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /pv [get]
func (p *PersistentVolumeApiHandler) GetPVList(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	res, err := p.svc.GetPVList(ctx.Request.Context(), keyword)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, "get persistent volume list", res)
}

// @Summary     	create persistent volume
// @Description.markdown createpv
// @Tags         persistent volume
// @Accept       json
// @Produce      json
// @Param PersistentVolume body object true "The configs of the persistent volume"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /pv [post]
func (p *PersistentVolumeApiHandler) CreatePV(ctx *gin.Context) {
	req := persistentvolume.NewPersistentVolumeReq()
	if err := ctx.ShouldBind(req); err != nil {
		response.Failed(ctx, err)
		return
	}
	res, err := p.svc.CreatePV(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, "create persistent volume", res)
}

// @Summary      delete persistent volume
// @Description	 delete persistent volume
// @Tags         persistent volume
// @Accept       json
// @Produce      json
// @Param name query string true "Delete the volume based on the name, required"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /pv [delete]
func (p *PersistentVolumeApiHandler) DeletePV(ctx *gin.Context) {
	name := ctx.Query("name")
	err := p.svc.DeletePV(ctx.Request.Context(), name)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, "delete persistent volume", name)
}

// @Summary      get persistent volume claim list
// @Description	 get persistent volume claim list
// @Tags         persistent volume
// @Accept       json
// @Produce      json
// @Param namespace query string false "Retrieve the volume claim list based on the namespace, not required"
// @Param name query string false "Retrieve the volume claim list based on the name, not required"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /pv/pvc [get]
func (p *PersistentVolumeApiHandler) GetPVCList(ctx *gin.Context) {
	name := ctx.Query("name")
	namespace := ctx.Query("namespace")
	l, err := p.svc.GetPVCList(ctx.Request.Context(), namespace, name)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, "get persistent volume claim list", l)
}

// @Summary     	create persistent volume claim
// @Description.markdown createpvc
// @Tags         persistent volume
// @Accept       json
// @Produce      json
// @Param PersistentVolumeClaim body object true "The configs of the persistent volume claim"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /pv/pvc [post]
func (p *PersistentVolumeApiHandler) CreatePVC(ctx *gin.Context) {
	req := persistentvolume.NewCreatePersistentVolumeClaimReq()

	if err := ctx.ShouldBind(req); err != nil {
		response.Failed(ctx, err)
		return
	}

	res, err := p.svc.CreatePVC(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, "create persistent volume claim", res)
}

// @Summary      delete persistent volume claim
// @Description	 delete persistent volume claim
// @Tags         persistent volume
// @Accept       json
// @Produce      json
// @Param namespace query string true "Delete the volume claim based on the namespace, required"
// @Param name query string true "Delete the volume claim based on the name, required"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /pv/pvc [delete]
func (p *PersistentVolumeApiHandler) DeletePVC(ctx *gin.Context) {
	name := ctx.Query("name")
	namespace := ctx.Query("namespace")

	err := p.svc.DeletePVC(ctx.Request.Context(), namespace, name)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, "delete persistent volume claim", name)
}
