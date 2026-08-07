// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalogintegrationopencatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/catalogintegrationopencatalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CatalogIntegrationOpenCatalogRestConfigOutputReference interface {
	cdktn.ComplexObject
	AccessDelegationMode() *string
	SetAccessDelegationMode(val *string)
	AccessDelegationModeInput() *string
	CatalogApiType() *string
	SetCatalogApiType(val *string)
	CatalogApiTypeInput() *string
	CatalogName() *string
	SetCatalogName(val *string)
	CatalogNameInput() *string
	CatalogUri() *string
	SetCatalogUri(val *string)
	CatalogUriInput() *string
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
	InternalValue() *CatalogIntegrationOpenCatalogRestConfig
	SetInternalValue(val *CatalogIntegrationOpenCatalogRestConfig)
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
	ResetAccessDelegationMode()
	ResetCatalogApiType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CatalogIntegrationOpenCatalogRestConfigOutputReference
type jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) AccessDelegationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessDelegationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) AccessDelegationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessDelegationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) CatalogApiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogApiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) CatalogApiTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogApiTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) CatalogName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) CatalogNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) CatalogUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) CatalogUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) InternalValue() *CatalogIntegrationOpenCatalogRestConfig {
	var returns *CatalogIntegrationOpenCatalogRestConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCatalogIntegrationOpenCatalogRestConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CatalogIntegrationOpenCatalogRestConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCatalogIntegrationOpenCatalogRestConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.catalogIntegrationOpenCatalog.CatalogIntegrationOpenCatalogRestConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCatalogIntegrationOpenCatalogRestConfigOutputReference_Override(c CatalogIntegrationOpenCatalogRestConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.catalogIntegrationOpenCatalog.CatalogIntegrationOpenCatalogRestConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference)SetAccessDelegationMode(val *string) {
	if err := j.validateSetAccessDelegationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessDelegationMode",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference)SetCatalogApiType(val *string) {
	if err := j.validateSetCatalogApiTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogApiType",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference)SetCatalogName(val *string) {
	if err := j.validateSetCatalogNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogName",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference)SetCatalogUri(val *string) {
	if err := j.validateSetCatalogUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogUri",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference)SetInternalValue(val *CatalogIntegrationOpenCatalogRestConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) ResetAccessDelegationMode() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessDelegationMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) ResetCatalogApiType() {
	_jsii_.InvokeVoid(
		c,
		"resetCatalogApiType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationOpenCatalogRestConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

