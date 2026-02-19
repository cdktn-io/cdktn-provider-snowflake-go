// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package task


type TaskTargetCompletionInterval struct {
	// Specifies the target completion interval in hours. (conflicts with `minutes` and `seconds`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/task#hours Task#hours}
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Specifies the target completion interval in minutes. (conflicts with `hours` and `seconds`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/task#minutes Task#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Specifies the target completion interval in seconds. (conflicts with `hours` and `minutes`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/task#seconds Task#seconds}
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
}

