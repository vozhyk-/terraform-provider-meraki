resource "meraki_network_firmware_upgrades_staged_event" "example" {
  network_id                                          = "L_123456"
  products_switch_next_upgrade_to_version_id          = "1234"
  products_switch_catalyst_next_upgrade_to_version_id = "4321"
  stages = [
    {
      group_id                 = "1234"
      milestones_scheduled_for = "2018-02-11T00:00:00Z"
    }
  ]
}
