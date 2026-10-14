// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package openflowdeploymentbyoc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpenflowDeploymentByocConfig struct {
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
	// Specifies the identifier for the Openflow deployment;
	//
	// must be unique for the account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#name OpenflowDeploymentByoc#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies whether the deployment's VPC is created by Snowflake or supplied by you.
	//
	// Valid values are (case-insensitive): `MANAGED` | `PROVIDED`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#vpc_type OpenflowDeploymentByoc#vpc_type}
	VpcType *string `field:"required" json:"vpcType" yaml:"vpcType"`
	// Specifies a comment for the Openflow deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#comment OpenflowDeploymentByoc#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies a custom hostname for ingress into the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#custom_ingress_hostname OpenflowDeploymentByoc#custom_ingress_hostname}
	CustomIngressHostname *string `field:"optional" json:"customIngressHostname" yaml:"customIngressHostname"`
	// A free-text alias for the deployment. Shown in the Openflow UI in place of the deployment's identifier when set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#display_name OpenflowDeploymentByoc#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Specifies the name of the event table for logging messages from stored procedures and UDFs contained by the object with which the event table is associated.
	//
	// Associating an event table with a database is available in [Enterprise Edition or higher](https://docs.snowflake.com/en/user-guide/intro-editions). For more information, check [EVENT_TABLE docs](https://docs.snowflake.com/en/sql-reference/parameters#event-table).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#event_table OpenflowDeploymentByoc#event_table}
	EventTable *string `field:"optional" json:"eventTable" yaml:"eventTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#id OpenflowDeploymentByoc#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#timeouts OpenflowDeploymentByoc#timeouts}
	Timeouts *OpenflowDeploymentByocTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether the deployment is reached over private link.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#use_private_link OpenflowDeploymentByoc#use_private_link}
	UsePrivateLink *string `field:"optional" json:"usePrivateLink" yaml:"usePrivateLink"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether user authentication is performed over private link.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#use_user_auth_over_privatelink OpenflowDeploymentByoc#use_user_auth_over_privatelink}
	UseUserAuthOverPrivatelink *string `field:"optional" json:"useUserAuthOverPrivatelink" yaml:"useUserAuthOverPrivatelink"`
}

