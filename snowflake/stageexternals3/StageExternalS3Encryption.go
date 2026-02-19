// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3


type StageExternalS3Encryption struct {
	// aws_cse block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#aws_cse StageExternalS3#aws_cse}
	AwsCse *StageExternalS3EncryptionAwsCse `field:"optional" json:"awsCse" yaml:"awsCse"`
	// aws_sse_kms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#aws_sse_kms StageExternalS3#aws_sse_kms}
	AwsSseKms *StageExternalS3EncryptionAwsSseKms `field:"optional" json:"awsSseKms" yaml:"awsSseKms"`
	// aws_sse_s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#aws_sse_s3 StageExternalS3#aws_sse_s3}
	AwsSseS3 *StageExternalS3EncryptionAwsSseS3 `field:"optional" json:"awsSseS3" yaml:"awsSseS3"`
	// none block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#none StageExternalS3#none}
	None *StageExternalS3EncryptionNone `field:"optional" json:"none" yaml:"none"`
}

