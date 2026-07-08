// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtablefromfiles

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IcebergTableFromFilesConfig struct {
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
	// The database in which to create the Iceberg table.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_files#database IcebergTableFromFiles#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the relative path of the Iceberg metadata file in the external volume.
	//
	// Cannot be changed after creation. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_files#metadata_file_path IcebergTableFromFiles#metadata_file_path}
	MetadataFilePath *string `field:"required" json:"metadataFilePath" yaml:"metadataFilePath"`
	// Specifies the identifier for the Iceberg table;
	//
	// must be unique for the schema in which the Iceberg table is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_files#name IcebergTableFromFiles#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the Iceberg table.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_files#schema IcebergTableFromFiles#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the identifier for the catalog integration to use for the Iceberg table.
	//
	// If not specified, the account-level default is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_files#catalog IcebergTableFromFiles#catalog}
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Specifies a comment for the Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_files#comment IcebergTableFromFiles#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Specifies the identifier for the external volume where the Iceberg table stores its metadata files and data in Parquet format.
	//
	// If not specified, the account-level default is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_files#external_volume IcebergTableFromFiles#external_volume}
	ExternalVolume *string `field:"optional" json:"externalVolume" yaml:"externalVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_files#id IcebergTableFromFiles#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies whether to replace invalid UTF-8 characters with the Unicode replacement character (`�`) in query results for an Iceberg table.
	//
	// For more information, check [REPLACE_INVALID_CHARACTERS docs](https://docs.snowflake.com/en/sql-reference/parameters#replace-invalid-characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_files#replace_invalid_characters IcebergTableFromFiles#replace_invalid_characters}
	ReplaceInvalidCharacters interface{} `field:"optional" json:"replaceInvalidCharacters" yaml:"replaceInvalidCharacters"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_files#timeouts IcebergTableFromFiles#timeouts}
	Timeouts *IcebergTableFromFilesTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

