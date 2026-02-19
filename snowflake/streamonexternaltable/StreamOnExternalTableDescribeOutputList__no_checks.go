// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package streamonexternaltable

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StreamOnExternalTableDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StreamOnExternalTableDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StreamOnExternalTableDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StreamOnExternalTableDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StreamOnExternalTableDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StreamOnExternalTableDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStreamOnExternalTableDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

