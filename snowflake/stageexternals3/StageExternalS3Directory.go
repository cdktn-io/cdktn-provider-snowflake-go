// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3


type StageExternalS3Directory struct {
	// Specifies whether to enable a directory table on the external stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/stage_external_s3#enable StageExternalS3#enable}
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether Snowflake should enable triggering automatic refreshes of the directory table metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/stage_external_s3#auto_refresh StageExternalS3#auto_refresh}
	AutoRefresh *string `field:"optional" json:"autoRefresh" yaml:"autoRefresh"`
	// Specifies the AWS SNS topic ARN used for directory table auto-refresh notifications.
	//
	// Changing this field causes resource recreation (ForceNew). External change detection for this field is not yet supported and will be addressed in a future update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/stage_external_s3#aws_sns_topic StageExternalS3#aws_sns_topic}
	AwsSnsTopic *string `field:"optional" json:"awsSnsTopic" yaml:"awsSnsTopic"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether to automatically refresh the directory table metadata once, immediately after the stage is created.This field is used only when creating the object. Changes on this field are ignored after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/stage_external_s3#refresh_on_create StageExternalS3#refresh_on_create}
	RefreshOnCreate *string `field:"optional" json:"refreshOnCreate" yaml:"refreshOnCreate"`
}

