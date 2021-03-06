package Slack

import (
	"fmt"
	"github.com/LastSprint/GopherDesk/L10n"
	"github.com/LastSprint/GopherDesk/ThirdParty/Slack"
	"log"
)

const (
	ticketFormCallBackID              = "new-sa-ticket"
	ticketFormTitleBlockID            = "title"
	ticketBlockTitleElementActionId   = "message"
	ticketFormDescriptionBlockID      = "description"
	ticketFormPriorityBlockID         = "priority"
	ticketFormPriorityElementActionID = "type"

	TicketPriorityValueLow    = "1"
	TicketPriorityValueMedium = "2"
	TicketPriorityValueHigh   = "3"
)

type SlackService interface {
	SendResponseToUrl(responseUrl string, text string) error
	SendView(view *Slack.DialogView, triggerID string) error
	SendMessage(msg *Slack.NewMessage) error
}

type TicketSystemService interface {
	CreateNewTicket(creatorID, creatorName, title, description, priority string) (string, error)
}

type Service struct {
	SlackService
	TicketSystemService
}

func (s *Service) HandleError(command *SlashCommand, err error) error {
	text := fmt.Sprintf("An error iccured while processing command. \nPlease contact one of system administarator with error message:\n%s", err.Error())
	return s.SlackService.SendResponseToUrl(command.ResponseUrl, text)
}

func (s *Service) HandleCommand(command *SlashCommand) error {
	switch command.Command {
	case "/ticket":
		if err := s.SlackService.SendView(defaultDialog(), command.TriggerId); err != nil {
			return fmt.Errorf("coouldn't send slack view for command %v due to -> %w", command, err)
		}
	default:
		return fmt.Errorf("command %s isn't supported", command.Command)
	}

	return nil
}

func (s *Service) HandleForm(form *FormPayload) error {
	ticketId, err := s.TicketSystemService.CreateNewTicket(
		form.User.ID,
		form.User.Name,
		form.View.State.Values.Title.Message.Value,
		form.View.State.Values.Description.Message.Value,
		form.View.State.Values.Priority.Type.Option.Value,
	)

	if err != nil {
		return fmt.Errorf("couldn't create ticket due to %w", err)
	}

	go func(userId, ticketId string) {
		err := s.SlackService.SendMessage(&Slack.NewMessage{
			Channel:  userId,
			Markdown: true,
			Text:     fmt.Sprintf(":white_check_mark: ?????? ?????????????? ?????????? ?? ID `%s`", ticketId),
			Username: "Steve",
		})

		if err != nil {
			log.Printf("[ERR] Couldn't send ACK about creating ticket %s to user %s due to error %s", ticketId, userId, err.Error())
			return
		}
	}(form.User.ID, ticketId)

	return nil
}

func defaultDialog() *Slack.DialogView {
	return &Slack.DialogView{
		CallbackID: ticketFormCallBackID,
		Type:       Slack.DialogViewTypeModal,
		Submit: Slack.ViewLabel{
			Type:  Slack.ViewLabelTypePlainText,
			Text:  L10n.Print.Sprintf(L10n.SDFormSendKey),
			Emoji: true,
		},
		Close: Slack.ViewLabel{
			Type:  Slack.ViewLabelTypePlainText,
			Text:  L10n.Print.Sprintf(L10n.SDFormCancelKey),
			Emoji: true,
		},
		Title: Slack.ViewLabel{
			Type:  Slack.ViewLabelTypePlainText,
			Text:  L10n.Print.Sprintf(L10n.SDFormHeadTitleKey),
			Emoji: true,
		},
		Blocks: []Slack.BlockItem{
			{
				ID:   ticketFormTitleBlockID,
				Type: Slack.BlockItemTypeInput,
				Label: Slack.ViewLabel{
					Type: Slack.ViewLabelTypePlainText,
					Text: L10n.Print.Sprintf(L10n.SDFormFieldTitleNameKey),
				},
				Element: Slack.BlockElement{
					Type:        Slack.BlockElementTypePlainTextInput,
					ActionID:    ticketBlockTitleElementActionId,
					IsMultiline: false,
				},
			},
			{
				ID:   ticketFormPriorityBlockID,
				Type: Slack.BlockItemTypeInput,
				Label: Slack.ViewLabel{
					Type: Slack.ViewLabelTypePlainText,
					Text: L10n.Print.Sprintf(L10n.SDFormFieldPriorityKey),
				},
				Element: Slack.BlockElement{
					Type:     Slack.BlockElementTypeStaticSelect,
					ActionID: ticketFormPriorityElementActionID,
					Options: []Slack.BlockElementOption{
						{
							Text: Slack.ViewLabel{
								Type:  Slack.ViewLabelTypePlainText,
								Text:  L10n.Print.Sprintf(L10n.SDFormFieldPriorityLowKey),
								Emoji: false,
							},
							Value: TicketPriorityValueLow,
						},
						{
							Text: Slack.ViewLabel{
								Type:  Slack.ViewLabelTypePlainText,
								Text:  L10n.Print.Sprintf(L10n.SDFormFieldPriorityMediumKey),
								Emoji: false,
							},
							Value: TicketPriorityValueMedium,
						},
						{
							Text: Slack.ViewLabel{
								Type:  Slack.ViewLabelTypePlainText,
								Text:  L10n.Print.Sprintf(L10n.SDFormFieldPriorityHighKey),
								Emoji: false,
							},
							Value: TicketPriorityValueHigh,
						},
					},
				},
			},
			{
				ID:   ticketFormDescriptionBlockID,
				Type: Slack.BlockItemTypeInput,
				Label: Slack.ViewLabel{
					Type: Slack.ViewLabelTypePlainText,
					Text: L10n.Print.Sprintf(L10n.SDFormFieldDescriptionKey),
				},
				Element: Slack.BlockElement{
					Type:        Slack.BlockElementTypePlainTextInput,
					ActionID:    ticketBlockTitleElementActionId,
					IsMultiline: true,
				},
			},
		},
	}
}
