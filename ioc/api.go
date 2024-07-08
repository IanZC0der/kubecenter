package ioc

func DefaultApiHandlerContainer() *IocContainer {
	return ApiHandlerIocContainer
}

var ApiHandlerIocContainer = &IocContainer{
	store: map[string]IocObject{},
}
