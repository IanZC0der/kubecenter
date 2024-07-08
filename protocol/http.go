package protocol

import (
	"context"
	"fmt"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/middlewares"
	"net/http"

	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() *HttpServer {
	r := gin.Default()
	r.Use(middlewares.Cors)
	ioc.DefaultApiHandlerContainer().RouterRegistry(r.Group("/api/kubecenter"))
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
