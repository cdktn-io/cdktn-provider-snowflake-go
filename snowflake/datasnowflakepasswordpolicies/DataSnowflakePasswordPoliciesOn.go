// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakepasswordpolicies


type DataSnowflakePasswordPoliciesOn struct {
	// Returns records for the entire account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/data-sources/password_policies#account DataSnowflakePasswordPolicies#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// Returns records for the specified user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/data-sources/password_policies#user DataSnowflakePasswordPolicies#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

