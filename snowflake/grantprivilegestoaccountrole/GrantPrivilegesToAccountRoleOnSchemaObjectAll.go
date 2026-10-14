// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoaccountrole


type GrantPrivilegesToAccountRoleOnSchemaObjectAll struct {
	// The plural object type of the schema object on which privileges will be granted.
	//
	// For more information head over to [Snowflake documentation](https://docs.snowflake.com/en/sql-reference/sql/grant-privilege#required-parameters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/grant_privileges_to_account_role#object_type_plural GrantPrivilegesToAccountRole#object_type_plural}
	ObjectTypePlural *string `field:"required" json:"objectTypePlural" yaml:"objectTypePlural"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/grant_privileges_to_account_role#in_database GrantPrivilegesToAccountRole#in_database}.
	InDatabase *string `field:"optional" json:"inDatabase" yaml:"inDatabase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/grant_privileges_to_account_role#in_schema GrantPrivilegesToAccountRole#in_schema}.
	InSchema *string `field:"optional" json:"inSchema" yaml:"inSchema"`
}

