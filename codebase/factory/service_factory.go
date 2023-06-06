package factory

type ServiceFactory interface {
	GetApplication() AppServerFactory
	GetModules() []ModuleFactory
}
