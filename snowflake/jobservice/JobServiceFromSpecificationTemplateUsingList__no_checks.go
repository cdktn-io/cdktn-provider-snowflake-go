// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package jobservice

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobServiceFromSpecificationTemplateUsingList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (j *jsiiProxy_JobServiceFromSpecificationTemplateUsingList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobServiceFromSpecificationTemplateUsingList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobServiceFromSpecificationTemplateUsingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobServiceFromSpecificationTemplateUsingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobServiceFromSpecificationTemplateUsingList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobServiceFromSpecificationTemplateUsingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobServiceFromSpecificationTemplateUsingListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

