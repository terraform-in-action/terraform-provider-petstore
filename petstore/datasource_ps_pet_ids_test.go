package petstore

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPSPetIDsDataSource_basic(t *testing.T) {
	rInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPSPetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPSPetIDsDataSourceConfig(rInt),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.petstore_pet_ids.pets", "ids.#", "1"),
					resource.TestCheckResourceAttrPair(
						"petstore_pet.pet", "id",
						"data.petstore_pet_ids.pets", "ids.0",
					),
				),
			},
		},
	})
}

func testAccPSPetIDsDataSourceConfig(rInt int) string {
	return fmt.Sprintf(`	  
	resource "petstore_pet" "pet" {
		name    = "%d"
		species = "cat"
		age     = 3
	}

	data "petstore_pet_ids" "pets" {
		names = [petstore_pet.pet.name]
	}
	`, rInt)
}
