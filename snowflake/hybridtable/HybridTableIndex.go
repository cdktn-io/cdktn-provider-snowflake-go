// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hybridtable


type HybridTableIndex struct {
	// Index key columns, in order. Order is semantically meaningful.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#columns HybridTable#columns}
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
	// Name of the secondary index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#name HybridTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Columns included in the index payload via INCLUDE (...). Order carries no meaning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#include_columns HybridTable#include_columns}
	IncludeColumns *[]*string `field:"optional" json:"includeColumns" yaml:"includeColumns"`
}

