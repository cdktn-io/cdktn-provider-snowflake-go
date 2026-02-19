// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalazure


type StageExternalAzureCredentials struct {
	// Specifies the shared access signature (SAS) token for Azure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#azure_sas_token StageExternalAzure#azure_sas_token}
	AzureSasToken *string `field:"required" json:"azureSasToken" yaml:"azureSasToken"`
}

