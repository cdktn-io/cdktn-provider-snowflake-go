// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalazure


type StageExternalAzureEncryption struct {
	// azure_cse block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#azure_cse StageExternalAzure#azure_cse}
	AzureCse *StageExternalAzureEncryptionAzureCse `field:"optional" json:"azureCse" yaml:"azureCse"`
	// none block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#none StageExternalAzure#none}
	None *StageExternalAzureEncryptionNone `field:"optional" json:"none" yaml:"none"`
}

