// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalgcs


type StageExternalGcsFileFormat struct {
	// avro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_gcs#avro StageExternalGcs#avro}
	Avro *StageExternalGcsFileFormatAvro `field:"optional" json:"avro" yaml:"avro"`
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_gcs#csv StageExternalGcs#csv}
	Csv *StageExternalGcsFileFormatCsv `field:"optional" json:"csv" yaml:"csv"`
	// Fully qualified name of the file format (e.g., 'database.schema.format_name').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_gcs#format_name StageExternalGcs#format_name}
	FormatName *string `field:"optional" json:"formatName" yaml:"formatName"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_gcs#json StageExternalGcs#json}
	Json *StageExternalGcsFileFormatJson `field:"optional" json:"json" yaml:"json"`
	// orc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_gcs#orc StageExternalGcs#orc}
	Orc *StageExternalGcsFileFormatOrc `field:"optional" json:"orc" yaml:"orc"`
	// parquet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_gcs#parquet StageExternalGcs#parquet}
	Parquet *StageExternalGcsFileFormatParquet `field:"optional" json:"parquet" yaml:"parquet"`
	// xml block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_gcs#xml StageExternalGcs#xml}
	Xml *StageExternalGcsFileFormatXml `field:"optional" json:"xml" yaml:"xml"`
}

