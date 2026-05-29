// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cortexagent

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CortexAgentDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CortexAgentDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CortexAgentDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CortexAgentDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CortexAgentDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CortexAgentDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCortexAgentDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

