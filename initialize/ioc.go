package initialize

import (
	"fmt"
	"github.com/IanZC0der/kubecenter/ioc"
	"os"
)

func IOCInit() {
	if err := ioc.DefaultControllerContainer().Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := ioc.DefaultApiHandlerContainer().Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
