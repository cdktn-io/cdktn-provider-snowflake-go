// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package task

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskParametersTransactionDefaultIsolationLevelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TaskParametersTransactionDefaultIsolationLevelList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TaskParametersTransactionDefaultIsolationLevelList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TaskParametersTransactionDefaultIsolationLevelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TaskParametersTransactionDefaultIsolationLevelList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TaskParametersTransactionDefaultIsolationLevelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTaskParametersTransactionDefaultIsolationLevelListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

