package L10n

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var Print = message.NewPrinter(language.English)

var (
	SDFormSendKey      = message.Key("SDFormSendKey", "")
	SDFormCancelKey    = message.Key("SDFormCancelKey", "")
	SDFormHeadTitleKey = message.Key("SDFormHeadTitleKey", "")

	SDFormFieldTitleNameKey      = message.Key("SDFormFieldTitleNameKey", "")
	SDFormFieldPriorityKey       = message.Key("SDFormFieldPriorityKey", "")
	SDFormFieldPriorityLowKey    = message.Key("SDFormFieldPriorityLowKey", "")
	SDFormFieldPriorityMediumKey = message.Key("SDFormFieldPriorityMediumKey", "")
	SDFormFieldPriorityHighKey   = message.Key("SDFormFieldPriorityHighKey", "")

	SDFormFieldDescriptionKey = message.Key("SDFormFieldDescriptionKey", "")
)
