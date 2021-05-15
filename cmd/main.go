package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/BurntSushi/toml"
	"github.com/amartery/trinity-task/config"
	"github.com/amartery/trinity-task/internal/utility"
	"github.com/amartery/trinity-task/pkg/app/delivery/http"
	"github.com/amartery/trinity-task/pkg/app/repository/postgres"
	"github.com/amartery/trinity-task/pkg/app/usecase"
)

var (
	configPath string = "./config/conf.toml"
)

// @title Trinity-task
// @version 1.0
// @description API for statistic server

// @host localhost:8080
// @BasePath /

func main() {
	config := config.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	postgresCon, err := utility.CreatePostgresConnection(config.DataBaseURL)
	if err != nil {
		log.Fatal(err)
	}

	appRep := postgres.NewAppRepository(postgresCon)
	appUsecase := usecase.NewAppUsecase(appRep)
	handlers := http.NewAppHandlers(appUsecase)

	app := fiber.New()
	handlers.ConfigureRoutes(app)

	app.Listen(config.BindAddr)

	defer utility.Close(postgresCon)
}
