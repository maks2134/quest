// @title Quest API
// @version 1.0.0
// @description API для управления процедурами розыска посылок и оформления заявлений о повреждении или утрате
// @host localhost:8000
package main

import (
	"log"
	"tech-quest/internal/app"
	"tech-quest/internal/configs"
	"tech-quest/pkg/database"
)

func main() {
	configs.LoadConfig()
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	app.App()
}
