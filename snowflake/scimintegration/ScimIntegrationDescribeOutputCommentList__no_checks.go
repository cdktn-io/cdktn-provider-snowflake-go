// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package scimintegration

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScimIntegrationDescribeOutputCommentList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ScimIntegrationDescribeOutputCommentList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ScimIntegrationDescribeOutputCommentList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ScimIntegrationDescribeOutputCommentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ScimIntegrationDescribeOutputCommentList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ScimIntegrationDescribeOutputCommentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewScimIntegrationDescribeOutputCommentListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

