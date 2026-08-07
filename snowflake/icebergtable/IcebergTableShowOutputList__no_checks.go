// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package icebergtable

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IcebergTableShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IcebergTableShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IcebergTableShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IcebergTableShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IcebergTableShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IcebergTableShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIcebergTableShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

