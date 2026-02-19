// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package stageexternals3

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StageExternalS3DescribeOutputLocationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StageExternalS3DescribeOutputLocationList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StageExternalS3DescribeOutputLocationList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StageExternalS3DescribeOutputLocationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StageExternalS3DescribeOutputLocationList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StageExternalS3DescribeOutputLocationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStageExternalS3DescribeOutputLocationListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

