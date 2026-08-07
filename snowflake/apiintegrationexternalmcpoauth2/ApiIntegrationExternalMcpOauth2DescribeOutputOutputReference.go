// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apiintegrationexternalmcpoauth2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/apiintegrationexternalmcpoauth2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference interface {
	cdktn.ComplexObject
	AllowedPrefixes() *[]*string
	ApiProvider() *string
	BlockedPrefixes() *[]*string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *ApiIntegrationExternalMcpOauth2DescribeOutput
	SetInternalValue(val *ApiIntegrationExternalMcpOauth2DescribeOutput)
	OauthAccessTokenValidity() *float64
	OauthAllowedScopes() *[]*string
	OauthAssertionIssuer() *string
	OauthAuthorizationEndpoint() *string
	OauthClientAuthMethod() *string
	OauthClientId() *string
	OauthGrant() *string
	OauthRefreshTokenValidity() *float64
	OauthTokenEndpoint() *string
	OauthUsername() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserAuthType() *string
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

// The jsii proxy struct for ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference
type jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) AllowedPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) ApiProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) BlockedPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) Enabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) InternalValue() *ApiIntegrationExternalMcpOauth2DescribeOutput {
	var returns *ApiIntegrationExternalMcpOauth2DescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) OauthAccessTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthAccessTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) OauthAllowedScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthAllowedScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) OauthAssertionIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthAssertionIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) OauthAuthorizationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthAuthorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) OauthClientAuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientAuthMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) OauthClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) OauthGrant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthGrant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) OauthRefreshTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) OauthTokenEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) OauthUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) UserAuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAuthType",
		&returns,
	)
	return returns
}


func NewApiIntegrationExternalMcpOauth2DescribeOutputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewApiIntegrationExternalMcpOauth2DescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.apiIntegrationExternalMcpOauth2.ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApiIntegrationExternalMcpOauth2DescribeOutputOutputReference_Override(a ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.apiIntegrationExternalMcpOauth2.ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference)SetInternalValue(val *ApiIntegrationExternalMcpOauth2DescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApiIntegrationExternalMcpOauth2DescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

