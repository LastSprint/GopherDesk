package Trello

import (
	"crypto/sha1"
	"encoding/json"
	"github.com/LastSprint/GopherDesk/Api/Trello/Entries"
	"github.com/LastSprint/GopherDesk/Utils"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"log"
	"net/http"
)

type WebHookHandlerService interface {
	OnBoardChange(model *Entries.WebHookPayload) error
}

type Controller struct {
	Service                   WebHookHandlerService
	TrelloPayloadValidatorKey string
	TrelloCallbackUrl         string
}

func (c *Controller) Start(r chi.Router) {
	r.Post("/trello/webhook/boardChangeEvent", c.boardChangeWebHookHandler)
	r.Head("/trello/webhook/boardChangeEvent", c.methodAllowingHandler)
}

func (c *Controller) methodAllowingHandler(_ http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] Somebody %s called %s with head", r.Host, r.URL.String())
}

func (c *Controller) boardChangeWebHookHandler(_ http.ResponseWriter, r *http.Request) {

	digest := r.Header.Get("x-trello-webhook")

	if len(digest) == 0 {
		log.Printf("[WARN] somebody %s was pretending Trello (no digest in request)", r.Host)
		return
	}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("[ERR] couldn't read request %s body -> %s", r.URL.String(), err.Error())
		return
	}

	if Utils.ValidatePayload(data, c.TrelloPayloadValidatorKey, digest, c.TrelloCallbackUrl, sha1.New) {
		log.Printf("[WARN] somebody %s was pretending Trello (payload didn't pass validation)", r.Host)
		return
	}

	var payload Entries.WebHookPayload

	if err = json.Unmarshal(data, &payload); err != nil {
		log.Printf("[ERR] couldn't parse requets %s body %s due to %s", r.URL.String(), string(data), err.Error())
		return
	}

	go func(payload *Entries.WebHookPayload) {
		if err := c.Service.OnBoardChange(payload); err != nil {
			log.Printf("[ERR] Trello Payload Handler Service returned error -> %s\n", err.Error())
		}
	}(&payload)
}
