package Tickets

import (
	slack "github.com/LastSprint/GopherDesk/ThirdParty/Slack"
	trello "github.com/LastSprint/GopherDesk/ThirdParty/Trello"
)

type Trello interface {
	CreateNewCard(card *trello.NewCard) (*trello.Card, error)
}

type Slack interface {
	SendMessage(msg *slack.NewMessage) error
	GetUserById(id string) (*slack.User, error)
}

type Service struct {
	SlackRepo  Slack
	TrelloRepo Trello
}

func (s *Service) CreateNewTicket(ticket *CreateTicketEntity) error {
	return nil
	//user, err := s.SlackRepo.GetUserById(ticket.SenderId)
	//
	//if err != nil {
	//	return fmt.Errorf("couldn't get user by id -> %w", err)
	//}
	//
	//card := trello.NewCard{
	//	Name:        ticket.Title,
	//	Description: ticket.Description,
	//	Position:    0,
	//	ListId:      "",
	//	LabelIds:    ,
	//}
	//
	//s.TrelloRepo.CreateNewCard()
}

//func convertTikcetPriorityToListId(priority int) string {}
