// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package materializedview

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MaterializedViewTagList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MaterializedViewTagList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MaterializedViewTagList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MaterializedViewTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MaterializedViewTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaterializedViewTagList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MaterializedViewTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMaterializedViewTagListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

