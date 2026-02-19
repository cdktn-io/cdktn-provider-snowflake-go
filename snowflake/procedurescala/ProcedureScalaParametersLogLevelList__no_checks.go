// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package procedurescala

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedureScalaParametersLogLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedureScalaParametersLogLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedureScalaParametersLogLevelList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedureScalaParametersLogLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedureScalaParametersLogLevelList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedureScalaParametersLogLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedureScalaParametersLogLevelListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

