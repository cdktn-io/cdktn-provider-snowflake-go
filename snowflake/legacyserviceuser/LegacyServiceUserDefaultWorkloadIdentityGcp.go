// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package legacyserviceuser


type LegacyServiceUserDefaultWorkloadIdentityGcp struct {
	// The GCP service account subject identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/legacy_service_user#subject LegacyServiceUser#subject}
	Subject *string `field:"required" json:"subject" yaml:"subject"`
}

