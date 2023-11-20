package telegram

import (
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

const (
	basePathPrefix = "bot"

	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "can't prepare request")
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't do request")
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error while response body read")
	}

	return body, nil
}

func newBasePath(token string) string {
	return basePathPrefix + token
}
