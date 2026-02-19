// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package stageexternalgcs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StageExternalGcsDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StageExternalGcsDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StageExternalGcsDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StageExternalGcsDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StageExternalGcsDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StageExternalGcsDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStageExternalGcsDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

