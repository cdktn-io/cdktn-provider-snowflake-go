// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package externaltable

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExternalTableColumnList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_ExternalTableColumnList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ExternalTableColumnList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ExternalTableColumnList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ExternalTableColumnList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExternalTableColumnList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ExternalTableColumnList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewExternalTableColumnListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

