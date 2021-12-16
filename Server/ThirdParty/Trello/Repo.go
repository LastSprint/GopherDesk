package Trello

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Repo struct {
	Token string
}

func (r *Repo) CreateNewCard(card *NewCard) (*Card, error) {

	const urlString = "https://api.trello.com/1/cards"

	uri, err := url.Parse(urlString)

	values, err := query.Values(card)

	if err != nil {
		return nil, fmt.Errorf("couldn't convert card model %v to url-query string", card)
	}

	if err != nil {
		return nil, fmt.Errorf("couldn't create url from string %s -> %w", urlString, err)
	}

	uri.RawQuery = values.Encode()

	request, err := http.NewRequest(http.MethodPost, uri.String(), nil)

	request.Header.Set("authorization", r.Token)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, fmt.Errorf("couldn't sent request %s -> %w", request.URL.String(), err)
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("couldn't read response %s body -> %w", request.URL.String(), err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response on %s had status code %v with body %s", request.URL.String(), response.StatusCode, string(data))
	}

	resp := &Card{}

	if err = json.Unmarshal(data, resp); err != nil {
		return nil, fmt.Errorf("couldn't parse response %s on request %s -> %w", string(data), request.URL.String(), err)
	}

	return resp, nil
}
