// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package passwordpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PasswordPolicyConfig struct {
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
	// The database this password policy belongs to.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#database PasswordPolicy#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Identifier for the password policy;
	//
	// must be unique for your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#name PasswordPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema this password policy belongs to.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#schema PasswordPolicy#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Adds a comment or overwrites an existing comment for the password policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#comment PasswordPolicy#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Specifies the number of the most recent passwords that Snowflake stores.
	//
	// These stored passwords cannot be repeated when a user updates their password value. The current password value does not count towards the history. When you increase the history value, Snowflake saves the previous values. When you decrease the value, Snowflake saves the stored values up to that value that is set. For example, if the history value is 8 and you change the history value to 3, Snowflake stores the most recent 3 passwords and deletes the 5 older password values from the history.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#history PasswordPolicy#history}
	History *float64 `field:"optional" json:"history" yaml:"history"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#id PasswordPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// (Default: `false`) Prevent overwriting a previous password policy with the same name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#if_not_exists PasswordPolicy#if_not_exists}
	IfNotExists interface{} `field:"optional" json:"ifNotExists" yaml:"ifNotExists"`
	// Specifies the number of minutes the user account will be locked after exhausting the designated number of password retries (i.e. PASSWORD_MAX_RETRIES).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#lockout_time_mins PasswordPolicy#lockout_time_mins}
	LockoutTimeMins *float64 `field:"optional" json:"lockoutTimeMins" yaml:"lockoutTimeMins"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Specifies the maximum number of days before the password must be changed.
	//
	// A value of zero (i.e. 0) indicates that the password does not need to be changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#max_age_days PasswordPolicy#max_age_days}
	MaxAgeDays *float64 `field:"optional" json:"maxAgeDays" yaml:"maxAgeDays"`
	// Specifies the maximum number of characters the password must contain.
	//
	// This number must be greater than or equal to the sum of PASSWORD_MIN_LENGTH, PASSWORD_MIN_UPPER_CASE_CHARS, and PASSWORD_MIN_LOWER_CASE_CHARS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#max_length PasswordPolicy#max_length}
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// Specifies the maximum number of attempts to enter a password before being locked out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#max_retries PasswordPolicy#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Specifies the number of days the user must wait before a recently changed password can be changed again.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#min_age_days PasswordPolicy#min_age_days}
	MinAgeDays *float64 `field:"optional" json:"minAgeDays" yaml:"minAgeDays"`
	// Specifies the minimum number of characters the password must contain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#min_length PasswordPolicy#min_length}
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Specifies the minimum number of lowercase characters the password must contain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#min_lower_case_chars PasswordPolicy#min_lower_case_chars}
	MinLowerCaseChars *float64 `field:"optional" json:"minLowerCaseChars" yaml:"minLowerCaseChars"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Specifies the minimum number of numeric characters the password must contain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#min_numeric_chars PasswordPolicy#min_numeric_chars}
	MinNumericChars *float64 `field:"optional" json:"minNumericChars" yaml:"minNumericChars"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Specifies the minimum number of special characters the password must contain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#min_special_chars PasswordPolicy#min_special_chars}
	MinSpecialChars *float64 `field:"optional" json:"minSpecialChars" yaml:"minSpecialChars"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Specifies the minimum number of uppercase characters the password must contain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#min_upper_case_chars PasswordPolicy#min_upper_case_chars}
	MinUpperCaseChars *float64 `field:"optional" json:"minUpperCaseChars" yaml:"minUpperCaseChars"`
	// (Default: `false`) Whether to override a previous password policy with the same name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#or_replace PasswordPolicy#or_replace}
	OrReplace interface{} `field:"optional" json:"orReplace" yaml:"orReplace"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/password_policy#timeouts PasswordPolicy#timeouts}
	Timeouts *PasswordPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

