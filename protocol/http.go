package protocol

import (
	"context"
	"fmt"
	docs "github.com/IanZC0der/kubecenter/docs"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/gin-contrib/cors"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"net/http"

	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() *HttpServer {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.Host = global.APP().HttpAddress()
	docs.SwaggerInfo.BasePath = "/api/kubecenter/v1"

	//r.Use(middlewares.Cors)
	ioc.DefaultApiHandlerContainer().RouterRegistry(r.Group("/api/kubecenter/v1"))
	return &HttpServer{
		sver: &http.Server{
			Addr:    global.APP().HttpAddress(),
			Handler: r,
		},
	}
}

type HttpServer struct {
	sver *http.Server
}

func (s *HttpServer) Run() error {
	fmt.Printf("Server starts at %s\n", global.APP().HttpAddress())
	return s.sver.ListenAndServe()
}

func (s *HttpServer) Close(ctx context.Context) {
	s.sver.Shutdown(ctx)
}
