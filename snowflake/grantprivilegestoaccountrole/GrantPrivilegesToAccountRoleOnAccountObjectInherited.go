// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoaccountrole


type GrantPrivilegesToAccountRoleOnAccountObjectInherited struct {
	// The plural object type of the account object on which an inherited privilege will be granted.
	//
	// For more information head over to [Snowflake documentation](https://docs.snowflake.com/en/sql-reference/sql/grant-privilege#required-parameters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/grant_privileges_to_account_role#object_type_plural GrantPrivilegesToAccountRole#object_type_plural}
	ObjectTypePlural *string `field:"required" json:"objectTypePlural" yaml:"objectTypePlural"`
}

