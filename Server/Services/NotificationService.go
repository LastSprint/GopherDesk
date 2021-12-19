package Services

import (
	"fmt"
	"github.com/LastSprint/GopherDesk/Services/Models"
	"github.com/LastSprint/GopherDesk/ThirdParty/Trello"
)

type NotificationTransport interface {
	OnStatusChanged(toUser *Models.User, about *Models.Ticket, newStatusName string) error
	OnAssigneeChanged(toUser *Models.User, about *Models.Ticket, newAssigneeFullName string) error
	OnCommentAdded(toUser *Models.User, about *Models.Ticket, comment string) error
}

type TicketRepo interface {
	GetTicketByID(id string) (*Models.Ticket, error)
	GetStatusNameByID(id string, boardId string) (string, error)
}

type UserRepo interface {
	GetUserByID(id string) (*Models.User, error)
}

type NotificationService struct {
	TicketRepo
	UserRepo
	NotificationTransport
}

func (n *NotificationService) getTicketAndCreator(ticketID string) (*Models.Ticket, *Models.User, error) {
	ticket, err := n.TicketRepo.GetTicketByID(ticketID)

	if err != nil {
		return nil, nil, fmt.Errorf("couldn't get ticket by id %s due to error %w", ticketID, err)
	}

	user, err := n.UserRepo.GetUserByID(ticket.CreatorID)

	if err != nil {
		return nil, nil, fmt.Errorf("couldn't get user by id %s due to error %w", ticket.CreatorID, err)
	}

	return ticket, user, nil
}

func (n *NotificationService) NotifyCustomerThatStatusChanged(ticketID, newStatusID string) error {
	ticket, user, err := n.getTicketAndCreator(ticketID)

	if err != nil {
		return fmt.Errorf("couldn't get ticket by id %s due to error %w", ticketID, err)
	}

	statusName, err := n.TicketRepo.GetStatusNameByID(newStatusID, ticket.Board)

	if err != nil {
		return fmt.Errorf("couldn't get ticket status id %s due to error %w", newStatusID, err)
	}

	return n.NotificationTransport.OnStatusChanged(user, ticket, statusName)
}

func (n *NotificationService) NotifyCustomerThatAssigneeChanged(ticketID string, newAssignee Trello.Member) error {
	ticket, user, err := n.getTicketAndCreator(ticketID)

	if err != nil {
		return fmt.Errorf("couldn't get ticket by id %s due to error %w", ticketID, err)
	}

	return n.NotificationTransport.OnAssigneeChanged(user, ticket, newAssignee.FullName)
}

func (n *NotificationService) NotifyCustomerAboutComment(ticketID, newComment string) error {
	ticket, user, err := n.getTicketAndCreator(ticketID)

	if err != nil {
		return fmt.Errorf("couldn't get ticket by id %s due to error %w", ticketID, err)
	}

	return n.NotificationTransport.OnCommentAdded(user, ticket, newComment)
}
