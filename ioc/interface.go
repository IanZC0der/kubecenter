package ioc

type IocObject interface {
	Init() error

	Name() string
}
