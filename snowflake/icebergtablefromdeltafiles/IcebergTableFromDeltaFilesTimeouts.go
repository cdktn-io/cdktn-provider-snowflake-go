// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtablefromdeltafiles


type IcebergTableFromDeltaFilesTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#create IcebergTableFromDeltaFiles#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#delete IcebergTableFromDeltaFiles#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#read IcebergTableFromDeltaFiles#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/iceberg_table_from_delta_files#update IcebergTableFromDeltaFiles#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

