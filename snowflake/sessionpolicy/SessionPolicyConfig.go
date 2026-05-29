// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sessionpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SessionPolicyConfig struct {
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
	// The database in which to create the session policy.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/session_policy#database SessionPolicy#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the session policy.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/session_policy#name SessionPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the session policy.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/session_policy#schema SessionPolicy#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// allowed_secondary_roles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/session_policy#allowed_secondary_roles SessionPolicy#allowed_secondary_roles}
	AllowedSecondaryRoles *SessionPolicyAllowedSecondaryRoles `field:"optional" json:"allowedSecondaryRoles" yaml:"allowedSecondaryRoles"`
	// blocked_secondary_roles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/session_policy#blocked_secondary_roles SessionPolicy#blocked_secondary_roles}
	BlockedSecondaryRoles *SessionPolicyBlockedSecondaryRoles `field:"optional" json:"blockedSecondaryRoles" yaml:"blockedSecondaryRoles"`
	// Specifies a comment for the session policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/session_policy#comment SessionPolicy#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/session_policy#id SessionPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// For Snowflake clients and programmatic clients, specifies the number of minutes in which a session can be idle before users must authenticate to Snowflake again.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/session_policy#session_idle_timeout_mins SessionPolicy#session_idle_timeout_mins}
	SessionIdleTimeoutMins *float64 `field:"optional" json:"sessionIdleTimeoutMins" yaml:"sessionIdleTimeoutMins"`
	// For Snowsight, specifies the number of minutes in which a session can be idle before users must authenticate to Snowflake again.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/session_policy#session_ui_idle_timeout_mins SessionPolicy#session_ui_idle_timeout_mins}
	SessionUiIdleTimeoutMins *float64 `field:"optional" json:"sessionUiIdleTimeoutMins" yaml:"sessionUiIdleTimeoutMins"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.17.0/docs/resources/session_policy#timeouts SessionPolicy#timeouts}
	Timeouts *SessionPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

