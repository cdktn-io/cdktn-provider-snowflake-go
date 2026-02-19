// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalgcs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/stageexternalgcs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageExternalGcsFileFormatParquetOutputReference interface {
	cdktn.ComplexObject
	BinaryAsText() *string
	SetBinaryAsText(val *string)
	BinaryAsTextInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *StageExternalGcsFileFormatParquet
	SetInternalValue(val *StageExternalGcsFileFormatParquet)
	NullIf() *[]*string
	SetNullIf(val *[]*string)
	NullIfInput() *[]*string
	ReplaceInvalidCharacters() *string
	SetReplaceInvalidCharacters(val *string)
	ReplaceInvalidCharactersInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrimSpace() *string
	SetTrimSpace(val *string)
	TrimSpaceInput() *string
	UseLogicalType() *string
	SetUseLogicalType(val *string)
	UseLogicalTypeInput() *string
	UseVectorizedScanner() *string
	SetUseVectorizedScanner(val *string)
	UseVectorizedScannerInput() *string
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
	ResetBinaryAsText()
	ResetCompression()
	ResetNullIf()
	ResetReplaceInvalidCharacters()
	ResetTrimSpace()
	ResetUseLogicalType()
	ResetUseVectorizedScanner()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StageExternalGcsFileFormatParquetOutputReference
type jsiiProxy_StageExternalGcsFileFormatParquetOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) BinaryAsText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryAsText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) BinaryAsTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryAsTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) InternalValue() *StageExternalGcsFileFormatParquet {
	var returns *StageExternalGcsFileFormatParquet
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) NullIf() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nullIf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) NullIfInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nullIfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) ReplaceInvalidCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) ReplaceInvalidCharactersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceInvalidCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) TrimSpace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trimSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) TrimSpaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trimSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) UseLogicalType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useLogicalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) UseLogicalTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useLogicalTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) UseVectorizedScanner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useVectorizedScanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) UseVectorizedScannerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useVectorizedScannerInput",
		&returns,
	)
	return returns
}


func NewStageExternalGcsFileFormatParquetOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StageExternalGcsFileFormatParquetOutputReference {
	_init_.Initialize()

	if err := validateNewStageExternalGcsFileFormatParquetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StageExternalGcsFileFormatParquetOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalGcs.StageExternalGcsFileFormatParquetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStageExternalGcsFileFormatParquetOutputReference_Override(s StageExternalGcsFileFormatParquetOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalGcs.StageExternalGcsFileFormatParquetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference)SetBinaryAsText(val *string) {
	if err := j.validateSetBinaryAsTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryAsText",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference)SetInternalValue(val *StageExternalGcsFileFormatParquet) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference)SetNullIf(val *[]*string) {
	if err := j.validateSetNullIfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullIf",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference)SetReplaceInvalidCharacters(val *string) {
	if err := j.validateSetReplaceInvalidCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceInvalidCharacters",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference)SetTrimSpace(val *string) {
	if err := j.validateSetTrimSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trimSpace",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference)SetUseLogicalType(val *string) {
	if err := j.validateSetUseLogicalTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLogicalType",
		val,
	)
}

func (j *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference)SetUseVectorizedScanner(val *string) {
	if err := j.validateSetUseVectorizedScannerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useVectorizedScanner",
		val,
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) ResetBinaryAsText() {
	_jsii_.InvokeVoid(
		s,
		"resetBinaryAsText",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		s,
		"resetCompression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) ResetNullIf() {
	_jsii_.InvokeVoid(
		s,
		"resetNullIf",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) ResetReplaceInvalidCharacters() {
	_jsii_.InvokeVoid(
		s,
		"resetReplaceInvalidCharacters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) ResetTrimSpace() {
	_jsii_.InvokeVoid(
		s,
		"resetTrimSpace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) ResetUseLogicalType() {
	_jsii_.InvokeVoid(
		s,
		"resetUseLogicalType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) ResetUseVectorizedScanner() {
	_jsii_.InvokeVoid(
		s,
		"resetUseVectorizedScanner",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_StageExternalGcsFileFormatParquetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

