package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]func() (*schema.Provider, error)
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]func() (*schema.Provider, error){
		"scm": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("SCM_URL"); err == "" {
		t.Fatal("SCM_URL must be set for acceptance tests")
	}
	if err := os.Getenv("SCM_USERNAME"); err == "" {
		t.Fatal("SCM_USERNAME must be set for acceptance tests")
	}
	if err := os.Getenv("SCM_PASSWORD"); err == "" {
		t.Fatal("SCM_PASSWORD must be set for acceptance tests")
	}
}
