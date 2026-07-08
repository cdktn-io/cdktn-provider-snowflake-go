// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtablefromdeltafiles

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IcebergTableFromDeltaFilesConfig struct {
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
	// Specifies the relative path of the Delta table's directory in the external volume. Cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#base_location IcebergTableFromDeltaFiles#base_location}
	BaseLocation *string `field:"required" json:"baseLocation" yaml:"baseLocation"`
	// The database in which to create the Iceberg table.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#database IcebergTableFromDeltaFiles#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the Iceberg table;
	//
	// must be unique for the schema in which the Iceberg table is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#name IcebergTableFromDeltaFiles#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the Iceberg table.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#schema IcebergTableFromDeltaFiles#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether Snowflake should automatically refresh the Iceberg table metadata when new files are added to the Delta table's directory.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#auto_refresh IcebergTableFromDeltaFiles#auto_refresh}
	AutoRefresh *string `field:"optional" json:"autoRefresh" yaml:"autoRefresh"`
	// Specifies the identifier for the catalog integration to use for the Iceberg table.
	//
	// If not specified, the account-level default is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#catalog IcebergTableFromDeltaFiles#catalog}
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Specifies a comment for the Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#comment IcebergTableFromDeltaFiles#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the identifier for the external volume where the Iceberg table stores its metadata files and data in Parquet format.
	//
	// If not specified, the account-level default is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#external_volume IcebergTableFromDeltaFiles#external_volume}
	ExternalVolume *string `field:"optional" json:"externalVolume" yaml:"externalVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#id IcebergTableFromDeltaFiles#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies whether to replace invalid UTF-8 characters with the Unicode replacement character (`�`) in query results for an Iceberg table.
	//
	// For more information, check [REPLACE_INVALID_CHARACTERS docs](https://docs.snowflake.com/en/sql-reference/parameters#replace-invalid-characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#replace_invalid_characters IcebergTableFromDeltaFiles#replace_invalid_characters}
	ReplaceInvalidCharacters interface{} `field:"optional" json:"replaceInvalidCharacters" yaml:"replaceInvalidCharacters"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#timeouts IcebergTableFromDeltaFiles#timeouts}
	Timeouts *IcebergTableFromDeltaFilesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

