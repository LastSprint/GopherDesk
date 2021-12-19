package Slack

type SlashCommand struct {
	ApiAppID    string `url:"api_app_id"`
	ChannelID   string `url:"channel_id"`
	ChannelName string `url:"channel_name"`
	Command     string `url:"command"`
	ResponseUrl string `url:"response_url"`
	TeamDomain  string `url:"team_domain"`
	TeamId      string `url:"team_id"`
	Text        string `url:"text"`
	Token       string `url:"token"`
	TriggerId   string `url:"trigger_id"`
	UserID      string `url:"user_id"`
	UserName    string `url:"user_name"`
}

type MessageState struct {
	Message struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"message"`
}

type SelectorState struct {
	Type struct {
		Type   string `json:"type"`
		Option struct {
			Value string `json:"value"`
		} `json:"selected_option"`
	} `json:"type"`
}

type TicketState struct {
	Title       MessageState  `json:"title"`
	Priority    SelectorState `json:"priority"`
	Description MessageState  `json:"description"`
}

type FormPayload struct {
	Type string `json:"type"`
	User struct {
		ID       string `json:"id"`
		UserName string `json:"username"`
		Name     string `json:"name"`
	} `json:"user"`
	ApiAppID  string `json:"api_app_id"`
	Token     string `json:"token"`
	TriggerID string `json:"trigger_id"`
	View      struct {
		ID         string `json:"id"`
		TeamID     string `json:"team_id"`
		Type       string `json:"type"`
		CallbackId string `json:"callback_id"`
		State      struct {
			Values TicketState `json:"values"`
		} `json:"state"`
	}
}
