// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable


type IcebergTableAggregationPolicy struct {
	// Aggregation policy name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#policy_name IcebergTable#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Defines which columns uniquely identify an entity within the Iceberg table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table#entity_key IcebergTable#entity_key}
	EntityKey *[]*string `field:"optional" json:"entityKey" yaml:"entityKey"`
}

