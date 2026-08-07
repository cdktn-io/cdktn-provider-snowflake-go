// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package fileformatjson

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FileFormatJsonDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FileFormatJsonDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FileFormatJsonDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FileFormatJsonDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FileFormatJsonDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FileFormatJsonDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFileFormatJsonDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

