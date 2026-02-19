// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package streamontable

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StreamOnTableShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StreamOnTableShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StreamOnTableShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StreamOnTableShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StreamOnTableShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StreamOnTableShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStreamOnTableShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

