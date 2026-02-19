// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalgcs


type StageExternalGcsDirectory struct {
	// Specifies whether to enable a directory table on the external stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_gcs#enable StageExternalGcs#enable}
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether Snowflake should enable triggering automatic refreshes of the directory table metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_gcs#auto_refresh StageExternalGcs#auto_refresh}
	AutoRefresh *string `field:"optional" json:"autoRefresh" yaml:"autoRefresh"`
	// Specifies the name of the notification integration used to automatically refresh the directory table metadata.
	//
	// Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_gcs#notification_integration StageExternalGcs#notification_integration}
	NotificationIntegration *string `field:"optional" json:"notificationIntegration" yaml:"notificationIntegration"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether to automatically refresh the directory table metadata once, immediately after the stage is created.This field is used only when creating the object. Changes on this field are ignored after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_gcs#refresh_on_create StageExternalGcs#refresh_on_create}
	RefreshOnCreate *string `field:"optional" json:"refreshOnCreate" yaml:"refreshOnCreate"`
}

