// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apiintegrationgitrepositorytoken

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiIntegrationGitRepositoryTokenConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_git_repository_token#api_allowed_prefixes ApiIntegrationGitRepositoryToken#api_allowed_prefixes}
	ApiAllowedPrefixes *[]*string `field:"required" json:"apiAllowedPrefixes" yaml:"apiAllowedPrefixes"`
	// Specifies whether this API integration is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_git_repository_token#enabled ApiIntegrationGitRepositoryToken#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the identifier (i.e. name) for the integration. This value must be unique in your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_git_repository_token#name ApiIntegrationGitRepositoryToken#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// When set to true, all authentication secrets are allowed to be used when authenticating to the git repository.
	//
	// Conflicts with `no_allowed_authentication_secrets` and `allowed_authentication_secrets`. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_git_repository_token#all_allowed_authentication_secrets ApiIntegrationGitRepositoryToken#all_allowed_authentication_secrets}
	AllAllowedAuthenticationSecrets interface{} `field:"optional" json:"allAllowedAuthenticationSecrets" yaml:"allAllowedAuthenticationSecrets"`
	// A list of fully-qualified secret identifiers (database.schema.secret) allowed to be used when authenticating to the git repository. Conflicts with `all_allowed_authentication_secrets` and `no_allowed_authentication_secrets`. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_git_repository_token#allowed_authentication_secrets ApiIntegrationGitRepositoryToken#allowed_authentication_secrets}
	AllowedAuthenticationSecrets *[]*string `field:"optional" json:"allowedAuthenticationSecrets" yaml:"allowedAuthenticationSecrets"`
	// Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_git_repository_token#api_blocked_prefixes ApiIntegrationGitRepositoryToken#api_blocked_prefixes}
	ApiBlockedPrefixes *[]*string `field:"optional" json:"apiBlockedPrefixes" yaml:"apiBlockedPrefixes"`
	// Specifies a comment for the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_git_repository_token#comment ApiIntegrationGitRepositoryToken#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_git_repository_token#id ApiIntegrationGitRepositoryToken#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When set to true, no authentication secrets are allowed to be used when authenticating to the git repository.
	//
	// Conflicts with `all_allowed_authentication_secrets` and `allowed_authentication_secrets`. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_git_repository_token#no_allowed_authentication_secrets ApiIntegrationGitRepositoryToken#no_allowed_authentication_secrets}
	NoAllowedAuthenticationSecrets interface{} `field:"optional" json:"noAllowedAuthenticationSecrets" yaml:"noAllowedAuthenticationSecrets"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_git_repository_token#timeouts ApiIntegrationGitRepositoryToken#timeouts}
	Timeouts *ApiIntegrationGitRepositoryTokenTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

