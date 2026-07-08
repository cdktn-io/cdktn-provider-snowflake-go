// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apiintegrationazureapimanagement


type ApiIntegrationAzureApiManagementTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_azure_api_management#create ApiIntegrationAzureApiManagement#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_azure_api_management#delete ApiIntegrationAzureApiManagement#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_azure_api_management#read ApiIntegrationAzureApiManagement#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_azure_api_management#update ApiIntegrationAzureApiManagement#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

