package Trello

import (
	"encoding/json"
	"fmt"
	"github.com/LastSprint/GopherDesk/Api/Trello/Entries"
	"github.com/LastSprint/GopherDesk/ThirdParty/Trello"
	"strings"
)

type NotificationService interface {
	NotifyCustomerThatStatusChanged(cardID, newStatusID string) error
	NotifyCustomerThatAssigneeChanged(cardID string, newAssignee Trello.Member) error
	NotifyCustomerAboutComment(cardID, newComment string) error
}

type Service struct {
	NotificationService
}

func (s *Service) OnBoardChange(model *Entries.WebHookPayload) error {

	switch model.Action.Type {
	case Entries.ActionTypeCommentCard:
		var entr Entries.ActionDataOnCommentCardEntry
		if err := json.Unmarshal(model.Action.Data, &entr); err != nil {
			return fmt.Errorf("parse action data %s for type %s failed -> %w", model.Action.Type, string(model.Action.Data), err)
		}
		return s.onCommentCard(&entr)
	case Entries.ActionTypeAddMemberToCard:
		var entr Entries.ActionDataOnAddMemberToCardEntry
		if err := json.Unmarshal(model.Action.Data, &entr); err != nil {
			return fmt.Errorf("parse action data %s for type %s failed -> %w", model.Action.Type, string(model.Action.Data), err)
		}
		return s.onAddMemberToCard(model, &entr)
	case Entries.ActionTypeUpdateCard:
		var entr Entries.ActionDataOnUpdateCardEntry
		if err := json.Unmarshal(model.Action.Data, &entr); err != nil {
			return fmt.Errorf("parse action data %s for type %s failed -> %w", model.Action.Type, string(model.Action.Data), err)
		}
		return s.onUpdateCard(&entr)
	}

	return nil
}

func (s *Service) onUpdateCard(model *Entries.ActionDataOnUpdateCardEntry) error {
	if len(model.ListAfter.ID) == 0 {
		return fmt.Errorf("on Update Card. list after is nill. Status couldn't be changed to nil")
	}

	return s.NotificationService.NotifyCustomerThatStatusChanged(model.Card.ID, model.ListAfter.ID)
}

func (s *Service) onCommentCard(model *Entries.ActionDataOnCommentCardEntry) error {

	const trigger = "USER_MSG"

	if !strings.Contains(model.Text, trigger) {
		return nil
	}

	userMsg := strings.ReplaceAll(model.Text, trigger, "")

	if len(userMsg) == 0 {
		return fmt.Errorf("msg %s became empty after replacing trigger %s", model.Text, trigger)
	}

	return s.NotificationService.NotifyCustomerAboutComment(model.Card.ID, userMsg)
}

func (s *Service) onAddMemberToCard(payload *Entries.WebHookPayload, data *Entries.ActionDataOnAddMemberToCardEntry) error {
	return s.NotificationService.NotifyCustomerThatAssigneeChanged(data.Card.ID, payload.Action.Member)
}
