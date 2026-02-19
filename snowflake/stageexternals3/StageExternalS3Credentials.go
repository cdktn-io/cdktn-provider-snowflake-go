// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3


type StageExternalS3Credentials struct {
	// Specifies the AWS access key ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#aws_key_id StageExternalS3#aws_key_id}
	AwsKeyId *string `field:"optional" json:"awsKeyId" yaml:"awsKeyId"`
	// Specifies the AWS IAM role ARN to use for accessing the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#aws_role StageExternalS3#aws_role}
	AwsRole *string `field:"optional" json:"awsRole" yaml:"awsRole"`
	// Specifies the AWS secret access key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#aws_secret_key StageExternalS3#aws_secret_key}
	AwsSecretKey *string `field:"optional" json:"awsSecretKey" yaml:"awsSecretKey"`
	// Specifies the AWS session token for temporary credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#aws_token StageExternalS3#aws_token}
	AwsToken *string `field:"optional" json:"awsToken" yaml:"awsToken"`
}

