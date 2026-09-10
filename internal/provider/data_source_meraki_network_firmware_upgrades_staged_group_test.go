// Copyright © 2024 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceMerakiNetworkFirmwareUpgradesStagedGroup(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_staged_group.test", "description", "The description of the group"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_staged_group.test", "is_default", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_staged_group.test", "name", "My Staged Upgrade Group"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_staged_group.test", "assigned_devices_devices.0.name", "Device Name"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_staged_group.test", "assigned_devices_devices.0.serial", "Q234-ABCD-5678"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_staged_group.test", "assigned_devices_switch_stacks.0.id", "1234"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_staged_group.test", "assigned_devices_switch_stacks.0.name", "Stack Name"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiNetworkFirmwareUpgradesStagedGroupPrerequisitesConfig + testAccDataSourceMerakiNetworkFirmwareUpgradesStagedGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiNetworkFirmwareUpgradesStagedGroupPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance", "sensor", "camera"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiNetworkFirmwareUpgradesStagedGroupConfig() string {
	config := `resource "meraki_network_firmware_upgrades_staged_group" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  description = "The description of the group"` + "\n"
	config += `  is_default = false` + "\n"
	config += `  name = "My Staged Upgrade Group"` + "\n"
	config += `  assigned_devices_devices = [{` + "\n"
	config += `    name = "Device Name"` + "\n"
	config += `    serial = "Q234-ABCD-5678"` + "\n"
	config += `  }]` + "\n"
	config += `  assigned_devices_switch_stacks = [{` + "\n"
	config += `    id = "1234"` + "\n"
	config += `    name = "Stack Name"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_network_firmware_upgrades_staged_group" "test" {
			id = meraki_network_firmware_upgrades_staged_group.test.id
			network_id = meraki_network.test.id
			depends_on = [meraki_network_firmware_upgrades_staged_group.test]
		}
	`
	return config
}

func testAccNamedDataSourceMerakiNetworkFirmwareUpgradesStagedGroupConfig() string {
	config := `resource "meraki_network_firmware_upgrades_staged_group" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  description = "The description of the group"` + "\n"
	config += `  is_default = false` + "\n"
	config += `  name = "My Staged Upgrade Group"` + "\n"
	config += `  assigned_devices_devices = [{` + "\n"
	config += `    name = "Device Name"` + "\n"
	config += `    serial = "Q234-ABCD-5678"` + "\n"
	config += `  }]` + "\n"
	config += `  assigned_devices_switch_stacks = [{` + "\n"
	config += `    id = "1234"` + "\n"
	config += `    name = "Stack Name"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_network_firmware_upgrades_staged_group" "test" {
			name = meraki_network_firmware_upgrades_staged_group.test.name
			network_id = meraki_network.test.id
			depends_on = [meraki_network_firmware_upgrades_staged_group.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
