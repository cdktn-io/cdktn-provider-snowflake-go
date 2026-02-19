// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package view

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_ViewDataMetricFunctionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (v *jsiiProxy_ViewDataMetricFunctionList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_ViewDataMetricFunctionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ViewDataMetricFunctionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ViewDataMetricFunctionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ViewDataMetricFunctionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ViewDataMetricFunctionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewViewDataMetricFunctionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

