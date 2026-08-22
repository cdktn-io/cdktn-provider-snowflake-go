// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package hybridtable

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HybridTableForeignKeyConstraintList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (h *jsiiProxy_HybridTableForeignKeyConstraintList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HybridTableForeignKeyConstraintList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HybridTableForeignKeyConstraintList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HybridTableForeignKeyConstraintList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HybridTableForeignKeyConstraintList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HybridTableForeignKeyConstraintList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHybridTableForeignKeyConstraintListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

