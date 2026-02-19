// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3


type StageExternalS3FileFormatJson struct {
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to allow duplicate object field names (only the last one will be preserved).
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#allow_duplicate StageExternalS3#allow_duplicate}
	AllowDuplicate *string `field:"optional" json:"allowDuplicate" yaml:"allowDuplicate"`
	// Defines the encoding format for binary input or output. Valid values: `HEX` | `BASE64` | `UTF8`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#binary_format StageExternalS3#binary_format}
	BinaryFormat *string `field:"optional" json:"binaryFormat" yaml:"binaryFormat"`
	// Specifies the compression format.
	//
	// Valid values: `AUTO` | `GZIP` | `BZ2` | `BROTLI` | `ZSTD` | `DEFLATE` | `RAW_DEFLATE` | `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#compression StageExternalS3#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Defines the format of date values in the data files. Use `AUTO` to have Snowflake auto-detect the format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#date_format StageExternalS3#date_format}
	DateFormat *string `field:"optional" json:"dateFormat" yaml:"dateFormat"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that enables parsing of octal numbers.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#enable_octal StageExternalS3#enable_octal}
	EnableOctal *string `field:"optional" json:"enableOctal" yaml:"enableOctal"`
	// Specifies the extension for files unloaded to a stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#file_extension StageExternalS3#file_extension}
	FileExtension *string `field:"optional" json:"fileExtension" yaml:"fileExtension"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether UTF-8 encoding errors produce error conditions.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#ignore_utf8_errors StageExternalS3#ignore_utf8_errors}
	IgnoreUtf8Errors *string `field:"optional" json:"ignoreUtf8Errors" yaml:"ignoreUtf8Errors"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to allow multiple records on a single line.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#multi_line StageExternalS3#multi_line}
	MultiLine *string `field:"optional" json:"multiLine" yaml:"multiLine"`
	// String used to convert to and from SQL NULL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#null_if StageExternalS3#null_if}
	NullIf *[]*string `field:"optional" json:"nullIf" yaml:"nullIf"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to replace invalid UTF-8 characters with the Unicode replacement character.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#replace_invalid_characters StageExternalS3#replace_invalid_characters}
	ReplaceInvalidCharacters *string `field:"optional" json:"replaceInvalidCharacters" yaml:"replaceInvalidCharacters"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to skip the BOM (byte order mark) if present in a data file.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#skip_byte_order_mark StageExternalS3#skip_byte_order_mark}
	SkipByteOrderMark *string `field:"optional" json:"skipByteOrderMark" yaml:"skipByteOrderMark"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that instructs the JSON parser to remove object fields or array elements containing null values.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#strip_null_values StageExternalS3#strip_null_values}
	StripNullValues *string `field:"optional" json:"stripNullValues" yaml:"stripNullValues"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that instructs the JSON parser to remove outer brackets.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#strip_outer_array StageExternalS3#strip_outer_array}
	StripOuterArray *string `field:"optional" json:"stripOuterArray" yaml:"stripOuterArray"`
	// Defines the format of time values in the data files. Use `AUTO` to have Snowflake auto-detect the format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#time_format StageExternalS3#time_format}
	TimeFormat *string `field:"optional" json:"timeFormat" yaml:"timeFormat"`
	// Defines the format of timestamp values in the data files. Use `AUTO` to have Snowflake auto-detect the format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#timestamp_format StageExternalS3#timestamp_format}
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to remove white space from fields.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#trim_space StageExternalS3#trim_space}
	TrimSpace *string `field:"optional" json:"trimSpace" yaml:"trimSpace"`
}

