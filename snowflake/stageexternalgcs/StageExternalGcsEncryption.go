// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalgcs


type StageExternalGcsEncryption struct {
	// gcs_sse_kms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_gcs#gcs_sse_kms StageExternalGcs#gcs_sse_kms}
	GcsSseKms *StageExternalGcsEncryptionGcsSseKms `field:"optional" json:"gcsSseKms" yaml:"gcsSseKms"`
	// none block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_gcs#none StageExternalGcs#none}
	None *StageExternalGcsEncryptionNone `field:"optional" json:"none" yaml:"none"`
}

