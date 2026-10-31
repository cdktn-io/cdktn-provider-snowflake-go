// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package intelligencecortexagentattachment


type IntelligenceCortexAgentAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/intelligence_cortex_agent_attachment#create IntelligenceCortexAgentAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/intelligence_cortex_agent_attachment#delete IntelligenceCortexAgentAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/intelligence_cortex_agent_attachment#read IntelligenceCortexAgentAttachment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/intelligence_cortex_agent_attachment#update IntelligenceCortexAgentAttachment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

