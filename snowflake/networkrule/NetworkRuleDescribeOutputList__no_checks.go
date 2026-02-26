// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkrule

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkRuleDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkRuleDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkRuleDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkRuleDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkRuleDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkRuleDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkRuleDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

