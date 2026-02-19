// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apiauthenticationintegrationwithclientcredentials

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/apiauthenticationintegrationwithclientcredentials/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList interface {
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
	Get(index *float64) ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList
type jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList {
	_init_.Initialize()

	if err := validateNewApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.apiAuthenticationIntegrationWithClientCredentials.ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList_Override(a ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.apiAuthenticationIntegrationWithClientCredentials.ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList) Get(index *float64) ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiAuthenticationIntegrationWithClientCredentialsDescribeOutputOauthTokenEndpointList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

