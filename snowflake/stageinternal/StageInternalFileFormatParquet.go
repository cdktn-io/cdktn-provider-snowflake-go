// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageinternal


type StageInternalFileFormatParquet struct {
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to interpret columns with no defined logical data type as UTF-8 text.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#binary_as_text StageInternal#binary_as_text}
	BinaryAsText *string `field:"optional" json:"binaryAsText" yaml:"binaryAsText"`
	// Specifies the compression format. Valid values: `AUTO` | `LZO` | `SNAPPY` | `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#compression StageInternal#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// String used to convert to and from SQL NULL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#null_if StageInternal#null_if}
	NullIf *[]*string `field:"optional" json:"nullIf" yaml:"nullIf"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to replace invalid UTF-8 characters with the Unicode replacement character.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#replace_invalid_characters StageInternal#replace_invalid_characters}
	ReplaceInvalidCharacters *string `field:"optional" json:"replaceInvalidCharacters" yaml:"replaceInvalidCharacters"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to remove white space from fields.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#trim_space StageInternal#trim_space}
	TrimSpace *string `field:"optional" json:"trimSpace" yaml:"trimSpace"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to use Parquet logical types when loading data.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#use_logical_type StageInternal#use_logical_type}
	UseLogicalType *string `field:"optional" json:"useLogicalType" yaml:"useLogicalType"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to use a vectorized scanner for loading Parquet files.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#use_vectorized_scanner StageInternal#use_vectorized_scanner}
	UseVectorizedScanner *string `field:"optional" json:"useVectorizedScanner" yaml:"useVectorizedScanner"`
}

