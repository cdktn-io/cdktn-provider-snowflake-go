// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package proceduresql

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedureSqlParametersMetricLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedureSqlParametersMetricLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedureSqlParametersMetricLevelList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedureSqlParametersMetricLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedureSqlParametersMetricLevelList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedureSqlParametersMetricLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedureSqlParametersMetricLevelListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

