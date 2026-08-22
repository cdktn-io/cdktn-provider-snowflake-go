// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hybridtable


type HybridTableUniqueConstraint struct {
	// The column(s) the constraint applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#columns HybridTable#columns}
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
	// Name of the constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#name HybridTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

