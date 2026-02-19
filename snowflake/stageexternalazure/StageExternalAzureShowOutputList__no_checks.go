// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package stageexternalazure

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StageExternalAzureShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StageExternalAzureShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StageExternalAzureShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StageExternalAzureShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StageExternalAzureShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StageExternalAzureShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStageExternalAzureShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

