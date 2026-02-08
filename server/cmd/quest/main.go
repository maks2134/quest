// @title Quest API
// @version 1.0.0
// @description API для управления процедурами розыска посылок и оформления заявлений о повреждении или утрате
// @host localhost:8000
// @BasePath /api/v1
package main

import (
	"log"
	"tech-quest/pkg/database"

	"tech-quest/internal/app"
	"tech-quest/internal/configs"
)

func main() {
	configs.LoadConfig()
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	app.App()
}
