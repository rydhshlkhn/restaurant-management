package usecase

import (
	"context"
	"time"

	"github.com/rydhshlkhn/restaurant-management/internal/modules/menu/domain"
	"github.com/rydhshlkhn/restaurant-management/pkg/helper"
	sharedDomain "github.com/rydhshlkhn/restaurant-management/pkg/shared/domain"
	"github.com/rydhshlkhn/restaurant-management/pkg/shared/repository"
)

type menuUsecaseImpl struct {
	repoSQL repository.RepoSQL
}

func NewMenuUsecase() MenuUsecase {
	return &menuUsecaseImpl{
		repoSQL: repository.GetSharedRepoSQL(),
	}
}

func (uc *menuUsecaseImpl) GetAllMenu(ctx context.Context, filter *domain.MenuFilter) (result []domain.MenuResponse, meta helper.Meta, err error) {
	menus, err := uc.repoSQL.MenuRepo().FetchAll(ctx, filter)
	if err != nil {
		return
	}
	count := uc.repoSQL.MenuRepo().Count(ctx, filter)
	meta = helper.NewMeta(filter.Page, filter.Limit, count)
	for _, menu := range menus {
		result = append(result, domain.MenuResponse{
			ID:        menu.ID,
			Name:      menu.Name,
			Category:  menu.Category,
			StartDate: menu.StartDate,
			EndDate:   menu.EndDate,
			CreatedAt: menu.CreatedAt,
			UpdatedAt: menu.UpdatedAt,
			DeletedAt: menu.DeletedAt,
		})
	}
	return
}

func (uc *menuUsecaseImpl) GetDetailMenu(ctx context.Context, id string) (result domain.MenuResponse, err error) {
	var data sharedDomain.Menu
	filter := domain.MenuFilter{ID: &id}
	data, err = uc.repoSQL.MenuRepo().Find(ctx, &filter)
	if err != nil {
		return
	}

	result.ID = data.ID
	result.Name = data.Name
	result.Category = data.Category
	result.StartDate = data.StartDate
	result.EndDate = data.EndDate
	result.CreatedAt = data.CreatedAt
	result.UpdatedAt = data.UpdatedAt
	result.DeletedAt = data.DeletedAt

	return
}

func (uc *menuUsecaseImpl) CreateMenu(ctx context.Context, menuModel domain.MenuCreateOrUpdateModel) (*sharedDomain.Menu, error) {
	start_date, err := helper.DateStringToDatetime(menuModel.StartDate)
	if err != nil {
		return &sharedDomain.Menu{}, err
	}
	end_date, err := helper.DateStringToDatetime(menuModel.EndDate)
	if err != nil {
		return &sharedDomain.Menu{}, err
	}

	menu := sharedDomain.Menu{
		Name:      menuModel.Name,
		Category:  menuModel.Category,
		StartDate: &start_date,
		EndDate:   &end_date,
	}
	return uc.repoSQL.MenuRepo().Save(ctx, &menu)
}

func (uc *menuUsecaseImpl) UpdateMenu(ctx context.Context, menuModel domain.MenuCreateOrUpdateModel) (*sharedDomain.Menu, error) {
	var menu sharedDomain.Menu
	filter := domain.MenuFilter{ID: &menuModel.ID}
	menu, err := uc.repoSQL.MenuRepo().Find(ctx, &filter)
	if err != nil {
		return &sharedDomain.Menu{}, err
	}

	start_date, err := helper.DateStringToDatetime(menuModel.StartDate)
	if err != nil {
		return &sharedDomain.Menu{}, err
	}
	end_date, err := helper.DateStringToDatetime(menuModel.EndDate)
	if err != nil {
		return &sharedDomain.Menu{}, err
	}

	menu.Name = menuModel.Name
	menu.Category = menuModel.Category
	menu.StartDate = &start_date
	menu.EndDate = &end_date

	return uc.repoSQL.MenuRepo().Save(ctx, &menu)
}

func (uc *menuUsecaseImpl) DeleteMenu(ctx context.Context, id string) (err error) {
	filter := domain.MenuFilter{ID: &id}
	menu, err := uc.repoSQL.MenuRepo().Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	menu.DeletedAt = &deleted_at
	_, err = uc.repoSQL.MenuRepo().Save(ctx, &menu)
	return
}
