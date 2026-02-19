// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rowaccesspolicy

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RowAccessPolicyShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RowAccessPolicyShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RowAccessPolicyShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RowAccessPolicyShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RowAccessPolicyShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RowAccessPolicyShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRowAccessPolicyShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

