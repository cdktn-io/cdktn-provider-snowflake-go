// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package catalogintegrationobjectstorage

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CatalogIntegrationObjectStorageDescribeOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CatalogIntegrationObjectStorageDescribeOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CatalogIntegrationObjectStorageDescribeOutputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CatalogIntegrationObjectStorageDescribeOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CatalogIntegrationObjectStorageDescribeOutputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CatalogIntegrationObjectStorageDescribeOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCatalogIntegrationObjectStorageDescribeOutputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

