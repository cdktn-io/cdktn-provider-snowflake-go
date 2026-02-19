// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkPolicyDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicyDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicyDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicyDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicyDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicyDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkPolicyDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

