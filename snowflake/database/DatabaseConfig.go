// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package database

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatabaseConfig struct {
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
	// Specifies the identifier for the database;
	//
	// must be unique for your account. As a best practice for [Database Replication and Failover](https://docs.snowflake.com/en/user-guide/db-replication-intro), it is recommended to give each secondary database the same name as its primary database. This practice supports referencing fully-qualified objects (i.e. '<db>.<schema>.<object>') by other objects in the same database, such as querying a fully-qualified table name in a view. If a secondary database has a different name from the primary database, then these object references would break in the secondary database. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#name Database#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The parameter that specifies the default catalog to use for Iceberg tables. For more information, check [CATALOG docs](https://docs.snowflake.com/en/sql-reference/parameters#catalog).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#catalog Database#catalog}
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Specifies a comment for the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#comment Database#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the number of days for which Time Travel actions (CLONE and UNDROP) can be performed on the database, as well as specifying the default Time Travel retention time for all schemas created in the database.
	//
	// For more details, see [Understanding & Using Time Travel](https://docs.snowflake.com/en/user-guide/data-time-travel). For more information, check [DATA_RETENTION_TIME_IN_DAYS docs](https://docs.snowflake.com/en/sql-reference/parameters#data-retention-time-in-days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#data_retention_time_in_days Database#data_retention_time_in_days}
	DataRetentionTimeInDays *float64 `field:"optional" json:"dataRetentionTimeInDays" yaml:"dataRetentionTimeInDays"`
	// Specifies a default collation specification for all schemas and tables added to the database.
	//
	// It can be overridden on schema or table level. For more information, see [collation specification](https://docs.snowflake.com/en/sql-reference/collation#label-collation-specification). For more information, check [DEFAULT_DDL_COLLATION docs](https://docs.snowflake.com/en/sql-reference/parameters#default-ddl-collation).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#default_ddl_collation Database#default_ddl_collation}
	DefaultDdlCollation *string `field:"optional" json:"defaultDdlCollation" yaml:"defaultDdlCollation"`
	// Sets the preferred CPU compute pool used for Notebooks on CPU Container Runtime. For more information, check [DEFAULT_NOTEBOOK_COMPUTE_POOL_CPU docs](https://docs.snowflake.com/en/sql-reference/parameters#default-notebook-compute-pool-cpu).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#default_notebook_compute_pool_cpu Database#default_notebook_compute_pool_cpu}
	DefaultNotebookComputePoolCpu *string `field:"optional" json:"defaultNotebookComputePoolCpu" yaml:"defaultNotebookComputePoolCpu"`
	// Sets the preferred GPU compute pool used for Notebooks on GPU Container Runtime. For more information, check [DEFAULT_NOTEBOOK_COMPUTE_POOL_GPU docs](https://docs.snowflake.com/en/sql-reference/parameters#default-notebook-compute-pool-gpu).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#default_notebook_compute_pool_gpu Database#default_notebook_compute_pool_gpu}
	DefaultNotebookComputePoolGpu *string `field:"optional" json:"defaultNotebookComputePoolGpu" yaml:"defaultNotebookComputePoolGpu"`
	// Specifies whether to drop public schema on creation or not.
	//
	// Modifying the parameter after database is already created won't have any effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#drop_public_schema_on_creation Database#drop_public_schema_on_creation}
	DropPublicSchemaOnCreation interface{} `field:"optional" json:"dropPublicSchemaOnCreation" yaml:"dropPublicSchemaOnCreation"`
	// If true, enables stdout/stderr fast path logging for anonymous stored procedures. For more information, check [ENABLE_CONSOLE_OUTPUT docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-console-output).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#enable_console_output Database#enable_console_output}
	EnableConsoleOutput interface{} `field:"optional" json:"enableConsoleOutput" yaml:"enableConsoleOutput"`
	// The parameter that specifies the default external volume to use for Iceberg tables. For more information, check [EXTERNAL_VOLUME docs](https://docs.snowflake.com/en/sql-reference/parameters#external-volume).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#external_volume Database#external_volume}
	ExternalVolume *string `field:"optional" json:"externalVolume" yaml:"externalVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#id Database#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the database as transient.
	//
	// Transient databases do not have a Fail-safe period so they do not incur additional storage costs once they leave Time Travel; however, this means they are also not protected by Fail-safe in the event of a data loss.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#is_transient Database#is_transient}
	IsTransient interface{} `field:"optional" json:"isTransient" yaml:"isTransient"`
	// Specifies the severity level of log events (rows with record type EVENT) that should be ingested and made available in the active event table.
	//
	// Log events at the specified level (and at more severe levels) are ingested. Valid values are (case-insensitive): `TRACE` | `DEBUG` | `INFO` | `WARN` | `ERROR` | `FATAL` | `OFF`. For more information, check [LOG_EVENT_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#log-event-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#log_event_level Database#log_event_level}
	LogEventLevel *string `field:"optional" json:"logEventLevel" yaml:"logEventLevel"`
	// Specifies the severity level of messages that should be ingested and made available in the active event table.
	//
	// Messages at the specified level (and at more severe levels) are ingested. Valid values are (case-insensitive): `TRACE` | `DEBUG` | `INFO` | `WARN` | `ERROR` | `FATAL` | `OFF`. For more information, check [LOG_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#log-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#log_level Database#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Object parameter that specifies the maximum number of days for which Snowflake can extend the data retention period for tables in the database to prevent streams on the tables from becoming stale.
	//
	// For more information, check [MAX_DATA_EXTENSION_TIME_IN_DAYS docs](https://docs.snowflake.com/en/sql-reference/parameters#max-data-extension-time-in-days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#max_data_extension_time_in_days Database#max_data_extension_time_in_days}
	MaxDataExtensionTimeInDays *float64 `field:"optional" json:"maxDataExtensionTimeInDays" yaml:"maxDataExtensionTimeInDays"`
	// If true, the case of quoted identifiers is ignored. For more information, check [QUOTED_IDENTIFIERS_IGNORE_CASE docs](https://docs.snowflake.com/en/sql-reference/parameters#quoted-identifiers-ignore-case).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#quoted_identifiers_ignore_case Database#quoted_identifiers_ignore_case}
	QuotedIdentifiersIgnoreCase interface{} `field:"optional" json:"quotedIdentifiersIgnoreCase" yaml:"quotedIdentifiersIgnoreCase"`
	// Specifies whether to replace invalid UTF-8 characters with the Unicode replacement character in query results for an Iceberg table.
	//
	// You can only set this parameter for tables that use an external Iceberg catalog. For more information, check [REPLACE_INVALID_CHARACTERS docs](https://docs.snowflake.com/en/sql-reference/parameters#replace-invalid-characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#replace_invalid_characters Database#replace_invalid_characters}
	ReplaceInvalidCharacters interface{} `field:"optional" json:"replaceInvalidCharacters" yaml:"replaceInvalidCharacters"`
	// replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#replication Database#replication}
	Replication *DatabaseReplication `field:"optional" json:"replication" yaml:"replication"`
	// The storage serialization policy for Iceberg tables that use Snowflake as the catalog.
	//
	// COMPATIBLE: Snowflake performs encoding and compression of data files that ensures interoperability with third-party compute engines. OPTIMIZED: Snowflake performs encoding and compression of data files that ensures the best table performance within Snowflake. Valid values are (case-insensitive): `COMPATIBLE` | `OPTIMIZED`. For more information, check [STORAGE_SERIALIZATION_POLICY docs](https://docs.snowflake.com/en/sql-reference/parameters#storage-serialization-policy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#storage_serialization_policy Database#storage_serialization_policy}
	StorageSerializationPolicy *string `field:"optional" json:"storageSerializationPolicy" yaml:"storageSerializationPolicy"`
	// How many times a task must fail in a row before it is automatically suspended.
	//
	// 0 disables auto-suspending. For more information, check [SUSPEND_TASK_AFTER_NUM_FAILURES docs](https://docs.snowflake.com/en/sql-reference/parameters#suspend-task-after-num-failures).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#suspend_task_after_num_failures Database#suspend_task_after_num_failures}
	SuspendTaskAfterNumFailures *float64 `field:"optional" json:"suspendTaskAfterNumFailures" yaml:"suspendTaskAfterNumFailures"`
	// Maximum automatic retries allowed for a user task. For more information, check [TASK_AUTO_RETRY_ATTEMPTS docs](https://docs.snowflake.com/en/sql-reference/parameters#task-auto-retry-attempts).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#task_auto_retry_attempts Database#task_auto_retry_attempts}
	TaskAutoRetryAttempts *float64 `field:"optional" json:"taskAutoRetryAttempts" yaml:"taskAutoRetryAttempts"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#timeouts Database#timeouts}
	Timeouts *DatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Controls how trace events are ingested into the event table.
	//
	// Valid values are (case-insensitive): `ALWAYS` | `ON_EVENT` | `PROPAGATE` | `OFF`. For more information, check [TRACE_LEVEL docs](https://docs.snowflake.com/en/sql-reference/parameters#trace-level).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#trace_level Database#trace_level}
	TraceLevel *string `field:"optional" json:"traceLevel" yaml:"traceLevel"`
	// The initial size of warehouse to use for managed warehouses in the absence of history.
	//
	// Valid values are (case-insensitive): `XSMALL` | `SMALL` | `MEDIUM` | `LARGE` | `XLARGE` | `XXLARGE` | `XXXLARGE` | `X4LARGE` | `X5LARGE` | `X6LARGE`. For more information, check [USER_TASK_MANAGED_INITIAL_WAREHOUSE_SIZE docs](https://docs.snowflake.com/en/sql-reference/parameters#user-task-managed-initial-warehouse-size).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#user_task_managed_initial_warehouse_size Database#user_task_managed_initial_warehouse_size}
	UserTaskManagedInitialWarehouseSize *string `field:"optional" json:"userTaskManagedInitialWarehouseSize" yaml:"userTaskManagedInitialWarehouseSize"`
	// Minimum amount of time between Triggered Task executions in seconds. For more information, check [USER_TASK_MINIMUM_TRIGGER_INTERVAL_IN_SECONDS docs](https://docs.snowflake.com/en/sql-reference/parameters#user-task-minimum-trigger-interval-in-seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#user_task_minimum_trigger_interval_in_seconds Database#user_task_minimum_trigger_interval_in_seconds}
	UserTaskMinimumTriggerIntervalInSeconds *float64 `field:"optional" json:"userTaskMinimumTriggerIntervalInSeconds" yaml:"userTaskMinimumTriggerIntervalInSeconds"`
	// User task execution timeout in milliseconds. For more information, check [USER_TASK_TIMEOUT_MS docs](https://docs.snowflake.com/en/sql-reference/parameters#user-task-timeout-ms).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/resources/database#user_task_timeout_ms Database#user_task_timeout_ms}
	UserTaskTimeoutMs *float64 `field:"optional" json:"userTaskTimeoutMs" yaml:"userTaskTimeoutMs"`
}

