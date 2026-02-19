// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoaccountrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GrantPrivilegesToAccountRoleConfig struct {
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
	// The fully qualified name of the account role to which privileges will be granted.
	//
	// For more information about this resource, see [docs](./account_role).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/grant_privileges_to_account_role#account_role_name GrantPrivilegesToAccountRole#account_role_name}
	AccountRoleName *string `field:"required" json:"accountRoleName" yaml:"accountRoleName"`
	// (Default: `false`) Grant all privileges on the account role.
	//
	// When all privileges cannot be granted, the provider returns a warning, which is aligned with the Snowsight behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/grant_privileges_to_account_role#all_privileges GrantPrivilegesToAccountRole#all_privileges}
	AllPrivileges interface{} `field:"optional" json:"allPrivileges" yaml:"allPrivileges"`
	// (Default: `false`) If true, the resource will always produce a “plan” and on “apply” it will re-grant defined privileges.
	//
	// It is supposed to be used only in “grant privileges on all X’s in database / schema Y” or “grant all privileges to X” scenarios to make sure that every new object in a given database / schema is granted by the account role and every new privilege is granted to the database role. Important note: this flag is not compliant with the Terraform assumptions of the config being eventually convergent (producing an empty plan).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/grant_privileges_to_account_role#always_apply GrantPrivilegesToAccountRole#always_apply}
	AlwaysApply interface{} `field:"optional" json:"alwaysApply" yaml:"alwaysApply"`
	// (Default: ``) This is a helper field and should not be set.
	//
	// Its main purpose is to help to achieve the functionality described by the always_apply field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/grant_privileges_to_account_role#always_apply_trigger GrantPrivilegesToAccountRole#always_apply_trigger}
	AlwaysApplyTrigger *string `field:"optional" json:"alwaysApplyTrigger" yaml:"alwaysApplyTrigger"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/grant_privileges_to_account_role#id GrantPrivilegesToAccountRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// (Default: `false`) If true, the privileges will be granted on the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/grant_privileges_to_account_role#on_account GrantPrivilegesToAccountRole#on_account}
	OnAccount interface{} `field:"optional" json:"onAccount" yaml:"onAccount"`
	// on_account_object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/grant_privileges_to_account_role#on_account_object GrantPrivilegesToAccountRole#on_account_object}
	OnAccountObject *GrantPrivilegesToAccountRoleOnAccountObject `field:"optional" json:"onAccountObject" yaml:"onAccountObject"`
	// on_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/grant_privileges_to_account_role#on_schema GrantPrivilegesToAccountRole#on_schema}
	OnSchema *GrantPrivilegesToAccountRoleOnSchema `field:"optional" json:"onSchema" yaml:"onSchema"`
	// on_schema_object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/grant_privileges_to_account_role#on_schema_object GrantPrivilegesToAccountRole#on_schema_object}
	OnSchemaObject *GrantPrivilegesToAccountRoleOnSchemaObject `field:"optional" json:"onSchemaObject" yaml:"onSchemaObject"`
	// The privileges to grant on the account role. This field is case-sensitive; use only upper-case privileges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/grant_privileges_to_account_role#privileges GrantPrivilegesToAccountRole#privileges}
	Privileges *[]*string `field:"optional" json:"privileges" yaml:"privileges"`
	// (Default: `false`) If true, the resource will revoke all privileges that are not explicitly defined in the config making it a central source of truth for the privileges granted on an object to an account role.
	//
	// If false, the resource will be only concerned with the privileges that are explicitly defined in the config. The potential privilege removals will be planned only after second `terraform apply` run, after setting the flag in resource configuration. This means, the flag update doesn't revoke immediately any externally granted privileges. This is a Terraform limitation, and two steps are needed to properly show the potential privilege changes (e.g., revoking privileges not specified in the configuration) in the plan. External privileges will be detected regardless of their grant option. The parameter can be only used when `GRANTS_STRICT_PRIVILEGE_MANAGEMENT` option is specified in provider block in the [`experimental_features_enabled`](https://registry.terraform.io/providers/snowflakedb/snowflake/latest/docs#experimental_features_enabled-1) field. Regular and future grants are treated separately, meaning, more resources need to be defined to control regular and future grants for a given object and role (and for a given database or schema they're defined in for future grants). See our [Strict privilege management](https://registry.terraform.io/providers/snowflakedb/snowflake/latest/docs/guides/strict_privilege_management) guide for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/grant_privileges_to_account_role#strict_privilege_management GrantPrivilegesToAccountRole#strict_privilege_management}
	StrictPrivilegeManagement interface{} `field:"optional" json:"strictPrivilegeManagement" yaml:"strictPrivilegeManagement"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/grant_privileges_to_account_role#timeouts GrantPrivilegesToAccountRole#timeouts}
	Timeouts *GrantPrivilegesToAccountRoleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// (Default: `false`) Specifies whether the grantee can grant the privileges to other users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/grant_privileges_to_account_role#with_grant_option GrantPrivilegesToAccountRole#with_grant_option}
	WithGrantOption interface{} `field:"optional" json:"withGrantOption" yaml:"withGrantOption"`
}

