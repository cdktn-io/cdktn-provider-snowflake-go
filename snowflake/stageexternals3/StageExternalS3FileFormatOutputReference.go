// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/stageexternals3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageExternalS3FileFormatOutputReference interface {
	cdktn.ComplexObject
	Avro() StageExternalS3FileFormatAvroOutputReference
	AvroInput() *StageExternalS3FileFormatAvro
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
	Csv() StageExternalS3FileFormatCsvOutputReference
	CsvInput() *StageExternalS3FileFormatCsv
	FormatName() *string
	SetFormatName(val *string)
	FormatNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *StageExternalS3FileFormat
	SetInternalValue(val *StageExternalS3FileFormat)
	Json() StageExternalS3FileFormatJsonOutputReference
	JsonInput() *StageExternalS3FileFormatJson
	Orc() StageExternalS3FileFormatOrcOutputReference
	OrcInput() *StageExternalS3FileFormatOrc
	Parquet() StageExternalS3FileFormatParquetOutputReference
	ParquetInput() *StageExternalS3FileFormatParquet
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Xml() StageExternalS3FileFormatXmlOutputReference
	XmlInput() *StageExternalS3FileFormatXml
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
	PutAvro(value *StageExternalS3FileFormatAvro)
	PutCsv(value *StageExternalS3FileFormatCsv)
	PutJson(value *StageExternalS3FileFormatJson)
	PutOrc(value *StageExternalS3FileFormatOrc)
	PutParquet(value *StageExternalS3FileFormatParquet)
	PutXml(value *StageExternalS3FileFormatXml)
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

// The jsii proxy struct for StageExternalS3FileFormatOutputReference
type jsiiProxy_StageExternalS3FileFormatOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) Avro() StageExternalS3FileFormatAvroOutputReference {
	var returns StageExternalS3FileFormatAvroOutputReference
	_jsii_.Get(
		j,
		"avro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) AvroInput() *StageExternalS3FileFormatAvro {
	var returns *StageExternalS3FileFormatAvro
	_jsii_.Get(
		j,
		"avroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) Csv() StageExternalS3FileFormatCsvOutputReference {
	var returns StageExternalS3FileFormatCsvOutputReference
	_jsii_.Get(
		j,
		"csv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) CsvInput() *StageExternalS3FileFormatCsv {
	var returns *StageExternalS3FileFormatCsv
	_jsii_.Get(
		j,
		"csvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) FormatName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) FormatNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) InternalValue() *StageExternalS3FileFormat {
	var returns *StageExternalS3FileFormat
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) Json() StageExternalS3FileFormatJsonOutputReference {
	var returns StageExternalS3FileFormatJsonOutputReference
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) JsonInput() *StageExternalS3FileFormatJson {
	var returns *StageExternalS3FileFormatJson
	_jsii_.Get(
		j,
		"jsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) Orc() StageExternalS3FileFormatOrcOutputReference {
	var returns StageExternalS3FileFormatOrcOutputReference
	_jsii_.Get(
		j,
		"orc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) OrcInput() *StageExternalS3FileFormatOrc {
	var returns *StageExternalS3FileFormatOrc
	_jsii_.Get(
		j,
		"orcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) Parquet() StageExternalS3FileFormatParquetOutputReference {
	var returns StageExternalS3FileFormatParquetOutputReference
	_jsii_.Get(
		j,
		"parquet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) ParquetInput() *StageExternalS3FileFormatParquet {
	var returns *StageExternalS3FileFormatParquet
	_jsii_.Get(
		j,
		"parquetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) Xml() StageExternalS3FileFormatXmlOutputReference {
	var returns StageExternalS3FileFormatXmlOutputReference
	_jsii_.Get(
		j,
		"xml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference) XmlInput() *StageExternalS3FileFormatXml {
	var returns *StageExternalS3FileFormatXml
	_jsii_.Get(
		j,
		"xmlInput",
		&returns,
	)
	return returns
}


func NewStageExternalS3FileFormatOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StageExternalS3FileFormatOutputReference {
	_init_.Initialize()

	if err := validateNewStageExternalS3FileFormatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StageExternalS3FileFormatOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStageExternalS3FileFormatOutputReference_Override(s StageExternalS3FileFormatOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference)SetFormatName(val *string) {
	if err := j.validateSetFormatNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formatName",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference)SetInternalValue(val *StageExternalS3FileFormat) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3FileFormatOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) PutAvro(value *StageExternalS3FileFormatAvro) {
	if err := s.validatePutAvroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAvro",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) PutCsv(value *StageExternalS3FileFormatCsv) {
	if err := s.validatePutCsvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCsv",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) PutJson(value *StageExternalS3FileFormatJson) {
	if err := s.validatePutJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putJson",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) PutOrc(value *StageExternalS3FileFormatOrc) {
	if err := s.validatePutOrcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOrc",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) PutParquet(value *StageExternalS3FileFormatParquet) {
	if err := s.validatePutParquetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putParquet",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) PutXml(value *StageExternalS3FileFormatXml) {
	if err := s.validatePutXmlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putXml",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) ResetAvro() {
	_jsii_.InvokeVoid(
		s,
		"resetAvro",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) ResetCsv() {
	_jsii_.InvokeVoid(
		s,
		"resetCsv",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) ResetFormatName() {
	_jsii_.InvokeVoid(
		s,
		"resetFormatName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) ResetJson() {
	_jsii_.InvokeVoid(
		s,
		"resetJson",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) ResetOrc() {
	_jsii_.InvokeVoid(
		s,
		"resetOrc",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) ResetParquet() {
	_jsii_.InvokeVoid(
		s,
		"resetParquet",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) ResetXml() {
	_jsii_.InvokeVoid(
		s,
		"resetXml",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_StageExternalS3FileFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

