package Entries

import (
	"encoding/json"
	"time"
)

type ActionType string

const (
	ActionTypeCommentCard = "commentCard"
	ActionTypeDeleteCard = "deleteCard"
	ActionTypeUpdateCard = "updateCard"
	ActionTypeAddMemberToCard = "addMemberToCard"
)

type MemberEntry struct {
	ID string `json:"id"`
	ActivityBlocked bool `json:"activityBlocked"`
	AvatarHash string `json:"avatarHash"`
	AvatarUrl string `json:"avatarUrl"`
	FullName string `json:"fullName"`
	MemberReferrerID string `json:"idMemberReferrer"`
	Initials string `json:"initials"`
	Username string `json:"username"`
}

type ActionEntry struct {
	ID string `json:"id"`
	MemberCreatorID string `json:"idMemberCreator"`
	Type string `json:"type"`
	Date time.Time
	Data json.RawMessage `json:"data"`
	MemberCreator MemberEntry `json:"memberCreator"`
	// Member can be set when type is kind of addMemberToCard or somehting
	Member MemberEntry `json:"member"`
}

type WebHookPayload struct {
	Action ActionEntry `json:"action"`
}