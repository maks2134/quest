package container

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"log"
	"tech-quest/internal/handlers"
	"tech-quest/internal/repository"
	"tech-quest/internal/services"
	"tech-quest/pkg/database"
)

type Services struct {
	ProcedureService *services.ProcedureService
}

func (c *Container) NewServices() *Services {
	return &Services{
		ProcedureService: services.NewProcedureService(c.repo.ProcedureRepository),
	}
}

type Handlers struct {
	ProcedureHandler *handlers.ProcedureHandler
}

func (c *Container) NewHandlers() *Handlers {
	return &Handlers{
		ProcedureHandler: handlers.NewProcedureHandler(c.services.ProcedureService),
	}
}

type Repository struct {
	ProcedureRepository *repository.ProcedureRepository
}

func (c *Container) NewRepository() *Repository {
	return &Repository{
		ProcedureRepository: repository.NewProcedureRepository(c.db),
	}
}

type Container struct {
	router   fiber.Router
	db       *sqlx.DB
	services *Services
	handlers *Handlers
	repo     *Repository
}

func NewContainer(router fiber.Router) *Container {
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	c := Container{
		router: router,
		db:     db,
	}
	c.repo = c.NewRepository()
	c.services = c.NewServices()
	c.handlers = c.NewHandlers()
	return &c
}

func (c *Container) Router() fiber.Router {
	return c.router
}

func (c *Container) Handlers() *Handlers {
	return c.handlers
}
