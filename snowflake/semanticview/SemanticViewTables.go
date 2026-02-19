// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package semanticview


type SemanticViewTables struct {
	// Specifies an alias for a logical table in the semantic view.
	//
	// This field is case-sensitive - the provider uses double quotes to wrap it when sending the SQL to Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/semantic_view#table_alias SemanticView#table_alias}
	TableAlias *string `field:"required" json:"tableAlias" yaml:"tableAlias"`
	// Specifies an identifier for the logical table.
	//
	// Example: `"\"<db_name>\".\"<schema_name>\".\"<table_name>\""`. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/semantic_view#table_name SemanticView#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Specifies a comment for the logical table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/semantic_view#comment SemanticView#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Definitions of primary keys in the logical table.
	//
	// This field is case-sensitive - the provider uses double quotes to wrap it when sending the SQL to Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/semantic_view#primary_key SemanticView#primary_key}
	PrimaryKey *[]*string `field:"optional" json:"primaryKey" yaml:"primaryKey"`
	// List of synonyms for the logical table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/semantic_view#synonym SemanticView#synonym}
	Synonym *[]*string `field:"optional" json:"synonym" yaml:"synonym"`
	// unique block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/semantic_view#unique SemanticView#unique}
	Unique interface{} `field:"optional" json:"unique" yaml:"unique"`
}

