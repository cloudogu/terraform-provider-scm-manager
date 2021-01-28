package scm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const (
	failedToCreateRequestError    = "failed to create new request"
	failedToCreateRequestUrlError = "failed to create request url"
	apiRepositoriesURL            = "/api/v2/repositories/"
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
	requestURL, err := UrlJoin(c.config.URL, apiRepositoriesURL)
	if err != nil {
		return errors.Wrap(err, failedToCreateRequestUrlError)
	}
	return c.setRepository(ctx, repo, http.MethodPost, requestURL)
}

func (c *Client) GetRepository(ctx context.Context, name string) (Repository, error) {
	requestURL, err := UrlJoin(c.config.URL, apiRepositoriesURL, name)
	if err != nil {
		return Repository{}, errors.Wrap(err, failedToCreateRequestUrlError)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return Repository{}, errors.Wrap(err, failedToCreateRequestError)
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
	requestURL, err := UrlJoin(c.config.URL, apiRepositoriesURL, name)
	if err != nil {
		return errors.Wrap(err, failedToCreateRequestUrlError)
	}
	return c.setRepository(ctx, repo, http.MethodPut, requestURL)
}

func (c *Client) DeleteRepository(ctx context.Context, name string) error {
	requestURL, err := UrlJoin(c.config.URL, apiRepositoriesURL, name)
	if err != nil {
		return errors.Wrap(err, failedToCreateRequestUrlError)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, requestURL, nil)
	if err != nil {
		return errors.Wrap(err, failedToCreateRequestError)
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ImportRepository(ctx context.Context, repo Repository) error {
	requestURL, err := UrlJoin(c.config.URL, apiRepositoriesURL, "/import/", repo.Type, "/url")
	if err != nil {
		return errors.Wrap(err, failedToCreateRequestUrlError)
	}
	return c.setRepository(ctx, repo, http.MethodPost, requestURL)
}

func (c *Client) setRepository(ctx context.Context, repo Repository, method string, url string) error {
	b, err := json.Marshal(&repo)
	if err != nil {
		return errors.Wrap(err, "failed to marshal repository")
	}

	buffer := bytes.NewBuffer(b)

	req, err := http.NewRequestWithContext(ctx, method, url, buffer)
	if err != nil {
		return errors.Wrap(err, failedToCreateRequestError)
	}

	req.Header.Set("Content-Type", "application/vnd.scmm-repository+json;v=2")

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
