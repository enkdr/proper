package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	"html/template"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
	Tmpl     *template.Template
}
