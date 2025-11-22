// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package semanticview

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/semanticview/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SemanticViewRelationshipsOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ReferencedRelationshipColumns() *[]*string
	SetReferencedRelationshipColumns(val *[]*string)
	ReferencedRelationshipColumnsInput() *[]*string
	ReferencedTableNameOrAlias() SemanticViewRelationshipsReferencedTableNameOrAliasOutputReference
	ReferencedTableNameOrAliasInput() *SemanticViewRelationshipsReferencedTableNameOrAlias
	RelationshipColumns() *[]*string
	SetRelationshipColumns(val *[]*string)
	RelationshipColumnsInput() *[]*string
	RelationshipIdentifier() *string
	SetRelationshipIdentifier(val *string)
	RelationshipIdentifierInput() *string
	TableNameOrAlias() SemanticViewRelationshipsTableNameOrAliasOutputReference
	TableNameOrAliasInput() *SemanticViewRelationshipsTableNameOrAlias
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutReferencedTableNameOrAlias(value *SemanticViewRelationshipsReferencedTableNameOrAlias)
	PutTableNameOrAlias(value *SemanticViewRelationshipsTableNameOrAlias)
	ResetReferencedRelationshipColumns()
	ResetRelationshipIdentifier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SemanticViewRelationshipsOutputReference
type jsiiProxy_SemanticViewRelationshipsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) ReferencedRelationshipColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"referencedRelationshipColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) ReferencedRelationshipColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"referencedRelationshipColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) ReferencedTableNameOrAlias() SemanticViewRelationshipsReferencedTableNameOrAliasOutputReference {
	var returns SemanticViewRelationshipsReferencedTableNameOrAliasOutputReference
	_jsii_.Get(
		j,
		"referencedTableNameOrAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) ReferencedTableNameOrAliasInput() *SemanticViewRelationshipsReferencedTableNameOrAlias {
	var returns *SemanticViewRelationshipsReferencedTableNameOrAlias
	_jsii_.Get(
		j,
		"referencedTableNameOrAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) RelationshipColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"relationshipColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) RelationshipColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"relationshipColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) RelationshipIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationshipIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) RelationshipIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationshipIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) TableNameOrAlias() SemanticViewRelationshipsTableNameOrAliasOutputReference {
	var returns SemanticViewRelationshipsTableNameOrAliasOutputReference
	_jsii_.Get(
		j,
		"tableNameOrAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) TableNameOrAliasInput() *SemanticViewRelationshipsTableNameOrAlias {
	var returns *SemanticViewRelationshipsTableNameOrAlias
	_jsii_.Get(
		j,
		"tableNameOrAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSemanticViewRelationshipsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SemanticViewRelationshipsOutputReference {
	_init_.Initialize()

	if err := validateNewSemanticViewRelationshipsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SemanticViewRelationshipsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.semanticView.SemanticViewRelationshipsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSemanticViewRelationshipsOutputReference_Override(s SemanticViewRelationshipsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.semanticView.SemanticViewRelationshipsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference)SetReferencedRelationshipColumns(val *[]*string) {
	if err := j.validateSetReferencedRelationshipColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referencedRelationshipColumns",
		val,
	)
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference)SetRelationshipColumns(val *[]*string) {
	if err := j.validateSetRelationshipColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relationshipColumns",
		val,
	)
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference)SetRelationshipIdentifier(val *string) {
	if err := j.validateSetRelationshipIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relationshipIdentifier",
		val,
	)
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SemanticViewRelationshipsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) PutReferencedTableNameOrAlias(value *SemanticViewRelationshipsReferencedTableNameOrAlias) {
	if err := s.validatePutReferencedTableNameOrAliasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReferencedTableNameOrAlias",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) PutTableNameOrAlias(value *SemanticViewRelationshipsTableNameOrAlias) {
	if err := s.validatePutTableNameOrAliasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTableNameOrAlias",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) ResetReferencedRelationshipColumns() {
	_jsii_.InvokeVoid(
		s,
		"resetReferencedRelationshipColumns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) ResetRelationshipIdentifier() {
	_jsii_.InvokeVoid(
		s,
		"resetRelationshipIdentifier",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SemanticViewRelationshipsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

