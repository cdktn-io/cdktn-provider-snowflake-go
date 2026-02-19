// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakestages

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/datasnowflakestages/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference interface {
	cdktn.ComplexObject
	Avro() DataSnowflakeStagesStagesDescribeOutputFileFormatAvroList
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
	Csv() DataSnowflakeStagesStagesDescribeOutputFileFormatCsvList
	FormatName() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataSnowflakeStagesStagesDescribeOutputFileFormat
	SetInternalValue(val *DataSnowflakeStagesStagesDescribeOutputFileFormat)
	Json() DataSnowflakeStagesStagesDescribeOutputFileFormatJsonList
	Orc() DataSnowflakeStagesStagesDescribeOutputFileFormatOrcList
	Parquet() DataSnowflakeStagesStagesDescribeOutputFileFormatParquetList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Xml() DataSnowflakeStagesStagesDescribeOutputFileFormatXmlList
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

// The jsii proxy struct for DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference
type jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) Avro() DataSnowflakeStagesStagesDescribeOutputFileFormatAvroList {
	var returns DataSnowflakeStagesStagesDescribeOutputFileFormatAvroList
	_jsii_.Get(
		j,
		"avro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) Csv() DataSnowflakeStagesStagesDescribeOutputFileFormatCsvList {
	var returns DataSnowflakeStagesStagesDescribeOutputFileFormatCsvList
	_jsii_.Get(
		j,
		"csv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) FormatName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) InternalValue() *DataSnowflakeStagesStagesDescribeOutputFileFormat {
	var returns *DataSnowflakeStagesStagesDescribeOutputFileFormat
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) Json() DataSnowflakeStagesStagesDescribeOutputFileFormatJsonList {
	var returns DataSnowflakeStagesStagesDescribeOutputFileFormatJsonList
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) Orc() DataSnowflakeStagesStagesDescribeOutputFileFormatOrcList {
	var returns DataSnowflakeStagesStagesDescribeOutputFileFormatOrcList
	_jsii_.Get(
		j,
		"orc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) Parquet() DataSnowflakeStagesStagesDescribeOutputFileFormatParquetList {
	var returns DataSnowflakeStagesStagesDescribeOutputFileFormatParquetList
	_jsii_.Get(
		j,
		"parquet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) Xml() DataSnowflakeStagesStagesDescribeOutputFileFormatXmlList {
	var returns DataSnowflakeStagesStagesDescribeOutputFileFormatXmlList
	_jsii_.Get(
		j,
		"xml",
		&returns,
	)
	return returns
}


func NewDataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeStagesStagesDescribeOutputFileFormatOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeStages.DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference_Override(d DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeStages.DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference)SetInternalValue(val *DataSnowflakeStagesStagesDescribeOutputFileFormat) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeStagesStagesDescribeOutputFileFormatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

