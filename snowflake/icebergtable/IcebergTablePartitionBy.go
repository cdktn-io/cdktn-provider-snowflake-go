// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable


type IcebergTablePartitionBy struct {
	// bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#bucket IcebergTable#bucket}
	Bucket *IcebergTablePartitionByBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// Partitions the table by the day component of the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#day IcebergTable#day}
	Day *string `field:"optional" json:"day" yaml:"day"`
	// Partitions the table by the hour component of the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#hour IcebergTable#hour}
	Hour *string `field:"optional" json:"hour" yaml:"hour"`
	// Name of the column to use as-is for partitioning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#identity IcebergTable#identity}
	Identity *string `field:"optional" json:"identity" yaml:"identity"`
	// Partitions the table by the month component of the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#month IcebergTable#month}
	Month *string `field:"optional" json:"month" yaml:"month"`
	// truncate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#truncate IcebergTable#truncate}
	Truncate *IcebergTablePartitionByTruncate `field:"optional" json:"truncate" yaml:"truncate"`
	// Partitions the table by the year component of the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#year IcebergTable#year}
	Year *string `field:"optional" json:"year" yaml:"year"`
}

