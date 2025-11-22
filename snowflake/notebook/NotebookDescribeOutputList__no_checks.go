// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package notebook

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NotebookDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NotebookDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NotebookDescribeOutputList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NotebookDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NotebookDescribeOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NotebookDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNotebookDescribeOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

