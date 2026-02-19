// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionsql

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionSqlParametersMetricLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionSqlParametersMetricLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionSqlParametersMetricLevelList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionSqlParametersMetricLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionSqlParametersMetricLevelList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionSqlParametersMetricLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionSqlParametersMetricLevelListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

