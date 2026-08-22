// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hybridtable


type HybridTableColumnDefault struct {
	// A constant default value for the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#constant HybridTable#constant}
	Constant *string `field:"optional" json:"constant" yaml:"constant"`
	// A SQL expression default value for the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#expression HybridTable#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The default sequence for the column (uses NEXTVAL).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#sequence HybridTable#sequence}
	Sequence *string `field:"optional" json:"sequence" yaml:"sequence"`
}

