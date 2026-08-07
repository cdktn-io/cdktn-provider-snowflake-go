// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtablefromdeltafiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/icebergtablefromdeltafiles/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IcebergTableFromDeltaFilesShowOutputOutputReference interface {
	cdktn.ComplexObject
	AutoRefreshStatus() IcebergTableFromDeltaFilesShowOutputAutoRefreshStatusList
	BaseLocation() *string
	CanWriteMetadata() cdktn.IResolvable
	CatalogName() *string
	CatalogNamespace() *string
	CatalogSyncName() *string
	CatalogTableName() *string
	Comment() *string
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
	CreatedOn() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CurrentPartitionSpecId() *float64
	DatabaseName() *string
	ExternalVolumeName() *string
	// Experimental.
	Fqn() *string
	IcebergTableFormatVersion() *float64
	IcebergTableType() *string
	InternalValue() *IcebergTableFromDeltaFilesShowOutput
	SetInternalValue(val *IcebergTableFromDeltaFilesShowOutput)
	Name() *string
	NameMapping() *string
	Owner() *string
	OwnerRoleType() *string
	PartitionSpecs() IcebergTableFromDeltaFilesShowOutputPartitionSpecsList
	SchemaName() *string
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

// The jsii proxy struct for IcebergTableFromDeltaFilesShowOutputOutputReference
type jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) AutoRefreshStatus() IcebergTableFromDeltaFilesShowOutputAutoRefreshStatusList {
	var returns IcebergTableFromDeltaFilesShowOutputAutoRefreshStatusList
	_jsii_.Get(
		j,
		"autoRefreshStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) BaseLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) CanWriteMetadata() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"canWriteMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) CatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) CatalogNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) CatalogSyncName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogSyncName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) CatalogTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) CurrentPartitionSpecId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"currentPartitionSpecId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) ExternalVolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalVolumeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) IcebergTableFormatVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icebergTableFormatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) IcebergTableType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icebergTableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) InternalValue() *IcebergTableFromDeltaFilesShowOutput {
	var returns *IcebergTableFromDeltaFilesShowOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) NameMapping() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) OwnerRoleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerRoleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) PartitionSpecs() IcebergTableFromDeltaFilesShowOutputPartitionSpecsList {
	var returns IcebergTableFromDeltaFilesShowOutputPartitionSpecsList
	_jsii_.Get(
		j,
		"partitionSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIcebergTableFromDeltaFilesShowOutputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IcebergTableFromDeltaFilesShowOutputOutputReference {
	_init_.Initialize()

	if err := validateNewIcebergTableFromDeltaFilesShowOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.icebergTableFromDeltaFiles.IcebergTableFromDeltaFilesShowOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIcebergTableFromDeltaFilesShowOutputOutputReference_Override(i IcebergTableFromDeltaFilesShowOutputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.icebergTableFromDeltaFiles.IcebergTableFromDeltaFilesShowOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference)SetInternalValue(val *IcebergTableFromDeltaFilesShowOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableFromDeltaFilesShowOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

