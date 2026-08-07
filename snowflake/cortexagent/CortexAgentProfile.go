// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cortexagent


type CortexAgentProfile struct {
	// Specifies an avatar image file name or identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/cortex_agent#avatar CortexAgent#avatar}
	Avatar *string `field:"optional" json:"avatar" yaml:"avatar"`
	// Specifies a color theme for the Cortex agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/cortex_agent#color CortexAgent#color}
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Specifies a display name for the Cortex agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/cortex_agent#display_name CortexAgent#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

