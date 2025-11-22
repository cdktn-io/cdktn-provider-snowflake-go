// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notebook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotebookConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The database in which to create the notebook.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/notebook#database Notebook#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the notebook;
	//
	// must be unique for the schema in which the notebook is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/notebook#name Notebook#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the notebook.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/notebook#schema Notebook#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies a comment for the notebook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/notebook#comment Notebook#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/notebook#from Notebook#from}
	From interface{} `field:"optional" json:"from" yaml:"from"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/notebook#id Notebook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the number of seconds of idle time before the notebook is shut down automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/notebook#idle_auto_shutdown_time_seconds Notebook#idle_auto_shutdown_time_seconds}
	IdleAutoShutdownTimeSeconds *float64 `field:"optional" json:"idleAutoShutdownTimeSeconds" yaml:"idleAutoShutdownTimeSeconds"`
	// Specifies a user-specified identifier for the notebook file name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/notebook#main_file Notebook#main_file}
	MainFile *string `field:"optional" json:"mainFile" yaml:"mainFile"`
	// Specifies the warehouse where SQL queries in the notebook are run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/notebook#query_warehouse Notebook#query_warehouse}
	QueryWarehouse *string `field:"optional" json:"queryWarehouse" yaml:"queryWarehouse"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/notebook#timeouts Notebook#timeouts}
	Timeouts *NotebookTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Specifies the warehouse that runs the notebook kernel and python code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/notebook#warehouse Notebook#warehouse}
	Warehouse *string `field:"optional" json:"warehouse" yaml:"warehouse"`
}

