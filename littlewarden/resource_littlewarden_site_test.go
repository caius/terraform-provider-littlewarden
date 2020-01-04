package littlewarden

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccLittlewardenSite(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccLittlewardenSiteBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("littlewarden_site.example_test", "id"),
				),
			},
		},
	})
}

var testAccLittlewardenSiteBasic = `
resource "littlewarden_site" "example_test" {}
`
