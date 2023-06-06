package usecase

import (
	"context"

	"github.com/rydhshlkhn/restaurant-management/pkg/shared/domain"
	"github.com/rydhshlkhn/restaurant-management/pkg/shared/repository"
)

type FoodUsecase interface {
	GetAllFood(ctx context.Context) domain.Food
}

type foodUsecaseImpl struct {
	repoSQL repository.RepoSQL
}

func NewFoodUsecase() FoodUsecase {
	return &foodUsecaseImpl{
		repoSQL: repository.GetSharedRepoSQL(),
	}
}

func (u *foodUsecaseImpl) GetAllFood(ctx context.Context) (result domain.Food) {
	result = *u.repoSQL.FoodRepo().FetchAll(ctx)
	return
}
