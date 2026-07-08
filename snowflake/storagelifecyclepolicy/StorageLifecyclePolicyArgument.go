// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagelifecyclepolicy


type StorageLifecyclePolicyArgument struct {
	// The argument name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/storage_lifecycle_policy#name StorageLifecyclePolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The argument type. For more information about data types, check [Snowflake docs](https://docs.snowflake.com/en/sql-reference/intro-summary-data-types).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/storage_lifecycle_policy#type StorageLifecyclePolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

