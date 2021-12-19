package Slack

import (
	"github.com/LastSprint/GopherDesk/Services"
	"github.com/LastSprint/GopherDesk/ThirdParty/Slack"
	"github.com/LastSprint/GopherDesk/ThirdParty/Trello"
)

func AssembleTrelloController(signInSecret, slackToken, trelloToken, trelloApiKey string) *Controller {
	srv := &Service{
		SlackService: &Slack.Repo{Token: slackToken},
		TicketSystemService: &Services.TrelloService{
			TrelloRepo: &Trello.Repo{
				Token:  trelloToken,
				ApiKey: trelloApiKey,
			},
		},
	}

	return &Controller{
		SigInSecret:       signInSecret,
		CommandHandler:    srv,
		FormAnswerHandler: srv,
	}
}
