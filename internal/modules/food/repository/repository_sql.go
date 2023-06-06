package repository

import (
	"context"

	"github.com/rydhshlkhn/restaurant-management/pkg/shared/domain"
	"gorm.io/gorm"
)

type FoodRepository interface {
	FetchAll(ctx context.Context) *domain.Food
}

type foodRepoSQL struct {
	readDB, writeDB *gorm.DB
}

func NewFoodRepoSQL(readDB, writeDB *gorm.DB) FoodRepository {
	return &foodRepoSQL{
		readDB:  readDB,
		writeDB: writeDB,
	}
}

func (r *foodRepoSQL) FetchAll(ctx context.Context) *domain.Food {
	return &domain.Food{
		ID:    "asdfasdf",
		Name:  "",
		Price: 0,
		Qty:   0,
	}
}
