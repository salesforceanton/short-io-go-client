package shortener

import (
	"errors"
	"net/http"
	"time"

	"github.com/salesforceanton/short-io-go-client/config"
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
		return nil, errors.New("Error with config initialization")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			// Add auth middleware
		},
		config: cfg,
	}, nil
}

func (c *Client) shortenLink(link string) (ShortenLinkResponseItem, error) {
	// url := "https://api.short.io/links"

	// payload := strings.NewReader("{\"domain\":\"ah0e.short.gy\",\"originalURL\":\"http://valeron.macron@gmail.com\"}")

	// req, _ := http.NewRequest("POST", url, payload)

	// req.Header.Add("accept", "application/json")
	// req.Header.Add("content-type", "application/json")
	// req.Header.Add("Authorization", "sk_ipjnYDD1sMoF9rkj")

	// res, _ := http.DefaultClient.Do(req)

	// defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)

	// fmt.Println(string(body))

}

func (c *Client) shortenLinks(links []string) ([]ShortenLinkResponseItem, error) {

}
