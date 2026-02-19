// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package stageexternalgcs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StageExternalGcsShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StageExternalGcsShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StageExternalGcsShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StageExternalGcsShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StageExternalGcsShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StageExternalGcsShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStageExternalGcsShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

