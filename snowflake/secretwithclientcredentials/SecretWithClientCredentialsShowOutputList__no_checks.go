// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package secretwithclientcredentials

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretWithClientCredentialsShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SecretWithClientCredentialsShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecretWithClientCredentialsShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecretWithClientCredentialsShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretWithClientCredentialsShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecretWithClientCredentialsShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecretWithClientCredentialsShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

