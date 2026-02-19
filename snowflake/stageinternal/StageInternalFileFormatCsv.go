// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageinternal


type StageInternalFileFormatCsv struct {
	// Defines the encoding format for binary input or output. Valid values: `HEX` | `BASE64` | `UTF8`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#binary_format StageInternal#binary_format}
	BinaryFormat *string `field:"optional" json:"binaryFormat" yaml:"binaryFormat"`
	// Specifies the compression format.
	//
	// Valid values: `AUTO` | `GZIP` | `BZ2` | `BROTLI` | `ZSTD` | `DEFLATE` | `RAW_DEFLATE` | `NONE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#compression StageInternal#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// Defines the format of date values in the data files. Use `AUTO` to have Snowflake auto-detect the format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#date_format StageInternal#date_format}
	DateFormat *string `field:"optional" json:"dateFormat" yaml:"dateFormat"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to insert SQL NULL for empty fields in an input file.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#empty_field_as_null StageInternal#empty_field_as_null}
	EmptyFieldAsNull *string `field:"optional" json:"emptyFieldAsNull" yaml:"emptyFieldAsNull"`
	// Specifies the character set of the source data when loading data into a table.
	//
	// Valid values: `BIG5` | `EUCJP` | `EUCKR` | `GB18030` | `IBM420` | `IBM424` | `ISO2022CN` | `ISO2022JP` | `ISO2022KR` | `ISO88591` | `ISO88592` | `ISO88595` | `ISO88596` | `ISO88597` | `ISO88598` | `ISO88599` | `ISO885915` | `KOI8R` | `SHIFTJIS` | `UTF8` | `UTF16` | `UTF16BE` | `UTF16LE` | `UTF32` | `UTF32BE` | `UTF32LE` | `WINDOWS1250` | `WINDOWS1251` | `WINDOWS1252` | `WINDOWS1253` | `WINDOWS1254` | `WINDOWS1255` | `WINDOWS1256`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#encoding StageInternal#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to generate a parsing error if the number of delimited columns in an input file does not match the number of columns in the corresponding table.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#error_on_column_count_mismatch StageInternal#error_on_column_count_mismatch}
	ErrorOnColumnCountMismatch *string `field:"optional" json:"errorOnColumnCountMismatch" yaml:"errorOnColumnCountMismatch"`
	// Single character string used as the escape character for field values.
	//
	// Use `NONE` to specify no escape character. NOTE: This value may be not imported properly from Snowflake. Snowflake returns escaped values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#escape StageInternal#escape}
	Escape *string `field:"optional" json:"escape" yaml:"escape"`
	// Single character string used as the escape character for unenclosed field values only.
	//
	// Use `NONE` to specify no escape character. NOTE: This value may be not imported properly from Snowflake. Snowflake returns escaped values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#escape_unenclosed_field StageInternal#escape_unenclosed_field}
	EscapeUnenclosedField *string `field:"optional" json:"escapeUnenclosedField" yaml:"escapeUnenclosedField"`
	// One or more singlebyte or multibyte characters that separate fields in an input file.
	//
	// Use `NONE` to specify no delimiter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#field_delimiter StageInternal#field_delimiter}
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// Character used to enclose strings. Use `NONE` to specify no enclosure character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#field_optionally_enclosed_by StageInternal#field_optionally_enclosed_by}
	FieldOptionallyEnclosedBy *string `field:"optional" json:"fieldOptionallyEnclosedBy" yaml:"fieldOptionallyEnclosedBy"`
	// Specifies the extension for files unloaded to a stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#file_extension StageInternal#file_extension}
	FileExtension *string `field:"optional" json:"fileExtension" yaml:"fileExtension"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to parse CSV files containing multiple records on a single line.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#multi_line StageInternal#multi_line}
	MultiLine *string `field:"optional" json:"multiLine" yaml:"multiLine"`
	// String used to convert to and from SQL NULL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#null_if StageInternal#null_if}
	NullIf *[]*string `field:"optional" json:"nullIf" yaml:"nullIf"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to use the first row headers in the data files to determine column names.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#parse_header StageInternal#parse_header}
	ParseHeader *string `field:"optional" json:"parseHeader" yaml:"parseHeader"`
	// One or more singlebyte or multibyte characters that separate records in an input file.
	//
	// Use `NONE` to specify no delimiter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#record_delimiter StageInternal#record_delimiter}
	RecordDelimiter *string `field:"optional" json:"recordDelimiter" yaml:"recordDelimiter"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to replace invalid UTF-8 characters with the Unicode replacement character.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#replace_invalid_characters StageInternal#replace_invalid_characters}
	ReplaceInvalidCharacters *string `field:"optional" json:"replaceInvalidCharacters" yaml:"replaceInvalidCharacters"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies to skip any blank lines encountered in the data files.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#skip_blank_lines StageInternal#skip_blank_lines}
	SkipBlankLines *string `field:"optional" json:"skipBlankLines" yaml:"skipBlankLines"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to skip the BOM (byte order mark) if present in a data file.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#skip_byte_order_mark StageInternal#skip_byte_order_mark}
	SkipByteOrderMark *string `field:"optional" json:"skipByteOrderMark" yaml:"skipByteOrderMark"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Number of lines at the start of the file to skip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#skip_header StageInternal#skip_header}
	SkipHeader *float64 `field:"optional" json:"skipHeader" yaml:"skipHeader"`
	// Defines the format of time values in the data files. Use `AUTO` to have Snowflake auto-detect the format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#time_format StageInternal#time_format}
	TimeFormat *string `field:"optional" json:"timeFormat" yaml:"timeFormat"`
	// Defines the format of timestamp values in the data files. Use `AUTO` to have Snowflake auto-detect the format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#timestamp_format StageInternal#timestamp_format}
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Boolean that specifies whether to remove white space from fields.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#trim_space StageInternal#trim_space}
	TrimSpace *string `field:"optional" json:"trimSpace" yaml:"trimSpace"`
}

