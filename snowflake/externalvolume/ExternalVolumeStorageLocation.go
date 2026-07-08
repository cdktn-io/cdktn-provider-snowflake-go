// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package externalvolume


type ExternalVolumeStorageLocation struct {
	// Specifies the base URL for your cloud storage location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/external_volume#storage_base_url ExternalVolume#storage_base_url}
	StorageBaseUrl *string `field:"required" json:"storageBaseUrl" yaml:"storageBaseUrl"`
	// Name of the storage location.
	//
	// Must be unique for the external volume. Do not use the name `terraform_provider_sentinel_storage_location` - this is reserved for the provider for performing update operations. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/external_volume#storage_location_name ExternalVolume#storage_location_name}
	StorageLocationName *string `field:"required" json:"storageLocationName" yaml:"storageLocationName"`
	// Specifies the cloud storage provider that stores your data files.
	//
	// Valid values are (case-insensitive): `GCS` | `AZURE` | `S3` | `S3GOV` | `S3COMPAT`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/external_volume#storage_provider ExternalVolume#storage_provider}
	StorageProvider *string `field:"required" json:"storageProvider" yaml:"storageProvider"`
	// Specifies the ID for your Office 365 tenant that the allowed and blocked storage accounts belong to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/external_volume#azure_tenant_id ExternalVolume#azure_tenant_id}
	AzureTenantId *string `field:"optional" json:"azureTenantId" yaml:"azureTenantId"`
	// Specifies the ID for the KMS-managed key used to encrypt files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/external_volume#encryption_kms_key_id ExternalVolume#encryption_kms_key_id}
	EncryptionKmsKeyId *string `field:"optional" json:"encryptionKmsKeyId" yaml:"encryptionKmsKeyId"`
	// Specifies the encryption type used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/external_volume#encryption_type ExternalVolume#encryption_type}
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// Specifies the access point ARN for the S3 bucket containing your data files.
	//
	// Only applicable for S3 and S3GOV storage providers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/external_volume#storage_aws_access_point_arn ExternalVolume#storage_aws_access_point_arn}
	StorageAwsAccessPointArn *string `field:"optional" json:"storageAwsAccessPointArn" yaml:"storageAwsAccessPointArn"`
	// External ID that Snowflake uses to establish a trust relationship with AWS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/external_volume#storage_aws_external_id ExternalVolume#storage_aws_external_id}
	StorageAwsExternalId *string `field:"optional" json:"storageAwsExternalId" yaml:"storageAwsExternalId"`
	// Specifies the AWS key ID for the S3-compatible storage location. Only applicable for S3COMPAT storage provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/external_volume#storage_aws_key_id ExternalVolume#storage_aws_key_id}
	StorageAwsKeyId *string `field:"optional" json:"storageAwsKeyId" yaml:"storageAwsKeyId"`
	// Specifies the case-sensitive Amazon Resource Name (ARN) of the AWS identity and access management (IAM) role that grants privileges on the S3 bucket containing your data files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/external_volume#storage_aws_role_arn ExternalVolume#storage_aws_role_arn}
	StorageAwsRoleArn *string `field:"optional" json:"storageAwsRoleArn" yaml:"storageAwsRoleArn"`
	// Specifies the AWS secret key for the S3-compatible storage location. Only applicable for S3COMPAT storage provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/external_volume#storage_aws_secret_key ExternalVolume#storage_aws_secret_key}
	StorageAwsSecretKey *string `field:"optional" json:"storageAwsSecretKey" yaml:"storageAwsSecretKey"`
	// Specifies the endpoint for the S3-compatible storage location. Only applicable for S3COMPAT storage provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/external_volume#storage_endpoint ExternalVolume#storage_endpoint}
	StorageEndpoint *string `field:"optional" json:"storageEndpoint" yaml:"storageEndpoint"`
	// (Default: fallback to Snowflake default - uses special value that cannot be set in the configuration manually (`default`)) Specifies whether to use a privatelink endpoint for the storage location.
	//
	// Only applicable for S3, S3GOV, and AZURE storage providers. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/external_volume#use_privatelink_endpoint ExternalVolume#use_privatelink_endpoint}
	UsePrivatelinkEndpoint *string `field:"optional" json:"usePrivatelinkEndpoint" yaml:"usePrivatelinkEndpoint"`
}

