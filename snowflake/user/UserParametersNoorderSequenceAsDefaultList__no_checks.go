// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package user

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserParametersNoorderSequenceAsDefaultList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserParametersNoorderSequenceAsDefaultList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserParametersNoorderSequenceAsDefaultList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserParametersNoorderSequenceAsDefaultList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserParametersNoorderSequenceAsDefaultList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserParametersNoorderSequenceAsDefaultList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserParametersNoorderSequenceAsDefaultListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

