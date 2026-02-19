// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package secretwithgenericstring

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretWithGenericStringDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SecretWithGenericStringDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecretWithGenericStringDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecretWithGenericStringDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretWithGenericStringDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecretWithGenericStringDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecretWithGenericStringDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

