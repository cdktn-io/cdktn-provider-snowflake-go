// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package warehouseinteractive

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WarehouseInteractiveConfig struct {
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
	// Identifier for the interactive warehouse;
	//
	// must be unique for your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#name WarehouseInteractive#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether to automatically resume an interactive warehouse when a SQL statement (e.g. query) is submitted to it. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#auto_resume WarehouseInteractive#auto_resume}
	AutoResume *string `field:"optional" json:"autoResume" yaml:"autoResume"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Specifies the number of seconds of inactivity after which an interactive warehouse is automatically suspended.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#auto_suspend WarehouseInteractive#auto_suspend}
	AutoSuspend *float64 `field:"optional" json:"autoSuspend" yaml:"autoSuspend"`
	// Specifies a comment for the interactive warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#comment WarehouseInteractive#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the name of the fallback warehouse for the interactive warehouse. For more information about this resource, see [docs](./warehouse).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#fallback_warehouse WarehouseInteractive#fallback_warehouse}
	FallbackWarehouse *string `field:"optional" json:"fallbackWarehouse" yaml:"fallbackWarehouse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#id WarehouseInteractive#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies whether the interactive warehouse is created initially in the ‘Suspended’ state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#initially_suspended WarehouseInteractive#initially_suspended}
	InitiallySuspended interface{} `field:"optional" json:"initiallySuspended" yaml:"initiallySuspended"`
	// Specifies the maximum number of server clusters for the interactive warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#max_cluster_count WarehouseInteractive#max_cluster_count}
	MaxClusterCount *float64 `field:"optional" json:"maxClusterCount" yaml:"maxClusterCount"`
	// Object parameter that specifies the concurrency level for SQL statements (i.e. queries and DML) executed by an interactive warehouse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#max_concurrency_level WarehouseInteractive#max_concurrency_level}
	MaxConcurrencyLevel *float64 `field:"optional" json:"maxConcurrencyLevel" yaml:"maxConcurrencyLevel"`
	// Specifies the minimum number of server clusters for the interactive warehouse (only applies to multi-cluster warehouses).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#min_cluster_count WarehouseInteractive#min_cluster_count}
	MinClusterCount *float64 `field:"optional" json:"minClusterCount" yaml:"minClusterCount"`
	// Specifies the name of a resource monitor that is explicitly assigned to the interactive warehouse.
	//
	// For more information about this resource, see [docs](./resource_monitor).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#resource_monitor WarehouseInteractive#resource_monitor}
	ResourceMonitor *string `field:"optional" json:"resourceMonitor" yaml:"resourceMonitor"`
	// Object parameter that specifies the time, in seconds, a SQL statement (query, DDL, DML, etc.) can be queued on an interactive warehouse before it is canceled by the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#statement_queued_timeout_in_seconds WarehouseInteractive#statement_queued_timeout_in_seconds}
	StatementQueuedTimeoutInSeconds *float64 `field:"optional" json:"statementQueuedTimeoutInSeconds" yaml:"statementQueuedTimeoutInSeconds"`
	// Specifies the time, in seconds, after which a running SQL statement (query, DDL, DML, etc.) is canceled by the system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#statement_timeout_in_seconds WarehouseInteractive#statement_timeout_in_seconds}
	StatementTimeoutInSeconds *float64 `field:"optional" json:"statementTimeoutInSeconds" yaml:"statementTimeoutInSeconds"`
	// Specifies the fully qualified names of the tables associated with the interactive warehouse.
	//
	// Changes are applied incrementally (ADD TABLES / DROP TABLES) rather than by full re-association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#tables WarehouseInteractive#tables}
	Tables *[]*string `field:"optional" json:"tables" yaml:"tables"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#timeouts WarehouseInteractive#timeouts}
	Timeouts *WarehouseInteractiveTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Specifies the size of the interactive warehouse.
	//
	// Valid values are (case-insensitive): `XSMALL` | `X-SMALL` | `SMALL` | `MEDIUM` | `LARGE` | `XLARGE` | `X-LARGE` | `XXLARGE` | `X2LARGE` | `2X-LARGE` | `XXXLARGE` | `X3LARGE` | `3X-LARGE` | `X4LARGE` | `4X-LARGE` | `X5LARGE` | `5X-LARGE` | `X6LARGE` | `6X-LARGE`. Note: changing the size briefly suspends and resumes the warehouse to apply the resize (an interactive warehouse cannot be resized while running); removing the size from config will result in the resource recreation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#warehouse_size WarehouseInteractive#warehouse_size}
	WarehouseSize *string `field:"optional" json:"warehouseSize" yaml:"warehouseSize"`
}

