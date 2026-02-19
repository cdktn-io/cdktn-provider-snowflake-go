// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package notebook

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NotebookFromList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NotebookFromList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NotebookFromList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NotebookFromList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NotebookFromList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NotebookFromList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NotebookFromList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNotebookFromListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

