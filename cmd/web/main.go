package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Kevonosdiaz/bnb-web/pkg/config"
	"github.com/Kevonosdiaz/bnb-web/pkg/handlers"
	"github.com/Kevonosdiaz/bnb-web/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const PORT_NO = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main app fn
func main() {
	// Change to true when in production
	app.InProduction = false

	// Create a session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	cache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not create template cache:", err)
	}

	app.TemplateCache = cache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PORT_NO))

	// Start the web server, with port being 8080
	srv := &http.Server{
		Addr:    PORT_NO,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
