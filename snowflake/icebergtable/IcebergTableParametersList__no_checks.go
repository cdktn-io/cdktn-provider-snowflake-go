// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package icebergtable

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IcebergTableParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IcebergTableParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IcebergTableParametersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IcebergTableParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IcebergTableParametersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IcebergTableParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIcebergTableParametersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

