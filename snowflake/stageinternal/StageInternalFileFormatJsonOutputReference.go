// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageinternal

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/stageinternal/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageInternalFileFormatJsonOutputReference interface {
	cdktn.ComplexObject
	AllowDuplicate() *string
	SetAllowDuplicate(val *string)
	AllowDuplicateInput() *string
	BinaryFormat() *string
	SetBinaryFormat(val *string)
	BinaryFormatInput() *string
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
	DateFormat() *string
	SetDateFormat(val *string)
	DateFormatInput() *string
	EnableOctal() *string
	SetEnableOctal(val *string)
	EnableOctalInput() *string
	FileExtension() *string
	SetFileExtension(val *string)
	FileExtensionInput() *string
	// Experimental.
	Fqn() *string
	IgnoreUtf8Errors() *string
	SetIgnoreUtf8Errors(val *string)
	IgnoreUtf8ErrorsInput() *string
	InternalValue() *StageInternalFileFormatJson
	SetInternalValue(val *StageInternalFileFormatJson)
	MultiLine() *string
	SetMultiLine(val *string)
	MultiLineInput() *string
	NullIf() *[]*string
	SetNullIf(val *[]*string)
	NullIfInput() *[]*string
	ReplaceInvalidCharacters() *string
	SetReplaceInvalidCharacters(val *string)
	ReplaceInvalidCharactersInput() *string
	SkipByteOrderMark() *string
	SetSkipByteOrderMark(val *string)
	SkipByteOrderMarkInput() *string
	StripNullValues() *string
	SetStripNullValues(val *string)
	StripNullValuesInput() *string
	StripOuterArray() *string
	SetStripOuterArray(val *string)
	StripOuterArrayInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeFormat() *string
	SetTimeFormat(val *string)
	TimeFormatInput() *string
	TimestampFormat() *string
	SetTimestampFormat(val *string)
	TimestampFormatInput() *string
	TrimSpace() *string
	SetTrimSpace(val *string)
	TrimSpaceInput() *string
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
	ResetAllowDuplicate()
	ResetBinaryFormat()
	ResetCompression()
	ResetDateFormat()
	ResetEnableOctal()
	ResetFileExtension()
	ResetIgnoreUtf8Errors()
	ResetMultiLine()
	ResetNullIf()
	ResetReplaceInvalidCharacters()
	ResetSkipByteOrderMark()
	ResetStripNullValues()
	ResetStripOuterArray()
	ResetTimeFormat()
	ResetTimestampFormat()
	ResetTrimSpace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StageInternalFileFormatJsonOutputReference
type jsiiProxy_StageInternalFileFormatJsonOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) AllowDuplicate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowDuplicate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) AllowDuplicateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowDuplicateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) BinaryFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) BinaryFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) DateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) DateFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) EnableOctal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enableOctal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) EnableOctalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enableOctalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) FileExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) FileExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileExtensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) IgnoreUtf8Errors() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreUtf8Errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) IgnoreUtf8ErrorsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreUtf8ErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) InternalValue() *StageInternalFileFormatJson {
	var returns *StageInternalFileFormatJson
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) MultiLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) MultiLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) NullIf() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nullIf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) NullIfInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nullIfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) ReplaceInvalidCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) ReplaceInvalidCharactersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) SkipByteOrderMark() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipByteOrderMark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) SkipByteOrderMarkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipByteOrderMarkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) StripNullValues() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stripNullValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) StripNullValuesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stripNullValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) StripOuterArray() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stripOuterArray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) StripOuterArrayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stripOuterArrayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) TimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) TimeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) TimestampFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) TimestampFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) TrimSpace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trimSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference) TrimSpaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trimSpaceInput",
		&returns,
	)
	return returns
}


func NewStageInternalFileFormatJsonOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StageInternalFileFormatJsonOutputReference {
	_init_.Initialize()

	if err := validateNewStageInternalFileFormatJsonOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StageInternalFileFormatJsonOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageInternal.StageInternalFileFormatJsonOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStageInternalFileFormatJsonOutputReference_Override(s StageInternalFileFormatJsonOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageInternal.StageInternalFileFormatJsonOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetAllowDuplicate(val *string) {
	if err := j.validateSetAllowDuplicateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowDuplicate",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetBinaryFormat(val *string) {
	if err := j.validateSetBinaryFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryFormat",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetDateFormat(val *string) {
	if err := j.validateSetDateFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateFormat",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetEnableOctal(val *string) {
	if err := j.validateSetEnableOctalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableOctal",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetFileExtension(val *string) {
	if err := j.validateSetFileExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileExtension",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetIgnoreUtf8Errors(val *string) {
	if err := j.validateSetIgnoreUtf8ErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreUtf8Errors",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetInternalValue(val *StageInternalFileFormatJson) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetMultiLine(val *string) {
	if err := j.validateSetMultiLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiLine",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetNullIf(val *[]*string) {
	if err := j.validateSetNullIfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullIf",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetReplaceInvalidCharacters(val *string) {
	if err := j.validateSetReplaceInvalidCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceInvalidCharacters",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetSkipByteOrderMark(val *string) {
	if err := j.validateSetSkipByteOrderMarkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipByteOrderMark",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetStripNullValues(val *string) {
	if err := j.validateSetStripNullValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stripNullValues",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetStripOuterArray(val *string) {
	if err := j.validateSetStripOuterArrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stripOuterArray",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetTimeFormat(val *string) {
	if err := j.validateSetTimeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeFormat",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetTimestampFormat(val *string) {
	if err := j.validateSetTimestampFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampFormat",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatJsonOutputReference)SetTrimSpace(val *string) {
	if err := j.validateSetTrimSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trimSpace",
		val,
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetAllowDuplicate() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowDuplicate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetBinaryFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetBinaryFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		s,
		"resetCompression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetDateFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetDateFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetEnableOctal() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableOctal",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetFileExtension() {
	_jsii_.InvokeVoid(
		s,
		"resetFileExtension",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetIgnoreUtf8Errors() {
	_jsii_.InvokeVoid(
		s,
		"resetIgnoreUtf8Errors",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetMultiLine() {
	_jsii_.InvokeVoid(
		s,
		"resetMultiLine",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetNullIf() {
	_jsii_.InvokeVoid(
		s,
		"resetNullIf",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetReplaceInvalidCharacters() {
	_jsii_.InvokeVoid(
		s,
		"resetReplaceInvalidCharacters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetSkipByteOrderMark() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipByteOrderMark",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetStripNullValues() {
	_jsii_.InvokeVoid(
		s,
		"resetStripNullValues",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetStripOuterArray() {
	_jsii_.InvokeVoid(
		s,
		"resetStripOuterArray",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetTimeFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetTimestampFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetTimestampFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ResetTrimSpace() {
	_jsii_.InvokeVoid(
		s,
		"resetTrimSpace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_StageInternalFileFormatJsonOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

