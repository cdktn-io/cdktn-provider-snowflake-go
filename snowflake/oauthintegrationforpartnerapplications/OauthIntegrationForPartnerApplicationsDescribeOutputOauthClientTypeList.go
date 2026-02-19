// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oauthintegrationforpartnerapplications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/oauthintegrationforpartnerapplications/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList interface {
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
	Get(index *float64) OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList
type jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewOauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList {
	_init_.Initialize()

	if err := validateNewOauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.oauthIntegrationForPartnerApplications.OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewOauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList_Override(o OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.oauthIntegrationForPartnerApplications.OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList) Get(index *float64) OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_OauthIntegrationForPartnerApplicationsDescribeOutputOauthClientTypeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

