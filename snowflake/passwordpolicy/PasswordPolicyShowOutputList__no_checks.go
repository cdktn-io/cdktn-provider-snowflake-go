// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package passwordpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PasswordPolicyShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PasswordPolicyShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PasswordPolicyShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PasswordPolicyShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PasswordPolicyShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PasswordPolicyShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPasswordPolicyShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

