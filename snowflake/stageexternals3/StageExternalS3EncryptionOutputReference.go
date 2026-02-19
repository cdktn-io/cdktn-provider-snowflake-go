// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/stageexternals3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageExternalS3EncryptionOutputReference interface {
	cdktn.ComplexObject
	AwsCse() StageExternalS3EncryptionAwsCseOutputReference
	AwsCseInput() *StageExternalS3EncryptionAwsCse
	AwsSseKms() StageExternalS3EncryptionAwsSseKmsOutputReference
	AwsSseKmsInput() *StageExternalS3EncryptionAwsSseKms
	AwsSseS3() StageExternalS3EncryptionAwsSseS3OutputReference
	AwsSseS3Input() *StageExternalS3EncryptionAwsSseS3
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
	InternalValue() *StageExternalS3Encryption
	SetInternalValue(val *StageExternalS3Encryption)
	None() StageExternalS3EncryptionNoneOutputReference
	NoneInput() *StageExternalS3EncryptionNone
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
	PutAwsCse(value *StageExternalS3EncryptionAwsCse)
	PutAwsSseKms(value *StageExternalS3EncryptionAwsSseKms)
	PutAwsSseS3(value *StageExternalS3EncryptionAwsSseS3)
	PutNone(value *StageExternalS3EncryptionNone)
	ResetAwsCse()
	ResetAwsSseKms()
	ResetAwsSseS3()
	ResetNone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StageExternalS3EncryptionOutputReference
type jsiiProxy_StageExternalS3EncryptionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) AwsCse() StageExternalS3EncryptionAwsCseOutputReference {
	var returns StageExternalS3EncryptionAwsCseOutputReference
	_jsii_.Get(
		j,
		"awsCse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) AwsCseInput() *StageExternalS3EncryptionAwsCse {
	var returns *StageExternalS3EncryptionAwsCse
	_jsii_.Get(
		j,
		"awsCseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) AwsSseKms() StageExternalS3EncryptionAwsSseKmsOutputReference {
	var returns StageExternalS3EncryptionAwsSseKmsOutputReference
	_jsii_.Get(
		j,
		"awsSseKms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) AwsSseKmsInput() *StageExternalS3EncryptionAwsSseKms {
	var returns *StageExternalS3EncryptionAwsSseKms
	_jsii_.Get(
		j,
		"awsSseKmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) AwsSseS3() StageExternalS3EncryptionAwsSseS3OutputReference {
	var returns StageExternalS3EncryptionAwsSseS3OutputReference
	_jsii_.Get(
		j,
		"awsSseS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) AwsSseS3Input() *StageExternalS3EncryptionAwsSseS3 {
	var returns *StageExternalS3EncryptionAwsSseS3
	_jsii_.Get(
		j,
		"awsSseS3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) InternalValue() *StageExternalS3Encryption {
	var returns *StageExternalS3Encryption
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) None() StageExternalS3EncryptionNoneOutputReference {
	var returns StageExternalS3EncryptionNoneOutputReference
	_jsii_.Get(
		j,
		"none",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) NoneInput() *StageExternalS3EncryptionNone {
	var returns *StageExternalS3EncryptionNone
	_jsii_.Get(
		j,
		"noneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStageExternalS3EncryptionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StageExternalS3EncryptionOutputReference {
	_init_.Initialize()

	if err := validateNewStageExternalS3EncryptionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StageExternalS3EncryptionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3EncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStageExternalS3EncryptionOutputReference_Override(s StageExternalS3EncryptionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3EncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference)SetInternalValue(val *StageExternalS3Encryption) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3EncryptionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) PutAwsCse(value *StageExternalS3EncryptionAwsCse) {
	if err := s.validatePutAwsCseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAwsCse",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) PutAwsSseKms(value *StageExternalS3EncryptionAwsSseKms) {
	if err := s.validatePutAwsSseKmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAwsSseKms",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) PutAwsSseS3(value *StageExternalS3EncryptionAwsSseS3) {
	if err := s.validatePutAwsSseS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAwsSseS3",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) PutNone(value *StageExternalS3EncryptionNone) {
	if err := s.validatePutNoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNone",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) ResetAwsCse() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsCse",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) ResetAwsSseKms() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsSseKms",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) ResetAwsSseS3() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsSseS3",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) ResetNone() {
	_jsii_.InvokeVoid(
		s,
		"resetNone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_StageExternalS3EncryptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

