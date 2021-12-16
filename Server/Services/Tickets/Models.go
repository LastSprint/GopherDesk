package Tickets

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
