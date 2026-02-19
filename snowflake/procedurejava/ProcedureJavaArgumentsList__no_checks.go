// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package procedurejava

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedureJavaArgumentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedureJavaArgumentsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedureJavaArgumentsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaArgumentsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaArgumentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaArgumentsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavaArgumentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedureJavaArgumentsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

