// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package openflowruntime


type OpenflowRuntimeTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_runtime#create OpenflowRuntime#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_runtime#delete OpenflowRuntime#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_runtime#read OpenflowRuntime#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_runtime#update OpenflowRuntime#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

