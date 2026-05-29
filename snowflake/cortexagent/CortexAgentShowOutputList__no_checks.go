// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cortexagent

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CortexAgentShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CortexAgentShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CortexAgentShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CortexAgentShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CortexAgentShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CortexAgentShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCortexAgentShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

