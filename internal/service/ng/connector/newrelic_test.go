package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnector_newrelic(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_connector.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_newrelic(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "newrelic.0.url", "https://newrelic.com/"),
					resource.TestCheckResourceAttr(resourceName, "newrelic.0.account_id", "nr_account_id"),
					resource.TestCheckResourceAttr(resourceName, "newrelic.0.api_key_ref", "account.acctest_sumo_access_key"),
					resource.TestCheckResourceAttr(resourceName, "newrelic.0.delegate_selectors.#", "1"),
				),
			},
			{
				Config: testAccResourceConnector_newrelic(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "newrelic.0.url", "https://newrelic.com/"),
					resource.TestCheckResourceAttr(resourceName, "newrelic.0.account_id", "nr_account_id"),
					resource.TestCheckResourceAttr(resourceName, "newrelic.0.api_key_ref", "account.acctest_sumo_access_key"),
					resource.TestCheckResourceAttr(resourceName, "newrelic.0.delegate_selectors.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceConnector_newrelic(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			newrelic {
				url = "https://newrelic.com/"
				delegate_selectors = ["harness-delegate"]
				account_id = "nr_account_id"
				api_key_ref = "account.acctest_sumo_access_key"
			}
		}
`, id, name)
}
