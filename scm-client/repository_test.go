package scm_client

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

var testConfig = Config{
	URL:      "http://localhost:8080",
	Username: "scmadmin",
	Password: "scmadmin",
}

var testRepo = Repository{
	NameSpace:   "testspace",
	Name:        "testrepo",
	Type:        "git",
	Description: "desc",
}

func TestClient_CreateRepository(t *testing.T) {
	c := NewClient(testConfig)

	err := c.CreateRepository(testRepo)

	require.NoError(t, err)
}

func TestClient_GetRepository(t *testing.T) {
	c := NewClient(testConfig)

	r, err := c.GetRepository(testRepo.GetID())
	require.NoError(t, err)

	assert.Equal(t, testRepo.NameSpace, r.NameSpace)
	assert.Equal(t, testRepo.Name, r.Name)
	assert.Equal(t, testRepo.Type, r.Type)
	assert.Equal(t, testRepo.Description, r.Description)
}
