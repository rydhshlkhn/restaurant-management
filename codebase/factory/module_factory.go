package factory

import "github.com/rydhshlkhn/restaurant-management/codebase/interfaces"

type ModuleFactory interface {
	RESTHandler() interfaces.RESTHandler
}