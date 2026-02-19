// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package procedurejava

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedureJavaParametersLogLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedureJavaParametersLogLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedureJavaParametersLogLevelList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaParametersLogLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaParametersLogLevelList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaParametersLogLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedureJavaParametersLogLevelListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

