// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package database

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseReplicationEnableToAccountList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseReplicationEnableToAccountList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatabaseReplicationEnableToAccountList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatabaseReplicationEnableToAccountList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatabaseReplicationEnableToAccountList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatabaseReplicationEnableToAccountList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatabaseReplicationEnableToAccountList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatabaseReplicationEnableToAccountListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

