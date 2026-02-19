// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3compatible

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/stageexternals3compatible/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageExternalS3CompatibleFileFormatOutputReference interface {
	cdktn.ComplexObject
	Avro() StageExternalS3CompatibleFileFormatAvroOutputReference
	AvroInput() *StageExternalS3CompatibleFileFormatAvro
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
	Csv() StageExternalS3CompatibleFileFormatCsvOutputReference
	CsvInput() *StageExternalS3CompatibleFileFormatCsv
	FormatName() *string
	SetFormatName(val *string)
	FormatNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *StageExternalS3CompatibleFileFormat
	SetInternalValue(val *StageExternalS3CompatibleFileFormat)
	Json() StageExternalS3CompatibleFileFormatJsonOutputReference
	JsonInput() *StageExternalS3CompatibleFileFormatJson
	Orc() StageExternalS3CompatibleFileFormatOrcOutputReference
	OrcInput() *StageExternalS3CompatibleFileFormatOrc
	Parquet() StageExternalS3CompatibleFileFormatParquetOutputReference
	ParquetInput() *StageExternalS3CompatibleFileFormatParquet
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Xml() StageExternalS3CompatibleFileFormatXmlOutputReference
	XmlInput() *StageExternalS3CompatibleFileFormatXml
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
	PutAvro(value *StageExternalS3CompatibleFileFormatAvro)
	PutCsv(value *StageExternalS3CompatibleFileFormatCsv)
	PutJson(value *StageExternalS3CompatibleFileFormatJson)
	PutOrc(value *StageExternalS3CompatibleFileFormatOrc)
	PutParquet(value *StageExternalS3CompatibleFileFormatParquet)
	PutXml(value *StageExternalS3CompatibleFileFormatXml)
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

// The jsii proxy struct for StageExternalS3CompatibleFileFormatOutputReference
type jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) Avro() StageExternalS3CompatibleFileFormatAvroOutputReference {
	var returns StageExternalS3CompatibleFileFormatAvroOutputReference
	_jsii_.Get(
		j,
		"avro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) AvroInput() *StageExternalS3CompatibleFileFormatAvro {
	var returns *StageExternalS3CompatibleFileFormatAvro
	_jsii_.Get(
		j,
		"avroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) Csv() StageExternalS3CompatibleFileFormatCsvOutputReference {
	var returns StageExternalS3CompatibleFileFormatCsvOutputReference
	_jsii_.Get(
		j,
		"csv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) CsvInput() *StageExternalS3CompatibleFileFormatCsv {
	var returns *StageExternalS3CompatibleFileFormatCsv
	_jsii_.Get(
		j,
		"csvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) FormatName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) FormatNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) InternalValue() *StageExternalS3CompatibleFileFormat {
	var returns *StageExternalS3CompatibleFileFormat
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) Json() StageExternalS3CompatibleFileFormatJsonOutputReference {
	var returns StageExternalS3CompatibleFileFormatJsonOutputReference
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) JsonInput() *StageExternalS3CompatibleFileFormatJson {
	var returns *StageExternalS3CompatibleFileFormatJson
	_jsii_.Get(
		j,
		"jsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) Orc() StageExternalS3CompatibleFileFormatOrcOutputReference {
	var returns StageExternalS3CompatibleFileFormatOrcOutputReference
	_jsii_.Get(
		j,
		"orc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) OrcInput() *StageExternalS3CompatibleFileFormatOrc {
	var returns *StageExternalS3CompatibleFileFormatOrc
	_jsii_.Get(
		j,
		"orcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) Parquet() StageExternalS3CompatibleFileFormatParquetOutputReference {
	var returns StageExternalS3CompatibleFileFormatParquetOutputReference
	_jsii_.Get(
		j,
		"parquet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) ParquetInput() *StageExternalS3CompatibleFileFormatParquet {
	var returns *StageExternalS3CompatibleFileFormatParquet
	_jsii_.Get(
		j,
		"parquetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) Xml() StageExternalS3CompatibleFileFormatXmlOutputReference {
	var returns StageExternalS3CompatibleFileFormatXmlOutputReference
	_jsii_.Get(
		j,
		"xml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) XmlInput() *StageExternalS3CompatibleFileFormatXml {
	var returns *StageExternalS3CompatibleFileFormatXml
	_jsii_.Get(
		j,
		"xmlInput",
		&returns,
	)
	return returns
}


func NewStageExternalS3CompatibleFileFormatOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StageExternalS3CompatibleFileFormatOutputReference {
	_init_.Initialize()

	if err := validateNewStageExternalS3CompatibleFileFormatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalS3Compatible.StageExternalS3CompatibleFileFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStageExternalS3CompatibleFileFormatOutputReference_Override(s StageExternalS3CompatibleFileFormatOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalS3Compatible.StageExternalS3CompatibleFileFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference)SetFormatName(val *string) {
	if err := j.validateSetFormatNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formatName",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference)SetInternalValue(val *StageExternalS3CompatibleFileFormat) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) PutAvro(value *StageExternalS3CompatibleFileFormatAvro) {
	if err := s.validatePutAvroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAvro",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) PutCsv(value *StageExternalS3CompatibleFileFormatCsv) {
	if err := s.validatePutCsvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCsv",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) PutJson(value *StageExternalS3CompatibleFileFormatJson) {
	if err := s.validatePutJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putJson",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) PutOrc(value *StageExternalS3CompatibleFileFormatOrc) {
	if err := s.validatePutOrcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOrc",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) PutParquet(value *StageExternalS3CompatibleFileFormatParquet) {
	if err := s.validatePutParquetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putParquet",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) PutXml(value *StageExternalS3CompatibleFileFormatXml) {
	if err := s.validatePutXmlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putXml",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) ResetAvro() {
	_jsii_.InvokeVoid(
		s,
		"resetAvro",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) ResetCsv() {
	_jsii_.InvokeVoid(
		s,
		"resetCsv",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) ResetFormatName() {
	_jsii_.InvokeVoid(
		s,
		"resetFormatName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) ResetJson() {
	_jsii_.InvokeVoid(
		s,
		"resetJson",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) ResetOrc() {
	_jsii_.InvokeVoid(
		s,
		"resetOrc",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) ResetParquet() {
	_jsii_.InvokeVoid(
		s,
		"resetParquet",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) ResetXml() {
	_jsii_.InvokeVoid(
		s,
		"resetXml",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_StageExternalS3CompatibleFileFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

