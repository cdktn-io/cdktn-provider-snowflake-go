// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package semanticview


type SemanticViewFacts struct {
	// Specifies a qualified name for the fact, including the table name and a unique identifier for the fact: `<table_alias>.<semantic_expression_name>`. Remember to wrap each part in double quotes like `"\"<table_alias>\".\"<semantic_expression_name>\""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/semantic_view#qualified_expression_name SemanticView#qualified_expression_name}
	QualifiedExpressionName *string `field:"required" json:"qualifiedExpressionName" yaml:"qualifiedExpressionName"`
	// The SQL expression used to compute the fact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/semantic_view#sql_expression SemanticView#sql_expression}
	SqlExpression *string `field:"required" json:"sqlExpression" yaml:"sqlExpression"`
	// Specifies a comment for the fact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/semantic_view#comment SemanticView#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// List of synonyms for the fact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/semantic_view#synonym SemanticView#synonym}
	Synonym *[]*string `field:"optional" json:"synonym" yaml:"synonym"`
}

