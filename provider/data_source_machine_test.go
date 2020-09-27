package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcemachine(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcemachine,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"danube_machine.foo", "hostname", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccDataSourcemachine = `
resource "danube_machine" "foo" {
  hostname = "bar"
}
`
