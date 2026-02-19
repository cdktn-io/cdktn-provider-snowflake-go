// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package semanticview

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SemanticViewDimensionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SemanticViewDimensionsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SemanticViewDimensionsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SemanticViewDimensionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SemanticViewDimensionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SemanticViewDimensionsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SemanticViewDimensionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSemanticViewDimensionsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

