package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alfredoxyanez/go_prisma_chi_example/database"
	"github.com/alfredoxyanez/go_prisma_chi_example/router"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func (app *Application) Serve() error {
	port := app.Config.Port
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}
	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	defer db.Client.Disconnect()

	config := Config{
		Port: os.Getenv("PORT"),
	}

	app := &Application{
		Config: config,
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
