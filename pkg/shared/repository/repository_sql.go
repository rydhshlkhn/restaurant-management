package repository

import (
	foodrepo "github.com/rydhshlkhn/restaurant-management/internal/modules/food/repository"
	menurepo "github.com/rydhshlkhn/restaurant-management/internal/modules/menu/repository"
	"gorm.io/gorm"
)

type RepoSQL interface {
	FoodRepo() foodrepo.FoodRepository
	MenuRepo() menurepo.MenuRepository
}

type repoSqlImpl struct {
	readDB, writeDB *gorm.DB
	foodRepo        foodrepo.FoodRepository
	menuRepo        menurepo.MenuRepository
}

var globalRepoSQL *repoSqlImpl

func SetSharedRepoSQL(readDB, writeDB *gorm.DB) {
	globalRepoSQL = new(repoSqlImpl)
	globalRepoSQL.readDB = readDB
	globalRepoSQL.writeDB = writeDB
	globalRepoSQL.foodRepo = foodrepo.NewFoodRepoSQL(readDB, writeDB)
	globalRepoSQL.menuRepo = menurepo.NewMenuRepoSQL(readDB, writeDB)
}

func GetSharedRepoSQL() RepoSQL {
	return globalRepoSQL
}

func (r *repoSqlImpl) FoodRepo() foodrepo.FoodRepository {
	return r.foodRepo
}

func (r *repoSqlImpl) MenuRepo() menurepo.MenuRepository {
	return r.menuRepo
}
