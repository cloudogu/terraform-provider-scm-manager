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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"import_url": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRepositoryCreate(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(Client)

	var diags diag.Diagnostics

	repo := repositoryFromState(d)

	if repo.ImportUrl != "" {
		err := client.ImportRepository(repo)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		err := client.CreateRepository(repo)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(repo.GetID())
	resourceRepositoryRead(ctx, d, i)

	return diags
}

func resourceRepositoryRead(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(Client)

	var diags diag.Diagnostics

	repoID := d.Id()

	repo, err := client.GetRepository(repoID)
	if err != nil {
		return diag.FromErr(err)
	}

	repositorySetToState(repo, d)

	return diags
}

func resourceRepositoryUpdate(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(Client)

	var diags diag.Diagnostics

	repoID := d.Id()
	repo := repositoryFromState(d)

	err := client.UpdateRepository(repoID, repo)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repo.GetID())
	repositorySetToState(repo, d)

	return diags
}

func resourceRepositoryDelete(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(Client)

	var diags diag.Diagnostics

	repoID := d.Id()

	err := client.DeleteRepository(repoID)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func repositorySetToState(repo scm.Repository, d *schema.ResourceData) {
	d.Set("namespace", repo.NameSpace)
	d.Set("name", repo.Name)
	d.Set("type", repo.Type)
	d.Set("description", repo.Description)
}

func repositoryFromState(d *schema.ResourceData) scm.Repository {
	repo := scm.Repository{}

	repo.NameSpace = d.Get("namespace").(string)
	repo.Name = d.Get("name").(string)
	repo.Type = d.Get("type").(string)
	repo.Description = d.Get("description").(string)
	repo.ImportUrl = d.Get("import_url").(string)

	return repo
}
