// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package legacyserviceuser

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LegacyServiceUserParametersAutocommitList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LegacyServiceUserParametersAutocommitList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LegacyServiceUserParametersAutocommitList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersAutocommitList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersAutocommitList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LegacyServiceUserParametersAutocommitList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLegacyServiceUserParametersAutocommitListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

