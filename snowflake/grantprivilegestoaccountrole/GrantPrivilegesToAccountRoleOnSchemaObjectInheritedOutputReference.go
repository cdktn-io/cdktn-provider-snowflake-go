// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grantprivilegestoaccountrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/grantprivilegestoaccountrole/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference interface {
	cdktn.ComplexObject
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
	InAccount() interface{}
	SetInAccount(val interface{})
	InAccountInput() interface{}
	InDatabase() *string
	SetInDatabase(val *string)
	InDatabaseInput() *string
	InSchema() *string
	SetInSchema(val *string)
	InSchemaInput() *string
	InternalValue() *GrantPrivilegesToAccountRoleOnSchemaObjectInherited
	SetInternalValue(val *GrantPrivilegesToAccountRoleOnSchemaObjectInherited)
	ObjectTypePlural() *string
	SetObjectTypePlural(val *string)
	ObjectTypePluralInput() *string
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
	ResetInAccount()
	ResetInDatabase()
	ResetInSchema()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference
type jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) InAccount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) InAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) InDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) InDatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) InSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) InSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) InternalValue() *GrantPrivilegesToAccountRoleOnSchemaObjectInherited {
	var returns *GrantPrivilegesToAccountRoleOnSchemaObjectInherited
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) ObjectTypePlural() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTypePlural",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) ObjectTypePluralInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectTypePluralInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference {
	_init_.Initialize()

	if err := validateNewGrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.grantPrivilegesToAccountRole.GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference_Override(g GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.grantPrivilegesToAccountRole.GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference)SetInAccount(val interface{}) {
	if err := j.validateSetInAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inAccount",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference)SetInDatabase(val *string) {
	if err := j.validateSetInDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inDatabase",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference)SetInSchema(val *string) {
	if err := j.validateSetInSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inSchema",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference)SetInternalValue(val *GrantPrivilegesToAccountRoleOnSchemaObjectInherited) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference)SetObjectTypePlural(val *string) {
	if err := j.validateSetObjectTypePluralParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectTypePlural",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) ResetInAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetInAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) ResetInDatabase() {
	_jsii_.InvokeVoid(
		g,
		"resetInDatabase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) ResetInSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetInSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GrantPrivilegesToAccountRoleOnSchemaObjectInheritedOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

