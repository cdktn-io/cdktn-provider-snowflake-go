// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package serviceuser

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceUserParametersSearchPathList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceUserParametersSearchPathList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceUserParametersSearchPathList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersSearchPathList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersSearchPathList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersSearchPathList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceUserParametersSearchPathListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

