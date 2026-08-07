// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tag

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TagConfig struct {
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
	// The database in which to create the tag.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/tag#database Tag#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Specifies the identifier for the tag;
	//
	// must be unique for the database in which the tag is created. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/tag#name Tag#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema in which to create the tag.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/tag#schema Tag#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// Set of allowed values for the tag (unordered).
	//
	// When specified, only these values can be assigned. When the `TAGS_ALLOW_EMPTY_ALLOWED_VALUES` experiment is enabled, removing this field from the configuration reverts the tag to accepting any value. Conflicts with `no_allowed_values` and `ordered_allowed_values`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/tag#allowed_values Tag#allowed_values}
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// Specifies a comment for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/tag#comment Tag#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/tag#id Tag#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of masking policies for the tag.
	//
	// A tag can support one masking policy for each data type. If masking policies are assigned to the tag, before dropping the tag, the provider automatically unassigns them. For more information about this resource, see [docs](./masking_policy).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/tag#masking_policies Tag#masking_policies}
	MaskingPolicies *[]*string `field:"optional" json:"maskingPolicies" yaml:"maskingPolicies"`
	// When set to true, the tag explicitly disallows any value from being assigned.
	//
	// This is different from omitting `allowed_values`, which means any value is accepted. Available only when the `TAGS_ALLOW_EMPTY_ALLOWED_VALUES` experiment is enabled. Conflicts with `allowed_values` and `ordered_allowed_values`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/tag#no_allowed_values Tag#no_allowed_values}
	NoAllowedValues interface{} `field:"optional" json:"noAllowedValues" yaml:"noAllowedValues"`
	// on_conflict block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/tag#on_conflict Tag#on_conflict}
	OnConflict *TagOnConflict `field:"optional" json:"onConflict" yaml:"onConflict"`
	// Ordered list of allowed values for the tag.
	//
	// The order is preserved in Snowflake and is significant when `on_conflict.allowed_values_sequence` is used — the first matching value in the sequence wins. Use this instead of `allowed_values` when order matters. Conflicts with `allowed_values` and `no_allowed_values`. Note: transitioning from `allowed_values` to `ordered_allowed_values` always plans an update-in-place for this field, even when the configured order already matches the order stored in Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/tag#ordered_allowed_values Tag#ordered_allowed_values}
	OrderedAllowedValues *[]*string `field:"optional" json:"orderedAllowedValues" yaml:"orderedAllowedValues"`
	// Specifies that the tag will be automatically propagated from source objects to target objects.
	//
	// See more about tag propagation in the [official documentation](https://docs.snowflake.com/en/user-guide/object-tagging/propagation). Valid options are: `NONE` | `ON_DEPENDENCY` | `ON_DATA_MOVEMENT` | `ON_DEPENDENCY_AND_DATA_MOVEMENT`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/tag#propagate Tag#propagate}
	Propagate *string `field:"optional" json:"propagate" yaml:"propagate"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/tag#timeouts Tag#timeouts}
	Timeouts *TagTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

