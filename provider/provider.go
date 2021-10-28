package provider

import (
	"context"
	"github.com/cloudogu/terraform-provider-scm-manager/scm"

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
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCM_URL", "http://localhost:8080/scm"),
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
			"skip_cert_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCM_SKIP_CERT_VERIFY", false),
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
	skipVerify := d.Get("skip_cert_verify").(bool)

	var url string

	uVal, ok := d.GetOk("url")
	if ok {
		url = uVal.(string)
	}

	client := scm.NewClient(scm.Config{
		URL:            url,
		Username:       username,
		Password:       password,
		SkipCertVerify: skipVerify,
	})

	return client, nil
}
