// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageinternal

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/stageinternal/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageInternalFileFormatOutputReference interface {
	cdktn.ComplexObject
	Avro() StageInternalFileFormatAvroOutputReference
	AvroInput() *StageInternalFileFormatAvro
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
	Csv() StageInternalFileFormatCsvOutputReference
	CsvInput() *StageInternalFileFormatCsv
	FormatName() *string
	SetFormatName(val *string)
	FormatNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *StageInternalFileFormat
	SetInternalValue(val *StageInternalFileFormat)
	Json() StageInternalFileFormatJsonOutputReference
	JsonInput() *StageInternalFileFormatJson
	Orc() StageInternalFileFormatOrcOutputReference
	OrcInput() *StageInternalFileFormatOrc
	Parquet() StageInternalFileFormatParquetOutputReference
	ParquetInput() *StageInternalFileFormatParquet
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Xml() StageInternalFileFormatXmlOutputReference
	XmlInput() *StageInternalFileFormatXml
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
	PutAvro(value *StageInternalFileFormatAvro)
	PutCsv(value *StageInternalFileFormatCsv)
	PutJson(value *StageInternalFileFormatJson)
	PutOrc(value *StageInternalFileFormatOrc)
	PutParquet(value *StageInternalFileFormatParquet)
	PutXml(value *StageInternalFileFormatXml)
	ResetAvro()
	ResetCsv()
	ResetFormatName()
	ResetJson()
	ResetOrc()
	ResetParquet()
	ResetXml()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StageInternalFileFormatOutputReference
type jsiiProxy_StageInternalFileFormatOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) Avro() StageInternalFileFormatAvroOutputReference {
	var returns StageInternalFileFormatAvroOutputReference
	_jsii_.Get(
		j,
		"avro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) AvroInput() *StageInternalFileFormatAvro {
	var returns *StageInternalFileFormatAvro
	_jsii_.Get(
		j,
		"avroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) Csv() StageInternalFileFormatCsvOutputReference {
	var returns StageInternalFileFormatCsvOutputReference
	_jsii_.Get(
		j,
		"csv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) CsvInput() *StageInternalFileFormatCsv {
	var returns *StageInternalFileFormatCsv
	_jsii_.Get(
		j,
		"csvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) FormatName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) FormatNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) InternalValue() *StageInternalFileFormat {
	var returns *StageInternalFileFormat
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) Json() StageInternalFileFormatJsonOutputReference {
	var returns StageInternalFileFormatJsonOutputReference
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) JsonInput() *StageInternalFileFormatJson {
	var returns *StageInternalFileFormatJson
	_jsii_.Get(
		j,
		"jsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) Orc() StageInternalFileFormatOrcOutputReference {
	var returns StageInternalFileFormatOrcOutputReference
	_jsii_.Get(
		j,
		"orc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) OrcInput() *StageInternalFileFormatOrc {
	var returns *StageInternalFileFormatOrc
	_jsii_.Get(
		j,
		"orcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) Parquet() StageInternalFileFormatParquetOutputReference {
	var returns StageInternalFileFormatParquetOutputReference
	_jsii_.Get(
		j,
		"parquet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) ParquetInput() *StageInternalFileFormatParquet {
	var returns *StageInternalFileFormatParquet
	_jsii_.Get(
		j,
		"parquetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) Xml() StageInternalFileFormatXmlOutputReference {
	var returns StageInternalFileFormatXmlOutputReference
	_jsii_.Get(
		j,
		"xml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference) XmlInput() *StageInternalFileFormatXml {
	var returns *StageInternalFileFormatXml
	_jsii_.Get(
		j,
		"xmlInput",
		&returns,
	)
	return returns
}


func NewStageInternalFileFormatOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StageInternalFileFormatOutputReference {
	_init_.Initialize()

	if err := validateNewStageInternalFileFormatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StageInternalFileFormatOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageInternal.StageInternalFileFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStageInternalFileFormatOutputReference_Override(s StageInternalFileFormatOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageInternal.StageInternalFileFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference)SetFormatName(val *string) {
	if err := j.validateSetFormatNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formatName",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference)SetInternalValue(val *StageInternalFileFormat) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StageInternalFileFormatOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StageInternalFileFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageInternalFileFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StageInternalFileFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StageInternalFileFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StageInternalFileFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StageInternalFileFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StageInternalFileFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StageInternalFileFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StageInternalFileFormatOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageInternalFileFormatOutputReference) PutAvro(value *StageInternalFileFormatAvro) {
	if err := s.validatePutAvroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAvro",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) PutCsv(value *StageInternalFileFormatCsv) {
	if err := s.validatePutCsvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCsv",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) PutJson(value *StageInternalFileFormatJson) {
	if err := s.validatePutJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putJson",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) PutOrc(value *StageInternalFileFormatOrc) {
	if err := s.validatePutOrcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOrc",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) PutParquet(value *StageInternalFileFormatParquet) {
	if err := s.validatePutParquetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putParquet",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) PutXml(value *StageInternalFileFormatXml) {
	if err := s.validatePutXmlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putXml",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) ResetAvro() {
	_jsii_.InvokeVoid(
		s,
		"resetAvro",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) ResetCsv() {
	_jsii_.InvokeVoid(
		s,
		"resetCsv",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) ResetFormatName() {
	_jsii_.InvokeVoid(
		s,
		"resetFormatName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) ResetJson() {
	_jsii_.InvokeVoid(
		s,
		"resetJson",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) ResetOrc() {
	_jsii_.InvokeVoid(
		s,
		"resetOrc",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) ResetParquet() {
	_jsii_.InvokeVoid(
		s,
		"resetParquet",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) ResetXml() {
	_jsii_.InvokeVoid(
		s,
		"resetXml",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageInternalFileFormatOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_StageInternalFileFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

