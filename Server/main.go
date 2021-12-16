package main

import (
	"github.com/LastSprint/GopherDesk/Api"
	"github.com/LastSprint/GopherDesk/Api/Slack"
	"github.com/LastSprint/GopherDesk/Api/Trello"
	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type config struct {
	SlackToken  string `env:"SLACK_TOKEN,unset"`
	TrelloToken string `env:"TRELLO_TOKEN,unset"`
}

func main() {
	cfg := config{}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Can't parse env config with error", err.Error())
		return
	}

	controller := Api.AssembleApi()
	trelloController := Trello.AssembleTrelloController()
	slackController := Slack.AssembleTrelloController()

	r := chi.NewRouter()

	r.Use(middleware.DefaultLogger)

	r.Route("/api/v1", func(r chi.Router) {
		controller.Start(r)
		trelloController.Start(r)
		slackController.Start(r)
	})

	log.Fatalln(http.ListenAndServe(":80", r))
}
