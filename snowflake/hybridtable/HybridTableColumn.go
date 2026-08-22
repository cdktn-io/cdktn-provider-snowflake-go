// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hybridtable


type HybridTableColumn struct {
	// Column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#name HybridTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Column type. See [Snowflake data types](https://docs.snowflake.com/en/sql-reference-data-types) for supported values. Example: VARCHAR(256), NUMBER(38,0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#type HybridTable#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Column collation specification, e.g. en-ci. Case-insensitive (en-ci and EN-CI are treated as equal).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#collate HybridTable#collate}
	Collate *string `field:"optional" json:"collate" yaml:"collate"`
	// Column-level comment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#comment HybridTable#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// default block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#default HybridTable#default}
	Default *HybridTableColumnDefault `field:"optional" json:"default" yaml:"default"`
	// Whether to restrict the column to NOT NULL values.
	//
	// Changing this on an existing column forces recreation. Primary key columns must set this to true because NOT NULL is implied by the primary key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/hybrid_table#not_null HybridTable#not_null}
	NotNull interface{} `field:"optional" json:"notNull" yaml:"notNull"`
}

