// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package hybridtable

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HybridTableShowKeysOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (h *jsiiProxy_HybridTableShowKeysOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HybridTableShowKeysOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HybridTableShowKeysOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HybridTableShowKeysOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HybridTableShowKeysOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHybridTableShowKeysOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

