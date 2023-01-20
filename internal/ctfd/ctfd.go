package ctfd

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrBadAuth = errors.New("ctfd: bad auth")

type Client struct {
	c       *http.Client
	urlBase string
	token   string
}

type Challenge struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// there are more fields, but we only care about id/name for now
}

func New(urlBase, token string) *Client {
	c := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	return &Client{
		c:       c,
		urlBase: urlBase,
		token:   token,
	}
}

func (c *Client) GetChallanges() ([]*Challenge, error) {
	req, _ := http.NewRequest(http.MethodGet, c.urlBase+"/api/v1/challenges", http.NoBody)
	req.Header.Add("Authorization", "Token "+c.token)
	req.Header.Add("Content-Type", "application/json") // should be accept, but their api is wierd

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 {
		return nil, ErrBadAuth
	}

	var bruh struct {
		Success bool
		Data    []*Challenge
	}
	err = json.NewDecoder(resp.Body).Decode(&bruh)
	if err != nil {
		return nil, err
	}

	return bruh.Data, nil
}
