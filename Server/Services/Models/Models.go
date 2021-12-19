package Models

import "time"

type TicketPriority int
type TicketBoard string

const (
	TicketPriorityLow    = 0
	TicketPriorityMedium = 1
	TicketPriorityHeight = 2

	TicketBoardSA = "SA"
)

type CreateTicketEntity struct {
	Title       string
	Description string
	Priority    TicketPriority
	SenderId    string
	Board       TicketBoard
}

type Ticket struct {
	ID               string
	CreatorID        string
	Title            string
	Description      string
	Priority         TicketPriority
	Board            string
	LastActivityDate time.Time
}

type NotificationEvent string

const (
	NotificationEventTicketStatusChanged = "StatusChanged"
	NotificationEventAssigneeChanged     = "AssigneeChanged"
	NotificationEventCommentAdded        = "CommentAdded"
)
