// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package streamlit

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StreamlitDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StreamlitDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StreamlitDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StreamlitDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StreamlitDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StreamlitDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStreamlitDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

