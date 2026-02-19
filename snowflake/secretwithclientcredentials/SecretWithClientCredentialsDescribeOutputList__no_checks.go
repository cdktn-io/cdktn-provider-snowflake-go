// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package secretwithclientcredentials

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretWithClientCredentialsDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SecretWithClientCredentialsDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecretWithClientCredentialsDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecretWithClientCredentialsDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretWithClientCredentialsDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecretWithClientCredentialsDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecretWithClientCredentialsDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

