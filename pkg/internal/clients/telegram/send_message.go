package telegram

import (
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return errors.Wrap(err, "error while message sending")
	}

	return nil
}
