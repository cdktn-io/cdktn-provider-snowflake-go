// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeicebergtables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/datasnowflakeicebergtables/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference interface {
	cdktn.ComplexObject
	Catalog() DataSnowflakeIcebergTablesIcebergTablesParametersCatalogList
	CatalogSync() DataSnowflakeIcebergTablesIcebergTablesParametersCatalogSyncList
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
	DataRetentionTimeInDays() DataSnowflakeIcebergTablesIcebergTablesParametersDataRetentionTimeInDaysList
	EnableDataCompaction() DataSnowflakeIcebergTablesIcebergTablesParametersEnableDataCompactionList
	EnableIcebergMergeOnRead() DataSnowflakeIcebergTablesIcebergTablesParametersEnableIcebergMergeOnReadList
	ExternalVolume() DataSnowflakeIcebergTablesIcebergTablesParametersExternalVolumeList
	// Experimental.
	Fqn() *string
	IcebergMergeOnReadBehavior() DataSnowflakeIcebergTablesIcebergTablesParametersIcebergMergeOnReadBehaviorList
	InternalValue() *DataSnowflakeIcebergTablesIcebergTablesParameters
	SetInternalValue(val *DataSnowflakeIcebergTablesIcebergTablesParameters)
	MaxDataExtensionTimeInDays() DataSnowflakeIcebergTablesIcebergTablesParametersMaxDataExtensionTimeInDaysList
	ReplaceInvalidCharacters() DataSnowflakeIcebergTablesIcebergTablesParametersReplaceInvalidCharactersList
	StorageSerializationPolicy() DataSnowflakeIcebergTablesIcebergTablesParametersStorageSerializationPolicyList
	TargetFileSize() DataSnowflakeIcebergTablesIcebergTablesParametersTargetFileSizeList
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

// The jsii proxy struct for DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference
type jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) Catalog() DataSnowflakeIcebergTablesIcebergTablesParametersCatalogList {
	var returns DataSnowflakeIcebergTablesIcebergTablesParametersCatalogList
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) CatalogSync() DataSnowflakeIcebergTablesIcebergTablesParametersCatalogSyncList {
	var returns DataSnowflakeIcebergTablesIcebergTablesParametersCatalogSyncList
	_jsii_.Get(
		j,
		"catalogSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) DataRetentionTimeInDays() DataSnowflakeIcebergTablesIcebergTablesParametersDataRetentionTimeInDaysList {
	var returns DataSnowflakeIcebergTablesIcebergTablesParametersDataRetentionTimeInDaysList
	_jsii_.Get(
		j,
		"dataRetentionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) EnableDataCompaction() DataSnowflakeIcebergTablesIcebergTablesParametersEnableDataCompactionList {
	var returns DataSnowflakeIcebergTablesIcebergTablesParametersEnableDataCompactionList
	_jsii_.Get(
		j,
		"enableDataCompaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) EnableIcebergMergeOnRead() DataSnowflakeIcebergTablesIcebergTablesParametersEnableIcebergMergeOnReadList {
	var returns DataSnowflakeIcebergTablesIcebergTablesParametersEnableIcebergMergeOnReadList
	_jsii_.Get(
		j,
		"enableIcebergMergeOnRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) ExternalVolume() DataSnowflakeIcebergTablesIcebergTablesParametersExternalVolumeList {
	var returns DataSnowflakeIcebergTablesIcebergTablesParametersExternalVolumeList
	_jsii_.Get(
		j,
		"externalVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) IcebergMergeOnReadBehavior() DataSnowflakeIcebergTablesIcebergTablesParametersIcebergMergeOnReadBehaviorList {
	var returns DataSnowflakeIcebergTablesIcebergTablesParametersIcebergMergeOnReadBehaviorList
	_jsii_.Get(
		j,
		"icebergMergeOnReadBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) InternalValue() *DataSnowflakeIcebergTablesIcebergTablesParameters {
	var returns *DataSnowflakeIcebergTablesIcebergTablesParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) MaxDataExtensionTimeInDays() DataSnowflakeIcebergTablesIcebergTablesParametersMaxDataExtensionTimeInDaysList {
	var returns DataSnowflakeIcebergTablesIcebergTablesParametersMaxDataExtensionTimeInDaysList
	_jsii_.Get(
		j,
		"maxDataExtensionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) ReplaceInvalidCharacters() DataSnowflakeIcebergTablesIcebergTablesParametersReplaceInvalidCharactersList {
	var returns DataSnowflakeIcebergTablesIcebergTablesParametersReplaceInvalidCharactersList
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) StorageSerializationPolicy() DataSnowflakeIcebergTablesIcebergTablesParametersStorageSerializationPolicyList {
	var returns DataSnowflakeIcebergTablesIcebergTablesParametersStorageSerializationPolicyList
	_jsii_.Get(
		j,
		"storageSerializationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) TargetFileSize() DataSnowflakeIcebergTablesIcebergTablesParametersTargetFileSizeList {
	var returns DataSnowflakeIcebergTablesIcebergTablesParametersTargetFileSizeList
	_jsii_.Get(
		j,
		"targetFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataSnowflakeIcebergTablesIcebergTablesParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeIcebergTablesIcebergTablesParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeIcebergTables.DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeIcebergTablesIcebergTablesParametersOutputReference_Override(d DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeIcebergTables.DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference)SetInternalValue(val *DataSnowflakeIcebergTablesIcebergTablesParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataSnowflakeIcebergTablesIcebergTablesParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

