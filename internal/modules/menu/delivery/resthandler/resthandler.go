package resthandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rydhshlkhn/restaurant-management/internal/modules/menu/domain"
	"github.com/rydhshlkhn/restaurant-management/pkg/helper"
	"github.com/rydhshlkhn/restaurant-management/pkg/shared/middleware"
	"github.com/rydhshlkhn/restaurant-management/pkg/shared/usecase"
)

type RestHandler struct {
	uc usecase.Usecase
}

func NewResthandler(uc usecase.Usecase) *RestHandler {
	return &RestHandler{uc: uc}
}

func (r *RestHandler) Route(app *fiber.App) {
	menu := app.Group("/menu", middleware.LoggerMiddleware)
	menu.Get("/", r.getAllFood)
	menu.Get("/:id", r.getDetailMenuByID)
	menu.Post("/", r.createMenu)
	menu.Put("/:id", r.updateMenu)
	menu.Delete("/:id", r.deleteMenu)
}

func (r *RestHandler) getAllFood(ctx *fiber.Ctx) error {
	queryParams := domain.NewMenuFilter()
	if err := ctx.QueryParser(queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid parameter",
			Data:    err,
		})
	}

	response, meta, err := r.uc.Menu().GetAllMenu(ctx.Context(), queryParams)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Fetch data failed: %s", err.Error()),
			Data:    response,
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(helper.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    response,
		Meta:    meta,
	})
}

func (r *RestHandler) getDetailMenuByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	menu, err := r.uc.Menu().GetDetailMenu(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Get detail menu failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(helper.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    menu,
	})
}

func (r *RestHandler) createMenu(ctx *fiber.Ctx) error {
	var request domain.MenuCreateOrUpdateModel
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
	}
	if err = request.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request",
			Errors:  helper.Extract(err),
		})
	}

	result, err := r.uc.Menu().CreateMenu(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Create Menu failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(helper.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

func (r *RestHandler) updateMenu(ctx *fiber.Ctx) error {
	var request domain.MenuCreateOrUpdateModel
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
	}
	if err := request.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request",
			Errors:  helper.Extract(err),
		})
	}

	request.ID = ctx.Params("id")
	result, err := r.uc.Menu().UpdateMenu(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Update Menu failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(helper.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

func (r *RestHandler) deleteMenu(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := r.uc.Menu().DeleteMenu(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.GeneralResponse{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("Delete Menu failed: %s", err.Error()),
		})
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(helper.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: fmt.Sprintf("Menu ID: %s has been deleted", id),
	})
}
