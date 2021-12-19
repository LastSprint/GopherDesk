package Entries

import (
	"encoding/json"
	"github.com/LastSprint/GopherDesk/ThirdParty/Trello"
	"time"
)

type ActionType string

const (
	ActionTypeCommentCard     = "commentCard"
	ActionTypeUpdateCard      = "updateCard"
	ActionTypeAddMemberToCard = "addMemberToCard"
)

type ActionEntry struct {
	ID              string `json:"id"`
	MemberCreatorID string `json:"idMemberCreator"`
	Type            string `json:"type"`
	Date            time.Time
	Data            json.RawMessage `json:"data"`
	MemberCreator   Trello.Member   `json:"memberCreator"`
	// Member can be set when type is kind of addMemberToCard or somehting
	Member Trello.Member `json:"member"`
}

type WebHookPayload struct {
	Action ActionEntry `json:"action"`
}
