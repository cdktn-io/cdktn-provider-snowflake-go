// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeexternalvolumes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/datasnowflakeexternalvolumes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference interface {
	cdktn.ComplexObject
	AzureStorageLocation() DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsAzureStorageLocationList
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EncryptionType() *string
	// Experimental.
	Fqn() *string
	GcsStorageLocation() DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsGcsStorageLocationList
	InternalValue() *DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocations
	SetInternalValue(val *DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocations)
	Name() *string
	S3CompatStorageLocation() DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3CompatStorageLocationList
	S3StorageLocation() DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationList
	StorageAllowedLocations() *[]*string
	StorageBaseUrl() *string
	StorageProvider() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference
type jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) AzureStorageLocation() DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsAzureStorageLocationList {
	var returns DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsAzureStorageLocationList
	_jsii_.Get(
		j,
		"azureStorageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) EncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) GcsStorageLocation() DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsGcsStorageLocationList {
	var returns DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsGcsStorageLocationList
	_jsii_.Get(
		j,
		"gcsStorageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) InternalValue() *DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocations {
	var returns *DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) S3CompatStorageLocation() DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3CompatStorageLocationList {
	var returns DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3CompatStorageLocationList
	_jsii_.Get(
		j,
		"s3CompatStorageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) S3StorageLocation() DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationList {
	var returns DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationList
	_jsii_.Get(
		j,
		"s3StorageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) StorageAllowedLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storageAllowedLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) StorageBaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBaseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) StorageProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeExternalVolumes.DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference_Override(d DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeExternalVolumes.DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference)SetInternalValue(val *DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

