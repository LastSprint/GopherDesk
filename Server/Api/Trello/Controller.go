package Trello

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"log"
	"net/http"
)

type Controller struct {
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

	if validatePayload(data, c.TrelloPayloadValidatorKey, digest, c.TrelloCallbackUrl) {
		log.Printf("[WARN] somebody %s was pretending Trello (payload didn't pass validation)", r.Host)
		return
	}

	log.Println("[INFO] the body:")
	log.Println(string(data))
}

func validatePayload(payload []byte, secret, digest, url string) bool {
	str := string(payload) + url

	var encoded []byte

	base64.StdEncoding.Encode([]byte(str), encoded)

	crp := hmac.New(sha1.New, []byte(secret))

	crp.Write(encoded)

	result := string(crp.Sum(nil))

	return result == digest
}
