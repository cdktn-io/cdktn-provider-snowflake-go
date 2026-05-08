// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sessionpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SessionPolicyDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicyDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SessionPolicyDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SessionPolicyDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SessionPolicyDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SessionPolicyDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSessionPolicyDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

