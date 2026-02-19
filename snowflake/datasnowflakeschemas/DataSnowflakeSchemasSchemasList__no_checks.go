// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datasnowflakeschemas

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSnowflakeSchemasSchemasList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataSnowflakeSchemasSchemasList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataSnowflakeSchemasSchemasList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeSchemasSchemasList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeSchemasSchemasList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataSnowflakeSchemasSchemasList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataSnowflakeSchemasSchemasListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

