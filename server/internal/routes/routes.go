package routes

import (
	"github.com/gofiber/fiber/v3"
	"tech-quest/internal/handlers"
)

func RegisterRoutes(router fiber.Router, procedureHandler *handlers.ProcedureHandler) {
	procedures := router.Group("/procedures")

	procedures.Get("/", procedureHandler.GetAll)
	procedures.Get("/:id", procedureHandler.GetByID)
	procedures.Get("/type/:type", procedureHandler.GetByType)
	procedures.Post("/", procedureHandler.Create)
	procedures.Put("/:id", procedureHandler.Update)
	procedures.Delete("/:id", procedureHandler.Delete)
}
