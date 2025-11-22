// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package semanticview


type SemanticViewMetrics struct {
	// semantic_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/semantic_view#semantic_expression SemanticView#semantic_expression}
	SemanticExpression *SemanticViewMetricsSemanticExpression `field:"optional" json:"semanticExpression" yaml:"semanticExpression"`
	// window_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/semantic_view#window_function SemanticView#window_function}
	WindowFunction *SemanticViewMetricsWindowFunction `field:"optional" json:"windowFunction" yaml:"windowFunction"`
}

