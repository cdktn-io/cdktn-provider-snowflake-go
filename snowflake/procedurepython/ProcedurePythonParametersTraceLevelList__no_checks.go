// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package procedurepython

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedurePythonParametersTraceLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedurePythonParametersTraceLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedurePythonParametersTraceLevelList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedurePythonParametersTraceLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedurePythonParametersTraceLevelList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedurePythonParametersTraceLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedurePythonParametersTraceLevelListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

