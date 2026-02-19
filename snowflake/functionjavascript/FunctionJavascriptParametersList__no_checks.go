// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionjavascript

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionJavascriptParametersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionJavascriptParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionJavascriptParametersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionJavascriptParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionJavascriptParametersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionJavascriptParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionJavascriptParametersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

