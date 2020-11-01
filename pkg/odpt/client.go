package odpt

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const baseUrl = "https://api-tokyochallenge.odpt.org/api/v4/"

type Client struct {
	token      string
	httpClient *http.Client
}

func NewClient(token string) *Client {
	if token == "" {
		panic("token required")
	}

	return &Client{
		token: token,
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (c *Client) buildURL(resource string, query map[string]string) url.URL {
	u := url.URL{
		Scheme: "https",
		Host:   "api-tokyochallenge.odpt.org",
		Path:   fmt.Sprintf("api/v4/%s", resource),
	}

	q := u.Query()

	q.Add("acl:consumerKey", c.token)

	if query != nil {
		for k, v := range query {
			q.Add(k, v)
		}
	}

	u.RawQuery = q.Encode()

	return u
}

func (c *Client) getBuses(ctx context.Context, query map[string]string) ([]*Bus, error) {
	u := c.buildURL("odpt:Bus", query)

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("ec.httpClient.Do: %w", err)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	bodyRaw, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("reading body failed: %w", err)
	}

	buses := make([]*Bus, 0)

	err = json.Unmarshal(bodyRaw, &buses)

	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return buses, nil
}

func (c *Client) GetAllBuses(ctx context.Context) ([]*Bus, error) {
	return c.getBuses(ctx, nil)
}

func (c *Client) GetBusesForRoute(ctx context.Context, route string) ([]*Bus, error) {
	return c.getBuses(ctx, map[string]string{
		"odpt:busroute": route,
	})
}
