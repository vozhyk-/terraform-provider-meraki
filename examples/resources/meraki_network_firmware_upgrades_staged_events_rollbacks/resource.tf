resource "meraki_network_firmware_upgrades_staged_events_rollbacks" "example" {
  network_id = "L_123456"
  reasons = [
    {
      category = "performance"
      comment  = "Network was slower with the upgrade"
    }
  ]
  stages = [
    {
      group_id                 = "1234"
      milestones_scheduled_for = "2018-02-11T00:00:00Z"
    }
  ]
}
