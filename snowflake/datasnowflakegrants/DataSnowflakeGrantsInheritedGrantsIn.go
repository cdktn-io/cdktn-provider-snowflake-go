// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakegrants


type DataSnowflakeGrantsInheritedGrantsIn struct {
	// Lists all inherited grants defined in the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/data-sources/grants#account DataSnowflakeGrants#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// Lists all inherited grants defined in the specified database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/data-sources/grants#database DataSnowflakeGrants#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Lists all inherited grants defined in the specified schema. Schema must be a fully qualified name ("&lt;db_name&gt;"."&lt;schema_name&gt;").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/data-sources/grants#schema DataSnowflakeGrants#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

