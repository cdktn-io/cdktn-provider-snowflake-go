// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package procedurepython

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedurePythonParametersMetricLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedurePythonParametersMetricLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedurePythonParametersMetricLevelList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedurePythonParametersMetricLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedurePythonParametersMetricLevelList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedurePythonParametersMetricLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedurePythonParametersMetricLevelListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

