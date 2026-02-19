// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storageintegration

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageIntegrationDescribeOutputStorageProviderList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegrationDescribeOutputStorageProviderList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageIntegrationDescribeOutputStorageProviderList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputStorageProviderList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputStorageProviderList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputStorageProviderList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageIntegrationDescribeOutputStorageProviderListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

