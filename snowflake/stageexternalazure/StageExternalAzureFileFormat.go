// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalazure


type StageExternalAzureFileFormat struct {
	// avro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#avro StageExternalAzure#avro}
	Avro *StageExternalAzureFileFormatAvro `field:"optional" json:"avro" yaml:"avro"`
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#csv StageExternalAzure#csv}
	Csv *StageExternalAzureFileFormatCsv `field:"optional" json:"csv" yaml:"csv"`
	// Fully qualified name of the file format (e.g., 'database.schema.format_name').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#format_name StageExternalAzure#format_name}
	FormatName *string `field:"optional" json:"formatName" yaml:"formatName"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#json StageExternalAzure#json}
	Json *StageExternalAzureFileFormatJson `field:"optional" json:"json" yaml:"json"`
	// orc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#orc StageExternalAzure#orc}
	Orc *StageExternalAzureFileFormatOrc `field:"optional" json:"orc" yaml:"orc"`
	// parquet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#parquet StageExternalAzure#parquet}
	Parquet *StageExternalAzureFileFormatParquet `field:"optional" json:"parquet" yaml:"parquet"`
	// xml block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/stage_external_azure#xml StageExternalAzure#xml}
	Xml *StageExternalAzureFileFormatXml `field:"optional" json:"xml" yaml:"xml"`
}

