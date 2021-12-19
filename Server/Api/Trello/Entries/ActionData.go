package Entries

type CardEntry struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IDShort   int    `json:"idShort"`
	ShortLink string `json:"shortLink"`
}

type BoardEntry struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortLink string `json:"shortLink"`
}

type ListEntry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ActionDataOnCommentCardEntry struct {
	Text string    `json:"text"`
	Card CardEntry `json:"card"`
	List ListEntry `json:"list"`
}

type ActionDataOnAddMemberToCardEntry struct {
	MemberID string    `json:"idMember"`
	Card     CardEntry `json:"card"`
	List     ListEntry `json:"list"`
}

type ActionDataOnUpdateCardEntry struct {
	Card       CardEntry `json:"card"`
	ListBefore ListEntry `json:"listBefore"`
	ListAfter  ListEntry `json:"listAfter"`
}
