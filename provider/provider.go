package provider

import (
	"context"

	"github.com/cloudogu/terraform-provider-scm/scm"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Client interface {
	CreateRepository(ctx context.Context, repo scm.Repository) error
	GetRepository(ctx context.Context, name string) (scm.Repository, error)
	UpdateRepository(ctx context.Context, name string, repo scm.Repository) error
	DeleteRepository(ctx context.Context, name string) error
	ImportRepository(ctx context.Context, repo scm.Repository) error
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCM_HOST", "http://localhost:8080/scm"),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCM_USERNAME", "scmadmin"),
			},
			"password": {
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
		host = hVal.(string)
	}

	client := scm.NewClient(scm.Config{
		URL:      host,
		Username: username,
		Password: password,
	})

	return client, nil
}
