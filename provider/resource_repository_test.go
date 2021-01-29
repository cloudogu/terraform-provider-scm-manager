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
				Config: basicRepositoryWithDescription("this is a testrepo"),
				Check:  resource.TestCheckResourceAttr("scm_repository.testrepo", "id", "scmadmin/testrepo"),
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

func basicRepositoryWithDescription(description string) string {
	return fmt.Sprintf(`
resource "scm_repository" "testrepo" {
  namespace = "scmadmin"
  name = "testrepo"
  type = "git"
  description = "%s"
  contact = "scmadmin@test.test"
}
`, description)
}
