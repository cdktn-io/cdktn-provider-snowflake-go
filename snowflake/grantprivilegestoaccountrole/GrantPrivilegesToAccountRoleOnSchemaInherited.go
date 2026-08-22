// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoaccountrole


type GrantPrivilegesToAccountRoleOnSchemaInherited struct {
	// If true, the inherited privilege will be granted on all schemas in the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/grant_privileges_to_account_role#in_account GrantPrivilegesToAccountRole#in_account}
	InAccount interface{} `field:"optional" json:"inAccount" yaml:"inAccount"`
	// The fully qualified name of the database in which the inherited privilege will be granted on all schemas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/grant_privileges_to_account_role#in_database GrantPrivilegesToAccountRole#in_database}
	InDatabase *string `field:"optional" json:"inDatabase" yaml:"inDatabase"`
}

