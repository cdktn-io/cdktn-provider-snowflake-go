// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package semanticview


type SemanticViewRelationships struct {
	// referenced_table_name_or_alias block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/semantic_view#referenced_table_name_or_alias SemanticView#referenced_table_name_or_alias}
	ReferencedTableNameOrAlias *SemanticViewRelationshipsReferencedTableNameOrAlias `field:"required" json:"referencedTableNameOrAlias" yaml:"referencedTableNameOrAlias"`
	// Specifies one or more columns in the first logical table that refers to columns in another logical table.
	//
	// Column names in this list are case-sensitive - the provider uses double quotes to wrap each of them when sending the SQL to Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/semantic_view#relationship_columns SemanticView#relationship_columns}
	RelationshipColumns *[]*string `field:"required" json:"relationshipColumns" yaml:"relationshipColumns"`
	// table_name_or_alias block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/semantic_view#table_name_or_alias SemanticView#table_name_or_alias}
	TableNameOrAlias *SemanticViewRelationshipsTableNameOrAlias `field:"required" json:"tableNameOrAlias" yaml:"tableNameOrAlias"`
	// Specifies one or more columns in the second logical table that are referred to by the first logical table.
	//
	// Column names in this list are case-sensitive - the provider uses double quotes to wrap each of them when sending the SQL to Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/semantic_view#referenced_relationship_columns SemanticView#referenced_relationship_columns}
	ReferencedRelationshipColumns *[]*string `field:"optional" json:"referencedRelationshipColumns" yaml:"referencedRelationshipColumns"`
	// Specifies an optional identifier for the relationship.
	//
	// This field is case-sensitive - the provider uses double quotes to wrap it when sending the SQL to Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/semantic_view#relationship_identifier SemanticView#relationship_identifier}
	RelationshipIdentifier *string `field:"optional" json:"relationshipIdentifier" yaml:"relationshipIdentifier"`
}

