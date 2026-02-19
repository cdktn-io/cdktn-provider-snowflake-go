// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storageintegrationazure

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StorageIntegrationAzureConfig struct {
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
	// Specifies the ID for your Office 365 tenant that the allowed and blocked storage accounts belong to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_azure#azure_tenant_id StorageIntegrationAzure#azure_tenant_id}
	AzureTenantId *string `field:"required" json:"azureTenantId" yaml:"azureTenantId"`
	// Specifies whether this storage integration is available for usage in stages.
	//
	// `TRUE` allows users to create new stages that reference this integration. Existing stages that reference this integration function normally. `FALSE` prevents users from creating new stages that reference this integration. Existing stages that reference this integration cannot access the storage location in the stage definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_azure#enabled StorageIntegrationAzure#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// String that specifies the identifier (i.e. name) for the integration; must be unique in your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_azure#name StorageIntegrationAzure#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Explicitly limits external stages that use the integration to reference one or more storage locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_azure#storage_allowed_locations StorageIntegrationAzure#storage_allowed_locations}
	StorageAllowedLocations *[]*string `field:"required" json:"storageAllowedLocations" yaml:"storageAllowedLocations"`
	// Specifies a comment for the storage integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_azure#comment StorageIntegrationAzure#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_azure#id StorageIntegrationAzure#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Explicitly prohibits external stages that use the integration from referencing one or more storage locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_azure#storage_blocked_locations StorageIntegrationAzure#storage_blocked_locations}
	StorageBlockedLocations *[]*string `field:"optional" json:"storageBlockedLocations" yaml:"storageBlockedLocations"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_azure#timeouts StorageIntegrationAzure#timeouts}
	Timeouts *StorageIntegrationAzureTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether to use outbound private connectivity to harden the security posture.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_azure#use_privatelink_endpoint StorageIntegrationAzure#use_privatelink_endpoint}
	UsePrivatelinkEndpoint *string `field:"optional" json:"usePrivatelinkEndpoint" yaml:"usePrivatelinkEndpoint"`
}

