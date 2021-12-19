package Services

import (
	"fmt"
	"github.com/LastSprint/GopherDesk/Services/Models"
	"github.com/LastSprint/GopherDesk/ThirdParty/Trello"
	"log"
	"strings"
)

const (
	CreatorIDKey     = "CreatorID"
	PriorityIDLow    = "619f84789e73d064c1b19aad"
	PriorityIDMedium = "619f8481db5ecb4e36c3868a"
	PriorityIDHigh   = "619f848da10bae822ff3fed5"

	TODOListID = "619f8459607bb41d8b1cbaa3"
)

type TrelloRepo interface {
	GetTicketByID(id string) (*Trello.Card, error)
	GetStatusNameByID(id string, boardID string) (string, error)
	CreateNewCard(card *Trello.NewCard) (*Trello.Card, error)
}

type TrelloService struct {
	TrelloRepo
}

func (t *TrelloService) GetTicketByID(id string) (*Models.Ticket, error) {
	ticket, err := t.TrelloRepo.GetTicketByID(id)

	if err != nil {
		return nil, err
	}

	creatorID, description, err := parseTicketInfoFromDescription(ticket.Description)

	if err != nil {
		return nil, fmt.Errorf("couldn't parse ticket %s description %s due to %w", ticket.Id, ticket.Description, err)
	}

	return &Models.Ticket{
		ID:          id,
		CreatorID:   creatorID,
		Title:       ticket.Name,
		Description: description,
		Priority:    0,
		Board:       ticket.BoardID,
	}, nil
}

func (t *TrelloService) GetStatusNameByID(id string, boardId string) (string, error) {
	return t.TrelloRepo.GetStatusNameByID(id, boardId)
}

func (t *TrelloService) CreateNewTicket(creatorID, creatorName, title, description, priority string) (string, error) {

	card, err := t.CreateNewCard(&Trello.NewCard{
		Name:        title,
		Description: createDescription(creatorID, creatorName, description),
		Position:    0,
		ListId:      TODOListID,
		LabelIds:    []string{convertPriority(priority)},
	})

	if err != nil {
		msg := fmt.Sprintf("couldn't create new card for creatorID: %s creatorName: %s title: %s description: %s priority: %s", creatorID, creatorName, title, description, priority)
		return "", fmt.Errorf("%s -> %w", msg, err)
	}

	return card.Id, nil
}

func convertPriority(entry string) string {
	switch entry {
	case "1":
		return PriorityIDLow
	case "2":
		return PriorityIDMedium
	case "3":
		return PriorityIDHigh
	}

	log.Printf("[ERR] got priority %s ad cant convert it to trello label id", entry)

	return ""
}

func parseTicketInfoFromDescription(rawDesc string) (creatorID string, description string, err error) {
	split := strings.Split(rawDesc, "\n")

	// message format:
	//
	// CreatorID: ....
	// CreatorName: ...
	// { EMPTY LINE }
	//

	if len(split) < 3 {
		return "", "", fmt.Errorf("raw description %s was corrupted", rawDesc)
	}

	if len(split[2]) != 0 {
		return "", "", fmt.Errorf("raw description %s was corrupted; Third line should be empty", rawDesc)
	}

	if creator := strings.Split(split[0], ":"); len(creator) == 2 && creator[0] == CreatorIDKey {
		creatorID = creator[1]
	} else {
		return "", "", fmt.Errorf("raw description %s was corrupted; couldn't read creator id", rawDesc)
	}

	description = strings.Join(split[3:], "\n")

	return strings.TrimSpace(creatorID), strings.TrimSpace(description), nil
}

func createDescription(creatorID, creatorName, description string) string {
	return fmt.Sprintf("CreatorID:%s\nCreatorName:%s\n\n%s", creatorID, creatorName, description)
}
