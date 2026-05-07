// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeexternalvolumes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v17/datasnowflakeexternalvolumes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference interface {
	cdktn.ComplexObject
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
	EncryptionKmsKeyId() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocation
	SetInternalValue(val *DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocation)
	StorageAwsAccessPointArn() *string
	StorageAwsExternalId() *string
	StorageAwsIamUserArn() *string
	StorageAwsRoleArn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UsePrivatelinkEndpoint() *string
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

// The jsii proxy struct for DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference
type jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) EncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) InternalValue() *DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocation {
	var returns *DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) StorageAwsAccessPointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAwsAccessPointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) StorageAwsExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAwsExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) StorageAwsIamUserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAwsIamUserArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) StorageAwsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAwsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) UsePrivatelinkEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePrivatelinkEndpoint",
		&returns,
	)
	return returns
}


func NewDataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeExternalVolumes.DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference_Override(d DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeExternalVolumes.DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference)SetInternalValue(val *DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataSnowflakeExternalVolumesExternalVolumesDescribeOutputStorageLocationsS3StorageLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

