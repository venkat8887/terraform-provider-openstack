package openstack

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccLBV1Monitor_importBasic(t *testing.T) {
	resourceName := "openstack_lb_monitor_v1.monitor_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckDeprecated(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLBV1MonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbV1MonitorBasic,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
