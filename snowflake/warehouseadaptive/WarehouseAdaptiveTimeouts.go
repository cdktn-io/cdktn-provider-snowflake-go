// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package warehouseadaptive


type WarehouseAdaptiveTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/warehouse_adaptive#create WarehouseAdaptive#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/warehouse_adaptive#delete WarehouseAdaptive#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/warehouse_adaptive#read WarehouseAdaptive#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.16.0/docs/resources/warehouse_adaptive#update WarehouseAdaptive#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

