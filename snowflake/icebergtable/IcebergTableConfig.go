// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IcebergTableConfig struct {
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
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#column IcebergTable#column}
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The database in which to create the Iceberg table.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#database IcebergTable#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the Iceberg table;
	//
	// must be unique for the schema in which the Iceberg table is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#name IcebergTable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the Iceberg table.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#schema IcebergTable#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// aggregation_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#aggregation_policy IcebergTable#aggregation_policy}
	AggregationPolicy *IcebergTableAggregationPolicy `field:"optional" json:"aggregationPolicy" yaml:"aggregationPolicy"`
	// The path to a directory where Snowflake can write data and metadata files for the Iceberg table.
	//
	// Specify a relative path from the table's `EXTERNAL_VOLUME` location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#base_location IcebergTable#base_location}
	BaseLocation *string `field:"optional" json:"baseLocation" yaml:"baseLocation"`
	// Specifies the identifier for the catalog integration to use for the Iceberg table.
	//
	// If not specified, the account-level default is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#catalog IcebergTable#catalog}
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Specifies the name of the catalog integration that Snowflake uses to automatically synchronize the Iceberg table with an external catalog.
	//
	// For more information, check [CATALOG_SYNC docs](https://docs.snowflake.com/en/sql-reference/parameters#catalog-sync).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#catalog_sync IcebergTable#catalog_sync}
	CatalogSync *string `field:"optional" json:"catalogSync" yaml:"catalogSync"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether to enable change tracking on the Iceberg table.
	//
	// Cannot be changed after creation. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#change_tracking IcebergTable#change_tracking}
	ChangeTracking *string `field:"optional" json:"changeTracking" yaml:"changeTracking"`
	// check_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#check_constraint IcebergTable#check_constraint}
	CheckConstraint interface{} `field:"optional" json:"checkConstraint" yaml:"checkConstraint"`
	// A list of one or more table columns/expressions to be used as clustering key(s) for the table.
	//
	// External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#cluster_by IcebergTable#cluster_by}
	ClusterBy *[]*string `field:"optional" json:"clusterBy" yaml:"clusterBy"`
	// Specifies a comment for the Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#comment IcebergTable#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the retention period for the Iceberg table so that Time Travel actions can be performed on historical data.
	//
	// For more information, check [DATA_RETENTION_TIME_IN_DAYS docs](https://docs.snowflake.com/en/sql-reference/parameters#data-retention-time-in-days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#data_retention_time_in_days IcebergTable#data_retention_time_in_days}
	DataRetentionTimeInDays *float64 `field:"optional" json:"dataRetentionTimeInDays" yaml:"dataRetentionTimeInDays"`
	// Specifies whether automatic background data compaction is enabled for the Iceberg table. For more information, check [ENABLE_DATA_COMPACTION docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-data-compaction).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#enable_data_compaction IcebergTable#enable_data_compaction}
	EnableDataCompaction interface{} `field:"optional" json:"enableDataCompaction" yaml:"enableDataCompaction"`
	// Specifies whether merge-on-read is enabled for the Iceberg table. For more information, check [ENABLE_ICEBERG_MERGE_ON_READ docs](https://docs.snowflake.com/en/sql-reference/parameters#enable-iceberg-merge-on-read).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#enable_iceberg_merge_on_read IcebergTable#enable_iceberg_merge_on_read}
	EnableIcebergMergeOnRead interface{} `field:"optional" json:"enableIcebergMergeOnRead" yaml:"enableIcebergMergeOnRead"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether error logging is enabled for the Iceberg table.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#error_logging IcebergTable#error_logging}
	ErrorLogging *string `field:"optional" json:"errorLogging" yaml:"errorLogging"`
	// Specifies the identifier for the external volume where the Iceberg table stores its metadata files and data in Parquet format.
	//
	// If not specified, the account-level default is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#external_volume IcebergTable#external_volume}
	ExternalVolume *string `field:"optional" json:"externalVolume" yaml:"externalVolume"`
	// foreign_key_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#foreign_key_constraint IcebergTable#foreign_key_constraint}
	ForeignKeyConstraint interface{} `field:"optional" json:"foreignKeyConstraint" yaml:"foreignKeyConstraint"`
	// Specifies the Iceberg table format version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#iceberg_version IcebergTable#iceberg_version}
	IcebergVersion *float64 `field:"optional" json:"icebergVersion" yaml:"icebergVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#id IcebergTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the maximum number of days for which Snowflake can extend the data retention period for the Iceberg table to prevent streams on the table from becoming stale.
	//
	// For more information, check [MAX_DATA_EXTENSION_TIME_IN_DAYS docs](https://docs.snowflake.com/en/sql-reference/parameters#max-data-extension-time-in-days).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#max_data_extension_time_in_days IcebergTable#max_data_extension_time_in_days}
	MaxDataExtensionTimeInDays *float64 `field:"optional" json:"maxDataExtensionTimeInDays" yaml:"maxDataExtensionTimeInDays"`
	// partition_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#partition_by IcebergTable#partition_by}
	PartitionBy interface{} `field:"optional" json:"partitionBy" yaml:"partitionBy"`
	// Specifies the storage layout for the Iceberg table's Parquet files.
	//
	// Valid values are: [FLAT HIERARCHICAL]. Cannot be changed after creation. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#path_layout IcebergTable#path_layout}
	PathLayout *string `field:"optional" json:"pathLayout" yaml:"pathLayout"`
	// primary_key_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#primary_key_constraint IcebergTable#primary_key_constraint}
	PrimaryKeyConstraint *IcebergTablePrimaryKeyConstraint `field:"optional" json:"primaryKeyConstraint" yaml:"primaryKeyConstraint"`
	// row_access_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#row_access_policy IcebergTable#row_access_policy}
	RowAccessPolicy *IcebergTableRowAccessPolicy `field:"optional" json:"rowAccessPolicy" yaml:"rowAccessPolicy"`
	// Specifies the storage serialization policy for the Iceberg table.
	//
	// Valid values are: [COMPATIBLE OPTIMIZED]. Cannot be changed after creation. For more information, check [STORAGE_SERIALIZATION_POLICY docs](https://docs.snowflake.com/en/sql-reference/parameters#storage-serialization-policy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#storage_serialization_policy IcebergTable#storage_serialization_policy}
	StorageSerializationPolicy *string `field:"optional" json:"storageSerializationPolicy" yaml:"storageSerializationPolicy"`
	// Specifies the target file size (in bytes) used when writing the Iceberg table's Parquet files.
	//
	// Valid values are: [AUTO 16MB 32MB 64MB 128MB]. For more information, check [TARGET_FILE_SIZE docs](https://docs.snowflake.com/en/sql-reference/parameters#target-file-size).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#target_file_size IcebergTable#target_file_size}
	TargetFileSize *string `field:"optional" json:"targetFileSize" yaml:"targetFileSize"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#timeouts IcebergTable#timeouts}
	Timeouts *IcebergTableTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// unique_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#unique_constraint IcebergTable#unique_constraint}
	UniqueConstraint interface{} `field:"optional" json:"uniqueConstraint" yaml:"uniqueConstraint"`
}

