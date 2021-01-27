// +build integration

package scm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

var testConfig = Config{
	URL:      "http://localhost:8080/scm",
	Username: "scmadmin",
	Password: "scmadmin",
}

var testRepo = Repository{
	NameSpace:   "scmadmin",
	Name:        "testrepo",
	Type:        "git",
	Description: "desc",
	Contact:     "test@test.test",
	ImportUrl:   "https://github.com/cloudogu/spring-petclinic",
}

func TestClient_CreateRepository(t *testing.T) {
	c := NewClient(testConfig)
	defer c.DeleteRepository(context.Background(), testRepo.GetID())

	err := c.CreateRepository(context.Background(), testRepo)

	require.NoError(t, err)
}

func TestClient_DeleteRepository(t *testing.T) {
	c := NewClient(testConfig)
	err := c.CreateRepository(context.Background(), testRepo)

	err = c.DeleteRepository(context.Background(), testRepo.GetID())
	require.NoError(t, err)

	_, err = c.GetRepository(context.Background(), testRepo.GetID())
	require.Error(t, err)
}

func TestClient_GetRepository(t *testing.T) {
	c := NewClient(testConfig)
	defer c.DeleteRepository(context.Background(), testRepo.GetID())
	err := c.CreateRepository(context.Background(), testRepo)

	r, err := c.GetRepository(context.Background(), testRepo.GetID())
	require.NoError(t, err)

	assert.Equal(t, testRepo.NameSpace, r.NameSpace)
	assert.Equal(t, testRepo.Name, r.Name)
	assert.Equal(t, testRepo.Type, r.Type)
	assert.Equal(t, testRepo.Description, r.Description)
}

func TestClient_UpdateRepository(t *testing.T) {
	c := NewClient(testConfig)
	defer c.DeleteRepository(context.Background(), testRepo.GetID())
	err := c.CreateRepository(context.Background(), testRepo)

	oldRepo, err := c.GetRepository(context.Background(), testRepo.GetID())
	require.NoError(t, err)
	updatedRepo := oldRepo
	updatedRepo.Description = "updated desc"

	err = c.UpdateRepository(context.Background(), testRepo.GetID(), updatedRepo)
	require.NoError(t, err)

	newRepo, err := c.GetRepository(context.Background(), testRepo.GetID())
	require.NoError(t, err)
	require.Equal(t, updatedRepo.Description, newRepo.Description)
}

func TestClient_ImportRepository(t *testing.T) {
	c := NewClient(testConfig)
	defer c.DeleteRepository(context.Background(), testRepo.GetID())

	err := c.ImportRepository(context.Background(), testRepo)

	require.NoError(t, err)
}
