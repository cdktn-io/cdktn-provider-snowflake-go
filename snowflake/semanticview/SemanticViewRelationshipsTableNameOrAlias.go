// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package semanticview


type SemanticViewRelationshipsTableNameOrAlias struct {
	// The alias used for the logical table, cannot be used in combination with the `table_name`.
	//
	// This field is case-sensitive - the provider uses double quotes to wrap it when sending the SQL to Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/semantic_view#table_alias SemanticView#table_alias}
	TableAlias *string `field:"optional" json:"tableAlias" yaml:"tableAlias"`
	// The name of the logical table, cannot be used in combination with the `table_alias`.
	//
	// This field is case-sensitive - the provider uses double quotes to wrap it when sending the SQL to Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/semantic_view#table_name SemanticView#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

