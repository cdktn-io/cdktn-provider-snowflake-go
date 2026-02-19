// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3


type StageExternalS3FileFormat struct {
	// avro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#avro StageExternalS3#avro}
	Avro *StageExternalS3FileFormatAvro `field:"optional" json:"avro" yaml:"avro"`
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#csv StageExternalS3#csv}
	Csv *StageExternalS3FileFormatCsv `field:"optional" json:"csv" yaml:"csv"`
	// Fully qualified name of the file format (e.g., 'database.schema.format_name').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#format_name StageExternalS3#format_name}
	FormatName *string `field:"optional" json:"formatName" yaml:"formatName"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#json StageExternalS3#json}
	Json *StageExternalS3FileFormatJson `field:"optional" json:"json" yaml:"json"`
	// orc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#orc StageExternalS3#orc}
	Orc *StageExternalS3FileFormatOrc `field:"optional" json:"orc" yaml:"orc"`
	// parquet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#parquet StageExternalS3#parquet}
	Parquet *StageExternalS3FileFormatParquet `field:"optional" json:"parquet" yaml:"parquet"`
	// xml block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3#xml StageExternalS3#xml}
	Xml *StageExternalS3FileFormatXml `field:"optional" json:"xml" yaml:"xml"`
}

