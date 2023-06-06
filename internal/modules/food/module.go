package food

import (
	"github.com/rydhshlkhn/restaurant-management/codebase/interfaces"
	"github.com/rydhshlkhn/restaurant-management/internal/modules/food/delivery/resthandler"
	"github.com/rydhshlkhn/restaurant-management/pkg/shared/usecase"
)

type Module struct {
	restHandler interfaces.RESTHandler
}

func NewModule() *Module {
	var mod Module
	mod.restHandler = resthandler.NewResthandler(usecase.GetSharedUsecase())
	return &mod
}

func (m *Module) RESTHandler() interfaces.RESTHandler {
	return m.restHandler
}
