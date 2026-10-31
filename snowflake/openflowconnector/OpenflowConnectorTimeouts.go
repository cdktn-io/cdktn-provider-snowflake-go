// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package openflowconnector


type OpenflowConnectorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_connector#create OpenflowConnector#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_connector#delete OpenflowConnector#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_connector#read OpenflowConnector#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/openflow_connector#update OpenflowConnector#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

