// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oauthintegrationforcustomclients

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/oauthintegrationforcustomclients/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktn.IInterpolatingParent
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList
type jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewOauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList {
	_init_.Initialize()

	if err := validateNewOauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.oauthIntegrationForCustomClients.OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewOauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList_Override(o OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.oauthIntegrationForCustomClients.OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := o.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		o,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList) Get(index *float64) OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputOauthAuthorizationEndpointList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

