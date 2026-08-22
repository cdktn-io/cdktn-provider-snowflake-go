// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoaccountrole


type GrantPrivilegesToAccountRoleOnSchemaObjectInherited struct {
	// The plural object type of the schema object on which an inherited privilege will be granted.
	//
	// Valid values are (case-insensitive): `AGENTS` | `AGGREGATION POLICIES` | `ALERTS` | `AUTHENTICATION POLICIES` | `CORTEX SEARCH SERVICES` | `DATA METRIC FUNCTIONS` | `DATASETS` | `DBT PROJECTS` | `DYNAMIC TABLES` | `EVENT TABLES` | `EXTERNAL TABLES` | `FILE FORMATS` | `FUNCTIONS` | `GIT REPOSITORIES` | `HYBRID TABLES` | `IMAGE REPOSITORIES` | `ICEBERG TABLES` | `INTERACTIVE TABLES` | `MASKING POLICIES` | `MATERIALIZED VIEWS` | `MCP SERVERS` | `MODELS` | `MODEL MONITORS` | `NETWORK RULES` | `NOTEBOOKS` | `ONLINE FEATURE TABLES` | `PACKAGES POLICIES` | `PASSWORD POLICIES` | `PIPES` | `PRIVACY POLICIES` | `PROCEDURES` | `PROJECTION POLICIES` | `ROW ACCESS POLICIES` | `SECRETS` | `SEMANTIC VIEWS` | `SERVICES` | `SESSION POLICIES` | `SEQUENCES` | `SNAPSHOTS` | `SNAPSHOT POLICIES` | `SNAPSHOT SETS` | `STAGES` | `STREAMS` | `STREAMLITS` | `TABLES` | `TAGS` | `TASKS` | `VIEWS` | `WORKSPACES`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/grant_privileges_to_account_role#object_type_plural GrantPrivilegesToAccountRole#object_type_plural}
	ObjectTypePlural *string `field:"required" json:"objectTypePlural" yaml:"objectTypePlural"`
	// If true, the inherited privilege will be granted on all objects of the given type in the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/grant_privileges_to_account_role#in_account GrantPrivilegesToAccountRole#in_account}
	InAccount interface{} `field:"optional" json:"inAccount" yaml:"inAccount"`
	// The fully qualified name of the database in which the inherited privilege will be granted on all objects of the given type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/grant_privileges_to_account_role#in_database GrantPrivilegesToAccountRole#in_database}
	InDatabase *string `field:"optional" json:"inDatabase" yaml:"inDatabase"`
	// The fully qualified name of the schema in which the inherited privilege will be granted on all objects of the given type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/grant_privileges_to_account_role#in_schema GrantPrivilegesToAccountRole#in_schema}
	InSchema *string `field:"optional" json:"inSchema" yaml:"inSchema"`
}

