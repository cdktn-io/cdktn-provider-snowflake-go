// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable


type IcebergTableRowAccessPolicy struct {
	// Defines which columns are affected by the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#on IcebergTable#on}
	On *[]*string `field:"required" json:"on" yaml:"on"`
	// Row access policy name. For more information about this resource, see [docs](./row_access_policy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#policy_name IcebergTable#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

