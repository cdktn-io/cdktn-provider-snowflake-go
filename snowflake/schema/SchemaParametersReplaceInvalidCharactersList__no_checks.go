// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package schema

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SchemaParametersReplaceInvalidCharactersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SchemaParametersReplaceInvalidCharactersList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SchemaParametersReplaceInvalidCharactersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SchemaParametersReplaceInvalidCharactersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SchemaParametersReplaceInvalidCharactersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SchemaParametersReplaceInvalidCharactersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSchemaParametersReplaceInvalidCharactersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

