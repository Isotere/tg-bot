package telegram

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, errors.Wrap(err, "get updates error")
	}

	var res UpdateOut

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, errors.Wrap(err, "cannot unmashal response data")
	}

	return res.Result, nil
}
