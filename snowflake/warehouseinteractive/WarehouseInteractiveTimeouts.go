// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package warehouseinteractive


type WarehouseInteractiveTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#create WarehouseInteractive#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#delete WarehouseInteractive#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#read WarehouseInteractive#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/warehouse_interactive#update WarehouseInteractive#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

