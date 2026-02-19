// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageinternal


type StageInternalDirectory struct {
	// Specifies whether to enable a directory table on the internal named stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#enable StageInternal#enable}
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether Snowflake should automatically refresh the directory table metadata when new or updated data files are available on the internal named stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#auto_refresh StageInternal#auto_refresh}
	AutoRefresh *string `field:"optional" json:"autoRefresh" yaml:"autoRefresh"`
}

