package main

import (
	"github.com/LastSprint/GopherDesk/Api/Slack"
	"github.com/LastSprint/GopherDesk/Api/Trello"
	"github.com/LastSprint/GopherDesk/L10n"
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

	L10nDirPath   string `env:"L10N_LOCALIZATION_FILES_DIR" envDefault:"L10n/locales"`
	CurrentLocale string `env:"CURRENT_LOCALE" envDefault:"en_US"`
}

func main() {
	cfg := config{}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Can't parse env config with error", err.Error())
		return
	}

	if err := L10n.Configure(cfg.L10nDirPath, cfg.CurrentLocale); err != nil {
		log.Fatalf("[ERR] Can't create localization from config due to error: %s", err.Error())
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
