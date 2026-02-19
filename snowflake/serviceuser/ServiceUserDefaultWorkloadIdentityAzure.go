// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package serviceuser


type ServiceUserDefaultWorkloadIdentityAzure struct {
	// The Azure issuer URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/service_user#issuer ServiceUser#issuer}
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The Azure subject identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/service_user#subject ServiceUser#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

