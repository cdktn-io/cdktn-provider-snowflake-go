// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package legacyserviceuser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/legacyserviceuser/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LegacyServiceUserDefaultWorkloadIdentityOutputReference interface {
	cdktn.ComplexObject
	Aws() LegacyServiceUserDefaultWorkloadIdentityAwsOutputReference
	AwsInput() *LegacyServiceUserDefaultWorkloadIdentityAws
	Azure() LegacyServiceUserDefaultWorkloadIdentityAzureOutputReference
	AzureInput() *LegacyServiceUserDefaultWorkloadIdentityAzure
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
	// Experimental.
	Fqn() *string
	Gcp() LegacyServiceUserDefaultWorkloadIdentityGcpOutputReference
	GcpInput() *LegacyServiceUserDefaultWorkloadIdentityGcp
	InternalValue() *LegacyServiceUserDefaultWorkloadIdentity
	SetInternalValue(val *LegacyServiceUserDefaultWorkloadIdentity)
	Oidc() LegacyServiceUserDefaultWorkloadIdentityOidcOutputReference
	OidcInput() *LegacyServiceUserDefaultWorkloadIdentityOidc
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
	PutAws(value *LegacyServiceUserDefaultWorkloadIdentityAws)
	PutAzure(value *LegacyServiceUserDefaultWorkloadIdentityAzure)
	PutGcp(value *LegacyServiceUserDefaultWorkloadIdentityGcp)
	PutOidc(value *LegacyServiceUserDefaultWorkloadIdentityOidc)
	ResetAws()
	ResetAzure()
	ResetGcp()
	ResetOidc()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LegacyServiceUserDefaultWorkloadIdentityOutputReference
type jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) Aws() LegacyServiceUserDefaultWorkloadIdentityAwsOutputReference {
	var returns LegacyServiceUserDefaultWorkloadIdentityAwsOutputReference
	_jsii_.Get(
		j,
		"aws",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) AwsInput() *LegacyServiceUserDefaultWorkloadIdentityAws {
	var returns *LegacyServiceUserDefaultWorkloadIdentityAws
	_jsii_.Get(
		j,
		"awsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) Azure() LegacyServiceUserDefaultWorkloadIdentityAzureOutputReference {
	var returns LegacyServiceUserDefaultWorkloadIdentityAzureOutputReference
	_jsii_.Get(
		j,
		"azure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) AzureInput() *LegacyServiceUserDefaultWorkloadIdentityAzure {
	var returns *LegacyServiceUserDefaultWorkloadIdentityAzure
	_jsii_.Get(
		j,
		"azureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) Gcp() LegacyServiceUserDefaultWorkloadIdentityGcpOutputReference {
	var returns LegacyServiceUserDefaultWorkloadIdentityGcpOutputReference
	_jsii_.Get(
		j,
		"gcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) GcpInput() *LegacyServiceUserDefaultWorkloadIdentityGcp {
	var returns *LegacyServiceUserDefaultWorkloadIdentityGcp
	_jsii_.Get(
		j,
		"gcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) InternalValue() *LegacyServiceUserDefaultWorkloadIdentity {
	var returns *LegacyServiceUserDefaultWorkloadIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) Oidc() LegacyServiceUserDefaultWorkloadIdentityOidcOutputReference {
	var returns LegacyServiceUserDefaultWorkloadIdentityOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) OidcInput() *LegacyServiceUserDefaultWorkloadIdentityOidc {
	var returns *LegacyServiceUserDefaultWorkloadIdentityOidc
	_jsii_.Get(
		j,
		"oidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLegacyServiceUserDefaultWorkloadIdentityOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LegacyServiceUserDefaultWorkloadIdentityOutputReference {
	_init_.Initialize()

	if err := validateNewLegacyServiceUserDefaultWorkloadIdentityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.legacyServiceUser.LegacyServiceUserDefaultWorkloadIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLegacyServiceUserDefaultWorkloadIdentityOutputReference_Override(l LegacyServiceUserDefaultWorkloadIdentityOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.legacyServiceUser.LegacyServiceUserDefaultWorkloadIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference)SetInternalValue(val *LegacyServiceUserDefaultWorkloadIdentity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) PutAws(value *LegacyServiceUserDefaultWorkloadIdentityAws) {
	if err := l.validatePutAwsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAws",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) PutAzure(value *LegacyServiceUserDefaultWorkloadIdentityAzure) {
	if err := l.validatePutAzureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAzure",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) PutGcp(value *LegacyServiceUserDefaultWorkloadIdentityGcp) {
	if err := l.validatePutGcpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putGcp",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) PutOidc(value *LegacyServiceUserDefaultWorkloadIdentityOidc) {
	if err := l.validatePutOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putOidc",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) ResetAws() {
	_jsii_.InvokeVoid(
		l,
		"resetAws",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) ResetAzure() {
	_jsii_.InvokeVoid(
		l,
		"resetAzure",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) ResetGcp() {
	_jsii_.InvokeVoid(
		l,
		"resetGcp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) ResetOidc() {
	_jsii_.InvokeVoid(
		l,
		"resetOidc",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LegacyServiceUserDefaultWorkloadIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

