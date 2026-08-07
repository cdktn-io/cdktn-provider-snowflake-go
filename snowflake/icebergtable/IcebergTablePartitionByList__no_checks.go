// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package icebergtable

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IcebergTablePartitionByList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IcebergTablePartitionByList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IcebergTablePartitionByList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IcebergTablePartitionByList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IcebergTablePartitionByList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IcebergTablePartitionByList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IcebergTablePartitionByList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIcebergTablePartitionByListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

