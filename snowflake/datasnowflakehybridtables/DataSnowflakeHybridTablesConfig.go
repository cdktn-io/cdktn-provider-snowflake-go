// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakehybridtables

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataSnowflakeHybridTablesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/data-sources/hybrid_tables#id DataSnowflakeHybridTables#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/data-sources/hybrid_tables#in DataSnowflakeHybridTables#in}
	In *DataSnowflakeHybridTablesIn `field:"optional" json:"in" yaml:"in"`
	// Filters the output with **case-insensitive** pattern, with support for SQL wildcard characters (`%` and `_`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/data-sources/hybrid_tables#like DataSnowflakeHybridTables#like}
	Like *string `field:"optional" json:"like" yaml:"like"`
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/data-sources/hybrid_tables#limit DataSnowflakeHybridTables#limit}
	Limit *DataSnowflakeHybridTablesLimit `field:"optional" json:"limit" yaml:"limit"`
	// Filters the output with **case-sensitive** characters indicating the beginning of the object name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/data-sources/hybrid_tables#starts_with DataSnowflakeHybridTables#starts_with}
	StartsWith *string `field:"optional" json:"startsWith" yaml:"startsWith"`
	// (Default: `true`) Runs DESC TABLE for each hybrid table returned by SHOW HYBRID TABLES.
	//
	// The output of describe is saved to the describe_output field. By default this value is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/data-sources/hybrid_tables#with_describe DataSnowflakeHybridTables#with_describe}
	WithDescribe interface{} `field:"optional" json:"withDescribe" yaml:"withDescribe"`
	// (Default: `false`) Runs SHOW INDEXES for each hybrid table returned by SHOW HYBRID TABLES.
	//
	// The output is saved to the show_indexes field. By default this value is set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/data-sources/hybrid_tables#with_indexes DataSnowflakeHybridTables#with_indexes}
	WithIndexes interface{} `field:"optional" json:"withIndexes" yaml:"withIndexes"`
	// (Default: `false`) Runs SHOW PRIMARY KEYS, SHOW UNIQUE KEYS, and SHOW IMPORTED KEYS for each hybrid table returned by SHOW HYBRID TABLES.
	//
	// The merged constraints are saved to the show_keys_output field, grouped by constraint name and ordered by kind, then by column names. By default this value is set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/data-sources/hybrid_tables#with_keys DataSnowflakeHybridTables#with_keys}
	WithKeys interface{} `field:"optional" json:"withKeys" yaml:"withKeys"`
	// (Default: `true`) Runs SHOW PARAMETERS FOR TABLE for each hybrid table returned by SHOW HYBRID TABLES.
	//
	// The output is saved to the parameters field. By default this value is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/data-sources/hybrid_tables#with_parameters DataSnowflakeHybridTables#with_parameters}
	WithParameters interface{} `field:"optional" json:"withParameters" yaml:"withParameters"`
}

