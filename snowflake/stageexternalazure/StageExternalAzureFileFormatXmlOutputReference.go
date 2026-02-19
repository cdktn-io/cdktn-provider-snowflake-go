// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalazure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/stageexternalazure/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageExternalAzureFileFormatXmlOutputReference interface {
	cdktn.ComplexObject
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
	Compression() *string
	SetCompression(val *string)
	CompressionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisableAutoConvert() *string
	SetDisableAutoConvert(val *string)
	DisableAutoConvertInput() *string
	// Experimental.
	Fqn() *string
	IgnoreUtf8Errors() *string
	SetIgnoreUtf8Errors(val *string)
	IgnoreUtf8ErrorsInput() *string
	InternalValue() *StageExternalAzureFileFormatXml
	SetInternalValue(val *StageExternalAzureFileFormatXml)
	PreserveSpace() *string
	SetPreserveSpace(val *string)
	PreserveSpaceInput() *string
	ReplaceInvalidCharacters() *string
	SetReplaceInvalidCharacters(val *string)
	ReplaceInvalidCharactersInput() *string
	SkipByteOrderMark() *string
	SetSkipByteOrderMark(val *string)
	SkipByteOrderMarkInput() *string
	StripOuterElement() *string
	SetStripOuterElement(val *string)
	StripOuterElementInput() *string
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
	ResetCompression()
	ResetDisableAutoConvert()
	ResetIgnoreUtf8Errors()
	ResetPreserveSpace()
	ResetReplaceInvalidCharacters()
	ResetSkipByteOrderMark()
	ResetStripOuterElement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StageExternalAzureFileFormatXmlOutputReference
type jsiiProxy_StageExternalAzureFileFormatXmlOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) DisableAutoConvert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disableAutoConvert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) DisableAutoConvertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disableAutoConvertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) IgnoreUtf8Errors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreUtf8Errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) IgnoreUtf8ErrorsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreUtf8ErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) InternalValue() *StageExternalAzureFileFormatXml {
	var returns *StageExternalAzureFileFormatXml
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) PreserveSpace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) PreserveSpaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) ReplaceInvalidCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) ReplaceInvalidCharactersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) SkipByteOrderMark() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipByteOrderMark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) SkipByteOrderMarkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipByteOrderMarkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) StripOuterElement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stripOuterElement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) StripOuterElementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stripOuterElementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStageExternalAzureFileFormatXmlOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StageExternalAzureFileFormatXmlOutputReference {
	_init_.Initialize()

	if err := validateNewStageExternalAzureFileFormatXmlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StageExternalAzureFileFormatXmlOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalAzure.StageExternalAzureFileFormatXmlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStageExternalAzureFileFormatXmlOutputReference_Override(s StageExternalAzureFileFormatXmlOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalAzure.StageExternalAzureFileFormatXmlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference)SetDisableAutoConvert(val *string) {
	if err := j.validateSetDisableAutoConvertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutoConvert",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference)SetIgnoreUtf8Errors(val *string) {
	if err := j.validateSetIgnoreUtf8ErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreUtf8Errors",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference)SetInternalValue(val *StageExternalAzureFileFormatXml) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference)SetPreserveSpace(val *string) {
	if err := j.validateSetPreserveSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveSpace",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference)SetReplaceInvalidCharacters(val *string) {
	if err := j.validateSetReplaceInvalidCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceInvalidCharacters",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference)SetSkipByteOrderMark(val *string) {
	if err := j.validateSetSkipByteOrderMarkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipByteOrderMark",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference)SetStripOuterElement(val *string) {
	if err := j.validateSetStripOuterElementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stripOuterElement",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		s,
		"resetCompression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) ResetDisableAutoConvert() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableAutoConvert",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) ResetIgnoreUtf8Errors() {
	_jsii_.InvokeVoid(
		s,
		"resetIgnoreUtf8Errors",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) ResetPreserveSpace() {
	_jsii_.InvokeVoid(
		s,
		"resetPreserveSpace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) ResetReplaceInvalidCharacters() {
	_jsii_.InvokeVoid(
		s,
		"resetReplaceInvalidCharacters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) ResetSkipByteOrderMark() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipByteOrderMark",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) ResetStripOuterElement() {
	_jsii_.InvokeVoid(
		s,
		"resetStripOuterElement",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_StageExternalAzureFileFormatXmlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

