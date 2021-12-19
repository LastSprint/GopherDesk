package Slack

type NewMessage struct {
	Channel   string        `json:"channel"`
	Blocks    []interface{} `json:"blocks"`
	IconEmoji string        `json:"icon_emoji"`
	IconUrl   string        `json:"icon_url"`
	LinkNames bool          `json:"link_names"`
	Markdown  bool          `json:"mrkdwn"`
	Text      string        `json:"text"`
	ThreadTs  string        `json:"thread_ts"`
	Username  string        `json:"username"`
}

type User struct {
	Name  string `json:"real_name"`
	Email string `json:"email"`
}
