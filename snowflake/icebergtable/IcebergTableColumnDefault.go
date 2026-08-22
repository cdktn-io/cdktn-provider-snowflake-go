// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable


type IcebergTableColumnDefault struct {
	// The default expression value for the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#expression IcebergTable#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

