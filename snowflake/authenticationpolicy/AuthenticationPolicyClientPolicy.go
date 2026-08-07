// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package authenticationpolicy


type AuthenticationPolicyClientPolicy struct {
	// The client or driver type.
	//
	// Valid values (case-insensitive): `JDBC_DRIVER` | `ODBC_DRIVER` | `PYTHON_DRIVER` | `JAVASCRIPT_DRIVER` | `C_DRIVER` | `GO_DRIVER` | `PHP_DRIVER` | `DOTNET_DRIVER` | `SQL_API` | `SNOWPIPE_STREAMING_CLIENT_SDK` | `PY_CORE` | `SPROC_PYTHON` | `PYTHON_SNOWPARK` | `SQL_ALCHEMY` | `SNOWPARK` | `SNOWFLAKE_CLIENT`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/authentication_policy#client_type AuthenticationPolicy#client_type}
	ClientType *string `field:"required" json:"clientType" yaml:"clientType"`
	// Minimum allowed version for this client/driver type (e.g. '1.14.1').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/resources/authentication_policy#minimum_version AuthenticationPolicy#minimum_version}
	MinimumVersion *string `field:"required" json:"minimumVersion" yaml:"minimumVersion"`
}

