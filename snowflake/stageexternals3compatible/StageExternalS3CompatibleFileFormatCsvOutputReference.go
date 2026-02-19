// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3compatible

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/stageexternals3compatible/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageExternalS3CompatibleFileFormatCsvOutputReference interface {
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
	InternalValue() *StageExternalS3CompatibleFileFormatCsv
	SetInternalValue(val *StageExternalS3CompatibleFileFormatCsv)
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

// The jsii proxy struct for StageExternalS3CompatibleFileFormatCsvOutputReference
type jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) BinaryFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) BinaryFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) DateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) DateFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) EmptyFieldAsNull() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyFieldAsNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) EmptyFieldAsNullInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyFieldAsNullInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ErrorOnColumnCountMismatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOnColumnCountMismatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ErrorOnColumnCountMismatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOnColumnCountMismatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) Escape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) EscapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) EscapeUnenclosedField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeUnenclosedField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) EscapeUnenclosedFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeUnenclosedFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) FieldDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) FieldDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) FieldOptionallyEnclosedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldOptionallyEnclosedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) FieldOptionallyEnclosedByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldOptionallyEnclosedByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) FileExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) FileExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileExtensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) InternalValue() *StageExternalS3CompatibleFileFormatCsv {
	var returns *StageExternalS3CompatibleFileFormatCsv
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) MultiLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) MultiLineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) NullIf() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nullIf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) NullIfInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nullIfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ParseHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parseHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ParseHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parseHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) RecordDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) RecordDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ReplaceInvalidCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ReplaceInvalidCharactersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) SkipBlankLines() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipBlankLines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) SkipBlankLinesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipBlankLinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) SkipByteOrderMark() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipByteOrderMark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) SkipByteOrderMarkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipByteOrderMarkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) SkipHeader() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skipHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) SkipHeaderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skipHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) TimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) TimeFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) TimestampFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) TimestampFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) TrimSpace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trimSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) TrimSpaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trimSpaceInput",
		&returns,
	)
	return returns
}


func NewStageExternalS3CompatibleFileFormatCsvOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StageExternalS3CompatibleFileFormatCsvOutputReference {
	_init_.Initialize()

	if err := validateNewStageExternalS3CompatibleFileFormatCsvOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalS3Compatible.StageExternalS3CompatibleFileFormatCsvOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStageExternalS3CompatibleFileFormatCsvOutputReference_Override(s StageExternalS3CompatibleFileFormatCsvOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalS3Compatible.StageExternalS3CompatibleFileFormatCsvOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetBinaryFormat(val *string) {
	if err := j.validateSetBinaryFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryFormat",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetDateFormat(val *string) {
	if err := j.validateSetDateFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dateFormat",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetEmptyFieldAsNull(val *string) {
	if err := j.validateSetEmptyFieldAsNullParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emptyFieldAsNull",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetErrorOnColumnCountMismatch(val *string) {
	if err := j.validateSetErrorOnColumnCountMismatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOnColumnCountMismatch",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetEscape(val *string) {
	if err := j.validateSetEscapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"escape",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetEscapeUnenclosedField(val *string) {
	if err := j.validateSetEscapeUnenclosedFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"escapeUnenclosedField",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetFieldDelimiter(val *string) {
	if err := j.validateSetFieldDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldDelimiter",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetFieldOptionallyEnclosedBy(val *string) {
	if err := j.validateSetFieldOptionallyEnclosedByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldOptionallyEnclosedBy",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetFileExtension(val *string) {
	if err := j.validateSetFileExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileExtension",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetInternalValue(val *StageExternalS3CompatibleFileFormatCsv) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetMultiLine(val *string) {
	if err := j.validateSetMultiLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiLine",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetNullIf(val *[]*string) {
	if err := j.validateSetNullIfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullIf",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetParseHeader(val *string) {
	if err := j.validateSetParseHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parseHeader",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetRecordDelimiter(val *string) {
	if err := j.validateSetRecordDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordDelimiter",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetReplaceInvalidCharacters(val *string) {
	if err := j.validateSetReplaceInvalidCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceInvalidCharacters",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetSkipBlankLines(val *string) {
	if err := j.validateSetSkipBlankLinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipBlankLines",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetSkipByteOrderMark(val *string) {
	if err := j.validateSetSkipByteOrderMarkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipByteOrderMark",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetSkipHeader(val *float64) {
	if err := j.validateSetSkipHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipHeader",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetTimeFormat(val *string) {
	if err := j.validateSetTimeFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeFormat",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetTimestampFormat(val *string) {
	if err := j.validateSetTimestampFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampFormat",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference)SetTrimSpace(val *string) {
	if err := j.validateSetTrimSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trimSpace",
		val,
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetBinaryFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetBinaryFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		s,
		"resetCompression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetDateFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetDateFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetEmptyFieldAsNull() {
	_jsii_.InvokeVoid(
		s,
		"resetEmptyFieldAsNull",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		s,
		"resetEncoding",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetErrorOnColumnCountMismatch() {
	_jsii_.InvokeVoid(
		s,
		"resetErrorOnColumnCountMismatch",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetEscape() {
	_jsii_.InvokeVoid(
		s,
		"resetEscape",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetEscapeUnenclosedField() {
	_jsii_.InvokeVoid(
		s,
		"resetEscapeUnenclosedField",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetFieldDelimiter() {
	_jsii_.InvokeVoid(
		s,
		"resetFieldDelimiter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetFieldOptionallyEnclosedBy() {
	_jsii_.InvokeVoid(
		s,
		"resetFieldOptionallyEnclosedBy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetFileExtension() {
	_jsii_.InvokeVoid(
		s,
		"resetFileExtension",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetMultiLine() {
	_jsii_.InvokeVoid(
		s,
		"resetMultiLine",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetNullIf() {
	_jsii_.InvokeVoid(
		s,
		"resetNullIf",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetParseHeader() {
	_jsii_.InvokeVoid(
		s,
		"resetParseHeader",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetRecordDelimiter() {
	_jsii_.InvokeVoid(
		s,
		"resetRecordDelimiter",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetReplaceInvalidCharacters() {
	_jsii_.InvokeVoid(
		s,
		"resetReplaceInvalidCharacters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetSkipBlankLines() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipBlankLines",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetSkipByteOrderMark() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipByteOrderMark",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetSkipHeader() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipHeader",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetTimeFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetTimestampFormat() {
	_jsii_.InvokeVoid(
		s,
		"resetTimestampFormat",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ResetTrimSpace() {
	_jsii_.InvokeVoid(
		s,
		"resetTrimSpace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatCsvOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

