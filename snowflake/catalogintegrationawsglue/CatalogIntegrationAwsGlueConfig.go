// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalogintegrationawsglue

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CatalogIntegrationAwsGlueConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_aws_glue#enabled CatalogIntegrationAwsGlue#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role to assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_aws_glue#glue_aws_role_arn CatalogIntegrationAwsGlue#glue_aws_role_arn}
	GlueAwsRoleArn *string `field:"required" json:"glueAwsRoleArn" yaml:"glueAwsRoleArn"`
	// Specifies the ID of your AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_aws_glue#glue_catalog_id CatalogIntegrationAwsGlue#glue_catalog_id}
	GlueCatalogId *string `field:"required" json:"glueCatalogId" yaml:"glueCatalogId"`
	// Specifies the identifier (i.e. name) of the catalog integration; must be unique in your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_aws_glue#name CatalogIntegrationAwsGlue#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the default AWS Glue Data Catalog namespace for all Iceberg tables that you associate with the catalog integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_aws_glue#catalog_namespace CatalogIntegrationAwsGlue#catalog_namespace}
	CatalogNamespace *string `field:"optional" json:"catalogNamespace" yaml:"catalogNamespace"`
	// (Default: ``) Specifies a comment for the catalog integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_aws_glue#comment CatalogIntegrationAwsGlue#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the AWS region of your AWS Glue Data Catalog.
	//
	// You must specify a value for this attribute if your Snowflake account is not hosted on AWS. Otherwise, the default region is the Snowflake deployment region for the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_aws_glue#glue_region CatalogIntegrationAwsGlue#glue_region}
	GlueRegion *string `field:"optional" json:"glueRegion" yaml:"glueRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_aws_glue#id CatalogIntegrationAwsGlue#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the number of seconds to wait between attempts to poll the external Iceberg catalog for metadata updates for automated refresh.
	//
	// For Delta-based tables, specifies the number of seconds to wait between attempts to poll your external cloud storage for new metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_aws_glue#refresh_interval_seconds CatalogIntegrationAwsGlue#refresh_interval_seconds}
	RefreshIntervalSeconds *float64 `field:"optional" json:"refreshIntervalSeconds" yaml:"refreshIntervalSeconds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_aws_glue#timeouts CatalogIntegrationAwsGlue#timeouts}
	Timeouts *CatalogIntegrationAwsGlueTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

