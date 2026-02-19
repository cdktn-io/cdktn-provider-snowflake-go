// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package serviceuser


type ServiceUserDefaultWorkloadIdentity struct {
	// aws block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/service_user#aws ServiceUser#aws}
	Aws *ServiceUserDefaultWorkloadIdentityAws `field:"optional" json:"aws" yaml:"aws"`
	// azure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/service_user#azure ServiceUser#azure}
	Azure *ServiceUserDefaultWorkloadIdentityAzure `field:"optional" json:"azure" yaml:"azure"`
	// gcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/service_user#gcp ServiceUser#gcp}
	Gcp *ServiceUserDefaultWorkloadIdentityGcp `field:"optional" json:"gcp" yaml:"gcp"`
	// oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/service_user#oidc ServiceUser#oidc}
	Oidc *ServiceUserDefaultWorkloadIdentityOidc `field:"optional" json:"oidc" yaml:"oidc"`
}

