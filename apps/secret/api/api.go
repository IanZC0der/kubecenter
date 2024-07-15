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
	response.Success(c, "get the config maps", maps)
}
