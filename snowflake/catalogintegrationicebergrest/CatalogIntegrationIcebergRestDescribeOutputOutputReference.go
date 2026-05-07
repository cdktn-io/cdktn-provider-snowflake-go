// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalogintegrationicebergrest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v17/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v17/catalogintegrationicebergrest/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CatalogIntegrationIcebergRestDescribeOutputOutputReference interface {
	cdktn.ComplexObject
	BearerRestAuthentication() CatalogIntegrationIcebergRestDescribeOutputBearerRestAuthenticationList
	CatalogNamespace() *string
	CatalogSource() *string
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
	Enabled() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *CatalogIntegrationIcebergRestDescribeOutput
	SetInternalValue(val *CatalogIntegrationIcebergRestDescribeOutput)
	OauthRestAuthentication() CatalogIntegrationIcebergRestDescribeOutputOauthRestAuthenticationList
	RefreshIntervalSeconds() *float64
	RestConfig() CatalogIntegrationIcebergRestDescribeOutputRestConfigList
	Sigv4RestAuthentication() CatalogIntegrationIcebergRestDescribeOutputSigv4RestAuthenticationList
	TableFormat() *string
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

// The jsii proxy struct for CatalogIntegrationIcebergRestDescribeOutputOutputReference
type jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) BearerRestAuthentication() CatalogIntegrationIcebergRestDescribeOutputBearerRestAuthenticationList {
	var returns CatalogIntegrationIcebergRestDescribeOutputBearerRestAuthenticationList
	_jsii_.Get(
		j,
		"bearerRestAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) CatalogNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) CatalogSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) Enabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) InternalValue() *CatalogIntegrationIcebergRestDescribeOutput {
	var returns *CatalogIntegrationIcebergRestDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) OauthRestAuthentication() CatalogIntegrationIcebergRestDescribeOutputOauthRestAuthenticationList {
	var returns CatalogIntegrationIcebergRestDescribeOutputOauthRestAuthenticationList
	_jsii_.Get(
		j,
		"oauthRestAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) RefreshIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) RestConfig() CatalogIntegrationIcebergRestDescribeOutputRestConfigList {
	var returns CatalogIntegrationIcebergRestDescribeOutputRestConfigList
	_jsii_.Get(
		j,
		"restConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) Sigv4RestAuthentication() CatalogIntegrationIcebergRestDescribeOutputSigv4RestAuthenticationList {
	var returns CatalogIntegrationIcebergRestDescribeOutputSigv4RestAuthenticationList
	_jsii_.Get(
		j,
		"sigv4RestAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) TableFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCatalogIntegrationIcebergRestDescribeOutputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CatalogIntegrationIcebergRestDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewCatalogIntegrationIcebergRestDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.catalogIntegrationIcebergRest.CatalogIntegrationIcebergRestDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCatalogIntegrationIcebergRestDescribeOutputOutputReference_Override(c CatalogIntegrationIcebergRestDescribeOutputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.catalogIntegrationIcebergRest.CatalogIntegrationIcebergRestDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference)SetInternalValue(val *CatalogIntegrationIcebergRestDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRestDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

