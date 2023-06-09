package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rydhshlkhn/restaurant-management/pkg/helper"
)

type MenuResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Category  string     `json:"category"`
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type MenuFilter struct {
	helper.Filter
	ID       *string `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
}

func NewMenuFilter() *MenuFilter {
	return &MenuFilter{
		Filter:   *helper.DefaultFilter(),
		ID:       new(string),
		Name:     "",
		Category: "",
	}
}

type MenuCreateOrUpdateModel struct {
	ID        string `json:"id" validate:"max=36"`
	Name      string `json:"name" validate:"required,min=1"`
	Category  string `json:"category" validate:"required,min=1"`
	StartDate string `json:"startDate" validate:"required,datetime=2006-01-02"`
	EndDate   string `json:"endDate" validate:"required,datetime=2006-01-02"`
}

func (f *MenuCreateOrUpdateModel) Validate() error {
	validate := validator.New()
	return validate.Struct(f)
}
