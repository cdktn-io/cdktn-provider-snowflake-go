// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package semanticview


type SemanticViewTablesUnique struct {
	// Unique key combinations in the logical table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/semantic_view#values SemanticView#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

