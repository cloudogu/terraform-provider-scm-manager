package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/cloudogu/terraform-provider-scm/scm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRepository() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRepositoryCreate,
		ReadContext:   resourceRepositoryRead,
		UpdateContext: resourceRepositoryUpdate,
		DeleteContext: resourceRepositoryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"git", "svn"}, false),
			},
			"creation_date": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contact": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"import_url": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"import_username": {
				Type:     schema.TypeString,
				Optional: true,
				StateFunc: func(i interface{}) string {
					return "excluded from state"
				},
			},
			"import_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				StateFunc: func(i interface{}) string {
					return "excluded from state"
				},
			},
			"last_modified": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRepositoryCreate(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(Client)

	repo := repositoryFromState(d)

	if repo.ImportUrl != "" {
		err := client.ImportRepository(ctx, repo)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		err := client.CreateRepository(ctx, repo)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(repo.GetID())
	return resourceRepositoryRead(ctx, d, i)
}

func resourceRepositoryRead(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(Client)

	repoID := d.Id()

	repo, err := client.GetRepository(ctx, repoID)
	if err != nil {
		return diag.FromErr(err)
	}

	return repositorySetToState(repo, d)
}

func resourceRepositoryUpdate(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(Client)

	repoID := d.Id()
	repo := repositoryFromState(d)

	err := client.UpdateRepository(ctx, repoID, repo)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repo.GetID())
	return resourceRepositoryRead(ctx, d, i)
}

func resourceRepositoryDelete(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(Client)

	var diags diag.Diagnostics

	repoID := d.Id()

	err := client.DeleteRepository(ctx, repoID)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func repositorySetToState(repo scm.Repository, d *schema.ResourceData) diag.Diagnostics {
	if err := d.Set("namespace", repo.NameSpace); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", repo.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("type", repo.Type); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", repo.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_modified", repo.LastModified); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("creation_date", repo.CreationDate); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("contact", repo.Contact); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func repositoryFromState(d *schema.ResourceData) scm.Repository {
	repo := scm.Repository{}

	repo.NameSpace = d.Get("namespace").(string)
	repo.Name = d.Get("name").(string)
	repo.Type = d.Get("type").(string)
	repo.Description = d.Get("description").(string)
	repo.Contact = d.Get("contact").(string)
	repo.ImportUrl = d.Get("import_url").(string)
	repo.ImportUsername = d.Get("import_username").(string)
	repo.ImportPassword = d.Get("import_password").(string)
	repo.LastModified = d.Get("last_modified").(string)
	repo.CreationDate = d.Get("creation_date").(string)

	return repo
}
