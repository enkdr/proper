package main

import (
	"proper/cmd"
)

func main() {
	cmd.Execute()
}

// import (
// 	"github.com/gorilla/handlers"
// 	"github.com/gorilla/mux"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"os"
// 	"proper/app"
// 	"proper/db"
// )

// func main() {

// 	database, err := db.CreateDatabase()
// 	if err != nil {
// 		log.Fatal("Database connection failed")
// 	}

// 	app := &app.App{
// 		Database: database,
// 		Router:   mux.NewRouter(),
// 		Tmpl:     template.Must(template.ParseGlob("templates/*.tmpl")),
// 	}

// 	app.InitializeRoutes()

// 	log.Fatal(http.ListenAndServe(":8000", handlers.
// 		LoggingHandler(os.Stdout, app.Router)))
// }
