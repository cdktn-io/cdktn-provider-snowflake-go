// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package usersessionpolicyattachment


type UserSessionPolicyAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/user_session_policy_attachment#create UserSessionPolicyAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/user_session_policy_attachment#delete UserSessionPolicyAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/user_session_policy_attachment#read UserSessionPolicyAttachment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/user_session_policy_attachment#update UserSessionPolicyAttachment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

