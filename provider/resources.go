// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package newrelic

import (
	"strings"
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/newrelic/terraform-provider-newrelic/newrelic"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "newrelic"
	// modules:
	mainMod       = "index"
	syntheticsMod = "Synthetics"
	insightsMod   = "Insights"
	pluginsMod    = "Plugins"
)

var namespaceMap = map[string]string{
	mainPkg: "NewRelic",
}

// makeMember manufactures a type token for the package and the given module and type.  It automatically
//// uses the NewRelic package and names the file by simply lower casing the resource's first character.
// first character.
func makeMember(moduleTitle string, mem string) tokens.ModuleMember {
	moduleName := strings.ToLower(moduleTitle)
	namespaceMap[moduleName] = moduleTitle
	fn := string(unicode.ToLower(rune(mem[0]))) + mem[1:]
	token := moduleName + "/" + fn
	return tokens.ModuleMember(mainPkg + ":" + token + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	return makeMember(mod, res)
}

// makeResource manufactures a standard resource token given a module and resource name.
func makeResource(mod string, res string) tokens.Type {
	return makeType(mod, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := newrelic.Provider().(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "newrelic",
		Description: "A Pulumi package for creating and managing New Relic resources.",
		Keywords:    []string{"pulumi", "new relic"},
		License:     "Apache-2.0",
		GitHubOrg:   "newrelic",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-newrelic",
		Config: map[string]*tfbridge.SchemaInfo{
			"account_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"NEW_RELIC_ACCOUNT_ID"},
				},
			},
			"api_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"NEW_RELIC_API_KEY"},
				},
			},
			"admin_api_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"NEW_RELIC_ADMIN_API_KEY"},
				},
			},
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"NEW_RELIC_REGION"},
					Value:   "US",
				},
			},
			"insights_insert_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"NEW_RELIC_INSIGHTS_INSERT_KEY"},
				},
			},
			"insights_insert_url": {
				Default: &tfbridge.DefaultInfo{
					Value:   "https://insights-collector.newrelic.com/v1/accounts",
					EnvVars: []string{"NEW_RELIC_INSIGHTS_INSERT_URL"},
				},
			},
			"insights_query_url": {
				Default: &tfbridge.DefaultInfo{
					Value:   "https://insights-api.newrelic.com/v1/accounts",
					EnvVars: []string{"NEW_RELIC_INSIGHTS_QUERY_URL"},
				},
			},
			"insecure_skip_verify": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"NEW_RELIC_API_SKIP_VERIFY"},
				},
			},
			"cacert_file": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"NEW_RELIC_API_CACERT"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"newrelic_alert_channel":          {Tok: makeResource(mainMod, "AlertChannel")},
			"newrelic_alert_condition":        {Tok: makeResource(mainMod, "AlertCondition")},
			"newrelic_alert_policy":           {Tok: makeResource(mainMod, "AlertPolicy")},
			"newrelic_alert_policy_channel":   {Tok: makeResource(mainMod, "AlertPolicyChannel")},
			"newrelic_dashboard":              {Tok: makeResource(mainMod, "Dashboard")},
			"newrelic_infra_alert_condition":  {Tok: makeResource(mainMod, "InfraAlertCondition")},
			"newrelic_nrql_alert_condition":   {Tok: makeResource(mainMod, "NrqlAlertCondition")},
			"newrelic_entity_tags":            {Tok: makeResource(mainMod, "EntityTags")},
			"newrelic_events_to_metrics_rule": {Tok: makeResource(mainMod, "EventsToMetricsRule")},

			"newrelic_synthetics_alert_condition":   {Tok: makeResource(syntheticsMod, "AlertCondition")},
			"newrelic_synthetics_monitor":           {Tok: makeResource(syntheticsMod, "Monitor")},
			"newrelic_synthetics_monitor_script":    {Tok: makeResource(syntheticsMod, "MonitorScript")},
			"newrelic_synthetics_label":             {Tok: makeResource(syntheticsMod, "Label")},
			"newrelic_synthetics_secure_credential": {Tok: makeResource(syntheticsMod, "SecureCredential")},
			"newrelic_synthetics_multilocation_alert_condition": {
				Tok: makeResource(syntheticsMod, "MultiLocationAlertCondition"),
			},

			"newrelic_insights_event": {Tok: makeResource(insightsMod, "Event")},

			"newrelic_plugins_alert_condition": {Tok: makeResource(pluginsMod, "AlertCondition")},
			"newrelic_workload":                {Tok: makeResource(pluginsMod, "Workload")},
			"newrelic_application_settings":    {Tok: makeResource(pluginsMod, "ApplicationSettings")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"newrelic_alert_channel":   {Tok: makeDataSource(mainMod, "getAlertChannel")},
			"newrelic_alert_policy":    {Tok: makeDataSource(mainMod, "getAlertPolicy")},
			"newrelic_application":     {Tok: makeDataSource(mainMod, "getApplication")},
			"newrelic_key_transaction": {Tok: makeDataSource(mainMod, "getKeyTransaction")},
			"newrelic_entity":          {Tok: makeDataSource(mainMod, "getEntity")},

			"newrelic_synthetics_monitor":           {Tok: makeDataSource(syntheticsMod, "getMonitor")},
			"newrelic_synthetics_secure_credential": {Tok: makeDataSource(syntheticsMod, "getSecureCredential")},

			"newrelic_plugin":           {Tok: makeDataSource(pluginsMod, "getPlugin")},
			"newrelic_plugin_component": {Tok: makeDataSource(pluginsMod, "getPluginComponent")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=2.0.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: namespaceMap,
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
