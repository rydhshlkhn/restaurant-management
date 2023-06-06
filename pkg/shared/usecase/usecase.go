package usecase

import (
	foodusecase "github.com/rydhshlkhn/restaurant-management/internal/modules/food/usecase"
	menuusecase "github.com/rydhshlkhn/restaurant-management/internal/modules/menu/usecase"
)

type Usecase interface {
	Food() foodusecase.FoodUsecase
	Menu() menuusecase.MenuUsecase
}

type usecaseImpl struct {
	foodusecase.FoodUsecase
	menuusecase.MenuUsecase
}

var usecaseInstance *usecaseImpl

func SetSharedUsecase() {
	usecaseInstance = new(usecaseImpl)
	usecaseInstance.FoodUsecase = foodusecase.NewFoodUsecase()
	usecaseInstance.MenuUsecase = menuusecase.NewMenuUsecase()
}

func GetSharedUsecase() Usecase {
	return usecaseInstance
}

func (u *usecaseImpl) Food() foodusecase.FoodUsecase {
	return u.FoodUsecase
}

func (u *usecaseImpl) Menu() menuusecase.MenuUsecase {
	return u.MenuUsecase
}
