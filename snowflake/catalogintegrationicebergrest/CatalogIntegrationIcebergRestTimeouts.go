// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalogintegrationicebergrest


type CatalogIntegrationIcebergRestTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_iceberg_rest#create CatalogIntegrationIcebergRest#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_iceberg_rest#delete CatalogIntegrationIcebergRest#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_iceberg_rest#read CatalogIntegrationIcebergRest#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_iceberg_rest#update CatalogIntegrationIcebergRest#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

