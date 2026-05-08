// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakenetworkrules


type DataSnowflakeNetworkRulesLimit struct {
	// The maximum number of rows to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/data-sources/network_rules#rows DataSnowflakeNetworkRules#rows}
	Rows *float64 `field:"required" json:"rows" yaml:"rows"`
	// Specifies a **case-sensitive** pattern that is used to match object name.
	//
	// After the first match, the limit on the number of rows will be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/data-sources/network_rules#from DataSnowflakeNetworkRules#from}
	From *string `field:"optional" json:"from" yaml:"from"`
}

