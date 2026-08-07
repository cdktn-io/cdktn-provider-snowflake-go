// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoaccountrole


type GrantPrivilegesToAccountRoleOnAccountObjectInherited struct {
	// The plural object type of the account object on which an inherited privilege will be granted.
	//
	// Valid values are (case-insensitive): `USERS` | `RESOURCE MONITORS` | `WAREHOUSES` | `COMPUTE POOLS` | `DATABASES` | `INTEGRATIONS` | `CONNECTIONS` | `FAILOVER GROUPS` | `REPLICATION GROUPS` | `EXTERNAL VOLUMES`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/grant_privileges_to_account_role#object_type_plural GrantPrivilegesToAccountRole#object_type_plural}
	ObjectTypePlural *string `field:"required" json:"objectTypePlural" yaml:"objectTypePlural"`
}

