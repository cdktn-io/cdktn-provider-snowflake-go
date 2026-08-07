// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package mcpserver

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_McpServer) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateImportFromParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateMoveToIdParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (m *jsiiProxy_McpServer) validatePutTimeoutsParameters(value *McpServerTimeouts) error {
	return nil
}

func (m *jsiiProxy_McpServer) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateMcpServer_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateMcpServer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMcpServer_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateMcpServer_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_McpServer) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_McpServer) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_McpServer) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_McpServer) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_McpServer) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_McpServer) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_McpServer) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_McpServer) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_McpServer) validateSetSchemaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_McpServer) validateSetSpecificationParameters(val *string) error {
	return nil
}

func validateNewMcpServerParameters(scope constructs.Construct, id *string, config *McpServerConfig) error {
	return nil
}

