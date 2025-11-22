// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notebook

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/notebook/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotebookDescribeOutputOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *NotebookDescribeOutput
	SetInternalValue(val *NotebookDescribeOutput)
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
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() *string
	UrlId() *string
	UserPackages() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NotebookDescribeOutputOutputReference
type jsiiProxy_NotebookDescribeOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) CodeWarehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) ComputePool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) DefaultPackages() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) DefaultVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) DefaultVersionAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) DefaultVersionGitCommitHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionGitCommitHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) DefaultVersionLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) DefaultVersionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) DefaultVersionSourceLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionSourceLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) ExternalAccessIntegrations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAccessIntegrations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) ExternalAccessSecrets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAccessSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) IdleAutoShutdownTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleAutoShutdownTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) ImportUrls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) InternalValue() *NotebookDescribeOutput {
	var returns *NotebookDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) LastVersionAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) LastVersionGitCommitHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionGitCommitHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) LastVersionLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) LastVersionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) LastVersionSourceLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastVersionSourceLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) LiveVersionLocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveVersionLocationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) MainFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) QueryWarehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) RuntimeEnvironmentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEnvironmentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) RuntimeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) UrlId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference) UserPackages() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPackages",
		&returns,
	)
	return returns
}


func NewNotebookDescribeOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NotebookDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewNotebookDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotebookDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.notebook.NotebookDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNotebookDescribeOutputOutputReference_Override(n NotebookDescribeOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.notebook.NotebookDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference)SetInternalValue(val *NotebookDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NotebookDescribeOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebookDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

