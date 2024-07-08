package ioc

import "github.com/gin-gonic/gin"

var controllerIocContainer = &IocContainer{
	store: map[string]IocObject{},
}

type IocContainer struct {
	store map[string]IocObject
}

func (c *IocContainer) Init() error {
	for _, obj := range c.store {
		if err := obj.Init(); err != nil {
			return err
		}
	}
	return nil
}

func DefaultControllerContainer() *IocContainer {
	return controllerIocContainer
}

func (c *IocContainer) Register(obj IocObject) {
	c.store[obj.Name()] = obj

}

func (c *IocContainer) Get(name string) IocObject {
	return c.store[name]
}

type GinApiHandler interface {
	Registry(gin.IRouter)
}

func (c *IocContainer) RouterRegistry(router gin.IRouter) {
	for _, obj := range c.store {
		if apiHandler, ok := obj.(GinApiHandler); ok {
			apiHandler.Registry(router)
		}
	}
}
