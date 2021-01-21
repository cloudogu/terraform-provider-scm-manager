package scm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRepository() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceScmCreate,
		ReadContext:   resourceScmRead,
		UpdateContext: resourceScmUpdate,
		DeleteContext: resourceScmDelete,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceScmRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return diag.Diagnostics{diag.Diagnostic{
		Severity:      diag.Error,
		Summary:       "resourceScmRead not implemented yet",
		Detail:        "resourceScmRead not implemented yet",
		AttributePath: nil,
	}}
}

func resourceScmCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return diag.Diagnostics{diag.Diagnostic{
		Severity:      diag.Error,
		Summary:       "resourceScmCreate not implemented yet",
		Detail:        "resourceScmCreate not implemented yet",
		AttributePath: nil,
	}}
}

func resourceScmUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return diag.Diagnostics{diag.Diagnostic{
		Severity:      diag.Error,
		Summary:       "resourceScmUpdate not implemented yet",
		Detail:        "resourceScmUpdate not implemented yet",
		AttributePath: nil,
	}}
}

func resourceScmDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return diag.Diagnostics{diag.Diagnostic{
		Severity:      diag.Error,
		Summary:       "resourceScmDelete not implemented yet",
		Detail:        "resourceScmDelete not implemented yet",
		AttributePath: nil,
	}}
}
