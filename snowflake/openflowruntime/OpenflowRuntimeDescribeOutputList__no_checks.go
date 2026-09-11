// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package openflowruntime

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OpenflowRuntimeDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OpenflowRuntimeDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OpenflowRuntimeDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OpenflowRuntimeDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OpenflowRuntimeDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OpenflowRuntimeDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOpenflowRuntimeDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

