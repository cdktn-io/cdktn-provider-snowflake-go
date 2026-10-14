// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package intelligencecortexagentattachment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IntelligenceCortexAgentAttachmentConfig struct {
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
	// Fully qualified name of the Cortex agent to attach.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using pipes (`|`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/intelligence_cortex_agent_attachment#cortex_agent_name IntelligenceCortexAgentAttachment#cortex_agent_name}
	CortexAgentName *string `field:"required" json:"cortexAgentName" yaml:"cortexAgentName"`
	// Fully qualified name of the Snowflake Intelligence object to attach the agent to.
	//
	// The account default object is SNOWFLAKE_INTELLIGENCE_OBJECT_DEFAULT. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using pipes (`|`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/intelligence_cortex_agent_attachment#snowflake_intelligence_name IntelligenceCortexAgentAttachment#snowflake_intelligence_name}
	SnowflakeIntelligenceName *string `field:"required" json:"snowflakeIntelligenceName" yaml:"snowflakeIntelligenceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/intelligence_cortex_agent_attachment#id IntelligenceCortexAgentAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/intelligence_cortex_agent_attachment#timeouts IntelligenceCortexAgentAttachment#timeouts}
	Timeouts *IntelligenceCortexAgentAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

