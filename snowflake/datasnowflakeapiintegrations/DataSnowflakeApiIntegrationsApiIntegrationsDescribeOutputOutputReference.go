// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeapiintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/datasnowflakeapiintegrations/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference interface {
	cdktn.ComplexObject
	AllowedAuthenticationSecrets() *string
	AllowedPrefixes() *[]*string
	ApiAwsExternalId() *string
	ApiAwsIamUserArn() *string
	ApiAwsRoleArn() *string
	ApiKey() *string
	ApiProvider() *string
	AzureAdApplicationId() *string
	AzureConsentUrl() *string
	AzureMultiTenantAppName() *string
	AzureTenantId() *string
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
	GoogleApiServiceAccount() *string
	GoogleAudience() *string
	InternalValue() *DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutput
	SetInternalValue(val *DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutput)
	OauthAccessTokenValidity() *float64
	OauthAllowedScopes() *[]*string
	OauthAssertionIssuer() *string
	OauthAuthorizationEndpoint() *string
	OauthClientAuthMethod() *string
	OauthClientId() *string
	OauthGrant() *string
	OauthRefreshTokenValidity() *float64
	OauthResourceUrl() *string
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
	TlsTrustedCertificates() *[]*string
	UsePrivatelinkEndpoint() cdktn.IResolvable
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

// The jsii proxy struct for DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference
type jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) AllowedAuthenticationSecrets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedAuthenticationSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) AllowedPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) ApiAwsExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiAwsExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) ApiAwsIamUserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiAwsIamUserArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) ApiAwsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiAwsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) ApiProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) AzureAdApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureAdApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) AzureConsentUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureConsentUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) AzureMultiTenantAppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureMultiTenantAppName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) AzureTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) BlockedPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"blockedPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) Enabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) GoogleApiServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleApiServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) GoogleAudience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAudience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) InternalValue() *DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutput {
	var returns *DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) OauthAccessTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthAccessTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) OauthAllowedScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthAllowedScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) OauthAssertionIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthAssertionIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) OauthAuthorizationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthAuthorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) OauthClientAuthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientAuthMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) OauthClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) OauthGrant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthGrant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) OauthRefreshTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) OauthResourceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthResourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) OauthTokenEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) OauthUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) TlsTrustedCertificates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tlsTrustedCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) UsePrivatelinkEndpoint() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"usePrivatelinkEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) UserAuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAuthType",
		&returns,
	)
	return returns
}


func NewDataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeApiIntegrations.DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference_Override(d DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeApiIntegrations.DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference)SetInternalValue(val *DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataSnowflakeApiIntegrationsApiIntegrationsDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

