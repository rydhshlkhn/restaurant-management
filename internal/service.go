package restaurantmanagement

import (
	"github.com/rydhshlkhn/restaurant-management/codebase/app/restserver"
	"github.com/rydhshlkhn/restaurant-management/codebase/factory"
	"github.com/rydhshlkhn/restaurant-management/configuration"
	"github.com/rydhshlkhn/restaurant-management/internal/modules/food"
	"github.com/rydhshlkhn/restaurant-management/internal/modules/menu"
	"github.com/rydhshlkhn/restaurant-management/pkg/shared/repository"
	"github.com/rydhshlkhn/restaurant-management/pkg/shared/usecase"
)

type Service struct {
	application factory.AppServerFactory
	module      []factory.ModuleFactory
}

func NewService() factory.ServiceFactory {
	readDB, writeDB, _ := configuration.NewDatabase(configuration.New())
	repository.SetSharedRepoSQL(readDB, writeDB)
	usecase.SetSharedUsecase()

	modules := []factory.ModuleFactory{
		food.NewModule(),
		menu.NewModule(),
	}
	service := &Service{
		module: modules,
	}
	service.application = restserver.NewServer(service)

	return service
}

func (s *Service) GetApplication() factory.AppServerFactory {
	return s.application
}

func (s *Service) GetModules() []factory.ModuleFactory {
	return s.module
}
