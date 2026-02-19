// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package functionpython

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionPythonImportsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FunctionPythonImportsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionPythonImportsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionPythonImportsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionPythonImportsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionPythonImportsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionPythonImportsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionPythonImportsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

