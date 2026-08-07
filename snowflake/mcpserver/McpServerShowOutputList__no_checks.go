// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package mcpserver

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_McpServerShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_McpServerShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_McpServerShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_McpServerShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_McpServerShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_McpServerShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMcpServerShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

