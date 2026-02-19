// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package authenticationpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AuthenticationPolicyShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AuthenticationPolicyShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AuthenticationPolicyShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AuthenticationPolicyShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AuthenticationPolicyShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AuthenticationPolicyShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAuthenticationPolicyShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

