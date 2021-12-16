package Slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Repo struct {
	Token string
}

func (r *Repo) SendMessage(msg *NewMessage) error {

	requestBody, err := json.Marshal(msg)

	if err != nil {
		return fmt.Errorf("couldn't marshall message %v for slack send -> %w", msg, err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://slack.com/api/chat.postMessage", bytes.NewBuffer(requestBody))

	req.Header.Set("authorization", r.Token)

	if err != nil {
		return fmt.Errorf("couldn't create request %s -> %w", req.URL.String(), err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return fmt.Errorf("couldn't send request %s -> %w", req.URL.String(), err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("couldn't read response %s body -> %w", req.URL.String(), err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response on %s had status code %v with body %s", req.URL.String(), resp.StatusCode, string(data))
	}

	return nil
}
