// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package procedurejavascript

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProcedureJavascriptShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ProcedureJavascriptShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ProcedureJavascriptShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavascriptShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavascriptShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ProcedureJavascriptShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewProcedureJavascriptShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

