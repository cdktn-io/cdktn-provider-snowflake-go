// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package externalaccessintegration

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalAccessIntegrationDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_ExternalAccessIntegrationDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ExternalAccessIntegrationDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ExternalAccessIntegrationDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalAccessIntegrationDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ExternalAccessIntegrationDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewExternalAccessIntegrationDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

