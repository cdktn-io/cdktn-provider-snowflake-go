// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package openflowdeploymentbyoc


type OpenflowDeploymentByocTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#create OpenflowDeploymentByoc#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#delete OpenflowDeploymentByoc#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#read OpenflowDeploymentByoc#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/openflow_deployment_byoc#update OpenflowDeploymentByoc#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

