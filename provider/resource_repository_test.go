package provider

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccRepositoryBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccCheckRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: basicRepositoryWithDescription("this is a testrepo", ""),
				Check:  resource.TestCheckResourceAttr("scm_repository.testrepo", "id", "scmadmin/testrepo"),
			},
		},
	})
}

func TestAccRepositoryUpdates(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccCheckRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: basicRepositoryWithDescription("this is a testrepo", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("scm_repository.testrepo", "id", "scmadmin/testrepo"),
					resource.TestCheckResourceAttr("scm_repository.testrepo", "last_modified", "")),
			},
			{
				Config: basicRepositoryWithDescription("this is new description", ""),
				Check:  resource.TestCheckResourceAttr("scm_repository.testrepo", "description", "this is new description"),
			},
			{
				Config: basicRepositoryWithDescription("this is new description2", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("scm_repository.testrepo", "description", "this is new description2"),
					resource.TestCheckResourceAttrSet("scm_repository.testrepo", "last_modified")),
			},
		},
	})
}

func TestAccRepositoryImport(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccCheckRepositoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: basicRepositoryWithDescription("this is a testrepo", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("scm_repository.testrepo", "id", "scmadmin/testrepo"),
					resource.TestCheckResourceAttr("scm_repository.testrepo", "last_modified", "")),
			},
			{
				Config: basicRepositoryWithDescription("this is new description", "import_url = \"https://github.com/cloudogu/spring-petclinic\""),
				// For now there is no real check whether the import was successful
				Check: resource.TestCheckResourceAttr("scm_repository.testrepo", "description", "this is new description"),
			},
		},
	})
}

func testAccCheckRepositoryDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "scm_repository" {
			continue
		}

		repoID := rs.Primary.ID

		_, err := c.GetRepository(context.Background(), repoID)
		if err == nil || !strings.Contains(err.Error(), "response had statuscode: 404") {
			return fmt.Errorf("repository %s was not destroyed correctly", repoID)
		}
	}

	return nil
}

func basicRepositoryWithDescription(description string, additionalFields string) string {
	return fmt.Sprintf(`
resource "scm_repository" "testrepo" {
  namespace = "scmadmin"
  name = "testrepo"
  type = "git"
  description = "%s"
  contact = "scmadmin@test.test"
  %s
}
`, description, additionalFields)
}
