package druid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yishanhe/druid-go-client/query"
)

type Client struct {
	URL string
}

func NewClient(url string) *Client {
	return &Client{url}
}

func (c *Client) Query(query query.Query) ([]byte, error) {
	var reqBody []byte
	reqBody, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.URL, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", resp.Status, string(body))
	}

	return body, err
}
