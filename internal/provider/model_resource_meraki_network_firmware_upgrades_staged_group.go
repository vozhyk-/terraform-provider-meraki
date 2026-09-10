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
	"context"
	"fmt"
	"net/url"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkFirmwareUpgradesStagedGroup struct {
	Id                          types.String                                                    `tfsdk:"id"`
	NetworkId                   types.String                                                    `tfsdk:"network_id"`
	Description                 types.String                                                    `tfsdk:"description"`
	IsDefault                   types.Bool                                                      `tfsdk:"is_default"`
	Name                        types.String                                                    `tfsdk:"name"`
	AssignedDevicesDevices      []NetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices      `tfsdk:"assigned_devices_devices"`
	AssignedDevicesSwitchStacks []NetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks `tfsdk:"assigned_devices_switch_stacks"`
}

type NetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices struct {
	Name   types.String `tfsdk:"name"`
	Serial types.String `tfsdk:"serial"`
}

type NetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type NetworkFirmwareUpgradesStagedGroupIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkFirmwareUpgradesStagedGroup) getPath() string {
	return fmt.Sprintf("/networks/%v/firmwareUpgrades/staged/groups", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkFirmwareUpgradesStagedGroup) toBody(ctx context.Context, state NetworkFirmwareUpgradesStagedGroup) string {
	body := ""
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.IsDefault.IsNull() {
		body, _ = sjson.Set(body, "isDefault", data.IsDefault.ValueBool())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if len(data.AssignedDevicesDevices) > 0 {
		body, _ = sjson.Set(body, "assignedDevices.devices", []interface{}{})
		for _, item := range data.AssignedDevicesDevices {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Serial.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serial", item.Serial.ValueString())
			}
			body, _ = sjson.SetRaw(body, "assignedDevices.devices.-1", itemBody)
		}
	}
	if len(data.AssignedDevicesSwitchStacks) > 0 {
		body, _ = sjson.Set(body, "assignedDevices.switchStacks", []interface{}{})
		for _, item := range data.AssignedDevicesSwitchStacks {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "assignedDevices.switchStacks.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkFirmwareUpgradesStagedGroup) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && value.Value() != nil {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("isDefault"); value.Exists() && value.Value() != nil {
		data.IsDefault = types.BoolValue(value.Bool())
	} else {
		data.IsDefault = types.BoolNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("assignedDevices.devices"); value.Exists() && value.Value() != nil {
		data.AssignedDevicesDevices = make([]NetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices{}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("serial"); value.Exists() && value.Value() != nil {
				data.Serial = types.StringValue(value.String())
			} else {
				data.Serial = types.StringNull()
			}
			(*parent).AssignedDevicesDevices = append((*parent).AssignedDevicesDevices, data)
			return true
		})
	}
	if value := res.Get("assignedDevices.switchStacks"); value.Exists() && value.Value() != nil {
		data.AssignedDevicesSwitchStacks = make([]NetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).AssignedDevicesSwitchStacks = append((*parent).AssignedDevicesSwitchStacks, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkFirmwareUpgradesStagedGroup) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("isDefault"); value.Exists() && !data.IsDefault.IsNull() {
		data.IsDefault = types.BoolValue(value.Bool())
	} else {
		data.IsDefault = types.BoolNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	for i := 0; i < len(data.AssignedDevicesDevices); i++ {
		keys := [...]string{"name", "serial"}
		keyValues := [...]string{data.AssignedDevicesDevices[i].Name.ValueString(), data.AssignedDevicesDevices[i].Serial.ValueString()}

		parent := &data
		data := (*parent).AssignedDevicesDevices[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("assignedDevices.devices").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing AssignedDevicesDevices[%d] = %+v",
				i,
				(*parent).AssignedDevicesDevices[i],
			))
			(*parent).AssignedDevicesDevices = slices.Delete((*parent).AssignedDevicesDevices, i, i+1)
			i--

			continue
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("serial"); value.Exists() && !data.Serial.IsNull() {
			data.Serial = types.StringValue(value.String())
		} else {
			data.Serial = types.StringNull()
		}
		(*parent).AssignedDevicesDevices[i] = data
	}
	for i := 0; i < len(data.AssignedDevicesSwitchStacks); i++ {
		keys := [...]string{"id", "name"}
		keyValues := [...]string{data.AssignedDevicesSwitchStacks[i].Id.ValueString(), data.AssignedDevicesSwitchStacks[i].Name.ValueString()}

		parent := &data
		data := (*parent).AssignedDevicesSwitchStacks[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("assignedDevices.switchStacks").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing AssignedDevicesSwitchStacks[%d] = %+v",
				i,
				(*parent).AssignedDevicesSwitchStacks[i],
			))
			(*parent).AssignedDevicesSwitchStacks = slices.Delete((*parent).AssignedDevicesSwitchStacks, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).AssignedDevicesSwitchStacks[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkFirmwareUpgradesStagedGroup) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *NetworkFirmwareUpgradesStagedGroupIdentity) toIdentity(ctx context.Context, plan *NetworkFirmwareUpgradesStagedGroup) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *NetworkFirmwareUpgradesStagedGroup) fromIdentity(ctx context.Context, identity *NetworkFirmwareUpgradesStagedGroupIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkFirmwareUpgradesStagedGroup) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
