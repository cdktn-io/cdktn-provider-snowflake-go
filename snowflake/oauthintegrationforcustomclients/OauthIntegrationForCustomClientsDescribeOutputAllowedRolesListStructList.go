// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oauthintegrationforcustomclients

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/oauthintegrationforcustomclients/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList
type jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewOauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList {
	_init_.Initialize()

	if err := validateNewOauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.oauthIntegrationForCustomClients.OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewOauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList_Override(o OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.oauthIntegrationForCustomClients.OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList) Get(index *float64) OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_OauthIntegrationForCustomClientsDescribeOutputAllowedRolesListStructList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

