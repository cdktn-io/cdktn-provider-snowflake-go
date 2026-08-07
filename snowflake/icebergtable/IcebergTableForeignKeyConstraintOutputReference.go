// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/icebergtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IcebergTableForeignKeyConstraintOutputReference interface {
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
	Match() *string
	SetMatch(val *string)
	MatchInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	OnDelete() *string
	SetOnDelete(val *string)
	OnDeleteInput() *string
	OnUpdate() *string
	SetOnUpdate(val *string)
	OnUpdateInput() *string
	RefColumn() *[]*string
	SetRefColumn(val *[]*string)
	RefColumnInput() *[]*string
	Rely() *string
	SetRely(val *string)
	RelyInput() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
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
	ResetMatch()
	ResetName()
	ResetOnDelete()
	ResetOnUpdate()
	ResetRefColumn()
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

// The jsii proxy struct for IcebergTableForeignKeyConstraintOutputReference
type jsiiProxy_IcebergTableForeignKeyConstraintOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) Column() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ColumnInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) Deferrable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deferrable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) DeferrableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deferrableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) Enable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) EnableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) Enforced() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforced",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) EnforcedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforcedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) InitiallyDeferred() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initiallyDeferred",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) InitiallyDeferredInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initiallyDeferredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) Match() *string {
	var returns *string
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) MatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) OnDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) OnDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) OnUpdate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) OnUpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) RefColumn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"refColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) RefColumnInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"refColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) Rely() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rely",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) RelyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) Validate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ValidateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validateInput",
		&returns,
	)
	return returns
}


func NewIcebergTableForeignKeyConstraintOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IcebergTableForeignKeyConstraintOutputReference {
	_init_.Initialize()

	if err := validateNewIcebergTableForeignKeyConstraintOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IcebergTableForeignKeyConstraintOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.icebergTable.IcebergTableForeignKeyConstraintOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIcebergTableForeignKeyConstraintOutputReference_Override(i IcebergTableForeignKeyConstraintOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.icebergTable.IcebergTableForeignKeyConstraintOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetColumn(val *[]*string) {
	if err := j.validateSetColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"column",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetDeferrable(val *string) {
	if err := j.validateSetDeferrableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deferrable",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetEnable(val *string) {
	if err := j.validateSetEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enable",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetEnforced(val *string) {
	if err := j.validateSetEnforcedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforced",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetInitiallyDeferred(val *string) {
	if err := j.validateSetInitiallyDeferredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initiallyDeferred",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetMatch(val *string) {
	if err := j.validateSetMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"match",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetOnDelete(val *string) {
	if err := j.validateSetOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDelete",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetOnUpdate(val *string) {
	if err := j.validateSetOnUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onUpdate",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetRefColumn(val *[]*string) {
	if err := j.validateSetRefColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refColumn",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetRely(val *string) {
	if err := j.validateSetRelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rely",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference)SetValidate(val *string) {
	if err := j.validateSetValidateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validate",
		val,
	)
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		i,
		"resetComment",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ResetDeferrable() {
	_jsii_.InvokeVoid(
		i,
		"resetDeferrable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ResetEnable() {
	_jsii_.InvokeVoid(
		i,
		"resetEnable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ResetEnforced() {
	_jsii_.InvokeVoid(
		i,
		"resetEnforced",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ResetInitiallyDeferred() {
	_jsii_.InvokeVoid(
		i,
		"resetInitiallyDeferred",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ResetMatch() {
	_jsii_.InvokeVoid(
		i,
		"resetMatch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		i,
		"resetName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ResetOnDelete() {
	_jsii_.InvokeVoid(
		i,
		"resetOnDelete",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ResetOnUpdate() {
	_jsii_.InvokeVoid(
		i,
		"resetOnUpdate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ResetRefColumn() {
	_jsii_.InvokeVoid(
		i,
		"resetRefColumn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ResetRely() {
	_jsii_.InvokeVoid(
		i,
		"resetRely",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ResetValidate() {
	_jsii_.InvokeVoid(
		i,
		"resetValidate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (i *jsiiProxy_IcebergTableForeignKeyConstraintOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

