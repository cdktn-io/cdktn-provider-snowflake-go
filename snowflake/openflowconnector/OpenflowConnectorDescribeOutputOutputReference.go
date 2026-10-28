// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package openflowconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v19/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v19/openflowconnector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpenflowConnectorDescribeOutputOutputReference interface {
	cdktn.ComplexObject
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
	ConnectorDefinition() *string
	ConnectorUrl() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultVersion() *string
	DefaultVersionAlias() *string
	DefaultVersionGitCommitHash() *string
	DefaultVersionLocationUri() *string
	DefaultVersionName() *string
	DefaultVersionSourceLocationUri() *string
	DisplayName() *string
	// Experimental.
	Fqn() *string
	InternalValue() *OpenflowConnectorDescribeOutput
	SetInternalValue(val *OpenflowConnectorDescribeOutput)
	LastVersionAlias() *string
	LastVersionGitCommitHash() *string
	LastVersionLocationUri() *string
	LastVersionName() *string
	LastVersionSourceLocationUri() *string
	LiveVersionLocationUri() *string
	Name() *string
	Owner() *string
	Runtime() *string
	Status() *string
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

// The jsii proxy struct for OpenflowConnectorDescribeOutputOutputReference
type jsiiProxy_OpenflowConnectorDescribeOutputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) ConnectorDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) ConnectorUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) DefaultVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) DefaultVersionAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) DefaultVersionGitCommitHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionGitCommitHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) DefaultVersionLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) DefaultVersionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) DefaultVersionSourceLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionSourceLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) InternalValue() *OpenflowConnectorDescribeOutput {
	var returns *OpenflowConnectorDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) LastVersionAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) LastVersionGitCommitHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionGitCommitHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) LastVersionLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) LastVersionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) LastVersionSourceLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionSourceLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) LiveVersionLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveVersionLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOpenflowConnectorDescribeOutputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OpenflowConnectorDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewOpenflowConnectorDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenflowConnectorDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.openflowConnector.OpenflowConnectorDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOpenflowConnectorDescribeOutputOutputReference_Override(o OpenflowConnectorDescribeOutputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.openflowConnector.OpenflowConnectorDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference)SetInternalValue(val *OpenflowConnectorDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenflowConnectorDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

