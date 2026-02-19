// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package serviceuser

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceUserParametersAbortDetachedQueryList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceUserParametersAbortDetachedQueryList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceUserParametersAbortDetachedQueryList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersAbortDetachedQueryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersAbortDetachedQueryList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceUserParametersAbortDetachedQueryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceUserParametersAbortDetachedQueryListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

