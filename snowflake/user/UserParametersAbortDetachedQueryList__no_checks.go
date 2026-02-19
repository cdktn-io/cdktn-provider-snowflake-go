// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package user

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserParametersAbortDetachedQueryList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserParametersAbortDetachedQueryList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserParametersAbortDetachedQueryList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserParametersAbortDetachedQueryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserParametersAbortDetachedQueryList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserParametersAbortDetachedQueryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserParametersAbortDetachedQueryListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

