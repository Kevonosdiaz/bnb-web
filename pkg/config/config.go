package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// "Global" config file that can be imported throughout the project
// Avoid importing unnecessary and external packages in this kind of distributed file

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
