// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable


type IcebergTableColumn struct {
	// Column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#name IcebergTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Column type, e.g. VARIANT. For a full list of column types, see [Summary of Data Types](https://docs.snowflake.com/en/sql-reference/intro-summary-data-types).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#type IcebergTable#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Column comment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#comment IcebergTable#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// default block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#default IcebergTable#default}
	Default *IcebergTableColumnDefault `field:"optional" json:"default" yaml:"default"`
	// masking_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#masking_policy IcebergTable#masking_policy}
	MaskingPolicy *IcebergTableColumnMaskingPolicy `field:"optional" json:"maskingPolicy" yaml:"maskingPolicy"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Whether to restrict the column to NOT NULL values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#not_null IcebergTable#not_null}
	NotNull *string `field:"optional" json:"notNull" yaml:"notNull"`
	// projection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#projection_policy IcebergTable#projection_policy}
	ProjectionPolicy *IcebergTableColumnProjectionPolicy `field:"optional" json:"projectionPolicy" yaml:"projectionPolicy"`
}

