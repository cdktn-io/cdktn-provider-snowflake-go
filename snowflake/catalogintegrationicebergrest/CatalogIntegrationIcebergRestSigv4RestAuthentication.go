// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalogintegrationicebergrest


type CatalogIntegrationIcebergRestSigv4RestAuthentication struct {
	// Specifies the Amazon Resource Name (ARN) for an IAM role that has permission to access your REST API in API Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/catalog_integration_iceberg_rest#sigv4_iam_role CatalogIntegrationIcebergRest#sigv4_iam_role}
	Sigv4IamRole *string `field:"required" json:"sigv4IamRole" yaml:"sigv4IamRole"`
	// Specifies an external ID that Snowflake uses to establish a trust relationship with AWS.
	//
	// If you don’t specify this parameter, Snowflake automatically generates a unique external ID when you create a catalog integration. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/catalog_integration_iceberg_rest#sigv4_external_id CatalogIntegrationIcebergRest#sigv4_external_id}
	Sigv4ExternalId *string `field:"optional" json:"sigv4ExternalId" yaml:"sigv4ExternalId"`
	// Specifies the AWS Region associated with your API in API Gateway.
	//
	// If you don’t specify this parameter, Snowflake uses the region in which your Snowflake account is deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/catalog_integration_iceberg_rest#sigv4_signing_region CatalogIntegrationIcebergRest#sigv4_signing_region}
	Sigv4SigningRegion *string `field:"optional" json:"sigv4SigningRegion" yaml:"sigv4SigningRegion"`
}

