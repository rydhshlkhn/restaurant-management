package interfaces

import "github.com/gofiber/fiber/v2"

type RESTHandler interface {
	Route(app *fiber.App)
}
