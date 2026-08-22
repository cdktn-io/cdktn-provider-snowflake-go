// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hybridtable


type HybridTableForeignKeyConstraint struct {
	// The local column(s) the foreign key is defined on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#columns HybridTable#columns}
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
	// The column(s) in the referenced table that the foreign key references.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#ref_columns HybridTable#ref_columns}
	RefColumns *[]*string `field:"required" json:"refColumns" yaml:"refColumns"`
	// The table that the foreign key references.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#table_name HybridTable#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Name of the constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#name HybridTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

