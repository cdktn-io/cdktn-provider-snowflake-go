// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalogintegrationicebergrest


type CatalogIntegrationIcebergRestRestConfig struct {
	// Specifies the endpoint URL for the catalog REST API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.15.0/docs/resources/catalog_integration_iceberg_rest#catalog_uri CatalogIntegrationIcebergRest#catalog_uri}
	CatalogUri *string `field:"required" json:"catalogUri" yaml:"catalogUri"`
	// Specifies the access delegation mode for accessing Iceberg table files in your external cloud storage.
	//
	// Valid values are (case-insensitive): `VENDED_CREDENTIALS` | `EXTERNAL_VOLUME_CREDENTIALS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.15.0/docs/resources/catalog_integration_iceberg_rest#access_delegation_mode CatalogIntegrationIcebergRest#access_delegation_mode}
	AccessDelegationMode *string `field:"optional" json:"accessDelegationMode" yaml:"accessDelegationMode"`
	// Specifies the connection type for the catalog API.
	//
	// Valid values are (case-insensitive): `PUBLIC` | `PRIVATE` | `AWS_API_GATEWAY` | `AWS_PRIVATE_API_GATEWAY` | `AWS_GLUE` | `AWS_PRIVATE_GLUE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.15.0/docs/resources/catalog_integration_iceberg_rest#catalog_api_type CatalogIntegrationIcebergRest#catalog_api_type}
	CatalogApiType *string `field:"optional" json:"catalogApiType" yaml:"catalogApiType"`
	// Specifies the catalog or identifier to request from your remote catalog service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.15.0/docs/resources/catalog_integration_iceberg_rest#catalog_name CatalogIntegrationIcebergRest#catalog_name}
	CatalogName *string `field:"optional" json:"catalogName" yaml:"catalogName"`
	// Specifies an optional prefix appended to all API routes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.15.0/docs/resources/catalog_integration_iceberg_rest#prefix CatalogIntegrationIcebergRest#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

