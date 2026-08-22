// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accountauthenticationpolicyattachment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AccountAuthenticationPolicyAttachmentConfig struct {
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
	// Fully qualified name of the authentication policy to apply to the current account.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using pipes (`|`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/account_authentication_policy_attachment#authentication_policy AccountAuthenticationPolicyAttachment#authentication_policy}
	AuthenticationPolicy *string `field:"required" json:"authenticationPolicy" yaml:"authenticationPolicy"`
	// If true, attaches the authentication policy to all person users in the current account.
	//
	// Conflicts with `for_all_service_users`. When neither field is set, the policy is attached account-wide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/account_authentication_policy_attachment#for_all_person_users AccountAuthenticationPolicyAttachment#for_all_person_users}
	ForAllPersonUsers interface{} `field:"optional" json:"forAllPersonUsers" yaml:"forAllPersonUsers"`
	// If true, attaches the authentication policy to all service users in the current account.
	//
	// Conflicts with `for_all_person_users`. When neither field is set, the policy is attached account-wide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/account_authentication_policy_attachment#for_all_service_users AccountAuthenticationPolicyAttachment#for_all_service_users}
	ForAllServiceUsers interface{} `field:"optional" json:"forAllServiceUsers" yaml:"forAllServiceUsers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/account_authentication_policy_attachment#id AccountAuthenticationPolicyAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/account_authentication_policy_attachment#timeouts AccountAuthenticationPolicyAttachment#timeouts}
	Timeouts *AccountAuthenticationPolicyAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

