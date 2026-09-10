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

type NetworkFirmwareUpgradesStagedEvent struct {
	Id                                           types.String                               `tfsdk:"id"`
	NetworkId                                    types.String                               `tfsdk:"network_id"`
	ProductsSwitchNextUpgradeToVersionId         types.String                               `tfsdk:"products_switch_next_upgrade_to_version_id"`
	ProductsSwitchCatalystNextUpgradeToVersionId types.String                               `tfsdk:"products_switch_catalyst_next_upgrade_to_version_id"`
	Stages                                       []NetworkFirmwareUpgradesStagedEventStages `tfsdk:"stages"`
}

type NetworkFirmwareUpgradesStagedEventStages struct {
	GroupId                types.String `tfsdk:"group_id"`
	MilestonesScheduledFor types.String `tfsdk:"milestones_scheduled_for"`
}

type NetworkFirmwareUpgradesStagedEventIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkFirmwareUpgradesStagedEvent) getPath() string {
	return fmt.Sprintf("/networks/%v/firmwareUpgrades/staged/events", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkFirmwareUpgradesStagedEvent) toBody(ctx context.Context, state NetworkFirmwareUpgradesStagedEvent) string {
	body := ""
	if !data.ProductsSwitchNextUpgradeToVersionId.IsNull() {
		body, _ = sjson.Set(body, "products.switch.nextUpgrade.toVersion.id", data.ProductsSwitchNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsSwitchCatalystNextUpgradeToVersionId.IsNull() {
		body, _ = sjson.Set(body, "products.switchCatalyst.nextUpgrade.toVersion.id", data.ProductsSwitchCatalystNextUpgradeToVersionId.ValueString())
	}
	{
		body, _ = sjson.Set(body, "stages", []interface{}{})
		for _, item := range data.Stages {
			itemBody := ""
			if !item.GroupId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "group.id", item.GroupId.ValueString())
			}
			if !item.MilestonesScheduledFor.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "milestones.scheduledFor", item.MilestonesScheduledFor.ValueString())
			}
			body, _ = sjson.SetRaw(body, "stages.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkFirmwareUpgradesStagedEvent) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("products.switch.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("stages"); value.Exists() && value.Value() != nil {
		data.Stages = make([]NetworkFirmwareUpgradesStagedEventStages, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesStagedEventStages{}
			if value := res.Get("group.id"); value.Exists() && value.Value() != nil {
				data.GroupId = types.StringValue(value.String())
			} else {
				data.GroupId = types.StringNull()
			}
			if value := res.Get("milestones.scheduledFor"); value.Exists() && value.Value() != nil {
				data.MilestonesScheduledFor = types.StringValue(value.String())
			} else {
				data.MilestonesScheduledFor = types.StringNull()
			}
			(*parent).Stages = append((*parent).Stages, data)
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
func (data *NetworkFirmwareUpgradesStagedEvent) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("products.switch.nextUpgrade.toVersion.id"); value.Exists() && !data.ProductsSwitchNextUpgradeToVersionId.IsNull() {
		data.ProductsSwitchNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.nextUpgrade.toVersion.id"); value.Exists() && !data.ProductsSwitchCatalystNextUpgradeToVersionId.IsNull() {
		data.ProductsSwitchCatalystNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystNextUpgradeToVersionId = types.StringNull()
	}
	for i := 0; i < len(data.Stages); i++ {
		keys := [...]string{"group.id", "milestones.scheduledFor"}
		keyValues := [...]string{data.Stages[i].GroupId.ValueString(), data.Stages[i].MilestonesScheduledFor.ValueString()}

		parent := &data
		data := (*parent).Stages[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("stages").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Stages[%d] = %+v",
				i,
				(*parent).Stages[i],
			))
			(*parent).Stages = slices.Delete((*parent).Stages, i, i+1)
			i--

			continue
		}
		if value := res.Get("group.id"); value.Exists() && !data.GroupId.IsNull() {
			data.GroupId = types.StringValue(value.String())
		} else {
			data.GroupId = types.StringNull()
		}
		if value := res.Get("milestones.scheduledFor"); value.Exists() && !data.MilestonesScheduledFor.IsNull() {
			data.MilestonesScheduledFor = types.StringValue(value.String())
		} else {
			data.MilestonesScheduledFor = types.StringNull()
		}
		(*parent).Stages[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkFirmwareUpgradesStagedEvent) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *NetworkFirmwareUpgradesStagedEventIdentity) toIdentity(ctx context.Context, plan *NetworkFirmwareUpgradesStagedEvent) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *NetworkFirmwareUpgradesStagedEvent) fromIdentity(ctx context.Context, identity *NetworkFirmwareUpgradesStagedEventIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkFirmwareUpgradesStagedEvent) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
