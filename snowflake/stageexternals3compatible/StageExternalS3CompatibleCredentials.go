// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3compatible


type StageExternalS3CompatibleCredentials struct {
	// Specifies the AWS access key ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#aws_key_id StageExternalS3Compatible#aws_key_id}
	AwsKeyId *string `field:"required" json:"awsKeyId" yaml:"awsKeyId"`
	// Specifies the AWS secret access key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#aws_secret_key StageExternalS3Compatible#aws_secret_key}
	AwsSecretKey *string `field:"required" json:"awsSecretKey" yaml:"awsSecretKey"`
}

