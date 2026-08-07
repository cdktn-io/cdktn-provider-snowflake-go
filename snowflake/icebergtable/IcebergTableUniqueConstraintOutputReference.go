// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/icebergtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IcebergTableUniqueConstraintOutputReference interface {
	cdktn.ComplexObject
	Column() *[]*string
	SetColumn(val *[]*string)
	ColumnInput() *[]*string
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	Deferrable() *string
	SetDeferrable(val *string)
	DeferrableInput() *string
	Enable() *string
	SetEnable(val *string)
	EnableInput() *string
	Enforced() *string
	SetEnforced(val *string)
	EnforcedInput() *string
	// Experimental.
	Fqn() *string
	InitiallyDeferred() *string
	SetInitiallyDeferred(val *string)
	InitiallyDeferredInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Rely() *string
	SetRely(val *string)
	RelyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Validate() *string
	SetValidate(val *string)
	ValidateInput() *string
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
	ResetComment()
	ResetDeferrable()
	ResetEnable()
	ResetEnforced()
	ResetInitiallyDeferred()
	ResetName()
	ResetRely()
	ResetValidate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IcebergTableUniqueConstraintOutputReference
type jsiiProxy_IcebergTableUniqueConstraintOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) Column() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ColumnInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) Deferrable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deferrable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) DeferrableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deferrableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) Enable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) EnableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) Enforced() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforced",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) EnforcedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforcedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) InitiallyDeferred() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initiallyDeferred",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) InitiallyDeferredInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initiallyDeferredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) Rely() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rely",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) RelyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) Validate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ValidateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validateInput",
		&returns,
	)
	return returns
}


func NewIcebergTableUniqueConstraintOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IcebergTableUniqueConstraintOutputReference {
	_init_.Initialize()

	if err := validateNewIcebergTableUniqueConstraintOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IcebergTableUniqueConstraintOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.icebergTable.IcebergTableUniqueConstraintOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIcebergTableUniqueConstraintOutputReference_Override(i IcebergTableUniqueConstraintOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.icebergTable.IcebergTableUniqueConstraintOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetColumn(val *[]*string) {
	if err := j.validateSetColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"column",
		val,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetDeferrable(val *string) {
	if err := j.validateSetDeferrableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deferrable",
		val,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetEnable(val *string) {
	if err := j.validateSetEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enable",
		val,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetEnforced(val *string) {
	if err := j.validateSetEnforcedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforced",
		val,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetInitiallyDeferred(val *string) {
	if err := j.validateSetInitiallyDeferredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initiallyDeferred",
		val,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetRely(val *string) {
	if err := j.validateSetRelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rely",
		val,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IcebergTableUniqueConstraintOutputReference)SetValidate(val *string) {
	if err := j.validateSetValidateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validate",
		val,
	)
}

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		i,
		"resetComment",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ResetDeferrable() {
	_jsii_.InvokeVoid(
		i,
		"resetDeferrable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ResetEnable() {
	_jsii_.InvokeVoid(
		i,
		"resetEnable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ResetEnforced() {
	_jsii_.InvokeVoid(
		i,
		"resetEnforced",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ResetInitiallyDeferred() {
	_jsii_.InvokeVoid(
		i,
		"resetInitiallyDeferred",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		i,
		"resetName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ResetRely() {
	_jsii_.InvokeVoid(
		i,
		"resetRely",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ResetValidate() {
	_jsii_.InvokeVoid(
		i,
		"resetValidate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IcebergTableUniqueConstraintOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

