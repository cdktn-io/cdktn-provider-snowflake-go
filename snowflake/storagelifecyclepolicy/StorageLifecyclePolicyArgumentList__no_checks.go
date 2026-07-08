// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storagelifecyclepolicy

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageLifecyclePolicyArgumentList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StorageLifecyclePolicyArgumentList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageLifecyclePolicyArgumentList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageLifecyclePolicyArgumentList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageLifecyclePolicyArgumentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageLifecyclePolicyArgumentList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageLifecyclePolicyArgumentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageLifecyclePolicyArgumentListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

