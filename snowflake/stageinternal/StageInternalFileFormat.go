// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageinternal


type StageInternalFileFormat struct {
	// avro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#avro StageInternal#avro}
	Avro *StageInternalFileFormatAvro `field:"optional" json:"avro" yaml:"avro"`
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#csv StageInternal#csv}
	Csv *StageInternalFileFormatCsv `field:"optional" json:"csv" yaml:"csv"`
	// Fully qualified name of the file format (e.g., 'database.schema.format_name').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#format_name StageInternal#format_name}
	FormatName *string `field:"optional" json:"formatName" yaml:"formatName"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#json StageInternal#json}
	Json *StageInternalFileFormatJson `field:"optional" json:"json" yaml:"json"`
	// orc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#orc StageInternal#orc}
	Orc *StageInternalFileFormatOrc `field:"optional" json:"orc" yaml:"orc"`
	// parquet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#parquet StageInternal#parquet}
	Parquet *StageInternalFileFormatParquet `field:"optional" json:"parquet" yaml:"parquet"`
	// xml block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_internal#xml StageInternal#xml}
	Xml *StageInternalFileFormatXml `field:"optional" json:"xml" yaml:"xml"`
}

