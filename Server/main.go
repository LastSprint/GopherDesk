package main

//go:generate go-localize -input localizations_src -output localizations

import (
	"github.com/LastSprint/GopherDesk/Api/Slack"
	"github.com/LastSprint/GopherDesk/Api/Trello"
	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type config struct {
	SlackToken        string `env:"SLACK_TOKEN,unset"`
	TrelloToken       string `env:"TRELLO_TOKEN,unset"`
	TrelloApiKey      string `env:"TRELLO_API_KEY,unset"`
	SlackSignInSecret string `env:"SLACK_SIGN_IN_KEY,unset"`

	TrelloCallbackUrl string `env:"TRELLO_CALLBACK_URL"`
}

func main() {
	cfg := config{}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Can't parse env config with error", err.Error())
		return
	}

	trelloController := Trello.AssembleTrelloController(cfg.SlackToken, cfg.TrelloApiKey, cfg.TrelloToken, cfg.TrelloCallbackUrl)
	slackController := Slack.AssembleTrelloController(cfg.SlackSignInSecret, cfg.SlackToken, cfg.TrelloToken, cfg.TrelloApiKey)

	r := chi.NewRouter()

	r.Use(middleware.DefaultLogger)

	r.Route("/api/v1", func(r chi.Router) {
		trelloController.Start(r)
		slackController.Start(r)
	})

	log.Fatalln(http.ListenAndServe(":80", r))
}
