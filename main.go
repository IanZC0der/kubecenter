package main

import (
	"fmt"
	_ "github.com/IanZC0der/kubecenter/apps"
	"github.com/IanZC0der/kubecenter/initialize"
	"github.com/IanZC0der/kubecenter/protocol"
	"os"
)

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
