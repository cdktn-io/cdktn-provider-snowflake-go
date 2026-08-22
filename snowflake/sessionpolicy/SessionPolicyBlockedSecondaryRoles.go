// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sessionpolicy


type SessionPolicyBlockedSecondaryRoles struct {
	// When true, disallows all secondary roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/session_policy#all SessionPolicy#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// When true, allows all secondary roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/session_policy#none SessionPolicy#none}
	None interface{} `field:"optional" json:"none" yaml:"none"`
	// Specifies roles to be blocked as secondary roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/session_policy#roles SessionPolicy#roles}
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
}

