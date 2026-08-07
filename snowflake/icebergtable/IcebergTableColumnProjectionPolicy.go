// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable


type IcebergTableColumnProjectionPolicy struct {
	// Projection policy name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#policy_name IcebergTable#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

