// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable


type IcebergTableForeignKeyConstraint struct {
	// The local column(s) the foreign key is defined on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#column IcebergTable#column}
	Column *[]*string `field:"required" json:"column" yaml:"column"`
	// The table that the foreign key references.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#table_name IcebergTable#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Constraint comment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#comment IcebergTable#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Whether the constraint is deferrable (`true`) or not deferrable (`false`).
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#deferrable IcebergTable#deferrable}
	Deferrable *string `field:"optional" json:"deferrable" yaml:"deferrable"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Whether the constraint is enabled (`true`) or disabled (`false`).
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#enable IcebergTable#enable}
	Enable *string `field:"optional" json:"enable" yaml:"enable"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Whether the constraint is enforced (`true`) or not enforced (`false`).
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#enforced IcebergTable#enforced}
	Enforced *string `field:"optional" json:"enforced" yaml:"enforced"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Whether the constraint is initially deferred (`true`) or initially immediate (`false`).
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#initially_deferred IcebergTable#initially_deferred}
	InitiallyDeferred *string `field:"optional" json:"initiallyDeferred" yaml:"initiallyDeferred"`
	// The match type for the foreign key. Valid values are: [FULL SIMPLE PARTIAL].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#match IcebergTable#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// Name of the constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#name IcebergTable#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the action to perform when the referenced primary/unique key is deleted.
	//
	// Valid values are: [CASCADE SET NULL SET DEFAULT RESTRICT NO ACTION].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#on_delete IcebergTable#on_delete}
	OnDelete *string `field:"optional" json:"onDelete" yaml:"onDelete"`
	// Specifies the action to perform when the referenced primary/unique key is updated.
	//
	// Valid values are: [CASCADE SET NULL SET DEFAULT RESTRICT NO ACTION].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#on_update IcebergTable#on_update}
	OnUpdate *string `field:"optional" json:"onUpdate" yaml:"onUpdate"`
	// The column(s) in the referenced table that the foreign key references.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#ref_column IcebergTable#ref_column}
	RefColumn *[]*string `field:"optional" json:"refColumn" yaml:"refColumn"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Whether a constraint in NOVALIDATE mode is taken into account (`true`) or not (`false`) during query rewrite.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#rely IcebergTable#rely}
	Rely *string `field:"optional" json:"rely" yaml:"rely"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Whether to validate existing data on the table when the constraint is created (`true`) or skip validation (`false`).
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/iceberg_table#validate IcebergTable#validate}
	Validate *string `field:"optional" json:"validate" yaml:"validate"`
}

