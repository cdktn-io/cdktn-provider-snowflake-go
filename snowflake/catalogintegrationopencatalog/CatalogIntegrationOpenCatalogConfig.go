// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalogintegrationopencatalog

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CatalogIntegrationOpenCatalogConfig struct {
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
	// Specifies whether the catalog integration is available for use for Iceberg tables.
	//
	// `true` allows users to create new Iceberg tables that reference this integration. Existing Iceberg tables that reference this integration function normally. `false` prevents users from creating new Iceberg tables that reference this integration. Existing Iceberg tables that reference this integration cannot access the catalog in the table definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/catalog_integration_open_catalog#enabled CatalogIntegrationOpenCatalog#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the identifier (i.e. name) of the catalog integration; must be unique in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/catalog_integration_open_catalog#name CatalogIntegrationOpenCatalog#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// rest_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/catalog_integration_open_catalog#rest_authentication CatalogIntegrationOpenCatalog#rest_authentication}
	RestAuthentication *CatalogIntegrationOpenCatalogRestAuthentication `field:"required" json:"restAuthentication" yaml:"restAuthentication"`
	// rest_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/catalog_integration_open_catalog#rest_config CatalogIntegrationOpenCatalog#rest_config}
	RestConfig *CatalogIntegrationOpenCatalogRestConfig `field:"required" json:"restConfig" yaml:"restConfig"`
	// Specifies the default Open Catalog namespace for all Iceberg tables that you associate with the catalog integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/catalog_integration_open_catalog#catalog_namespace CatalogIntegrationOpenCatalog#catalog_namespace}
	CatalogNamespace *string `field:"optional" json:"catalogNamespace" yaml:"catalogNamespace"`
	// (Default: ``) Specifies a comment for the catalog integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/catalog_integration_open_catalog#comment CatalogIntegrationOpenCatalog#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/catalog_integration_open_catalog#id CatalogIntegrationOpenCatalog#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the number of seconds to wait between attempts to poll the external Iceberg catalog for metadata updates for automated refresh.
	//
	// For Delta-based tables, specifies the number of seconds to wait between attempts to poll your external cloud storage for new metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/catalog_integration_open_catalog#refresh_interval_seconds CatalogIntegrationOpenCatalog#refresh_interval_seconds}
	RefreshIntervalSeconds *float64 `field:"optional" json:"refreshIntervalSeconds" yaml:"refreshIntervalSeconds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/catalog_integration_open_catalog#timeouts CatalogIntegrationOpenCatalog#timeouts}
	Timeouts *CatalogIntegrationOpenCatalogTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

