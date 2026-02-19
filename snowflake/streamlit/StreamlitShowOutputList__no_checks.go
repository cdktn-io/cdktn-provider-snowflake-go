// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package streamlit

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StreamlitShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StreamlitShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StreamlitShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StreamlitShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StreamlitShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StreamlitShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStreamlitShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

