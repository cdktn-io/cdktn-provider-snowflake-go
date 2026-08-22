// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable


type IcebergTablePartitionByBucket struct {
	// Name of the column to bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#column IcebergTable#column}
	Column *string `field:"required" json:"column" yaml:"column"`
	// Number of buckets to hash the column values into.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#num_buckets IcebergTable#num_buckets}
	NumBuckets *float64 `field:"required" json:"numBuckets" yaml:"numBuckets"`
}

