// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package semanticview

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/jsii"

	"github.com/cdktf/cdktf-provider-snowflake-go/snowflake/v15/semanticview/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SemanticViewMetricsWindowFunctionOutputReference interface {
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
	InternalValue() *SemanticViewMetricsWindowFunction
	SetInternalValue(val *SemanticViewMetricsWindowFunction)
	OverClause() SemanticViewMetricsWindowFunctionOverClauseOutputReference
	OverClauseInput() *SemanticViewMetricsWindowFunctionOverClause
	QualifiedExpressionName() *string
	SetQualifiedExpressionName(val *string)
	QualifiedExpressionNameInput() *string
	SqlExpression() *string
	SetSqlExpression(val *string)
	SqlExpressionInput() *string
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
	PutOverClause(value *SemanticViewMetricsWindowFunctionOverClause)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SemanticViewMetricsWindowFunctionOutputReference
type jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) InternalValue() *SemanticViewMetricsWindowFunction {
	var returns *SemanticViewMetricsWindowFunction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) OverClause() SemanticViewMetricsWindowFunctionOverClauseOutputReference {
	var returns SemanticViewMetricsWindowFunctionOverClauseOutputReference
	_jsii_.Get(
		j,
		"overClause",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) OverClauseInput() *SemanticViewMetricsWindowFunctionOverClause {
	var returns *SemanticViewMetricsWindowFunctionOverClause
	_jsii_.Get(
		j,
		"overClauseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) QualifiedExpressionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifiedExpressionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) QualifiedExpressionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifiedExpressionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) SqlExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) SqlExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSemanticViewMetricsWindowFunctionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SemanticViewMetricsWindowFunctionOutputReference {
	_init_.Initialize()

	if err := validateNewSemanticViewMetricsWindowFunctionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-snowflake.semanticView.SemanticViewMetricsWindowFunctionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSemanticViewMetricsWindowFunctionOutputReference_Override(s SemanticViewMetricsWindowFunctionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-snowflake.semanticView.SemanticViewMetricsWindowFunctionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference)SetInternalValue(val *SemanticViewMetricsWindowFunction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference)SetQualifiedExpressionName(val *string) {
	if err := j.validateSetQualifiedExpressionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qualifiedExpressionName",
		val,
	)
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference)SetSqlExpression(val *string) {
	if err := j.validateSetSqlExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlExpression",
		val,
	)
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) PutOverClause(value *SemanticViewMetricsWindowFunctionOverClause) {
	if err := s.validatePutOverClauseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOverClause",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SemanticViewMetricsWindowFunctionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

