// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cortexagent

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CortexAgentShowOutputProfileList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CortexAgentShowOutputProfileList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CortexAgentShowOutputProfileList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CortexAgentShowOutputProfileList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CortexAgentShowOutputProfileList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CortexAgentShowOutputProfileList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCortexAgentShowOutputProfileListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

