package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oscarracuna/go-clocko/pkg/config"
	"github.com/oscarracuna/go-clocko/pkg/handlers"
	"github.com/oscarracuna/go-clocko/pkg/render"
)

func main() {
	portNumber := "localhost:8484"

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Unable to create template cache.")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("App successfully started on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
