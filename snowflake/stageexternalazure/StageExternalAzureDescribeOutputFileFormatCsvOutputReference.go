// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalazure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/stageexternalazure/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageExternalAzureDescribeOutputFileFormatCsvOutputReference interface {
	cdktn.ComplexObject
	BinaryFormat() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DateFormat() *string
	EmptyFieldAsNull() cdktn.IResolvable
	Encoding() *string
	ErrorOnColumnCountMismatch() cdktn.IResolvable
	Escape() *string
	EscapeUnenclosedField() *string
	FieldDelimiter() *string
	FieldOptionallyEnclosedBy() *string
	FileExtension() *string
	// Experimental.
	Fqn() *string
	InternalValue() *StageExternalAzureDescribeOutputFileFormatCsv
	SetInternalValue(val *StageExternalAzureDescribeOutputFileFormatCsv)
	MultiLine() cdktn.IResolvable
	NullIf() *[]*string
	ParseHeader() cdktn.IResolvable
	RecordDelimiter() *string
	ReplaceInvalidCharacters() cdktn.IResolvable
	SkipBlankLines() cdktn.IResolvable
	SkipByteOrderMark() cdktn.IResolvable
	SkipHeader() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeFormat() *string
	TimestampFormat() *string
	TrimSpace() cdktn.IResolvable
	Type() *string
	ValidateUtf8() cdktn.IResolvable
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

// The jsii proxy struct for StageExternalAzureDescribeOutputFileFormatCsvOutputReference
type jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) BinaryFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) DateFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) EmptyFieldAsNull() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"emptyFieldAsNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) ErrorOnColumnCountMismatch() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"errorOnColumnCountMismatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) Escape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) EscapeUnenclosedField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeUnenclosedField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) FieldDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) FieldOptionallyEnclosedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldOptionallyEnclosedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) FileExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) InternalValue() *StageExternalAzureDescribeOutputFileFormatCsv {
	var returns *StageExternalAzureDescribeOutputFileFormatCsv
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) MultiLine() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"multiLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) NullIf() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nullIf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) ParseHeader() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"parseHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) RecordDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) ReplaceInvalidCharacters() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) SkipBlankLines() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"skipBlankLines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) SkipByteOrderMark() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"skipByteOrderMark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) SkipHeader() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skipHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) TimeFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) TimestampFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) TrimSpace() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"trimSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) ValidateUtf8() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"validateUtf8",
		&returns,
	)
	return returns
}


func NewStageExternalAzureDescribeOutputFileFormatCsvOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StageExternalAzureDescribeOutputFileFormatCsvOutputReference {
	_init_.Initialize()

	if err := validateNewStageExternalAzureDescribeOutputFileFormatCsvOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalAzure.StageExternalAzureDescribeOutputFileFormatCsvOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStageExternalAzureDescribeOutputFileFormatCsvOutputReference_Override(s StageExternalAzureDescribeOutputFileFormatCsvOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalAzure.StageExternalAzureDescribeOutputFileFormatCsvOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference)SetInternalValue(val *StageExternalAzureDescribeOutputFileFormatCsv) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_StageExternalAzureDescribeOutputFileFormatCsvOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

