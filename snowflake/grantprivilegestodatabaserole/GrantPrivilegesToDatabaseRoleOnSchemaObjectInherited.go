// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestodatabaserole


type GrantPrivilegesToDatabaseRoleOnSchemaObjectInherited struct {
	// The plural object type of the schema object on which privileges will be granted.
	//
	// For more information head over to [Snowflake documentation](https://docs.snowflake.com/en/sql-reference/sql/grant-privilege#required-parameters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/grant_privileges_to_database_role#object_type_plural GrantPrivilegesToDatabaseRole#object_type_plural}
	ObjectTypePlural *string `field:"required" json:"objectTypePlural" yaml:"objectTypePlural"`
	// The fully qualified name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/grant_privileges_to_database_role#in_database GrantPrivilegesToDatabaseRole#in_database}
	InDatabase *string `field:"optional" json:"inDatabase" yaml:"inDatabase"`
	// The fully qualified name of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/grant_privileges_to_database_role#in_schema GrantPrivilegesToDatabaseRole#in_schema}
	InSchema *string `field:"optional" json:"inSchema" yaml:"inSchema"`
}

