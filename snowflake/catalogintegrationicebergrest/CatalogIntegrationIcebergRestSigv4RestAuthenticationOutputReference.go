// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalogintegrationicebergrest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/catalogintegrationicebergrest/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference interface {
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
	InternalValue() *CatalogIntegrationIcebergRestSigv4RestAuthentication
	SetInternalValue(val *CatalogIntegrationIcebergRestSigv4RestAuthentication)
	Sigv4ExternalId() *string
	SetSigv4ExternalId(val *string)
	Sigv4ExternalIdInput() *string
	Sigv4IamRole() *string
	SetSigv4IamRole(val *string)
	Sigv4IamRoleInput() *string
	Sigv4SigningRegion() *string
	SetSigv4SigningRegion(val *string)
	Sigv4SigningRegionInput() *string
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
	ResetSigv4ExternalId()
	ResetSigv4SigningRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference
type jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) InternalValue() *CatalogIntegrationIcebergRestSigv4RestAuthentication {
	var returns *CatalogIntegrationIcebergRestSigv4RestAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) Sigv4ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sigv4ExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) Sigv4ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sigv4ExternalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) Sigv4IamRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sigv4IamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) Sigv4IamRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sigv4IamRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) Sigv4SigningRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sigv4SigningRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) Sigv4SigningRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sigv4SigningRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference {
	_init_.Initialize()

	if err := validateNewCatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.catalogIntegrationIcebergRest.CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference_Override(c CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.catalogIntegrationIcebergRest.CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference)SetInternalValue(val *CatalogIntegrationIcebergRestSigv4RestAuthentication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference)SetSigv4ExternalId(val *string) {
	if err := j.validateSetSigv4ExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sigv4ExternalId",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference)SetSigv4IamRole(val *string) {
	if err := j.validateSetSigv4IamRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sigv4IamRole",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference)SetSigv4SigningRegion(val *string) {
	if err := j.validateSetSigv4SigningRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sigv4SigningRegion",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) ResetSigv4ExternalId() {
	_jsii_.InvokeVoid(
		c,
		"resetSigv4ExternalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) ResetSigv4SigningRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetSigv4SigningRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

