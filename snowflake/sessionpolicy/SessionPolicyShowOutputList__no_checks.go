// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sessionpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SessionPolicyShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicyShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SessionPolicyShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SessionPolicyShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SessionPolicyShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SessionPolicyShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSessionPolicyShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

