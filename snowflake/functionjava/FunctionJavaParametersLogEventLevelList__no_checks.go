// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionjava

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionJavaParametersLogEventLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionJavaParametersLogEventLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionJavaParametersLogEventLevelList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionJavaParametersLogEventLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionJavaParametersLogEventLevelList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionJavaParametersLogEventLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionJavaParametersLogEventLevelListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

