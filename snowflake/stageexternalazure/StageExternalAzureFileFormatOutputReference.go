// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternalazure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/stageexternalazure/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type StageExternalAzureFileFormatOutputReference interface {
	cdktn.ComplexObject
	Avro() StageExternalAzureFileFormatAvroOutputReference
	AvroInput() *StageExternalAzureFileFormatAvro
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
	Csv() StageExternalAzureFileFormatCsvOutputReference
	CsvInput() *StageExternalAzureFileFormatCsv
	FormatName() *string
	SetFormatName(val *string)
	FormatNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *StageExternalAzureFileFormat
	SetInternalValue(val *StageExternalAzureFileFormat)
	Json() StageExternalAzureFileFormatJsonOutputReference
	JsonInput() *StageExternalAzureFileFormatJson
	Orc() StageExternalAzureFileFormatOrcOutputReference
	OrcInput() *StageExternalAzureFileFormatOrc
	Parquet() StageExternalAzureFileFormatParquetOutputReference
	ParquetInput() *StageExternalAzureFileFormatParquet
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Xml() StageExternalAzureFileFormatXmlOutputReference
	XmlInput() *StageExternalAzureFileFormatXml
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
	PutAvro(value *StageExternalAzureFileFormatAvro)
	PutCsv(value *StageExternalAzureFileFormatCsv)
	PutJson(value *StageExternalAzureFileFormatJson)
	PutOrc(value *StageExternalAzureFileFormatOrc)
	PutParquet(value *StageExternalAzureFileFormatParquet)
	PutXml(value *StageExternalAzureFileFormatXml)
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

// The jsii proxy struct for StageExternalAzureFileFormatOutputReference
type jsiiProxy_StageExternalAzureFileFormatOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) Avro() StageExternalAzureFileFormatAvroOutputReference {
	var returns StageExternalAzureFileFormatAvroOutputReference
	_jsii_.Get(
		j,
		"avro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) AvroInput() *StageExternalAzureFileFormatAvro {
	var returns *StageExternalAzureFileFormatAvro
	_jsii_.Get(
		j,
		"avroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) Csv() StageExternalAzureFileFormatCsvOutputReference {
	var returns StageExternalAzureFileFormatCsvOutputReference
	_jsii_.Get(
		j,
		"csv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) CsvInput() *StageExternalAzureFileFormatCsv {
	var returns *StageExternalAzureFileFormatCsv
	_jsii_.Get(
		j,
		"csvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) FormatName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) FormatNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) InternalValue() *StageExternalAzureFileFormat {
	var returns *StageExternalAzureFileFormat
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) Json() StageExternalAzureFileFormatJsonOutputReference {
	var returns StageExternalAzureFileFormatJsonOutputReference
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) JsonInput() *StageExternalAzureFileFormatJson {
	var returns *StageExternalAzureFileFormatJson
	_jsii_.Get(
		j,
		"jsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) Orc() StageExternalAzureFileFormatOrcOutputReference {
	var returns StageExternalAzureFileFormatOrcOutputReference
	_jsii_.Get(
		j,
		"orc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) OrcInput() *StageExternalAzureFileFormatOrc {
	var returns *StageExternalAzureFileFormatOrc
	_jsii_.Get(
		j,
		"orcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) Parquet() StageExternalAzureFileFormatParquetOutputReference {
	var returns StageExternalAzureFileFormatParquetOutputReference
	_jsii_.Get(
		j,
		"parquet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) ParquetInput() *StageExternalAzureFileFormatParquet {
	var returns *StageExternalAzureFileFormatParquet
	_jsii_.Get(
		j,
		"parquetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) Xml() StageExternalAzureFileFormatXmlOutputReference {
	var returns StageExternalAzureFileFormatXmlOutputReference
	_jsii_.Get(
		j,
		"xml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference) XmlInput() *StageExternalAzureFileFormatXml {
	var returns *StageExternalAzureFileFormatXml
	_jsii_.Get(
		j,
		"xmlInput",
		&returns,
	)
	return returns
}


func NewStageExternalAzureFileFormatOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) StageExternalAzureFileFormatOutputReference {
	_init_.Initialize()

	if err := validateNewStageExternalAzureFileFormatOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StageExternalAzureFileFormatOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalAzure.StageExternalAzureFileFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStageExternalAzureFileFormatOutputReference_Override(s StageExternalAzureFileFormatOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.stageExternalAzure.StageExternalAzureFileFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference)SetFormatName(val *string) {
	if err := j.validateSetFormatNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formatName",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference)SetInternalValue(val *StageExternalAzureFileFormat) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StageExternalAzureFileFormatOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) PutAvro(value *StageExternalAzureFileFormatAvro) {
	if err := s.validatePutAvroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAvro",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) PutCsv(value *StageExternalAzureFileFormatCsv) {
	if err := s.validatePutCsvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCsv",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) PutJson(value *StageExternalAzureFileFormatJson) {
	if err := s.validatePutJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putJson",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) PutOrc(value *StageExternalAzureFileFormatOrc) {
	if err := s.validatePutOrcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOrc",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) PutParquet(value *StageExternalAzureFileFormatParquet) {
	if err := s.validatePutParquetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putParquet",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) PutXml(value *StageExternalAzureFileFormatXml) {
	if err := s.validatePutXmlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putXml",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) ResetAvro() {
	_jsii_.InvokeVoid(
		s,
		"resetAvro",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) ResetCsv() {
	_jsii_.InvokeVoid(
		s,
		"resetCsv",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) ResetFormatName() {
	_jsii_.InvokeVoid(
		s,
		"resetFormatName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) ResetJson() {
	_jsii_.InvokeVoid(
		s,
		"resetJson",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) ResetOrc() {
	_jsii_.InvokeVoid(
		s,
		"resetOrc",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) ResetParquet() {
	_jsii_.InvokeVoid(
		s,
		"resetParquet",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) ResetXml() {
	_jsii_.InvokeVoid(
		s,
		"resetXml",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_StageExternalAzureFileFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

