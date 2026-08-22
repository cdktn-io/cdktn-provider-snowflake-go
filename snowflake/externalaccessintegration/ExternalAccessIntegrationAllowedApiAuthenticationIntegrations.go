// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package externalaccessintegration


type ExternalAccessIntegrationAllowedApiAuthenticationIntegrations struct {
	// Specifies the API authentication integrations allowed for authenticating to external locations. Conflicts with `none`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/external_access_integration#integrations ExternalAccessIntegration#integrations}
	Integrations *[]*string `field:"optional" json:"integrations" yaml:"integrations"`
	// When true, no API authentication integrations are allowed. Conflicts with `integrations`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/external_access_integration#none ExternalAccessIntegration#none}
	None interface{} `field:"optional" json:"none" yaml:"none"`
}

