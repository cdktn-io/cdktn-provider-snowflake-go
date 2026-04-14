// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tag


type TagOnConflict struct {
	// The order of the values in the ALLOWED_VALUES property of the tag determines which value is used when there is a conflict.
	//
	// External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.15.0/docs/resources/tag#allowed_values_sequence Tag#allowed_values_sequence}
	AllowedValuesSequence interface{} `field:"optional" json:"allowedValuesSequence" yaml:"allowedValuesSequence"`
	// Whenever there is a conflict, the value of tag is set to custom_value.
	//
	// If `allowed_values` are set, the value set in this field should be one of the values in the `allowed_values` list. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.15.0/docs/resources/tag#custom_value Tag#custom_value}
	CustomValue *string `field:"optional" json:"customValue" yaml:"customValue"`
}

