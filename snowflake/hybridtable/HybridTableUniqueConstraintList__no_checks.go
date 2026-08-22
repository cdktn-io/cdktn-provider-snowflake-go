// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package hybridtable

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HybridTableUniqueConstraintList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (h *jsiiProxy_HybridTableUniqueConstraintList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HybridTableUniqueConstraintList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HybridTableUniqueConstraintList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HybridTableUniqueConstraintList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HybridTableUniqueConstraintList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HybridTableUniqueConstraintList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHybridTableUniqueConstraintListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

