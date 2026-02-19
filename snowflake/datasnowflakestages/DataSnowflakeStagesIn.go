// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakestages


type DataSnowflakeStagesIn struct {
	// Returns records for the entire account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/data-sources/stages#account DataSnowflakeStages#account}
	Account interface{} `field:"optional" json:"account" yaml:"account"`
	// Returns records for the specified application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/data-sources/stages#application DataSnowflakeStages#application}
	Application *string `field:"optional" json:"application" yaml:"application"`
	// Returns records for the specified application package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/data-sources/stages#application_package DataSnowflakeStages#application_package}
	ApplicationPackage *string `field:"optional" json:"applicationPackage" yaml:"applicationPackage"`
	// Returns records for the current database in use or for a specified database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/data-sources/stages#database DataSnowflakeStages#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Returns records for the current schema in use or a specified schema. Use fully qualified name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/data-sources/stages#schema DataSnowflakeStages#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

