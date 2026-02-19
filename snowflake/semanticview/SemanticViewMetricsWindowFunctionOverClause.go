// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package semanticview


type SemanticViewMetricsWindowFunctionOverClause struct {
	// Specifies an order by clause.
	//
	// It must be a complete SQL expression, including any `[ ASC | DESC ] [ NULLS { FIRST | LAST } ]` modifiers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/semantic_view#order_by SemanticView#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
	// Specifies a partition by clause.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/semantic_view#partition_by SemanticView#partition_by}
	PartitionBy *string `field:"optional" json:"partitionBy" yaml:"partitionBy"`
	// Specifies a window frame clause.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/semantic_view#window_frame_clause SemanticView#window_frame_clause}
	WindowFrameClause *string `field:"optional" json:"windowFrameClause" yaml:"windowFrameClause"`
}

