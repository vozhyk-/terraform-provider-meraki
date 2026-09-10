resource "meraki_network_firmware_upgrades_staged_stage" "example" {
  network_id = "L_123456"
  _json = [
    {
      group_id = "1234"
    }
  ]
}
