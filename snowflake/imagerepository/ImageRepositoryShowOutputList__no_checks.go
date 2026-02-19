// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package imagerepository

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ImageRepositoryShowOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_ImageRepositoryShowOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_ImageRepositoryShowOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ImageRepositoryShowOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ImageRepositoryShowOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ImageRepositoryShowOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewImageRepositoryShowOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

