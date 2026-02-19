// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package task

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskParametersTaskAutoRetryAttemptsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TaskParametersTaskAutoRetryAttemptsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TaskParametersTaskAutoRetryAttemptsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TaskParametersTaskAutoRetryAttemptsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TaskParametersTaskAutoRetryAttemptsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TaskParametersTaskAutoRetryAttemptsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTaskParametersTaskAutoRetryAttemptsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

