package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andkolbe/go-practice/pkg/config"
	"github.com/andkolbe/go-practice/pkg/handlers"
	"github.com/andkolbe/go-practice/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting app on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
