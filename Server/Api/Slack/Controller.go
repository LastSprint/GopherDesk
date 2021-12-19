package Slack

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/LastSprint/GopherDesk/Utils"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type CommandHandler interface {
	HandleError(command *SlashCommand, err error) error
	HandleCommand(command *SlashCommand) error
}

type FormAnswerHandler interface {
	HandleForm(form *FormPayload) error
}

type Controller struct {
	SigInSecret string
	CommandHandler
	FormAnswerHandler
}

func (c *Controller) Start(r chi.Router) {
	r.Route("/slack", func(r chi.Router) {
		r.Post("/show-dialog-command", c.showDialogCommandHandler)
		r.Post("/webhooks/dialog-sent-handler", c.dialogSentHandler)
	})
}

func (c *Controller) showDialogCommandHandler(w http.ResponseWriter, r *http.Request) {
	signature := r.Header.Get("X-Slack-Signature")

	if len(signature) == 0 {
		log.Println("[WARN] Somebody pretending Slack. No signature in request", r.URL.String())
		http.Error(w, "Internal error #1", http.StatusInternalServerError)
		return
	}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("[ERR] couldn't read request %s body -> %s", r.URL.String(), err.Error())
		http.Error(w, "Internal error #2", http.StatusInternalServerError)
		return
	}

	if Utils.ValidatePayload(data, c.SigInSecret, signature, c.SigInSecret, sha256.New) {
		log.Printf("[WARN] somebody %s was pretending Slack (payload didn't pass validation)", r.Host)
		http.Error(w, "Internal error #3", http.StatusInternalServerError)
		return
	}

	form, err := url.ParseQuery(string(data))

	if err != nil {
		log.Printf("[ERR] request %s can't parse body %s due to error %s", r.URL.String(), string(data), err.Error())
		http.Error(w, "Internal error #3", http.StatusInternalServerError)
		return
	}

	command := SlashCommand{
		ApiAppID:    form.Get("api_app_id"),
		ChannelID:   form.Get("channel_id"),
		ChannelName: form.Get("channel_name"),
		Command:     form.Get("command"),
		ResponseUrl: form.Get("response_url"),
		TeamDomain:  form.Get("team_domain"),
		TeamId:      form.Get("team_id"),
		Text:        form.Get("text"),
		Token:       form.Get("token"),
		TriggerId:   form.Get("trigger_id"),
		UserID:      form.Get("user_id"),
		UserName:    form.Get("user_name"),
	}

	go func(cmd *SlashCommand) {
		if err = c.CommandHandler.HandleCommand(cmd); err != nil {
			log.Printf("[ERR] command handler failed -> %s", err.Error())
			if err = c.CommandHandler.HandleError(cmd, err); err != nil {
				log.Printf("[ERR] error handler failed -> %s", err.Error())
			}
			return
		}
	}(&command)
}

func (c *Controller) dialogSentHandler(w http.ResponseWriter, r *http.Request) {

	signature := r.Header.Get("X-Slack-Signature")

	if len(signature) == 0 {
		log.Println("[WARN] Somebody pretending Slack. No signature in request", r.URL.String())
		http.Error(w, "Internal error #10", http.StatusInternalServerError)
		return
	}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("[ERR] couldn't read request %s body -> %s", r.URL.String(), err.Error())
		http.Error(w, "Internal error #20", http.StatusInternalServerError)
		return
	}
	values, err := url.ParseQuery(string(data))

	if err != nil {
		log.Printf("[ERR] request %s can't parse body %s due to error %s", r.URL.String(), string(data), err.Error())
		http.Error(w, "Internal error #30", http.StatusInternalServerError)
		return
	}

	payload := values.Get("payload")

	if len(payload) == 0 {
		log.Printf("[ERR] request %s empty payload in body %s", r.URL.String(), string(data))
		http.Error(w, "Internal error #40", http.StatusInternalServerError)
		return
	}

	var form FormPayload

	if err = json.Unmarshal([]byte(payload), &form); err != nil {
		log.Printf("[ERR] couldn't unmarshall %s payload %s to json -> %s", r.URL.String(), payload, err.Error())
		http.Error(w, "Internal error #40", http.StatusInternalServerError)
	}

	if err = c.FormAnswerHandler.HandleForm(&form); err != nil {
		log.Printf("[ERR] handling form with payload %s failed -> %s", payload, err.Error())
		encodeFormError(w, "We couldn't handle the form. Please contact with SA", "title")
	}
}

func encodeFormError(w http.ResponseWriter, text, key string) {

	w.Header().Add("Content-Type", "application/json")

	object := struct {
		ResponseAction string            `json:"response_action"`
		Errors         map[string]string `json:"errors"`
	}{
		ResponseAction: "errors",
		Errors: map[string]string{
			key: text,
		},
	}

	if err := json.NewEncoder(w).Encode(object); err != nil {
		log.Println("[ERR] couldn't encode error object ->", err.Error())
	}
}
