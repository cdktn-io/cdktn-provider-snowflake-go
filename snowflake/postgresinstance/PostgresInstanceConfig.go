// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package postgresinstance

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PostgresInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Specifies the authentication authority for the Postgres instance. Valid values are (case-insensitive): `POSTGRES` | `POSTGRES_OR_SNOWFLAKE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/postgres_instance#authentication_authority PostgresInstance#authentication_authority}
	AuthenticationAuthority *string `field:"required" json:"authenticationAuthority" yaml:"authenticationAuthority"`
	// Specifies the compute family for the Postgres instance.
	//
	// Valid values are (case-insensitive): `STANDARD_M` | `STANDARD_L` | `STANDARD_XL` | `STANDARD_2XL` | `STANDARD_4XL` | `STANDARD_8XL` | `STANDARD_12XL` | `STANDARD_24XL` | `HIGHMEM_L` | `HIGHMEM_XL` | `HIGHMEM_2XL` | `HIGHMEM_4XL` | `HIGHMEM_8XL` | `HIGHMEM_12XL` | `HIGHMEM_16XL` | `HIGHMEM_24XL` | `HIGHMEM_32XL` | `HIGHMEM_48XL` | `BURST_XS` | `BURST_S` | `BURST_M`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/postgres_instance#compute_family PostgresInstance#compute_family}
	ComputeFamily *string `field:"required" json:"computeFamily" yaml:"computeFamily"`
	// Specifies the identifier for the Postgres instance;
	//
	// must be unique for your account. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/postgres_instance#name PostgresInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the Postgres version for the instance.
	//
	// Note that Snowflake does not allow downgrading; the version can only be upgraded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/postgres_instance#postgres_version PostgresInstance#postgres_version}
	PostgresVersion *float64 `field:"required" json:"postgresVersion" yaml:"postgresVersion"`
	// Specifies the storage size in GB for the Postgres instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/postgres_instance#storage_size_gb PostgresInstance#storage_size_gb}
	StorageSizeGb *float64 `field:"required" json:"storageSizeGb" yaml:"storageSizeGb"`
	// Specifies a comment for the Postgres instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/postgres_instance#comment PostgresInstance#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether the Postgres instance should be configured for high availability.
	//
	// Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/postgres_instance#high_availability PostgresInstance#high_availability}
	HighAvailability *string `field:"optional" json:"highAvailability" yaml:"highAvailability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/postgres_instance#id PostgresInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`-1`)) Specifies the hour (0-23 UTC) at which the maintenance window starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/postgres_instance#maintenance_window_start PostgresInstance#maintenance_window_start}
	MaintenanceWindowStart *float64 `field:"optional" json:"maintenanceWindowStart" yaml:"maintenanceWindowStart"`
	// Specifies the network policy to associate with the Postgres instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/postgres_instance#network_policy PostgresInstance#network_policy}
	NetworkPolicy *string `field:"optional" json:"networkPolicy" yaml:"networkPolicy"`
	// Specifies custom Postgres settings as a JSON string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/postgres_instance#postgres_settings PostgresInstance#postgres_settings}
	PostgresSettings *string `field:"optional" json:"postgresSettings" yaml:"postgresSettings"`
	// Specifies the storage integration for the Postgres instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/postgres_instance#storage_integration PostgresInstance#storage_integration}
	StorageIntegration *string `field:"optional" json:"storageIntegration" yaml:"storageIntegration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/postgres_instance#timeouts PostgresInstance#timeouts}
	Timeouts *PostgresInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

