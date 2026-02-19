// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package serviceuser


type ServiceUserDefaultWorkloadIdentityAws struct {
	// The ARN of the AWS IAM role to use for workload identity federation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/service_user#arn ServiceUser#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

