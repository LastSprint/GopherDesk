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

	req.Header.Set("Authorization", "Bearer "+r.Token)
	req.Header.Set("Content-Type", "application/json")

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

	var result struct {
		Ok      bool `json:"ok"`
		Profile User `json:"profile"`
	}

	if err = json.Unmarshal(data, &result); err != nil {
		return fmt.Errorf("response %s on %s couldn't be parsed -> %w", string(data), req.URL.String(), err)
	}

	if !result.Ok {
		return fmt.Errorf("request %s returned non success response %s", req.URL.String(), string(data))
	}

	return nil
}

func (r *Repo) GetUserByID(id string) (*User, error) {
	req, err := http.NewRequest(http.MethodPost, "https://slack.com/api/users.profile.get?user="+id, nil)

	if err != nil {
		return nil, fmt.Errorf("couldn't create request %s -> %w", req.URL.String(), err)
	}

	req.Header.Set("Authorization", "Bearer "+r.Token)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("couldn't send request %s -> %w", req.URL.String(), err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("couldn't read response %s body -> %w", req.URL.String(), err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response on %s had status code %v with body %s", req.URL.String(), resp.StatusCode, string(data))
	}

	var result struct {
		Ok      bool `json:"ok"`
		Profile User `json:"profile"`
	}

	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("response %s on %s couldn't be parsed -> %w", string(data), req.URL.String(), err)
	}

	if !result.Ok {
		return nil, fmt.Errorf("request %s returned non success response %s", req.URL.String(), string(data))
	}

	return &result.Profile, nil
}

func (r *Repo) SendResponseToUrl(responseUrl string, text string) error {
	data, err := json.Marshal(map[string]string{"text": text})

	if err != nil {
		return fmt.Errorf("can't marshall request for Slack.Repo.SendResponseToUrl -> %w", err)
	}

	response, err := http.Post(responseUrl, "application/json", bytes.NewBuffer(data))

	if err != nil {
		return fmt.Errorf("can't send message on response url -> %w", err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return fmt.Errorf("can't read SendResponseToUrl response body -> %w", err)
	}

	var res struct {
		Ok bool `json:"ok"`
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return fmt.Errorf("can't unmarshall SendResponseToUrl  response to json -> %w", err)
	}

	if !res.Ok {
		return fmt.Errorf("SendResponseToUrl request failed with response %s", string(body))
	}

	return nil
}

func (s *Repo) SendView(view *DialogView, triggerID string) error {
	requestData := map[string]interface{}{
		"trigger_id": triggerID,
		"view":       view,
	}

	requestPayload, err := json.Marshal(requestData)

	if err != nil {
		return fmt.Errorf("couldn't marshall send view request body -> %w", err)
	}

	uri := "https://slack.com/api/views.open"

	request, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(requestPayload))

	if err != nil {
		return fmt.Errorf("couldn't create request to %s -> %w", uri, err)
	}

	request.Header.Set("Authorization", "Bearer "+s.Token)
	request.Header.Set("Content-type", "application/json")

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return fmt.Errorf("couldn't send request %s -> %w", request.URL.String(), err)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("couldn't read response %s body -> %w", request.URL.String(), err)
	}

	var result struct {
		Ok      bool `json:"ok"`
		Profile User `json:"profile"`
	}

	if err = json.Unmarshal(data, &result); err != nil {
		return fmt.Errorf("response %s on %s couldn't be parsed -> %w", string(data), request.URL.String(), err)
	}

	if !result.Ok {
		return fmt.Errorf("request %s returned non success response %s", request.URL.String(), string(data))
	}

	return nil
}
