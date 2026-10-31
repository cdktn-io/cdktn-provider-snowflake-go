// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package openflowruntime

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpenflowRuntimeConfig struct {
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
	// The database in which to create the Openflow runtime.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_runtime#database OpenflowRuntime#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the Openflow deployment the runtime runs in.
	//
	// Snowflake has no ALTER for it, so changing it recreates the runtime. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_runtime#deployment OpenflowRuntime#deployment}
	Deployment *string `field:"required" json:"deployment" yaml:"deployment"`
	// Specifies the role the runtime executes as.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_runtime#execute_as_role OpenflowRuntime#execute_as_role}
	ExecuteAsRole *string `field:"required" json:"executeAsRole" yaml:"executeAsRole"`
	// Specifies the maximum number of nodes the runtime scales up to. For more information, check [CREATE OPENFLOW RUNTIME documentation](https://docs.snowflake.com/en/sql-reference/sql/create-openflow-runtime).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_runtime#max_nodes OpenflowRuntime#max_nodes}
	MaxNodes *float64 `field:"required" json:"maxNodes" yaml:"maxNodes"`
	// Specifies the minimum number of nodes the runtime scales down to. For more information, check [CREATE OPENFLOW RUNTIME documentation](https://docs.snowflake.com/en/sql-reference/sql/create-openflow-runtime).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_runtime#min_nodes OpenflowRuntime#min_nodes}
	MinNodes *float64 `field:"required" json:"minNodes" yaml:"minNodes"`
	// Specifies the identifier for the Openflow runtime;
	//
	// must be unique for the schema in which the runtime is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_runtime#name OpenflowRuntime#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the size of the runtime's nodes.
	//
	// Snowflake has no ALTER for it, so changing it recreates the runtime. Valid values are (case-insensitive): `SMALL` | `MEDIUM` | `LARGE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_runtime#node_type OpenflowRuntime#node_type}
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// The schema in which to create the Openflow runtime.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_runtime#schema OpenflowRuntime#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies a comment for the Openflow runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_runtime#comment OpenflowRuntime#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// A free-text alias for the runtime. Shown in the Openflow UI in place of the runtime's identifier when set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_runtime#display_name OpenflowRuntime#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Specifies the names of the external access integrations that allow the runtime egress to external data sources.
	//
	// Supported for Snowflake-managed deployments only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_runtime#external_access_integrations OpenflowRuntime#external_access_integrations}
	ExternalAccessIntegrations *[]*string `field:"optional" json:"externalAccessIntegrations" yaml:"externalAccessIntegrations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_runtime#id OpenflowRuntime#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_runtime#timeouts OpenflowRuntime#timeouts}
	Timeouts *OpenflowRuntimeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

