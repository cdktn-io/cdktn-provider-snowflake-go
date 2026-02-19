// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionjava

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionJavaArgumentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionJavaArgumentsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionJavaArgumentsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionJavaArgumentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionJavaArgumentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionJavaArgumentsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionJavaArgumentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionJavaArgumentsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

