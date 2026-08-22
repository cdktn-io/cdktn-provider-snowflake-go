// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package hybridtable

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HybridTableDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (h *jsiiProxy_HybridTableDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HybridTableDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HybridTableDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HybridTableDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HybridTableDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHybridTableDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

