// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalgcs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/stageexternalgcs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageExternalGcsFileFormatXmlOutputReference interface {
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
	InternalValue() *StageExternalGcsFileFormatXml
	SetInternalValue(val *StageExternalGcsFileFormatXml)
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

// The jsii proxy struct for StageExternalGcsFileFormatXmlOutputReference
type jsiiProxy_StageExternalGcsFileFormatXmlOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) DisableAutoConvert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disableAutoConvert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) DisableAutoConvertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"disableAutoConvertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) IgnoreUtf8Errors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreUtf8Errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) IgnoreUtf8ErrorsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreUtf8ErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) InternalValue() *StageExternalGcsFileFormatXml {
	var returns *StageExternalGcsFileFormatXml
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) PreserveSpace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) PreserveSpaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) ReplaceInvalidCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) ReplaceInvalidCharactersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) SkipByteOrderMark() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipByteOrderMark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) SkipByteOrderMarkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipByteOrderMarkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) StripOuterElement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stripOuterElement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) StripOuterElementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stripOuterElementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStageExternalGcsFileFormatXmlOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StageExternalGcsFileFormatXmlOutputReference {
	_init_.Initialize()

	if err := validateNewStageExternalGcsFileFormatXmlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StageExternalGcsFileFormatXmlOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalGcs.StageExternalGcsFileFormatXmlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStageExternalGcsFileFormatXmlOutputReference_Override(s StageExternalGcsFileFormatXmlOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalGcs.StageExternalGcsFileFormatXmlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference)SetDisableAutoConvert(val *string) {
	if err := j.validateSetDisableAutoConvertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutoConvert",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference)SetIgnoreUtf8Errors(val *string) {
	if err := j.validateSetIgnoreUtf8ErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreUtf8Errors",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference)SetInternalValue(val *StageExternalGcsFileFormatXml) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference)SetPreserveSpace(val *string) {
	if err := j.validateSetPreserveSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveSpace",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference)SetReplaceInvalidCharacters(val *string) {
	if err := j.validateSetReplaceInvalidCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceInvalidCharacters",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference)SetSkipByteOrderMark(val *string) {
	if err := j.validateSetSkipByteOrderMarkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipByteOrderMark",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference)SetStripOuterElement(val *string) {
	if err := j.validateSetStripOuterElementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stripOuterElement",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		s,
		"resetCompression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) ResetDisableAutoConvert() {
	_jsii_.InvokeVoid(
		s,
		"resetDisableAutoConvert",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) ResetIgnoreUtf8Errors() {
	_jsii_.InvokeVoid(
		s,
		"resetIgnoreUtf8Errors",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) ResetPreserveSpace() {
	_jsii_.InvokeVoid(
		s,
		"resetPreserveSpace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) ResetReplaceInvalidCharacters() {
	_jsii_.InvokeVoid(
		s,
		"resetReplaceInvalidCharacters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) ResetSkipByteOrderMark() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipByteOrderMark",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) ResetStripOuterElement() {
	_jsii_.InvokeVoid(
		s,
		"resetStripOuterElement",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_StageExternalGcsFileFormatXmlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

