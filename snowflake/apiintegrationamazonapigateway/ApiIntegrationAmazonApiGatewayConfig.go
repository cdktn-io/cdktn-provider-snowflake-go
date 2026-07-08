// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apiintegrationamazonapigateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiIntegrationAmazonApiGatewayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_amazon_api_gateway#api_allowed_prefixes ApiIntegrationAmazonApiGateway#api_allowed_prefixes}
	ApiAllowedPrefixes *[]*string `field:"required" json:"apiAllowedPrefixes" yaml:"apiAllowedPrefixes"`
	// The Amazon Resource Name (ARN) of the IAM role that grants Snowflake permission to call the API endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_amazon_api_gateway#api_aws_role_arn ApiIntegrationAmazonApiGateway#api_aws_role_arn}
	ApiAwsRoleArn *string `field:"required" json:"apiAwsRoleArn" yaml:"apiAwsRoleArn"`
	// Specifies the type of AWS gateway. Valid values are (case-insensitive): `aws_api_gateway` | `aws_private_api_gateway` | `aws_gov_api_gateway` | `aws_gov_private_api_gateway`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_amazon_api_gateway#api_provider ApiIntegrationAmazonApiGateway#api_provider}
	ApiProvider *string `field:"required" json:"apiProvider" yaml:"apiProvider"`
	// Specifies whether this API integration is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_amazon_api_gateway#enabled ApiIntegrationAmazonApiGateway#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the identifier (i.e. name) for the integration. This value must be unique in your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_amazon_api_gateway#name ApiIntegrationAmazonApiGateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Lists the endpoints and resources in the HTTPS proxy service that are not allowed to be called from Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_amazon_api_gateway#api_blocked_prefixes ApiIntegrationAmazonApiGateway#api_blocked_prefixes}
	ApiBlockedPrefixes *[]*string `field:"optional" json:"apiBlockedPrefixes" yaml:"apiBlockedPrefixes"`
	// Specifies the API key (secret) that Snowflake uses to authenticate when making calls to the proxy service.
	//
	// External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_amazon_api_gateway#api_key ApiIntegrationAmazonApiGateway#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Specifies a comment for the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_amazon_api_gateway#comment ApiIntegrationAmazonApiGateway#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_amazon_api_gateway#id ApiIntegrationAmazonApiGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_amazon_api_gateway#timeouts ApiIntegrationAmazonApiGateway#timeouts}
	Timeouts *ApiIntegrationAmazonApiGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

