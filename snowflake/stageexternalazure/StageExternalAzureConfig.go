// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalazure

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageExternalAzureConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#database StageExternalAzure#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the stage;
	//
	// must be unique for the database and schema in which the stage is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#name StageExternalAzure#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the stage.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#schema StageExternalAzure#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the URL for the Azure storage container (e.g., 'azure://account.blob.core.windows.net/container').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#url StageExternalAzure#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Specifies a comment for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#comment StageExternalAzure#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#credentials StageExternalAzure#credentials}
	Credentials *StageExternalAzureCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#directory StageExternalAzure#directory}
	Directory *StageExternalAzureDirectory `field:"optional" json:"directory" yaml:"directory"`
	// encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#encryption StageExternalAzure#encryption}
	Encryption *StageExternalAzureEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// file_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#file_format StageExternalAzure#file_format}
	FileFormat *StageExternalAzureFileFormat `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#id StageExternalAzure#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the name of the storage integration used to delegate authentication responsibility to a Snowflake identity.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#storage_integration StageExternalAzure#storage_integration}
	StorageIntegration *string `field:"optional" json:"storageIntegration" yaml:"storageIntegration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#timeouts StageExternalAzure#timeouts}
	Timeouts *StageExternalAzureTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether to use a private link endpoint for Azure storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#use_privatelink_endpoint StageExternalAzure#use_privatelink_endpoint}
	UsePrivatelinkEndpoint *string `field:"optional" json:"usePrivatelinkEndpoint" yaml:"usePrivatelinkEndpoint"`
}

