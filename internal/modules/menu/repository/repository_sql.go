package repository

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/rydhshlkhn/restaurant-management/internal/modules/menu/domain"
	sharedDomain "github.com/rydhshlkhn/restaurant-management/pkg/shared/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MenuRepository interface {
	FetchAll(ctx context.Context, filter *domain.MenuFilter) ([]sharedDomain.Menu, error)
	Count(ctx context.Context, filter *domain.MenuFilter) int
	Find(ctx context.Context, filter *domain.MenuFilter) (sharedDomain.Menu, error)
	Save(ctx context.Context, data *sharedDomain.Menu) (*sharedDomain.Menu, error)
}

type menuRepoSQL struct {
	readDB, writeDB *gorm.DB
}

func NewMenuRepoSQL(readDB, writeDB *gorm.DB) MenuRepository {
	return &menuRepoSQL{
		readDB:  readDB,
		writeDB: writeDB,
	}
}

func (r *menuRepoSQL) FetchAll(ctx context.Context, filter *domain.MenuFilter) (menus []sharedDomain.Menu, err error) {
	err = r.setFilter(r.readDB, filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	}).Limit(filter.Limit).Offset(filter.CalculateOffset()).Find(&menus).Error
	return
}

func (r *menuRepoSQL) Count(ctx context.Context, filter *domain.MenuFilter) (count int) {
	var total int64
	r.setFilter(r.readDB, filter).Model(&sharedDomain.Menu{}).Count(&total)
	count = int(total)
	return
}

func (r *menuRepoSQL) Find(ctx context.Context, filter *domain.MenuFilter) (result sharedDomain.Menu, err error) {
	err = r.setFilter(r.readDB, filter).First(&result).Error
	return
}

func (r *menuRepoSQL) Save(ctx context.Context, menu *sharedDomain.Menu) (*sharedDomain.Menu, error) {
	if menu.ID == "" {
		menu.ID = uuid.NewString()
		err := r.writeDB.Create(menu).Error
		return menu, err
	}
	err := r.writeDB.Save(menu).Error
	return menu, err
}

func (r *menuRepoSQL) setFilter(db *gorm.DB, filter *domain.MenuFilter) *gorm.DB {
	if *filter.ID != "" {
		db = db.Where("id = ?", *filter.ID)
	}

	if filter.Search != "" {
		db = db.Where("id ILIKE '%%' || ? || '%%' OR name ILIKE '%%' || ? || '%%' OR category ILIKE '%%' || ? || '%%'", filter.Search, filter.Search, filter.Search)
	}

	return db
}
