// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageExternalS3Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The database in which to create the stage.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#database StageExternalS3#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the stage;
	//
	// must be unique for the database and schema in which the stage is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#name StageExternalS3#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the stage.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#schema StageExternalS3#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the URL for the S3 bucket (e.g., 's3://bucket-name/path/').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#url StageExternalS3#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Specifies the ARN for an AWS S3 Access Point to use for data transfer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#aws_access_point_arn StageExternalS3#aws_access_point_arn}
	AwsAccessPointArn *string `field:"optional" json:"awsAccessPointArn" yaml:"awsAccessPointArn"`
	// Specifies a comment for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#comment StageExternalS3#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#credentials StageExternalS3#credentials}
	Credentials *StageExternalS3Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#directory StageExternalS3#directory}
	Directory *StageExternalS3Directory `field:"optional" json:"directory" yaml:"directory"`
	// encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#encryption StageExternalS3#encryption}
	Encryption *StageExternalS3Encryption `field:"optional" json:"encryption" yaml:"encryption"`
	// file_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#file_format StageExternalS3#file_format}
	FileFormat *StageExternalS3FileFormat `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#id StageExternalS3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the name of the storage integration used to delegate authentication responsibility to a Snowflake identity.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#storage_integration StageExternalS3#storage_integration}
	StorageIntegration *string `field:"optional" json:"storageIntegration" yaml:"storageIntegration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#timeouts StageExternalS3#timeouts}
	Timeouts *StageExternalS3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether to use a private link endpoint for S3 storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#use_privatelink_endpoint StageExternalS3#use_privatelink_endpoint}
	UsePrivatelinkEndpoint *string `field:"optional" json:"usePrivatelinkEndpoint" yaml:"usePrivatelinkEndpoint"`
}

