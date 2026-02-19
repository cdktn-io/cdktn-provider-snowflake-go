// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package semanticview

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SemanticViewTablesUniqueList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SemanticViewTablesUniqueList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SemanticViewTablesUniqueList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SemanticViewTablesUniqueList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SemanticViewTablesUniqueList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SemanticViewTablesUniqueList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SemanticViewTablesUniqueList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSemanticViewTablesUniqueListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

