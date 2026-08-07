// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagelifecyclepolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StorageLifecyclePolicyConfig struct {
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
	// argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/storage_lifecycle_policy#argument StorageLifecyclePolicy#argument}
	Argument interface{} `field:"required" json:"argument" yaml:"argument"`
	// Specifies the SQL expression.
	//
	// The expression can be any boolean-valued SQL expression. To mitigate permadiff on this field, the provider replaces blank characters with a space. This can lead to false positives in cases where a change in case or run of whitespace is semantically significant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/storage_lifecycle_policy#body StorageLifecyclePolicy#body}
	Body *string `field:"required" json:"body" yaml:"body"`
	// The database in which to create the storage lifecycle policy.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/storage_lifecycle_policy#database StorageLifecyclePolicy#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the storage lifecycle policy;
	//
	// must be unique for the database and schema in which the storage lifecycle policy is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/storage_lifecycle_policy#name StorageLifecyclePolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the storage lifecycle policy.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/storage_lifecycle_policy#schema StorageLifecyclePolicy#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Specifies the number of days to keep rows that match the policy expression in archive storage.
	//
	// If set, Snowflake moves the data into archive storage according to the value you select for archive_tier. If unset, Snowflake expires the rows from the table without archiving the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/storage_lifecycle_policy#archive_for_days StorageLifecyclePolicy#archive_for_days}
	ArchiveForDays *float64 `field:"optional" json:"archiveForDays" yaml:"archiveForDays"`
	// Specifies the type of storage tier to use for archiving rows.
	//
	// After you set the ARCHIVE_TIER for a policy, you can’t modify it. If you don’t specify this parameter, the policy is an expiration policy that deletes rows without archiving them. Valid values are (case-insensitive): `COOL` | `COLD`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/storage_lifecycle_policy#archive_tier StorageLifecyclePolicy#archive_tier}
	ArchiveTier *string `field:"optional" json:"archiveTier" yaml:"archiveTier"`
	// Specifies a comment for the storage lifecycle policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/storage_lifecycle_policy#comment StorageLifecyclePolicy#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/storage_lifecycle_policy#id StorageLifecyclePolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/storage_lifecycle_policy#timeouts StorageLifecyclePolicy#timeouts}
	Timeouts *StorageLifecyclePolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

