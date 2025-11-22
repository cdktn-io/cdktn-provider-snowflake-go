// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package task


type TaskSchedule struct {
	// Specifies an interval (in hours) of wait time inserted between runs of the task.
	//
	// Accepts positive integers. (conflicts with `seconds`, `minutes`, and `using_cron`)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/task#hours Task#hours}
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Specifies an interval (in minutes) of wait time inserted between runs of the task.
	//
	// Accepts positive integers. (conflicts with `seconds`, `hours`, and `using_cron`)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/task#minutes Task#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Specifies an interval (in seconds) of wait time inserted between runs of the task.
	//
	// Accepts positive integers. (conflicts with `minutes`, `hours`, and `using_cron`)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/task#seconds Task#seconds}
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
	// Specifies a cron expression and time zone for periodically running the task.
	//
	// Supports a subset of standard cron utility syntax. (conflicts with `seconds`, `minutes`, and `hours`)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.11.0/docs/resources/task#using_cron Task#using_cron}
	UsingCron *string `field:"optional" json:"usingCron" yaml:"usingCron"`
}

