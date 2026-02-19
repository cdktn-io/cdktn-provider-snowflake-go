// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package jobservice

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobServiceShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (j *jsiiProxy_JobServiceShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobServiceShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobServiceShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobServiceShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobServiceShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobServiceShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

