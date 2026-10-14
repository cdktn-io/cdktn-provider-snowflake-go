// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestodatabaserole


type GrantPrivilegesToDatabaseRoleOnSchema struct {
	// The fully qualified name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/grant_privileges_to_database_role#all_schemas_in_database GrantPrivilegesToDatabaseRole#all_schemas_in_database}
	AllSchemasInDatabase *string `field:"optional" json:"allSchemasInDatabase" yaml:"allSchemasInDatabase"`
	// The fully qualified name of the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/grant_privileges_to_database_role#future_schemas_in_database GrantPrivilegesToDatabaseRole#future_schemas_in_database}
	FutureSchemasInDatabase *string `field:"optional" json:"futureSchemasInDatabase" yaml:"futureSchemasInDatabase"`
	// Configures an inherited privilege to be granted on all current and future schemas in a database.
	//
	// See [Inherited grants](https://docs.snowflake.com/en/user-guide/inherited-grants-using) for more details. This field is enabled by default as part of the `INHERITED_GRANTS` experimental feature. You can opt out by listing `INHERITED_GRANTS` in [`experimental_features_disabled`](../#experimental_features_disabled-1) in the provider block. This experiment will be removed in v2.23.0 or v2.24.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/grant_privileges_to_database_role#inherited GrantPrivilegesToDatabaseRole#inherited}
	Inherited *string `field:"optional" json:"inherited" yaml:"inherited"`
	// The fully qualified name of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/grant_privileges_to_database_role#schema_name GrantPrivilegesToDatabaseRole#schema_name}
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
}

