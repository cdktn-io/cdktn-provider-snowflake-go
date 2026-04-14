// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package authenticationpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AuthenticationPolicyClientPolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AuthenticationPolicyClientPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AuthenticationPolicyClientPolicyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AuthenticationPolicyClientPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AuthenticationPolicyClientPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AuthenticationPolicyClientPolicyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AuthenticationPolicyClientPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAuthenticationPolicyClientPolicyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

