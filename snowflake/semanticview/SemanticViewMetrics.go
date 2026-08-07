// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package semanticview


type SemanticViewMetrics struct {
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether the metric is private.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/semantic_view#is_private SemanticView#is_private}
	IsPrivate *string `field:"optional" json:"isPrivate" yaml:"isPrivate"`
	// semantic_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/semantic_view#semantic_expression SemanticView#semantic_expression}
	SemanticExpression *SemanticViewMetricsSemanticExpression `field:"optional" json:"semanticExpression" yaml:"semanticExpression"`
	// window_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/semantic_view#window_function SemanticView#window_function}
	WindowFunction *SemanticViewMetricsWindowFunction `field:"optional" json:"windowFunction" yaml:"windowFunction"`
}

