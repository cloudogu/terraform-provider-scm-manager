package scm

import (
	"bytes"
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
	CreationDate string `json:"creationDate"`
	ImportUrl    string `json:"importUrl"`
}

func (r *Repository) GetID() string {
	return fmt.Sprintf("%s/%s", r.NameSpace, r.Name)
}

func (c *Client) CreateRepository(repo Repository) error {

	b, err := json.Marshal(&repo)
	if err != nil {
		return errors.Wrap(err, "failed to marshal repository")
	}

	buffer := bytes.NewBuffer(b)
	req, err := http.NewRequest("POST", c.config.URL+"/api/v2/repositories", buffer)
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

func (c *Client) GetRepository(name string) (Repository, error) {
	req, err := http.NewRequest("GET", c.config.URL+"/api/v2/repositories/"+name, nil)
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

func (c *Client) UpdateRepository(name string, repo Repository) error {
	b, err := json.Marshal(&repo)
	if err != nil {
		return errors.Wrap(err, "failed to marshal repository")
	}

	buffer := bytes.NewBuffer(b)
	req, err := http.NewRequest("PUT", c.config.URL+"/api/v2/repositories"+"/"+name, buffer)
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

func (c *Client) DeleteRepository(name string) error {
	req, err := http.NewRequest("DELETE", c.config.URL+"/api/v2/repositories/"+name, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create new request")
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ImportRepository(repo Repository) error {
	b, err := json.Marshal(&repo)
	if err != nil {
		return errors.Wrap(err, "failed to marshal repository")
	}

	buffer := bytes.NewBuffer(b)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v2/repositories/import/%s/url", c.config.URL, repo.Type), buffer)
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
