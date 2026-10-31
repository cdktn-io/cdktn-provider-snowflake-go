// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakehybridtables


type DataSnowflakeHybridTablesIn struct {
	// Returns records for the entire account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/data-sources/hybrid_tables#account DataSnowflakeHybridTables#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// Returns records for the current database in use or for a specified database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/data-sources/hybrid_tables#database DataSnowflakeHybridTables#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Returns records for the current schema in use or a specified schema. Use fully qualified name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.1/docs/data-sources/hybrid_tables#schema DataSnowflakeHybridTables#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

