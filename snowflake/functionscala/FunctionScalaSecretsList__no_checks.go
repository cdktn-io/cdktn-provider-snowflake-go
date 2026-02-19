// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionscala

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionScalaSecretsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionScalaSecretsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionScalaSecretsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaSecretsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaSecretsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaSecretsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaSecretsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionScalaSecretsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

