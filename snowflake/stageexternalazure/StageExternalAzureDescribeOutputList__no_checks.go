// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package stageexternalazure

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StageExternalAzureDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StageExternalAzureDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StageExternalAzureDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStageExternalAzureDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

