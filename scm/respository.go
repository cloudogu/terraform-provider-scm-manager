package scm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type Repository struct {
	NameSpace    string `json:"namespace"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	Contact      string `json:"contact"`
	CreationDate string `json:"creationDate"`
	ImportUrl    string `json:"importUrl"`
	LastModified string `json:"lastModified"`
}

func (r *Repository) GetID() string {
	return fmt.Sprintf("%s/%s", r.NameSpace, r.Name)
}

func (c *Client) CreateRepository(ctx context.Context, repo Repository) error {
	url := fmt.Sprintf("%s/api/v2/repositories", c.config.URL)
	return c.setRepository(ctx, repo, "POST", url)
}

func (c *Client) GetRepository(ctx context.Context, name string) (Repository, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.config.URL+"/api/v2/repositories/"+name, nil)
	if err != nil {
		return Repository{}, errors.Wrap(err, "failed to create new request")
	}

	body, err := c.doRequest(req)
	if err != nil {
		return Repository{}, err
	}

	repo := &Repository{}
	err = json.Unmarshal(body, repo)
	if err != nil {
		return Repository{}, errors.Wrap(err, "failed to unmarshal repository")
	}

	return *repo, nil
}

func (c *Client) UpdateRepository(ctx context.Context, name string, repo Repository) error {
	url := fmt.Sprintf("%s/api/v2/repositories/%s", c.config.URL, name)
	return c.setRepository(ctx, repo, "PUT", url)
}

func (c *Client) DeleteRepository(ctx context.Context, name string) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", c.config.URL+"/api/v2/repositories/"+name, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create new request")
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ImportRepository(ctx context.Context, repo Repository) error {
	url := fmt.Sprintf("%s/api/v2/repositories/import/%s/url", c.config.URL, repo.Type)
	return c.setRepository(ctx, repo, "POST", url)
}

func (c *Client) setRepository(ctx context.Context, repo Repository, method string, url string) error {
	b, err := json.Marshal(&repo)
	if err != nil {
		return errors.Wrap(err, "failed to marshal repository")
	}

	buffer := bytes.NewBuffer(b)

	req, err := http.NewRequestWithContext(ctx, method, url, buffer)
	if err != nil {
		return errors.Wrap(err, "failed to create new request")
	}

	req.Header.Set("Content-Type", "application/vnd.scmm-repository+json;v=2")

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
