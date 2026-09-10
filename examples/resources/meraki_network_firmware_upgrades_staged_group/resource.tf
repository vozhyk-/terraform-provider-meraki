resource "meraki_network_firmware_upgrades_staged_group" "example" {
  network_id  = "L_123456"
  description = "The description of the group"
  is_default  = false
  name        = "My Staged Upgrade Group"
  assigned_devices_devices = [
    {
      name   = "Device Name"
      serial = "Q234-ABCD-5678"
    }
  ]
  assigned_devices_switch_stacks = [
    {
      id   = "1234"
      name = "Stack Name"
    }
  ]
}
