// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package externalaccessintegration


type ExternalAccessIntegrationAllowedAuthenticationSecrets struct {
	// When true, all secrets in the account are allowed for authentication. Conflicts with `none` and `secrets`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/external_access_integration#all ExternalAccessIntegration#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// When true, no secrets are allowed for authentication. Conflicts with `all` and `secrets`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/external_access_integration#none ExternalAccessIntegration#none}
	None interface{} `field:"optional" json:"none" yaml:"none"`
	// Specifies the fully qualified identifiers of secrets allowed for authentication. Conflicts with `none` and `all`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/external_access_integration#secrets ExternalAccessIntegration#secrets}
	Secrets *[]*string `field:"optional" json:"secrets" yaml:"secrets"`
}

