// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package serviceuser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/serviceuser/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServiceUserDefaultWorkloadIdentityOutputReference interface {
	cdktn.ComplexObject
	Aws() ServiceUserDefaultWorkloadIdentityAwsOutputReference
	AwsInput() *ServiceUserDefaultWorkloadIdentityAws
	Azure() ServiceUserDefaultWorkloadIdentityAzureOutputReference
	AzureInput() *ServiceUserDefaultWorkloadIdentityAzure
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
	Gcp() ServiceUserDefaultWorkloadIdentityGcpOutputReference
	GcpInput() *ServiceUserDefaultWorkloadIdentityGcp
	InternalValue() *ServiceUserDefaultWorkloadIdentity
	SetInternalValue(val *ServiceUserDefaultWorkloadIdentity)
	Oidc() ServiceUserDefaultWorkloadIdentityOidcOutputReference
	OidcInput() *ServiceUserDefaultWorkloadIdentityOidc
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
	PutAws(value *ServiceUserDefaultWorkloadIdentityAws)
	PutAzure(value *ServiceUserDefaultWorkloadIdentityAzure)
	PutGcp(value *ServiceUserDefaultWorkloadIdentityGcp)
	PutOidc(value *ServiceUserDefaultWorkloadIdentityOidc)
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

// The jsii proxy struct for ServiceUserDefaultWorkloadIdentityOutputReference
type jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) Aws() ServiceUserDefaultWorkloadIdentityAwsOutputReference {
	var returns ServiceUserDefaultWorkloadIdentityAwsOutputReference
	_jsii_.Get(
		j,
		"aws",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) AwsInput() *ServiceUserDefaultWorkloadIdentityAws {
	var returns *ServiceUserDefaultWorkloadIdentityAws
	_jsii_.Get(
		j,
		"awsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) Azure() ServiceUserDefaultWorkloadIdentityAzureOutputReference {
	var returns ServiceUserDefaultWorkloadIdentityAzureOutputReference
	_jsii_.Get(
		j,
		"azure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) AzureInput() *ServiceUserDefaultWorkloadIdentityAzure {
	var returns *ServiceUserDefaultWorkloadIdentityAzure
	_jsii_.Get(
		j,
		"azureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) Gcp() ServiceUserDefaultWorkloadIdentityGcpOutputReference {
	var returns ServiceUserDefaultWorkloadIdentityGcpOutputReference
	_jsii_.Get(
		j,
		"gcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) GcpInput() *ServiceUserDefaultWorkloadIdentityGcp {
	var returns *ServiceUserDefaultWorkloadIdentityGcp
	_jsii_.Get(
		j,
		"gcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) InternalValue() *ServiceUserDefaultWorkloadIdentity {
	var returns *ServiceUserDefaultWorkloadIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) Oidc() ServiceUserDefaultWorkloadIdentityOidcOutputReference {
	var returns ServiceUserDefaultWorkloadIdentityOidcOutputReference
	_jsii_.Get(
		j,
		"oidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) OidcInput() *ServiceUserDefaultWorkloadIdentityOidc {
	var returns *ServiceUserDefaultWorkloadIdentityOidc
	_jsii_.Get(
		j,
		"oidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceUserDefaultWorkloadIdentityOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ServiceUserDefaultWorkloadIdentityOutputReference {
	_init_.Initialize()

	if err := validateNewServiceUserDefaultWorkloadIdentityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.serviceUser.ServiceUserDefaultWorkloadIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceUserDefaultWorkloadIdentityOutputReference_Override(s ServiceUserDefaultWorkloadIdentityOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.serviceUser.ServiceUserDefaultWorkloadIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference)SetInternalValue(val *ServiceUserDefaultWorkloadIdentity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) PutAws(value *ServiceUserDefaultWorkloadIdentityAws) {
	if err := s.validatePutAwsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAws",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) PutAzure(value *ServiceUserDefaultWorkloadIdentityAzure) {
	if err := s.validatePutAzureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzure",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) PutGcp(value *ServiceUserDefaultWorkloadIdentityGcp) {
	if err := s.validatePutGcpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGcp",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) PutOidc(value *ServiceUserDefaultWorkloadIdentityOidc) {
	if err := s.validatePutOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOidc",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) ResetAws() {
	_jsii_.InvokeVoid(
		s,
		"resetAws",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) ResetAzure() {
	_jsii_.InvokeVoid(
		s,
		"resetAzure",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) ResetGcp() {
	_jsii_.InvokeVoid(
		s,
		"resetGcp",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) ResetOidc() {
	_jsii_.InvokeVoid(
		s,
		"resetOidc",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceUserDefaultWorkloadIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

