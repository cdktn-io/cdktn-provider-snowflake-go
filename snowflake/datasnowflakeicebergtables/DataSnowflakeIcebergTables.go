// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeicebergtables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/datasnowflakeicebergtables/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/data-sources/iceberg_tables snowflake_iceberg_tables}.
type DataSnowflakeIcebergTables interface {
	cdktn.TerraformDataSource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IcebergTables() DataSnowflakeIcebergTablesIcebergTablesList
	Id() *string
	SetId(val *string)
	IdInput() *string
	In() DataSnowflakeIcebergTablesInOutputReference
	InInput() *DataSnowflakeIcebergTablesIn
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Like() *string
	SetLike(val *string)
	LikeInput() *string
	Limit() DataSnowflakeIcebergTablesLimitOutputReference
	LimitInput() *DataSnowflakeIcebergTablesLimit
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	StartsWith() *string
	SetStartsWith(val *string)
	StartsWithInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WithDescribe() interface{}
	SetWithDescribe(val interface{})
	WithDescribeInput() interface{}
	WithParameters() interface{}
	SetWithParameters(val interface{})
	WithParametersInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutIn(value *DataSnowflakeIcebergTablesIn)
	PutLimit(value *DataSnowflakeIcebergTablesLimit)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetId()
	ResetIn()
	ResetLike()
	ResetLimit()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStartsWith()
	ResetWithDescribe()
	ResetWithParameters()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for DataSnowflakeIcebergTables
type jsiiProxy_DataSnowflakeIcebergTables struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) IcebergTables() DataSnowflakeIcebergTablesIcebergTablesList {
	var returns DataSnowflakeIcebergTablesIcebergTablesList
	_jsii_.Get(
		j,
		"icebergTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) In() DataSnowflakeIcebergTablesInOutputReference {
	var returns DataSnowflakeIcebergTablesInOutputReference
	_jsii_.Get(
		j,
		"in",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) InInput() *DataSnowflakeIcebergTablesIn {
	var returns *DataSnowflakeIcebergTablesIn
	_jsii_.Get(
		j,
		"inInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) Like() *string {
	var returns *string
	_jsii_.Get(
		j,
		"like",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) LikeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"likeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) Limit() DataSnowflakeIcebergTablesLimitOutputReference {
	var returns DataSnowflakeIcebergTablesLimitOutputReference
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) LimitInput() *DataSnowflakeIcebergTablesLimit {
	var returns *DataSnowflakeIcebergTablesLimit
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) StartsWith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) StartsWithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) WithDescribe() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withDescribe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) WithDescribeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withDescribeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) WithParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeIcebergTables) WithParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withParametersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/data-sources/iceberg_tables snowflake_iceberg_tables} Data Source.
func NewDataSnowflakeIcebergTables(scope constructs.Construct, id *string, config *DataSnowflakeIcebergTablesConfig) DataSnowflakeIcebergTables {
	_init_.Initialize()

	if err := validateNewDataSnowflakeIcebergTablesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeIcebergTables{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeIcebergTables.DataSnowflakeIcebergTables",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.19.0/docs/data-sources/iceberg_tables snowflake_iceberg_tables} Data Source.
func NewDataSnowflakeIcebergTables_Override(d DataSnowflakeIcebergTables, scope constructs.Construct, id *string, config *DataSnowflakeIcebergTablesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeIcebergTables.DataSnowflakeIcebergTables",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTables)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTables)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTables)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTables)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTables)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTables)SetLike(val *string) {
	if err := j.validateSetLikeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"like",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTables)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTables)SetStartsWith(val *string) {
	if err := j.validateSetStartsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startsWith",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTables)SetWithDescribe(val interface{}) {
	if err := j.validateSetWithDescribeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withDescribe",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeIcebergTables)SetWithParameters(val interface{}) {
	if err := j.validateSetWithParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withParameters",
		val,
	)
}

// Generates CDKTN code for importing a DataSnowflakeIcebergTables resource upon running "cdktn plan <stack-name>".
func DataSnowflakeIcebergTables_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataSnowflakeIcebergTables_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.dataSnowflakeIcebergTables.DataSnowflakeIcebergTables",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DataSnowflakeIcebergTables_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataSnowflakeIcebergTables_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.dataSnowflakeIcebergTables.DataSnowflakeIcebergTables",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataSnowflakeIcebergTables_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataSnowflakeIcebergTables_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.dataSnowflakeIcebergTables.DataSnowflakeIcebergTables",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataSnowflakeIcebergTables_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataSnowflakeIcebergTables_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.dataSnowflakeIcebergTables.DataSnowflakeIcebergTables",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataSnowflakeIcebergTables_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-snowflake.dataSnowflakeIcebergTables.DataSnowflakeIcebergTables",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeIcebergTables) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeIcebergTables) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeIcebergTables) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeIcebergTables) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeIcebergTables) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeIcebergTables) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeIcebergTables) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeIcebergTables) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeIcebergTables) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeIcebergTables) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) PutIn(value *DataSnowflakeIcebergTablesIn) {
	if err := d.validatePutInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIn",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) PutLimit(value *DataSnowflakeIcebergTablesLimit) {
	if err := d.validatePutLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLimit",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) ResetIn() {
	_jsii_.InvokeVoid(
		d,
		"resetIn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) ResetLike() {
	_jsii_.InvokeVoid(
		d,
		"resetLike",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) ResetLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) ResetStartsWith() {
	_jsii_.InvokeVoid(
		d,
		"resetStartsWith",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) ResetWithDescribe() {
	_jsii_.InvokeVoid(
		d,
		"resetWithDescribe",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) ResetWithParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetWithParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeIcebergTables) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

