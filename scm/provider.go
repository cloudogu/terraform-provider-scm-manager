package scm

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
	//	ConfigureContextFunc: providerConfigure,
	}
}

/*  TODO: create scm client
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
}*/