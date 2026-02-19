// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storageintegration

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageIntegrationDescribeOutputCommentList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StorageIntegrationDescribeOutputCommentList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageIntegrationDescribeOutputCommentList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputCommentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputCommentList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageIntegrationDescribeOutputCommentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageIntegrationDescribeOutputCommentListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

