// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalazure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/stageexternalazure/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageExternalAzureFileFormatCsvOutputReference interface {
	cdktn.ComplexObject
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
	EmptyFieldAsNull() *string
	SetEmptyFieldAsNull(val *string)
	EmptyFieldAsNullInput() *string
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	ErrorOnColumnCountMismatch() *string
	SetErrorOnColumnCountMismatch(val *string)
	ErrorOnColumnCountMismatchInput() *string
	Escape() *string
	SetEscape(val *string)
	EscapeInput() *string
	EscapeUnenclosedField() *string
	SetEscapeUnenclosedField(val *string)
	EscapeUnenclosedFieldInput() *string
	FieldDelimiter() *string
	SetFieldDelimiter(val *string)
	FieldDelimiterInput() *string
	FieldOptionallyEnclosedBy() *string
	SetFieldOptionallyEnclosedBy(val *string)
	FieldOptionallyEnclosedByInput() *string
	FileExtension() *string
	SetFileExtension(val *string)
	FileExtensionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *StageExternalAzureFileFormatCsv
	SetInternalValue(val *StageExternalAzureFileFormatCsv)
	MultiLine() *string
	SetMultiLine(val *string)
	MultiLineInput() *string
	NullIf() *[]*string
	SetNullIf(val *[]*string)
	NullIfInput() *[]*string
	ParseHeader() *string
	SetParseHeader(val *string)
	ParseHeaderInput() *string
	RecordDelimiter() *string
	SetRecordDelimiter(val *string)
	RecordDelimiterInput() *string
	ReplaceInvalidCharacters() *string
	SetReplaceInvalidCharacters(val *string)
	ReplaceInvalidCharactersInput() *string
	SkipBlankLines() *string
	SetSkipBlankLines(val *string)
	SkipBlankLinesInput() *string
	SkipByteOrderMark() *string
	SetSkipByteOrderMark(val *string)
	SkipByteOrderMarkInput() *string
	SkipHeader() *float64
	SetSkipHeader(val *float64)
	SkipHeaderInput() *float64
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
	ResetBinaryFormat()
	ResetCompression()
	ResetDateFormat()
	ResetEmptyFieldAsNull()
	ResetEncoding()
	ResetErrorOnColumnCountMismatch()
	ResetEscape()
	ResetEscapeUnenclosedField()
	ResetFieldDelimiter()
	ResetFieldOptionallyEnclosedBy()
	ResetFileExtension()
	ResetMultiLine()
	ResetNullIf()
	ResetParseHeader()
	ResetRecordDelimiter()
	ResetReplaceInvalidCharacters()
	ResetSkipBlankLines()
	ResetSkipByteOrderMark()
	ResetSkipHeader()
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

// The jsii proxy struct for StageExternalAzureFileFormatCsvOutputReference
type jsiiProxy_StageExternalAzureFileFormatCsvOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) BinaryFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) BinaryFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) DateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) DateFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) EmptyFieldAsNull() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyFieldAsNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) EmptyFieldAsNullInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyFieldAsNullInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ErrorOnColumnCountMismatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOnColumnCountMismatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ErrorOnColumnCountMismatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOnColumnCountMismatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) Escape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) EscapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) EscapeUnenclosedField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeUnenclosedField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) EscapeUnenclosedFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeUnenclosedFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) FieldDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) FieldDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) FieldOptionallyEnclosedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldOptionallyEnclosedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) FieldOptionallyEnclosedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldOptionallyEnclosedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) FileExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) FileExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileExtensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) InternalValue() *StageExternalAzureFileFormatCsv {
	var returns *StageExternalAzureFileFormatCsv
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) MultiLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) MultiLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) NullIf() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nullIf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) NullIfInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nullIfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ParseHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parseHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ParseHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parseHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) RecordDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) RecordDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ReplaceInvalidCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ReplaceInvalidCharactersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) SkipBlankLines() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipBlankLines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) SkipBlankLinesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipBlankLinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) SkipByteOrderMark() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipByteOrderMark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) SkipByteOrderMarkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipByteOrderMarkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) SkipHeader() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skipHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) SkipHeaderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skipHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) TimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) TimeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) TimestampFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) TimestampFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) TrimSpace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trimSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) TrimSpaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trimSpaceInput",
		&returns,
	)
	return returns
}


func NewStageExternalAzureFileFormatCsvOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StageExternalAzureFileFormatCsvOutputReference {
	_init_.Initialize()

	if err := validateNewStageExternalAzureFileFormatCsvOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StageExternalAzureFileFormatCsvOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalAzure.StageExternalAzureFileFormatCsvOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStageExternalAzureFileFormatCsvOutputReference_Override(s StageExternalAzureFileFormatCsvOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalAzure.StageExternalAzureFileFormatCsvOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetBinaryFormat(val *string) {
	if err := j.validateSetBinaryFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryFormat",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetDateFormat(val *string) {
	if err := j.validateSetDateFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateFormat",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetEmptyFieldAsNull(val *string) {
	if err := j.validateSetEmptyFieldAsNullParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emptyFieldAsNull",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetErrorOnColumnCountMismatch(val *string) {
	if err := j.validateSetErrorOnColumnCountMismatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnColumnCountMismatch",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetEscape(val *string) {
	if err := j.validateSetEscapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"escape",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetEscapeUnenclosedField(val *string) {
	if err := j.validateSetEscapeUnenclosedFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"escapeUnenclosedField",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetFieldDelimiter(val *string) {
	if err := j.validateSetFieldDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldDelimiter",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetFieldOptionallyEnclosedBy(val *string) {
	if err := j.validateSetFieldOptionallyEnclosedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldOptionallyEnclosedBy",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetFileExtension(val *string) {
	if err := j.validateSetFileExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileExtension",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetInternalValue(val *StageExternalAzureFileFormatCsv) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetMultiLine(val *string) {
	if err := j.validateSetMultiLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiLine",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetNullIf(val *[]*string) {
	if err := j.validateSetNullIfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullIf",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetParseHeader(val *string) {
	if err := j.validateSetParseHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parseHeader",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetRecordDelimiter(val *string) {
	if err := j.validateSetRecordDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordDelimiter",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetReplaceInvalidCharacters(val *string) {
	if err := j.validateSetReplaceInvalidCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceInvalidCharacters",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetSkipBlankLines(val *string) {
	if err := j.validateSetSkipBlankLinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipBlankLines",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetSkipByteOrderMark(val *string) {
	if err := j.validateSetSkipByteOrderMarkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipByteOrderMark",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetSkipHeader(val *float64) {
	if err := j.validateSetSkipHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipHeader",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetTimeFormat(val *string) {
	if err := j.validateSetTimeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeFormat",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetTimestampFormat(val *string) {
	if err := j.validateSetTimestampFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampFormat",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference)SetTrimSpace(val *string) {
	if err := j.validateSetTrimSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trimSpace",
		val,
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetBinaryFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetBinaryFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		s,
		"resetCompression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetDateFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetDateFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetEmptyFieldAsNull() {
	_jsii_.InvokeVoid(
		s,
		"resetEmptyFieldAsNull",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		s,
		"resetEncoding",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetErrorOnColumnCountMismatch() {
	_jsii_.InvokeVoid(
		s,
		"resetErrorOnColumnCountMismatch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetEscape() {
	_jsii_.InvokeVoid(
		s,
		"resetEscape",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetEscapeUnenclosedField() {
	_jsii_.InvokeVoid(
		s,
		"resetEscapeUnenclosedField",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetFieldDelimiter() {
	_jsii_.InvokeVoid(
		s,
		"resetFieldDelimiter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetFieldOptionallyEnclosedBy() {
	_jsii_.InvokeVoid(
		s,
		"resetFieldOptionallyEnclosedBy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetFileExtension() {
	_jsii_.InvokeVoid(
		s,
		"resetFileExtension",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetMultiLine() {
	_jsii_.InvokeVoid(
		s,
		"resetMultiLine",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetNullIf() {
	_jsii_.InvokeVoid(
		s,
		"resetNullIf",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetParseHeader() {
	_jsii_.InvokeVoid(
		s,
		"resetParseHeader",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetRecordDelimiter() {
	_jsii_.InvokeVoid(
		s,
		"resetRecordDelimiter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetReplaceInvalidCharacters() {
	_jsii_.InvokeVoid(
		s,
		"resetReplaceInvalidCharacters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetSkipBlankLines() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipBlankLines",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetSkipByteOrderMark() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipByteOrderMark",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetSkipHeader() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipHeader",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetTimeFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetTimestampFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetTimestampFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ResetTrimSpace() {
	_jsii_.InvokeVoid(
		s,
		"resetTrimSpace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_StageExternalAzureFileFormatCsvOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

