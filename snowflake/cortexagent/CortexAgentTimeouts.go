// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cortexagent


type CortexAgentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/cortex_agent#create CortexAgent#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/cortex_agent#delete CortexAgent#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/cortex_agent#read CortexAgent#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/cortex_agent#update CortexAgent#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

