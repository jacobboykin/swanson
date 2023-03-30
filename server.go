package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

const (
	swansonKindEnvVar = "SWANSON_KIND"

	swansonKindChaos = "chaos"
	swansonKindHappy = "happy"
	swansonKindSad   = "sad"
)

type swanson struct {
	Kind string
}

func (app *application) serve(_ *cli.Context) error {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveTemplate)

	app.logger.Infof("Starting HTTP server at http://localhost%s", app.config.listenAddress)
	return http.ListenAndServe(app.config.listenAddress, nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "swanson.html")

	tmpl, err := template.ParseFiles(lp)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	swansonKind := swansonKindChaos
	swansonKindEnvVarValue := os.Getenv(swansonKindEnvVar)
	switch swansonKindEnvVarValue {
	case swansonKindHappy:
		swansonKind = swansonKindHappy
	case swansonKindSad:
		swansonKind = swansonKindSad
	}

	s := swanson{Kind: swansonKind}

	err = tmpl.ExecuteTemplate(w, "home", s)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
