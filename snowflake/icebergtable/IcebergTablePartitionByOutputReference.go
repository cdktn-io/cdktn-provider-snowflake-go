// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/icebergtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IcebergTablePartitionByOutputReference interface {
	cdktn.ComplexObject
	Bucket() IcebergTablePartitionByBucketOutputReference
	BucketInput() *IcebergTablePartitionByBucket
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
	Day() *string
	SetDay(val *string)
	DayInput() *string
	// Experimental.
	Fqn() *string
	Hour() *string
	SetHour(val *string)
	HourInput() *string
	Identity() *string
	SetIdentity(val *string)
	IdentityInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Month() *string
	SetMonth(val *string)
	MonthInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Truncate() IcebergTablePartitionByTruncateOutputReference
	TruncateInput() *IcebergTablePartitionByTruncate
	Year() *string
	SetYear(val *string)
	YearInput() *string
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
	PutBucket(value *IcebergTablePartitionByBucket)
	PutTruncate(value *IcebergTablePartitionByTruncate)
	ResetBucket()
	ResetDay()
	ResetHour()
	ResetIdentity()
	ResetMonth()
	ResetTruncate()
	ResetYear()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IcebergTablePartitionByOutputReference
type jsiiProxy_IcebergTablePartitionByOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) Bucket() IcebergTablePartitionByBucketOutputReference {
	var returns IcebergTablePartitionByBucketOutputReference
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) BucketInput() *IcebergTablePartitionByBucket {
	var returns *IcebergTablePartitionByBucket
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) Day() *string {
	var returns *string
	_jsii_.Get(
		j,
		"day",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) DayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) Hour() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) HourInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) Identity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) IdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) Month() *string {
	var returns *string
	_jsii_.Get(
		j,
		"month",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) MonthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) Truncate() IcebergTablePartitionByTruncateOutputReference {
	var returns IcebergTablePartitionByTruncateOutputReference
	_jsii_.Get(
		j,
		"truncate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) TruncateInput() *IcebergTablePartitionByTruncate {
	var returns *IcebergTablePartitionByTruncate
	_jsii_.Get(
		j,
		"truncateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) Year() *string {
	var returns *string
	_jsii_.Get(
		j,
		"year",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference) YearInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"yearInput",
		&returns,
	)
	return returns
}


func NewIcebergTablePartitionByOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IcebergTablePartitionByOutputReference {
	_init_.Initialize()

	if err := validateNewIcebergTablePartitionByOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IcebergTablePartitionByOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.icebergTable.IcebergTablePartitionByOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIcebergTablePartitionByOutputReference_Override(i IcebergTablePartitionByOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.icebergTable.IcebergTablePartitionByOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference)SetDay(val *string) {
	if err := j.validateSetDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"day",
		val,
	)
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference)SetHour(val *string) {
	if err := j.validateSetHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hour",
		val,
	)
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference)SetIdentity(val *string) {
	if err := j.validateSetIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identity",
		val,
	)
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference)SetMonth(val *string) {
	if err := j.validateSetMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"month",
		val,
	)
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IcebergTablePartitionByOutputReference)SetYear(val *string) {
	if err := j.validateSetYearParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"year",
		val,
	)
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) PutBucket(value *IcebergTablePartitionByBucket) {
	if err := i.validatePutBucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putBucket",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) PutTruncate(value *IcebergTablePartitionByTruncate) {
	if err := i.validatePutTruncateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTruncate",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) ResetBucket() {
	_jsii_.InvokeVoid(
		i,
		"resetBucket",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) ResetDay() {
	_jsii_.InvokeVoid(
		i,
		"resetDay",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) ResetHour() {
	_jsii_.InvokeVoid(
		i,
		"resetHour",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) ResetIdentity() {
	_jsii_.InvokeVoid(
		i,
		"resetIdentity",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) ResetMonth() {
	_jsii_.InvokeVoid(
		i,
		"resetMonth",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) ResetTruncate() {
	_jsii_.InvokeVoid(
		i,
		"resetTruncate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) ResetYear() {
	_jsii_.InvokeVoid(
		i,
		"resetYear",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTablePartitionByOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

