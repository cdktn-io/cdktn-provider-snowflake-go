// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakesessionpolicies


type DataSnowflakeSessionPoliciesOn struct {
	// Returns records for the entire account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/data-sources/session_policies#account DataSnowflakeSessionPolicies#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// Returns records for the specified user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/data-sources/session_policies#user DataSnowflakeSessionPolicies#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

