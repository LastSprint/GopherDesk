package Trello

import (
	"github.com/LastSprint/GopherDesk/Services"
	"github.com/LastSprint/GopherDesk/ThirdParty/Slack"
	"github.com/LastSprint/GopherDesk/ThirdParty/Trello"
)

func AssembleTrelloController(slackToken, trelloApiKey, trelloToken string, fullCallbackUrl string) *Controller {

	slackService := &Services.SlackService{
		Slack: &Slack.Repo{
			Token: slackToken,
		},
	}

	return &Controller{
		Service: &Service{
			NotificationService: &Services.NotificationService{
				TicketRepo: &Services.TrelloService{
					TrelloRepo: &Trello.Repo{
						Token:  trelloToken,
						ApiKey: trelloApiKey,
					},
				},
				UserRepo:              slackService,
				NotificationTransport: slackService,
			},
		},
		TrelloPayloadValidatorKey: trelloApiKey,
		TrelloCallbackUrl:         fullCallbackUrl,
	}
}
