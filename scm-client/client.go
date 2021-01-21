package scm_client

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	config Config
}

type Config struct {
	URL      string
	Username string
	Password string
}

func NewClient(config Config) *Client {
	return &Client{config}
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.config.Username, c.config.Password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to do request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read body")
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("response had statuscode: %d", resp.StatusCode)
	}

	return body, err
}
