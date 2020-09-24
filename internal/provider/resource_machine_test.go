package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourcemachine(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcemachine,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"danube_machine.foo", "name", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccResourcemachine = `
resource "danube_machine" "foo" {
  name = "bar"
}
`
