// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3compatible


type StageExternalS3CompatibleFileFormat struct {
	// avro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#avro StageExternalS3Compatible#avro}
	Avro *StageExternalS3CompatibleFileFormatAvro `field:"optional" json:"avro" yaml:"avro"`
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#csv StageExternalS3Compatible#csv}
	Csv *StageExternalS3CompatibleFileFormatCsv `field:"optional" json:"csv" yaml:"csv"`
	// Fully qualified name of the file format (e.g., 'database.schema.format_name').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#format_name StageExternalS3Compatible#format_name}
	FormatName *string `field:"optional" json:"formatName" yaml:"formatName"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#json StageExternalS3Compatible#json}
	Json *StageExternalS3CompatibleFileFormatJson `field:"optional" json:"json" yaml:"json"`
	// orc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#orc StageExternalS3Compatible#orc}
	Orc *StageExternalS3CompatibleFileFormatOrc `field:"optional" json:"orc" yaml:"orc"`
	// parquet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#parquet StageExternalS3Compatible#parquet}
	Parquet *StageExternalS3CompatibleFileFormatParquet `field:"optional" json:"parquet" yaml:"parquet"`
	// xml block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_s3_compatible#xml StageExternalS3Compatible#xml}
	Xml *StageExternalS3CompatibleFileFormatXml `field:"optional" json:"xml" yaml:"xml"`
}

