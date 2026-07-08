// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionpython

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionPythonParametersLogEventLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionPythonParametersLogEventLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionPythonParametersLogEventLevelList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionPythonParametersLogEventLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionPythonParametersLogEventLevelList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionPythonParametersLogEventLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionPythonParametersLogEventLevelListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

