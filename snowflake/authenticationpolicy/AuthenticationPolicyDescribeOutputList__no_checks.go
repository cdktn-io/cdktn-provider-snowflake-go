// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package authenticationpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AuthenticationPolicyDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AuthenticationPolicyDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AuthenticationPolicyDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AuthenticationPolicyDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AuthenticationPolicyDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AuthenticationPolicyDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAuthenticationPolicyDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

