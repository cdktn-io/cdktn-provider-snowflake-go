// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable


type IcebergTablePartitionByTruncate struct {
	// Name of the column to truncate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#column IcebergTable#column}
	Column *string `field:"required" json:"column" yaml:"column"`
	// Width to truncate the column value to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#width IcebergTable#width}
	Width *float64 `field:"required" json:"width" yaml:"width"`
}

