package api

import (
	"github.com/IanZC0der/kubecenter/apps/secret"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/response"
	"github.com/gin-gonic/gin"
)

type SecretApiHandler struct {
	svc secret.Service
}

var _ ioc.GinApiHandler = &SecretApiHandler{}

func init() {
	ioc.DefaultApiHandlerContainer().Register(&SecretApiHandler{})
}

func (s *SecretApiHandler) Init() error {
	s.svc = ioc.DefaultControllerContainer().Get(secret.AppName).(secret.Service)
	return nil
}

func (s *SecretApiHandler) Name() string {
	return secret.AppName
}

func (s *SecretApiHandler) Registry(router gin.IRouter) {
	v1 := router.Group("secrets")
	v1.GET("", s.GetSecrets)
	v1.GET("/detail", s.GetSecretDetail)
	v1.POST("", s.CreateOrUpdateSecret)

}

// @Summary      get secret list
// @Description	 get secret list
// @Tags         secret
// @Accept       json
// @Produce      json
// @Param namespace query string false "Retrieve the secret list based on the namespace, not required"
// @Param keyword query string false "Retrieve the secret list based on the keyword, not required"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /secrets [get]
func (s *SecretApiHandler) GetSecrets(c *gin.Context) {
	namespace := c.Query("namespace")
	keyword := c.Query("keyword")
	maps, err := s.svc.GetSecretsList(c.Request.Context(), namespace, keyword)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "get the secret list", maps)
}

// @Summary      get secret detail
// @Description	 get secret detail
// @Tags         secret
// @Accept       json
// @Produce      json
// @Param namespace query string true "Retrieve the secret detail based on the namespace"
// @Param name query string true "Retrieve the secret detail based on the name"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /secrets/detail [get]
func (s *SecretApiHandler) GetSecretDetail(c *gin.Context) {
	namespace := c.Query("namespace")
	name := c.Query("name")
	cMap, err := s.svc.GetSecretDetail(c.Request.Context(), namespace, name)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "get the secret detail", cMap)
}

// @Summary      create/update secret
// @Description.markdown updatesecret
// @Tags         secret
// @Accept       json
// @Produce      json
// @Param secret body object true "the configs of the secret"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /secrets [post]
func (s *SecretApiHandler) CreateOrUpdateSecret(c *gin.Context) {
	newSecret := secret.NewSecret()
	if err := c.BindJSON(newSecret); err != nil {
		response.Failed(c, err)
		return
	}
	k8sSecret, msg, err := s.svc.UpdateSecret(c.Request.Context(), newSecret)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, msg, k8sSecret)
}
