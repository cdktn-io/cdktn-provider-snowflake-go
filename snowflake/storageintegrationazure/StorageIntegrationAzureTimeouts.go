// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storageintegrationazure


type StorageIntegrationAzureTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_azure#create StorageIntegrationAzure#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_azure#delete StorageIntegrationAzure#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_azure#read StorageIntegrationAzure#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_azure#update StorageIntegrationAzure#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

