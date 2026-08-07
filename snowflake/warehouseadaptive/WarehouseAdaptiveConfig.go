// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package warehouseadaptive

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WarehouseAdaptiveConfig struct {
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
	// Identifier for the adaptive warehouse;
	//
	// must be unique for your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_adaptive#name WarehouseAdaptive#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies a comment for the adaptive warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_adaptive#comment WarehouseAdaptive#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_adaptive#id WarehouseAdaptive#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the maximum query performance level for the adaptive warehouse.
	//
	// Determines the initial compute capacity. Valid values are (case-insensitive): `XSMALL` | `SMALL` | `MEDIUM` | `LARGE` | `XLARGE` | `XXLARGE` | `XXXLARGE` | `X4LARGE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_adaptive#max_query_performance_level WarehouseAdaptive#max_query_performance_level}
	MaxQueryPerformanceLevel *string `field:"optional" json:"maxQueryPerformanceLevel" yaml:"maxQueryPerformanceLevel"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Specifies the query throughput multiplier for the adaptive warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_adaptive#query_throughput_multiplier WarehouseAdaptive#query_throughput_multiplier}
	QueryThroughputMultiplier *float64 `field:"optional" json:"queryThroughputMultiplier" yaml:"queryThroughputMultiplier"`
	// Specifies the name of a resource monitor that is explicitly assigned to the adaptive warehouse.
	//
	// For more information about this resource, see [docs](./resource_monitor).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_adaptive#resource_monitor WarehouseAdaptive#resource_monitor}
	ResourceMonitor *string `field:"optional" json:"resourceMonitor" yaml:"resourceMonitor"`
	// Object parameter that specifies the time, in seconds, a SQL statement (query, DDL, DML, etc.) can be queued on a warehouse before it is canceled by the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_adaptive#statement_queued_timeout_in_seconds WarehouseAdaptive#statement_queued_timeout_in_seconds}
	StatementQueuedTimeoutInSeconds *float64 `field:"optional" json:"statementQueuedTimeoutInSeconds" yaml:"statementQueuedTimeoutInSeconds"`
	// Specifies the time, in seconds, after which a running SQL statement (query, DDL, DML, etc.) is canceled by the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_adaptive#statement_timeout_in_seconds WarehouseAdaptive#statement_timeout_in_seconds}
	StatementTimeoutInSeconds *float64 `field:"optional" json:"statementTimeoutInSeconds" yaml:"statementTimeoutInSeconds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_adaptive#timeouts WarehouseAdaptive#timeouts}
	Timeouts *WarehouseAdaptiveTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

