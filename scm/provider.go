package scm

import (
	"context"

	scm_client "github.com/cloudogu/terraform-provider-scm/scm-client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Client interface {
	CreateRepository(repo scm_client.Repository) error
	GetRepository(name string) (scm_client.Repository, error)
	UpdateRepository(name string, repo scm_client.Repository) error
	DeleteRepository(name string) error
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCM_HOST", "http://localhost:8080/scm"),
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCM_USERNAME", "scmadmin"),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SCM_PASSWORD", "scmadmin"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"scm_repository": resourceRepository(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	var host string

	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = tempHost
	}

	client := scm_client.NewClient(scm_client.Config{
		URL:      host,
		Username: username,
		Password: password,
	})

	return client, nil
}
