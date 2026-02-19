// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3compatible

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageExternalS3CompatibleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#database StageExternalS3Compatible#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the endpoint for the S3-compatible storage provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#endpoint StageExternalS3Compatible#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Specifies the identifier for the stage;
	//
	// must be unique for the database and schema in which the stage is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#name StageExternalS3Compatible#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the stage.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#schema StageExternalS3Compatible#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the URL for the S3-compatible storage location (e.g., 's3compat://bucket/path/').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#url StageExternalS3Compatible#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Specifies a comment for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#comment StageExternalS3Compatible#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#credentials StageExternalS3Compatible#credentials}
	Credentials *StageExternalS3CompatibleCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#directory StageExternalS3Compatible#directory}
	Directory *StageExternalS3CompatibleDirectory `field:"optional" json:"directory" yaml:"directory"`
	// file_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#file_format StageExternalS3Compatible#file_format}
	FileFormat *StageExternalS3CompatibleFileFormat `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#id StageExternalS3Compatible#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#timeouts StageExternalS3Compatible#timeouts}
	Timeouts *StageExternalS3CompatibleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

