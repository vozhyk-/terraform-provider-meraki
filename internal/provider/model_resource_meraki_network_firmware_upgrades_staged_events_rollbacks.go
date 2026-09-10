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

type NetworkFirmwareUpgradesStagedEventsRollbacks struct {
	Id        types.String                                          `tfsdk:"id"`
	NetworkId types.String                                          `tfsdk:"network_id"`
	Reasons   []NetworkFirmwareUpgradesStagedEventsRollbacksReasons `tfsdk:"reasons"`
	Stages    []NetworkFirmwareUpgradesStagedEventsRollbacksStages  `tfsdk:"stages"`
}

type NetworkFirmwareUpgradesStagedEventsRollbacksReasons struct {
	Category types.String `tfsdk:"category"`
	Comment  types.String `tfsdk:"comment"`
}

type NetworkFirmwareUpgradesStagedEventsRollbacksStages struct {
	GroupId                types.String `tfsdk:"group_id"`
	MilestonesScheduledFor types.String `tfsdk:"milestones_scheduled_for"`
}

type NetworkFirmwareUpgradesStagedEventsRollbacksIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkFirmwareUpgradesStagedEventsRollbacks) getPath() string {
	return fmt.Sprintf("/networks/%v/firmwareUpgrades/staged/events/rollbacks", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkFirmwareUpgradesStagedEventsRollbacks) toBody(ctx context.Context, state NetworkFirmwareUpgradesStagedEventsRollbacks) string {
	body := ""
	if len(data.Reasons) > 0 {
		body, _ = sjson.Set(body, "reasons", []interface{}{})
		for _, item := range data.Reasons {
			itemBody := ""
			if !item.Category.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "category", item.Category.ValueString())
			}
			if !item.Comment.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "comment", item.Comment.ValueString())
			}
			body, _ = sjson.SetRaw(body, "reasons.-1", itemBody)
		}
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

func (data *NetworkFirmwareUpgradesStagedEventsRollbacks) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("reasons"); value.Exists() && value.Value() != nil {
		data.Reasons = make([]NetworkFirmwareUpgradesStagedEventsRollbacksReasons, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesStagedEventsRollbacksReasons{}
			if value := res.Get("category"); value.Exists() && value.Value() != nil {
				data.Category = types.StringValue(value.String())
			} else {
				data.Category = types.StringNull()
			}
			if value := res.Get("comment"); value.Exists() && value.Value() != nil {
				data.Comment = types.StringValue(value.String())
			} else {
				data.Comment = types.StringNull()
			}
			(*parent).Reasons = append((*parent).Reasons, data)
			return true
		})
	}
	if value := res.Get("stages"); value.Exists() && value.Value() != nil {
		data.Stages = make([]NetworkFirmwareUpgradesStagedEventsRollbacksStages, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesStagedEventsRollbacksStages{}
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
func (data *NetworkFirmwareUpgradesStagedEventsRollbacks) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Reasons); i++ {
		keys := [...]string{"category", "comment"}
		keyValues := [...]string{data.Reasons[i].Category.ValueString(), data.Reasons[i].Comment.ValueString()}

		parent := &data
		data := (*parent).Reasons[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("reasons").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Reasons[%d] = %+v",
				i,
				(*parent).Reasons[i],
			))
			(*parent).Reasons = slices.Delete((*parent).Reasons, i, i+1)
			i--

			continue
		}
		if value := res.Get("category"); value.Exists() && !data.Category.IsNull() {
			data.Category = types.StringValue(value.String())
		} else {
			data.Category = types.StringNull()
		}
		if value := res.Get("comment"); value.Exists() && !data.Comment.IsNull() {
			data.Comment = types.StringValue(value.String())
		} else {
			data.Comment = types.StringNull()
		}
		(*parent).Reasons[i] = data
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
func (data *NetworkFirmwareUpgradesStagedEventsRollbacks) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *NetworkFirmwareUpgradesStagedEventsRollbacksIdentity) toIdentity(ctx context.Context, plan *NetworkFirmwareUpgradesStagedEventsRollbacks) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *NetworkFirmwareUpgradesStagedEventsRollbacks) fromIdentity(ctx context.Context, identity *NetworkFirmwareUpgradesStagedEventsRollbacksIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkFirmwareUpgradesStagedEventsRollbacks) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
