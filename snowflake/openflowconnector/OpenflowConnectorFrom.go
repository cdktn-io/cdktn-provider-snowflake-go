// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package openflowconnector


type OpenflowConnectorFrom struct {
	// Catalog definition ID for the connector type, for example `OPENFLOW_POSTGRES_CDC`.
	//
	// List the available IDs with the `snowflake_openflow_connector_definitions` data source. A connector created this way is a draft: it settles on STOPPED and stays there until a configuration version is committed, which this resource does not do.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.21.0/docs/resources/openflow_connector#definition OpenflowConnector#definition}
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// Path to the bundle within the stage. The bundle's root is used when omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.21.0/docs/resources/openflow_connector#path OpenflowConnector#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Identifier of a stage holding a complete configuration bundle, which is how a connector arrives already configured and able to start without a commit.
	//
	// A git repository stage works here too.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.21.0/docs/resources/openflow_connector#stage OpenflowConnector#stage}
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

