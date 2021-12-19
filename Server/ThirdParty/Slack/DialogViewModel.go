package Slack

type DialogViewType string
type ViewLabelType string
type BlockItemType string
type BlockElementType string

const (
	DialogViewTypeModal = "modal"

	ViewLabelTypePlainText = "plain_text"

	BlockItemTypeInput = "input"

	BlockElementTypePlainTextInput = "plain_text_input"
	BlockElementTypeStaticSelect   = "static_select"
)

type DialogView struct {
	CallbackID string         `json:"callback_id"`
	Type       DialogViewType `json:"type"`
	Submit     ViewLabel      `json:"submit,omitempty"`
	Close      ViewLabel      `json:"close,omitempty"`
	Title      ViewLabel      `json:"title,omitempty"`
	Blocks     []BlockItem    `json:"blocks,omitempty"`
}

type ViewLabel struct {
	Type  ViewLabelType `json:"type,omitempty"`
	Text  string        `json:"text,omitempty"`
	Emoji bool          `json:"emoji,omitempty"`
}

type BlockItem struct {
	ID      string        `json:"block_id"`
	Type    BlockItemType `json:"type"`
	Label   ViewLabel     `json:"label"`
	Element BlockElement  `json:"element"`
}

type BlockElementOption struct {
	// Text is value which user will choose
	Text ViewLabel `json:"text"`
	// Value is the value which will be sent as selected (it's a key for Text)
	Value string `json:"value"`
}

type BlockElement struct {
	Type     BlockElementType `json:"type"`
	ActionID string           `json:"action_id,omitempty"`
	// Options is used only for select* elements (like BlockElementTypeStaticSelect)
	Options []BlockElementOption `json:"options,omitempty"`
	// IsMultiline could be used only for text input elements (like BlockElementTypePlainTextInput)
	IsMultiline bool `json:"multiline,omitempty"`
}
