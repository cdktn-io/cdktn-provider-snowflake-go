// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grantownership


type GrantOwnershipOnFuture struct {
	// Specifies the type of object in plural form on which you are transferring ownership.
	//
	// For more information head over to [Snowflake documentation](https://docs.snowflake.com/en/sql-reference/sql/grant-ownership#required-parameters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/grant_ownership#object_type_plural GrantOwnership#object_type_plural}
	ObjectTypePlural *string `field:"required" json:"objectTypePlural" yaml:"objectTypePlural"`
	// The fully qualified name of the database. For more information about this resource, see [docs](./database).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/grant_ownership#in_database GrantOwnership#in_database}
	InDatabase *string `field:"optional" json:"inDatabase" yaml:"inDatabase"`
	// The fully qualified name of the schema. For more information about this resource, see [docs](./schema).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/grant_ownership#in_schema GrantOwnership#in_schema}
	InSchema *string `field:"optional" json:"inSchema" yaml:"inSchema"`
}

