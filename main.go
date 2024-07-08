package main

import (
	"fmt"
	_ "github.com/IanZC0der/kubecenter/apps"
	_ "github.com/IanZC0der/kubecenter/docs"
	"github.com/IanZC0der/kubecenter/initialize"
	"github.com/IanZC0der/kubecenter/protocol"
	"os"
)

// @title           Kubecenter API doc
// @version         1.0
// @description     This webapp is for managing the k8s resources

// @contact.name   Ian Zhang
// @BasePath  /api/kubecenter
func main() {
	//panic(initialize.Viper)
	if err := initialize.LoadConfigFromEnv(); err != nil {
		fmt.Println(err)
		return
	}
	initialize.K8S()
	initialize.IOCInit()
	httpServer := protocol.NewHttpServer()

	if err := httpServer.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
