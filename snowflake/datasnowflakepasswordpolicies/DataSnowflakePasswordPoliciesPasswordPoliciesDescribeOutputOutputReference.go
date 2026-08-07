// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakepasswordpolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/datasnowflakepasswordpolicies/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference interface {
	cdktn.ComplexObject
	Comment() *string
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
	DatabaseName() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutput
	SetInternalValue(val *DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutput)
	Name() *string
	Owner() *string
	PasswordHistory() *float64
	PasswordLockoutTimeMins() *float64
	PasswordMaxAgeDays() *float64
	PasswordMaxLength() *float64
	PasswordMaxRetries() *float64
	PasswordMinAgeDays() *float64
	PasswordMinLength() *float64
	PasswordMinLowerCaseChars() *float64
	PasswordMinNumericChars() *float64
	PasswordMinSpecialChars() *float64
	PasswordMinUpperCaseChars() *float64
	SchemaName() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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

// The jsii proxy struct for DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference
type jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) InternalValue() *DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutput {
	var returns *DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) PasswordHistory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordHistory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) PasswordLockoutTimeMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordLockoutTimeMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) PasswordMaxAgeDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMaxAgeDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) PasswordMaxLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMaxLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) PasswordMaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMaxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) PasswordMinAgeDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinAgeDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) PasswordMinLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) PasswordMinLowerCaseChars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinLowerCaseChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) PasswordMinNumericChars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinNumericChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) PasswordMinSpecialChars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinSpecialChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) PasswordMinUpperCaseChars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinUpperCaseChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakePasswordPolicies.DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference_Override(d DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakePasswordPolicies.DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference)SetInternalValue(val *DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataSnowflakePasswordPoliciesPasswordPoliciesDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

