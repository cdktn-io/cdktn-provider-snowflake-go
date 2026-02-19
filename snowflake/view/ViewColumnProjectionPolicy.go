// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package view


type ViewColumnProjectionPolicy struct {
	// Specifies the projection policy to set on a column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/view#policy_name View#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
}

