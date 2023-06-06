package usecase

import (
	"context"

	"github.com/rydhshlkhn/restaurant-management/internal/modules/menu/domain"
	"github.com/rydhshlkhn/restaurant-management/pkg/helper"
	sharedDomain "github.com/rydhshlkhn/restaurant-management/pkg/shared/domain"
)

type MenuUsecase interface {
	GetAllMenu(ctx context.Context, filter *domain.MenuFilter) (result []domain.MenuResponse, meta helper.Meta, err error)
	GetDetailMenu(ctx context.Context, id string) (result domain.MenuResponse, err error)
	CreateMenu(ctx context.Context, data domain.MenuCreateOrUpdateModel) (*sharedDomain.Menu, error)
	UpdateMenu(ctx context.Context, data domain.MenuCreateOrUpdateModel) (*sharedDomain.Menu, error)
	DeleteMenu(ctx context.Context, id string) error
}
