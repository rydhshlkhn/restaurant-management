package resthandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rydhshlkhn/restaurant-management/pkg/shared/usecase"
)

type RestHandler struct {
	uc usecase.Usecase
}

func NewResthandler(uc usecase.Usecase) *RestHandler {
	return &RestHandler{uc: uc}
}

func (r *RestHandler) Route(app *fiber.App) {
	food := app.Group("/food")
	food.Get("", r.FindAll)
}

func (r *RestHandler) FindAll(c *fiber.Ctx) error {
	food := r.uc.Food().GetAllFood(c.Context())
	return c.Status(fiber.StatusOK).JSON(food)
}
