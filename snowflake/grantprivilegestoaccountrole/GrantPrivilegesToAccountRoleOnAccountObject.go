// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoaccountrole


type GrantPrivilegesToAccountRoleOnAccountObject struct {
	// inherited block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.21.0/docs/resources/grant_privileges_to_account_role#inherited GrantPrivilegesToAccountRole#inherited}
	Inherited *GrantPrivilegesToAccountRoleOnAccountObjectInherited `field:"optional" json:"inherited" yaml:"inherited"`
	// The fully qualified name of the object on which privileges will be granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.21.0/docs/resources/grant_privileges_to_account_role#object_name GrantPrivilegesToAccountRole#object_name}
	ObjectName *string `field:"optional" json:"objectName" yaml:"objectName"`
	// The object type of the account object on which privileges will be granted.
	//
	// For more information head over to [Snowflake documentation](https://docs.snowflake.com/en/sql-reference/sql/grant-privilege#required-parameters). Known examples (case-insensitive): `USER` | `RESOURCE MONITOR` | `WAREHOUSE` | `COMPUTE POOL` | `DATABASE` | `INTEGRATION` | `CONNECTION` | `FAILOVER GROUP` | `REPLICATION GROUP` | `EXTERNAL VOLUME` | `SNOWFLAKE INTELLIGENCE`. Snowflake validates the type at apply time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.21.0/docs/resources/grant_privileges_to_account_role#object_type GrantPrivilegesToAccountRole#object_type}
	ObjectType *string `field:"optional" json:"objectType" yaml:"objectType"`
}

