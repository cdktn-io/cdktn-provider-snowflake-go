// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package serviceuser


type ServiceUserDefaultWorkloadIdentityAws struct {
	// The ARN of the AWS IAM role to use for workload identity federation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/service_user#arn ServiceUser#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The AWS issuer URL. Required for JWT-based (GetWebIdentityToken) workload identity federation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/service_user#issuer ServiceUser#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
}

