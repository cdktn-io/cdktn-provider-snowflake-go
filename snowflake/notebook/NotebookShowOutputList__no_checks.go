// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package notebook

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NotebookShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NotebookShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NotebookShowOutputList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NotebookShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NotebookShowOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NotebookShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNotebookShowOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

