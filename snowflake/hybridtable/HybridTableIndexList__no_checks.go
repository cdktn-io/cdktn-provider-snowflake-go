// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package hybridtable

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HybridTableIndexList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (h *jsiiProxy_HybridTableIndexList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HybridTableIndexList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HybridTableIndexList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HybridTableIndexList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HybridTableIndexList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HybridTableIndexList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHybridTableIndexListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

