package Trello

import "time"

const (
	LabelIdLow    = "619f84789e73d064c1b19aad"
	LabelIdMedium = "619f8481db5ecb4e36c3868a"
	LabelIdHeight = "619f420e2dd3d7ffdb1d7e2d"

	ListIdToDo = "619f8459607bb41d8b1cbaa3"
)

type NewCard struct {
	Name        string   `url:"name" json:"name"`
	Description string   `url:"desc" json:"desc"`
	Position    float32  `url:"pos" json:"pos"`
	ListId      string   `url:"idList" json:"idList"`
	LabelIds    []string `url:"idLabels" json:"idLabels"`
}

type Card struct {
	NewCard
	Id               string    `json:"id"`
	Closed           bool      `json:"closed"`
	CreationMethod   string    `json:"creationMethod"`
	DateLastActivity time.Time `json:"dateLastActivity"` // date-time
	Labels           []string  `json:"labels"`
	ShortLink        string    `json:"shortLink"`
}
