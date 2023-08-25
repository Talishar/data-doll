package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	// register the migration files from the migrations package
	_ "talishar/data-doll/migrations"
)

var devMode bool

func main() {
	// pocketbase, the most important bit of the app.
	app := pocketbase.New()

	log.Println("🚀 Running the server")

	// Define a flag for dev mode, eagerly parse the arguments. In dev mode automigrate is enabled.
	// enable auto creation of migration files when making collection changes in the Admin UI
	app.RootCmd.PersistentFlags().BoolVar(&devMode, "dev", false, "run in dev mode")
	app.RootCmd.ParseFlags(os.Args[1:])
	if devMode {
		log.Println("🖥️ Running in dev mode")
		migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
			Automigrate: true,
		})
	}

	// add a route to say 'hello world'
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/hello", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello world!")
		}, apis.ActivityLogger(app), apis.RequireAdminAuth())
		return nil
	})

	log.Println("🤖 Data Doll is online")

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
