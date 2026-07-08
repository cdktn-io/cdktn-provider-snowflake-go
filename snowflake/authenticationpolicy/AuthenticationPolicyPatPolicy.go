// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package authenticationpolicy


type AuthenticationPolicyPatPolicy struct {
	// Specifies the default expiration time (in days) for a programmatic access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/authentication_policy#default_expiry_in_days AuthenticationPolicy#default_expiry_in_days}
	DefaultExpiryInDays *float64 `field:"optional" json:"defaultExpiryInDays" yaml:"defaultExpiryInDays"`
	// Specifies the maximum number of days that can be set for the expiration time for a programmatic access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/authentication_policy#max_expiry_in_days AuthenticationPolicy#max_expiry_in_days}
	MaxExpiryInDays *float64 `field:"optional" json:"maxExpiryInDays" yaml:"maxExpiryInDays"`
	// Specifies the network policy evaluation for the PAT. Valid values are: `ENFORCED_REQUIRED` | `ENFORCED_NOT_REQUIRED` | `NOT_ENFORCED`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/authentication_policy#network_policy_evaluation AuthenticationPolicy#network_policy_evaluation}
	NetworkPolicyEvaluation *string `field:"optional" json:"networkPolicyEvaluation" yaml:"networkPolicyEvaluation"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) If true, when you generate a programmatic access token for a service user, you must restrict the use of that token to a specific role.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/authentication_policy#require_role_restriction_for_service_users AuthenticationPolicy#require_role_restriction_for_service_users}
	RequireRoleRestrictionForServiceUsers *string `field:"optional" json:"requireRoleRestrictionForServiceUsers" yaml:"requireRoleRestrictionForServiceUsers"`
}

