// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apiintegrationexternalmcpoauth2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiIntegrationExternalMcpOauth2Config struct {
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
	// Explicitly limits external functions that use the integration to reference one or more HTTPS proxy service and remote service endpoints and resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_external_mcp_oauth2#api_allowed_prefixes ApiIntegrationExternalMcpOauth2#api_allowed_prefixes}
	ApiAllowedPrefixes *[]*string `field:"required" json:"apiAllowedPrefixes" yaml:"apiAllowedPrefixes"`
	// Specifies whether this API integration is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_external_mcp_oauth2#enabled ApiIntegrationExternalMcpOauth2#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the identifier (i.e. name) for the integration. This value must be unique in your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_external_mcp_oauth2#name ApiIntegrationExternalMcpOauth2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the OAuth 2.0 authorization endpoint URL for the MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_external_mcp_oauth2#oauth_authorization_endpoint ApiIntegrationExternalMcpOauth2#oauth_authorization_endpoint}
	OauthAuthorizationEndpoint *string `field:"required" json:"oauthAuthorizationEndpoint" yaml:"oauthAuthorizationEndpoint"`
	// Specifies the OAuth 2.0 client ID for the MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_external_mcp_oauth2#oauth_client_id ApiIntegrationExternalMcpOauth2#oauth_client_id}
	OauthClientId *string `field:"required" json:"oauthClientId" yaml:"oauthClientId"`
	// Specifies the OAuth 2.0 client secret for the MCP server. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_external_mcp_oauth2#oauth_client_secret ApiIntegrationExternalMcpOauth2#oauth_client_secret}
	OauthClientSecret *string `field:"required" json:"oauthClientSecret" yaml:"oauthClientSecret"`
	// Specifies the OAuth 2.0 token endpoint URL for the MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_external_mcp_oauth2#oauth_token_endpoint ApiIntegrationExternalMcpOauth2#oauth_token_endpoint}
	OauthTokenEndpoint *string `field:"required" json:"oauthTokenEndpoint" yaml:"oauthTokenEndpoint"`
	// Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_external_mcp_oauth2#api_blocked_prefixes ApiIntegrationExternalMcpOauth2#api_blocked_prefixes}
	ApiBlockedPrefixes *[]*string `field:"optional" json:"apiBlockedPrefixes" yaml:"apiBlockedPrefixes"`
	// Specifies a comment for the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_external_mcp_oauth2#comment ApiIntegrationExternalMcpOauth2#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_external_mcp_oauth2#id ApiIntegrationExternalMcpOauth2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the OAuth 2.0 client authentication method. Valid values are (case-insensitive): `CLIENT_SECRET_BASIC` | `CLIENT_SECRET_POST`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_external_mcp_oauth2#oauth_client_auth_method ApiIntegrationExternalMcpOauth2#oauth_client_auth_method}
	OauthClientAuthMethod *string `field:"optional" json:"oauthClientAuthMethod" yaml:"oauthClientAuthMethod"`
	// Specifies the validity period (in seconds) for refresh tokens issued by the MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_external_mcp_oauth2#oauth_refresh_token_validity ApiIntegrationExternalMcpOauth2#oauth_refresh_token_validity}
	OauthRefreshTokenValidity *float64 `field:"optional" json:"oauthRefreshTokenValidity" yaml:"oauthRefreshTokenValidity"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_external_mcp_oauth2#timeouts ApiIntegrationExternalMcpOauth2#timeouts}
	Timeouts *ApiIntegrationExternalMcpOauth2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

