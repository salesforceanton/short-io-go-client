package shortener

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/salesforceanton/short-io-go-client/config"
)

const (
	BASE_ENDPOINT       = "https://api.short.io"
	LINKS_ENDPOINT      = "links"
	CONTENT_TYPE        = "application/json"
	AUTH_HEADER         = "Authorization"
	ACCEPT_HEADER       = "Accept"
	CONTENT_TYPE_HEADER = "Content-type"
)

type Client struct {
	client *http.Client
	config *config.Config
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("Timeout for Client callouts cannot be equals 0")
	}

	cfg, err := config.New()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error with config initialization: %s", err.Error()))
	}

	return &Client{
		client: &http.Client{
			Timeout:   timeout,
			Transport: NewAuthRoundTripper(cfg.Token),
		},
		config: cfg,
	}, nil
}

func (c *Client) ShortenLink(link string) (string, error) {
	// Prepare Http request
	requestBody, err := json.Marshal(c.mapUrlToRequest(link))
	if err != nil {
		return "", errors.New("Request in incorrect type")
	}

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", BASE_ENDPOINT, LINKS_ENDPOINT),
		bytes.NewBuffer(requestBody),
	)

	req.Header.Add(ACCEPT_HEADER, CONTENT_TYPE)
	req.Header.Add(CONTENT_TYPE_HEADER, CONTENT_TYPE)

	// Perform http callout and process result
	res, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var result ShortenLinkResponseItem
	if err = json.Unmarshal(body, &result); err != nil {
		return "", errors.New("Error with parsing response body")
	}

	// To make a clear result for this method return only shortened link without any other info
	return result.ShortURL, nil
}

func (c *Client) ShortenLinks(links []string) ([]ShortenLinkResponseItem, error) {
	// Prepare Http request
	requestBody, err := json.Marshal(c.mapUrlsToBulkRequest(links))
	if err != nil {
		return nil, errors.New("Request in incorrect type")
	}

	req, _ := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/%s", BASE_ENDPOINT, LINKS_ENDPOINT),
		bytes.NewBuffer(requestBody),
	)

	req.Header.Add(ACCEPT_HEADER, CONTENT_TYPE)
	req.Header.Add(CONTENT_TYPE_HEADER, CONTENT_TYPE)

	// Perform http callout and process result
	res, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var result []ShortenLinkResponseItem
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("Error with parsing response body")
	}

	return result, nil
}
