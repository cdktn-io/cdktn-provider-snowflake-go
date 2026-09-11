// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package openflowconnector

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OpenflowConnectorDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOpenflowConnectorDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

