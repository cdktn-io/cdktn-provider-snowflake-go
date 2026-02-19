// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageinternal


type StageInternalEncryption struct {
	// snowflake_full block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#snowflake_full StageInternal#snowflake_full}
	SnowflakeFull *StageInternalEncryptionSnowflakeFull `field:"optional" json:"snowflakeFull" yaml:"snowflakeFull"`
	// snowflake_sse block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#snowflake_sse StageInternal#snowflake_sse}
	SnowflakeSse *StageInternalEncryptionSnowflakeSse `field:"optional" json:"snowflakeSse" yaml:"snowflakeSse"`
}

