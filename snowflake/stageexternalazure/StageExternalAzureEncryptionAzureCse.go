// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalazure


type StageExternalAzureEncryptionAzureCse struct {
	// Specifies the 128-bit or 256-bit client-side master key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#master_key StageExternalAzure#master_key}
	MasterKey *string `field:"required" json:"masterKey" yaml:"masterKey"`
}

