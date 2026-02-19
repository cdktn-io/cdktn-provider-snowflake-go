// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionjavascript

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionJavascriptArgumentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionJavascriptArgumentsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionJavascriptArgumentsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionJavascriptArgumentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionJavascriptArgumentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionJavascriptArgumentsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionJavascriptArgumentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionJavascriptArgumentsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

