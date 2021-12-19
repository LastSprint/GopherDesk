package Services

import (
	"fmt"
	"github.com/LastSprint/GopherDesk/Services/Models"
	"github.com/LastSprint/GopherDesk/ThirdParty/Slack"
)

type SlackRepo interface {
	SendMessage(msg *Slack.NewMessage) error
	GetUserByID(id string) (*Slack.User, error)
}

type SlackService struct {
	Slack SlackRepo
}

func (s *SlackService) OnStatusChanged(toUser *Models.User, about *Models.Ticket, newStatusName string) error {

	statusString := newStatusName
	title := fmt.Sprintf("Тикет \"%s\"", about.Title)

	switch newStatusName {
	case "TODO":
		statusString = "`To Do`"
		title = "→ :waiting: " + title
	case "IN PROGRESS":
		statusString = "`In Progress`"
		title = "→ :construction: " + title
	case "BLOCKED":
		statusString = "`Blocked`"
		title = "→ :no_entry: " + title
	case "DONE":
		statusString = "`Done`"
		title = "→ :white_check_mark: " + title
	}

	msg := &Slack.NewMessage{
		Channel: toUser.ID,
		Blocks: []interface{}{
			map[string]interface{}{
				"type": "header",
				"text": map[string]string{
					"type": "plain_text",
					"text": title,
				},
			},
			map[string]interface{}{
				"type": "divider",
			},
			map[string]interface{}{
				"type": "section",
				"text": map[string]string{
					"type": "mrkdwn",
					"text": fmt.Sprintf("Статус изменен на %s", statusString),
				},
			},
			map[string]interface{}{
				"type": "divider",
			},
			map[string]interface{}{
				"type": "context",
				"elements": []map[string]string{
					{
						"type": "mrkdwn",
						"text": fmt.Sprintf("ID тикета: %s \nОписание:\n%s", about.ID, about.Description),
					},
				},
			},
		},
	}

	return s.Slack.SendMessage(msg)
}

func (s *SlackService) OnAssigneeChanged(toUser *Models.User, about *Models.Ticket, newAssigneeFullName string) error {
	msg := &Slack.NewMessage{
		Channel: toUser.ID,
		Blocks: []interface{}{
			map[string]interface{}{
				"type": "header",
				"text": map[string]string{
					"type": "plain_text",
					"text": fmt.Sprintf(":construction_worker: Тикет \"%s\"", about.Title),
				},
			},
			map[string]interface{}{
				"type": "divider",
			},
			map[string]interface{}{
				"type": "section",
				"text": map[string]string{
					"type": "mrkdwn",
					"text": fmt.Sprintf("Исполнителем назначен `%s`", newAssigneeFullName),
				},
			},
			map[string]interface{}{
				"type": "divider",
			},
			map[string]interface{}{
				"type": "context",
				"elements": []map[string]string{
					{
						"type": "mrkdwn",
						"text": fmt.Sprintf("ID тикета: %s \nОписание:\n%s", about.ID, about.Description),
					},
				},
			},
		},
	}

	return s.Slack.SendMessage(msg)
}

func (s *SlackService) OnCommentAdded(toUser *Models.User, about *Models.Ticket, comment string) error {
	msg := &Slack.NewMessage{
		Channel: toUser.ID,
		Blocks: []interface{}{
			map[string]interface{}{
				"type": "header",
				"text": map[string]string{
					"type": "plain_text",
					"text": fmt.Sprintf(":speech_balloon: Тикет \"%s\"", about.Title),
				},
			},
			map[string]interface{}{
				"type": "divider",
			},
			map[string]interface{}{
				"type": "section",
				"text": map[string]string{
					"type": "mrkdwn",
					"text": fmt.Sprintf("Новый комментарий:\n\n%s", comment),
				},
			},
			map[string]interface{}{
				"type": "divider",
			},
			map[string]interface{}{
				"type": "context",
				"elements": []map[string]string{
					{
						"type": "mrkdwn",
						"text": fmt.Sprintf("ID тикета: %s \nОписание:\n%s", about.ID, about.Description),
					},
				},
			},
		},
	}

	return s.Slack.SendMessage(msg)
}

func (s *SlackService) GetUserByID(id string) (*Models.User, error) {
	user, err := s.Slack.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	return &Models.User{
		FullName: user.Name,
		Email:    user.Email,
		ID:       id,
	}, nil
}
