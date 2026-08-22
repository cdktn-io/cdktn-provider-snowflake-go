// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable


type IcebergTableColumnMaskingPolicy struct {
	// Masking policy name. For more information about this resource, see [docs](./masking_policy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#policy_name IcebergTable#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Specifies the arguments to pass into the conditional masking policy SQL expression, in order.
	//
	// The first column in the list specifies the column for the policy conditions to mask or tokenize the data and must match the column to which the masking policy is set. The additional columns specify the columns to evaluate to determine whether to mask or tokenize the data in each row of the query result when a query is made on the first column. If the USING clause is omitted, Snowflake treats the conditional masking policy as a normal masking policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#using IcebergTable#using}
	Using *[]*string `field:"optional" json:"using" yaml:"using"`
}

