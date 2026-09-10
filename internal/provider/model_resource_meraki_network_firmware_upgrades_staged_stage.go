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

type NetworkFirmwareUpgradesStagedStage struct {
	Id        types.String                             `tfsdk:"id"`
	NetworkId types.String                             `tfsdk:"network_id"`
	Json      []NetworkFirmwareUpgradesStagedStageJson `tfsdk:"_json"`
}

type NetworkFirmwareUpgradesStagedStageJson struct {
	GroupId types.String `tfsdk:"group_id"`
}

type NetworkFirmwareUpgradesStagedStageIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkFirmwareUpgradesStagedStage) getPath() string {
	return fmt.Sprintf("/networks/%v/firmwareUpgrades/staged/stages", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkFirmwareUpgradesStagedStage) toBody(ctx context.Context, state NetworkFirmwareUpgradesStagedStage) string {
	body := ""
	if len(data.Json) > 0 {
		body, _ = sjson.Set(body, "_json", []interface{}{})
		for _, item := range data.Json {
			itemBody := ""
			if !item.GroupId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "group.id", item.GroupId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "_json.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPreservingNulls

// toBodyPreservingNulls walks the same writable-attribute schema as toBody but
// reads directly from the raw API response (gjson) instead of from the
// Terraform model. Unlike toBody, it preserves attributes that the API
// explicitly returned as `null` (emitting them as JSON `null` rather than
// dropping them). This is used by the singleton restoreOriginalStateOnDestroy
// path so that explicit-null fields captured during Create are restored on
// Delete. Keep this method in sync with toBody — both walk the same
// `.Attributes` schema and must agree on which fields are writable.
func (data NetworkFirmwareUpgradesStagedStage) toBodyPreservingNulls(ctx context.Context, res meraki.Res) string {
	body := ""
	if value := res.Get("_json"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "_json", "null")
		} else {
			body, _ = sjson.Set(body, "_json", []interface{}{})
			parent := &body
			value.ForEach(func(k, res gjson.Result) bool {
				body := ""
				if value := res.Get("group.id"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "group.id", "null")
					} else {
						body, _ = sjson.Set(body, "group.id", value.String())
					}
				}
				*parent, _ = sjson.SetRaw(*parent, "_json.-1", body)
				return true
			})
		}
	}
	return body
}

// End of section. //template:end toBodyPreservingNulls

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkFirmwareUpgradesStagedStage) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("_json"); value.Exists() && value.Value() != nil {
		data.Json = make([]NetworkFirmwareUpgradesStagedStageJson, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesStagedStageJson{}
			if value := res.Get("group.id"); value.Exists() && value.Value() != nil {
				data.GroupId = types.StringValue(value.String())
			} else {
				data.GroupId = types.StringNull()
			}
			(*parent).Json = append((*parent).Json, data)
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
func (data *NetworkFirmwareUpgradesStagedStage) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Json); i++ {
		keys := [...]string{"group.id"}
		keyValues := [...]string{data.Json[i].GroupId.ValueString()}

		parent := &data
		data := (*parent).Json[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("_json").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Json[%d] = %+v",
				i,
				(*parent).Json[i],
			))
			(*parent).Json = slices.Delete((*parent).Json, i, i+1)
			i--

			continue
		}
		if value := res.Get("group.id"); value.Exists() && !data.GroupId.IsNull() {
			data.GroupId = types.StringValue(value.String())
		} else {
			data.GroupId = types.StringNull()
		}
		(*parent).Json[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkFirmwareUpgradesStagedStage) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *NetworkFirmwareUpgradesStagedStageIdentity) toIdentity(ctx context.Context, plan *NetworkFirmwareUpgradesStagedStage) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *NetworkFirmwareUpgradesStagedStage) fromIdentity(ctx context.Context, identity *NetworkFirmwareUpgradesStagedStageIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkFirmwareUpgradesStagedStage) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
