// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package passwordpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v17/passwordpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PasswordPolicyDescribeOutputOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *PasswordPolicyDescribeOutput
	SetInternalValue(val *PasswordPolicyDescribeOutput)
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

// The jsii proxy struct for PasswordPolicyDescribeOutputOutputReference
type jsiiProxy_PasswordPolicyDescribeOutputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) InternalValue() *PasswordPolicyDescribeOutput {
	var returns *PasswordPolicyDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) PasswordHistory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordHistory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) PasswordLockoutTimeMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordLockoutTimeMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) PasswordMaxAgeDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMaxAgeDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) PasswordMaxLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMaxLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) PasswordMaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMaxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) PasswordMinAgeDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinAgeDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) PasswordMinLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) PasswordMinLowerCaseChars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinLowerCaseChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) PasswordMinNumericChars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinNumericChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) PasswordMinSpecialChars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinSpecialChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) PasswordMinUpperCaseChars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordMinUpperCaseChars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPasswordPolicyDescribeOutputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PasswordPolicyDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewPasswordPolicyDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PasswordPolicyDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.passwordPolicy.PasswordPolicyDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPasswordPolicyDescribeOutputOutputReference_Override(p PasswordPolicyDescribeOutputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.passwordPolicy.PasswordPolicyDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference)SetInternalValue(val *PasswordPolicyDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PasswordPolicyDescribeOutputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PasswordPolicyDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

