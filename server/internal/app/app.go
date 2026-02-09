package app

import (
	"strings"
	"tech-quest/internal/configs"
	"tech-quest/internal/container"
	"tech-quest/internal/routes"
	"tech-quest/pkg/errors"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recover2 "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func App() {
	appConfig := fiber.Config{ErrorHandler: errors.HandlerErrorFormatter, StrictRouting: false}
	app := fiber.New(appConfig)
	cfg := configs.Configs
	app.Use(recover2.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Split(cfg.CORSAllowOrigins, ","),
		AllowMethods: strings.Split(cfg.CORSAllowMethods, ","),
		AllowHeaders: strings.Split(cfg.CORSAllowHeaders, ","),
		ExposeHeaders: []string{
			"Content-Length",
			"Content-Type",
		},
		AllowCredentials: true,
		MaxAge:           cfg.CORSMaxAge,
	}))
	app.Use("/swagger", static.New("./swagger-ui"))
	app.Use("/docs", static.New("./docs"))
	app.Use(logger.New(logger.Config{
		Format: `${time} | ${status} | ${latency} | ${method} | ${url} | ${ip} | ${bytesSent}` + "\n",
	}))
	api := app.Group("api/v1/")
	c := container.NewContainer(api)
	routes.RegisterRoutes(api, c.Handlers().ProcedureHandler)
	if err := app.Listen(":8000"); err != nil {
		panic(err)
	}
}
