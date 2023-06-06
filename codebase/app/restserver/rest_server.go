package restserver

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rydhshlkhn/restaurant-management/codebase/factory"
)

type restServer struct {
	serverEngine *fiber.App
}

func NewServer(service factory.ServiceFactory) factory.AppServerFactory {
	server := &restServer{
		serverEngine: fiber.New(),
	}

	for _, m := range service.GetModules() {
		m.RESTHandler().Route(server.serverEngine)
	}

	return server
}

func (r *restServer) Serve() {
	log.Fatal(
		r.serverEngine.Listen(":8001"),
	)
}
