// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package hybridtable

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HybridTableShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (h *jsiiProxy_HybridTableShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HybridTableShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HybridTableShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HybridTableShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HybridTableShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHybridTableShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

