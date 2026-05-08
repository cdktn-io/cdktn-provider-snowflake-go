// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sessionpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SessionPolicy) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validatePutAllowedSecondaryRolesParameters(value *SessionPolicyAllowedSecondaryRoles) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validatePutBlockedSecondaryRolesParameters(value *SessionPolicyBlockedSecondaryRoles) error {
	return nil
}

func (s *jsiiProxy_SessionPolicy) validatePutTimeoutsParameters(value *SessionPolicyTimeouts) error {
	return nil
}

func validateSessionPolicy_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSessionPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSessionPolicy_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSessionPolicy_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SessionPolicy) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SessionPolicy) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SessionPolicy) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SessionPolicy) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SessionPolicy) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SessionPolicy) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_SessionPolicy) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SessionPolicy) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_SessionPolicy) validateSetSchemaParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SessionPolicy) validateSetSessionIdleTimeoutMinsParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_SessionPolicy) validateSetSessionUiIdleTimeoutMinsParameters(val *float64) error {
	return nil
}

func validateNewSessionPolicyParameters(scope constructs.Construct, id *string, config *SessionPolicyConfig) error {
	return nil
}

