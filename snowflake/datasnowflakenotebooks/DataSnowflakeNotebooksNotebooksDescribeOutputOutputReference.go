// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakenotebooks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/datasnowflakenotebooks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference interface {
	cdktn.ComplexObject
	CodeWarehouse() *string
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
	ComputePool() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultPackages() *string
	DefaultVersion() *string
	DefaultVersionAlias() *string
	DefaultVersionGitCommitHash() *string
	DefaultVersionLocationUri() *string
	DefaultVersionName() *string
	DefaultVersionSourceLocationUri() *string
	ExternalAccessIntegrations() *string
	ExternalAccessSecrets() *string
	// Experimental.
	Fqn() *string
	IdleAutoShutdownTimeSeconds() *float64
	ImportUrls() *string
	InternalValue() *DataSnowflakeNotebooksNotebooksDescribeOutput
	SetInternalValue(val *DataSnowflakeNotebooksNotebooksDescribeOutput)
	LastVersionAlias() *string
	LastVersionGitCommitHash() *string
	LastVersionLocationUri() *string
	LastVersionName() *string
	LastVersionSourceLocationUri() *string
	LiveVersionLocationUri() *string
	MainFile() *string
	Name() *string
	Owner() *string
	QueryWarehouse() *string
	RuntimeEnvironmentVersion() *string
	RuntimeName() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Title() *string
	UrlId() *string
	UserPackages() *string
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

// The jsii proxy struct for DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference
type jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) CodeWarehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) ComputePool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) DefaultPackages() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) DefaultVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) DefaultVersionAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) DefaultVersionGitCommitHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionGitCommitHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) DefaultVersionLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) DefaultVersionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) DefaultVersionSourceLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionSourceLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) ExternalAccessIntegrations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAccessIntegrations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) ExternalAccessSecrets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAccessSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) IdleAutoShutdownTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleAutoShutdownTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) ImportUrls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) InternalValue() *DataSnowflakeNotebooksNotebooksDescribeOutput {
	var returns *DataSnowflakeNotebooksNotebooksDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) LastVersionAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) LastVersionGitCommitHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionGitCommitHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) LastVersionLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) LastVersionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) LastVersionSourceLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionSourceLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) LiveVersionLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveVersionLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) MainFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) QueryWarehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) RuntimeEnvironmentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEnvironmentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) RuntimeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) UrlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) UserPackages() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPackages",
		&returns,
	)
	return returns
}


func NewDataSnowflakeNotebooksNotebooksDescribeOutputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeNotebooksNotebooksDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeNotebooks.DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeNotebooksNotebooksDescribeOutputOutputReference_Override(d DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeNotebooks.DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference)SetInternalValue(val *DataSnowflakeNotebooksNotebooksDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataSnowflakeNotebooksNotebooksDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

