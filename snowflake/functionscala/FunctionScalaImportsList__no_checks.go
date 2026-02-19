// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionscala

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionScalaImportsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionScalaImportsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionScalaImportsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaImportsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaImportsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaImportsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionScalaImportsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionScalaImportsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

