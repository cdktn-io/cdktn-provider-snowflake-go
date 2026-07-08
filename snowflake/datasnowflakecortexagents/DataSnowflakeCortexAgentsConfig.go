// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakecortexagents

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataSnowflakeCortexAgentsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/data-sources/cortex_agents#id DataSnowflakeCortexAgents#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/data-sources/cortex_agents#in DataSnowflakeCortexAgents#in}
	In *DataSnowflakeCortexAgentsIn `field:"optional" json:"in" yaml:"in"`
	// Filters the output with **case-insensitive** pattern, with support for SQL wildcard characters (`%` and `_`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/data-sources/cortex_agents#like DataSnowflakeCortexAgents#like}
	Like *string `field:"optional" json:"like" yaml:"like"`
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/data-sources/cortex_agents#limit DataSnowflakeCortexAgents#limit}
	Limit *DataSnowflakeCortexAgentsLimit `field:"optional" json:"limit" yaml:"limit"`
	// Filters the output with **case-sensitive** characters indicating the beginning of the object name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/data-sources/cortex_agents#starts_with DataSnowflakeCortexAgents#starts_with}
	StartsWith *string `field:"optional" json:"startsWith" yaml:"startsWith"`
	// (Default: `true`) Runs DESC AGENT for each object returned by SHOW AGENTS.
	//
	// The output of describe is saved to the describe_output field. By default this value is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/data-sources/cortex_agents#with_describe DataSnowflakeCortexAgents#with_describe}
	WithDescribe interface{} `field:"optional" json:"withDescribe" yaml:"withDescribe"`
}

