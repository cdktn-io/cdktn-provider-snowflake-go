// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tablestoragelifecyclepolicyattachment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TableStorageLifecyclePolicyAttachmentConfig struct {
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
	// List of the columns the storage lifecycle policy applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/table_storage_lifecycle_policy_attachment#on TableStorageLifecyclePolicyAttachment#on}
	On *[]*string `field:"required" json:"on" yaml:"on"`
	// Fully qualified name of the storage lifecycle policy to attach to the table.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using pipes (`|`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/table_storage_lifecycle_policy_attachment#storage_lifecycle_policy_name TableStorageLifecyclePolicyAttachment#storage_lifecycle_policy_name}
	StorageLifecyclePolicyName *string `field:"required" json:"storageLifecyclePolicyName" yaml:"storageLifecyclePolicyName"`
	// Fully qualified name of the table (or dynamic table) the storage lifecycle policy is attached to.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using pipes (`|`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/table_storage_lifecycle_policy_attachment#table_name TableStorageLifecyclePolicyAttachment#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Specifies the type of the table referenced in `table_name`. Valid values are (case-insensitive): `TABLE` | `DYNAMIC_TABLE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/table_storage_lifecycle_policy_attachment#table_type TableStorageLifecyclePolicyAttachment#table_type}
	TableType *string `field:"required" json:"tableType" yaml:"tableType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/table_storage_lifecycle_policy_attachment#id TableStorageLifecyclePolicyAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/table_storage_lifecycle_policy_attachment#timeouts TableStorageLifecyclePolicyAttachment#timeouts}
	Timeouts *TableStorageLifecyclePolicyAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

