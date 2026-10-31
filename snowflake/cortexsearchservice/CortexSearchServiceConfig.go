// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cortexsearchservice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CortexSearchServiceConfig struct {
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
	// The database in which to create the Cortex search service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#database CortexSearchService#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the name of the Cortex search service.
	//
	// The name must be unique for the schema in which the service is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#name CortexSearchService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the column to use as the search column for the Cortex search service; must be a text value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#on CortexSearchService#on}
	On *string `field:"required" json:"on" yaml:"on"`
	// Specifies the query to use to populate the Cortex search service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#query CortexSearchService#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The schema in which to create the Cortex search service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#schema CortexSearchService#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the maximum target lag time for the Cortex search service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#target_lag CortexSearchService#target_lag}
	TargetLag *string `field:"required" json:"targetLag" yaml:"targetLag"`
	// The warehouse in which to create the Cortex search service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#warehouse CortexSearchService#warehouse}
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
	// Specifies the list of columns in the base table to enable filtering on when issuing queries to the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#attributes CortexSearchService#attributes}
	Attributes *[]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// Specifies the number of seconds of inactivity after which the Cortex search service automatically suspends its serving compute, releasing the resources.
	//
	// Snowflake requires a minimum of 1800 (30 minutes). If unset, Snowflake defaults to no automatic suspension. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#auto_suspend CortexSearchService#auto_suspend}
	AutoSuspend *float64 `field:"optional" json:"autoSuspend" yaml:"autoSuspend"`
	// Specifies a comment for the Cortex search service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#comment CortexSearchService#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the embedding model to use for the Cortex search service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#embedding_model CortexSearchService#embedding_model}
	EmbeddingModel *string `field:"optional" json:"embeddingModel" yaml:"embeddingModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#id CortexSearchService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the column(s) in the base table that uniquely identify each row, used to enable optimized (incremental) refreshes when the underlying data changes.
	//
	// All primary key columns must use the TEXT data type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#primary_key CortexSearchService#primary_key}
	PrimaryKey *[]*string `field:"optional" json:"primaryKey" yaml:"primaryKey"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/resources/cortex_search_service#timeouts CortexSearchService#timeouts}
	Timeouts *CortexSearchServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

