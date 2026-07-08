// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalogintegrationicebergrest


type CatalogIntegrationIcebergRestBearerRestAuthentication struct {
	// The bearer token for the identity provider.
	//
	// External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/catalog_integration_iceberg_rest#bearer_token CatalogIntegrationIcebergRest#bearer_token}
	BearerToken *string `field:"required" json:"bearerToken" yaml:"bearerToken"`
}

