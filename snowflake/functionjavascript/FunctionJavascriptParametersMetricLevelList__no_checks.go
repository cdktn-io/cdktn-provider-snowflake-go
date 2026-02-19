// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionjavascript

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionJavascriptParametersMetricLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionJavascriptParametersMetricLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionJavascriptParametersMetricLevelList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionJavascriptParametersMetricLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionJavascriptParametersMetricLevelList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionJavascriptParametersMetricLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionJavascriptParametersMetricLevelListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

