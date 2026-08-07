// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apiintegrationgitrepositoryoauth2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiIntegrationGitRepositoryOauth2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#api_allowed_prefixes ApiIntegrationGitRepositoryOauth2#api_allowed_prefixes}
	ApiAllowedPrefixes *[]*string `field:"required" json:"apiAllowedPrefixes" yaml:"apiAllowedPrefixes"`
	// Specifies whether this API integration is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#enabled ApiIntegrationGitRepositoryOauth2#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the identifier (i.e. name) for the integration. This value must be unique in your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#name ApiIntegrationGitRepositoryOauth2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The OAuth 2.0 authorization endpoint for the Git repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#oauth_authorization_endpoint ApiIntegrationGitRepositoryOauth2#oauth_authorization_endpoint}
	OauthAuthorizationEndpoint *string `field:"required" json:"oauthAuthorizationEndpoint" yaml:"oauthAuthorizationEndpoint"`
	// The client ID for the OAuth 2.0 application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#oauth_client_id ApiIntegrationGitRepositoryOauth2#oauth_client_id}
	OauthClientId *string `field:"required" json:"oauthClientId" yaml:"oauthClientId"`
	// The client secret for the OAuth 2.0 application. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#oauth_client_secret ApiIntegrationGitRepositoryOauth2#oauth_client_secret}
	OauthClientSecret *string `field:"required" json:"oauthClientSecret" yaml:"oauthClientSecret"`
	// The OAuth 2.0 token endpoint for the Git repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#oauth_token_endpoint ApiIntegrationGitRepositoryOauth2#oauth_token_endpoint}
	OauthTokenEndpoint *string `field:"required" json:"oauthTokenEndpoint" yaml:"oauthTokenEndpoint"`
	// Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#api_blocked_prefixes ApiIntegrationGitRepositoryOauth2#api_blocked_prefixes}
	ApiBlockedPrefixes *[]*string `field:"optional" json:"apiBlockedPrefixes" yaml:"apiBlockedPrefixes"`
	// Specifies a comment for the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#comment ApiIntegrationGitRepositoryOauth2#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#id ApiIntegrationGitRepositoryOauth2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the validity period (in seconds) for the OAuth 2.0 access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#oauth_access_token_validity ApiIntegrationGitRepositoryOauth2#oauth_access_token_validity}
	OauthAccessTokenValidity *float64 `field:"optional" json:"oauthAccessTokenValidity" yaml:"oauthAccessTokenValidity"`
	// Specifies a list of scopes to use when making a request from the OAuth by a role with USAGE on the integration.
	//
	// Valid values are (case-insensitive): `read_api` | `read_repository` | `write_repository`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#oauth_allowed_scopes ApiIntegrationGitRepositoryOauth2#oauth_allowed_scopes}
	OauthAllowedScopes *[]*string `field:"optional" json:"oauthAllowedScopes" yaml:"oauthAllowedScopes"`
	// Specifies the validity period (in seconds) for the OAuth 2.0 refresh token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#oauth_refresh_token_validity ApiIntegrationGitRepositoryOauth2#oauth_refresh_token_validity}
	OauthRefreshTokenValidity *float64 `field:"optional" json:"oauthRefreshTokenValidity" yaml:"oauthRefreshTokenValidity"`
	// Specifies the username to authenticate with the Git repository using OAuth 2.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#oauth_username ApiIntegrationGitRepositoryOauth2#oauth_username}
	OauthUsername *string `field:"optional" json:"oauthUsername" yaml:"oauthUsername"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/api_integration_git_repository_oauth2#timeouts ApiIntegrationGitRepositoryOauth2#timeouts}
	Timeouts *ApiIntegrationGitRepositoryOauth2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

