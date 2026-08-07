// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalogintegrationicebergrest


type CatalogIntegrationIcebergRestOauthRestAuthentication struct {
	// Specifies one or more scopes for the OAuth token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/catalog_integration_iceberg_rest#oauth_allowed_scopes CatalogIntegrationIcebergRest#oauth_allowed_scopes}
	OauthAllowedScopes *[]*string `field:"required" json:"oauthAllowedScopes" yaml:"oauthAllowedScopes"`
	// Specifies the client ID of the OAuth2 credential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/catalog_integration_iceberg_rest#oauth_client_id CatalogIntegrationIcebergRest#oauth_client_id}
	OauthClientId *string `field:"required" json:"oauthClientId" yaml:"oauthClientId"`
	// Specifies the secret of the OAuth2 credential.
	//
	// External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/catalog_integration_iceberg_rest#oauth_client_secret CatalogIntegrationIcebergRest#oauth_client_secret}
	OauthClientSecret *string `field:"required" json:"oauthClientSecret" yaml:"oauthClientSecret"`
	// Specifies URL for the third-party identity provider.
	//
	// If not specified, Snowflake assumes the remote catalog provider is the identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/catalog_integration_iceberg_rest#oauth_token_uri CatalogIntegrationIcebergRest#oauth_token_uri}
	OauthTokenUri *string `field:"optional" json:"oauthTokenUri" yaml:"oauthTokenUri"`
}

