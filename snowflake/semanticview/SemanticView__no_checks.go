// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package semanticview

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SemanticView) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validatePutDimensionsParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validatePutFactsParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validatePutMetricsParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validatePutRelationshipsParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validatePutTablesParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_SemanticView) validatePutTimeoutsParameters(value *SemanticViewTimeouts) error {
	return nil
}

func validateSemanticView_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSemanticView_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSemanticView_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSemanticView_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SemanticView) validateSetCommentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SemanticView) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SemanticView) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SemanticView) validateSetDatabaseParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SemanticView) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SemanticView) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_SemanticView) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SemanticView) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_SemanticView) validateSetSchemaParameters(val *string) error {
	return nil
}

func validateNewSemanticViewParameters(scope constructs.Construct, id *string, config *SemanticViewConfig) error {
	return nil
}

