// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalogintegrationopencatalog


type CatalogIntegrationOpenCatalogRestConfig struct {
	// Specifies the name of the catalog to use in Open Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/catalog_integration_open_catalog#catalog_name CatalogIntegrationOpenCatalog#catalog_name}
	CatalogName *string `field:"required" json:"catalogName" yaml:"catalogName"`
	// Specifies Open Catalog account URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/catalog_integration_open_catalog#catalog_uri CatalogIntegrationOpenCatalog#catalog_uri}
	CatalogUri *string `field:"required" json:"catalogUri" yaml:"catalogUri"`
	// Specifies the access delegation mode for accessing Iceberg table files in your external cloud storage.
	//
	// Valid values are (case-insensitive): `VENDED_CREDENTIALS` | `EXTERNAL_VOLUME_CREDENTIALS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/catalog_integration_open_catalog#access_delegation_mode CatalogIntegrationOpenCatalog#access_delegation_mode}
	AccessDelegationMode *string `field:"optional" json:"accessDelegationMode" yaml:"accessDelegationMode"`
	// Specifies how Snowflake connects to Open Catalog.
	//
	// Valid values are (case-insensitive): `PUBLIC` | `PRIVATE` | `AWS_API_GATEWAY` | `AWS_PRIVATE_API_GATEWAY` | `AWS_GLUE` | `AWS_PRIVATE_GLUE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/catalog_integration_open_catalog#catalog_api_type CatalogIntegrationOpenCatalog#catalog_api_type}
	CatalogApiType *string `field:"optional" json:"catalogApiType" yaml:"catalogApiType"`
}

