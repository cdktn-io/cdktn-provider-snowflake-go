// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package legacyserviceuser


type LegacyServiceUserDefaultWorkloadIdentity struct {
	// aws block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/legacy_service_user#aws LegacyServiceUser#aws}
	Aws *LegacyServiceUserDefaultWorkloadIdentityAws `field:"optional" json:"aws" yaml:"aws"`
	// azure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/legacy_service_user#azure LegacyServiceUser#azure}
	Azure *LegacyServiceUserDefaultWorkloadIdentityAzure `field:"optional" json:"azure" yaml:"azure"`
	// gcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/legacy_service_user#gcp LegacyServiceUser#gcp}
	Gcp *LegacyServiceUserDefaultWorkloadIdentityGcp `field:"optional" json:"gcp" yaml:"gcp"`
	// oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/legacy_service_user#oidc LegacyServiceUser#oidc}
	Oidc *LegacyServiceUserDefaultWorkloadIdentityOidc `field:"optional" json:"oidc" yaml:"oidc"`
}

