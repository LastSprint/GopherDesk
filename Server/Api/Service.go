package Api

import (
	slack "github.com/LastSprint/GopherDesk/ThirdParty/Slack"
	trello "github.com/LastSprint/GopherDesk/ThirdParty/Trello"
)

type Trello interface {
	CreateNewCard(card *trello.NewCard) (*trello.Card, error)
}

type Slack interface {
	SendMessage(msg *slack.NewMessage) error
}

type Service struct {
	SlackRepo  Slack
	TrelloRepo Trello
}

func (s *Service) CreateNewTicket() {

}
