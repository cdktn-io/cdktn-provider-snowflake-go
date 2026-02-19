// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionsql

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionSqlParametersTraceLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionSqlParametersTraceLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionSqlParametersTraceLevelList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionSqlParametersTraceLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionSqlParametersTraceLevelList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionSqlParametersTraceLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionSqlParametersTraceLevelListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

