// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakehybridtables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/datasnowflakehybridtables/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/data-sources/hybrid_tables snowflake_hybrid_tables}.
type DataSnowflakeHybridTables interface {
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
	HybridTables() DataSnowflakeHybridTablesHybridTablesList
	Id() *string
	SetId(val *string)
	IdInput() *string
	In() DataSnowflakeHybridTablesInOutputReference
	InInput() *DataSnowflakeHybridTablesIn
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Like() *string
	SetLike(val *string)
	LikeInput() *string
	Limit() DataSnowflakeHybridTablesLimitOutputReference
	LimitInput() *DataSnowflakeHybridTablesLimit
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
	WithIndexes() interface{}
	SetWithIndexes(val interface{})
	WithIndexesInput() interface{}
	WithKeys() interface{}
	SetWithKeys(val interface{})
	WithKeysInput() interface{}
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
	PutIn(value *DataSnowflakeHybridTablesIn)
	PutLimit(value *DataSnowflakeHybridTablesLimit)
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
	ResetWithIndexes()
	ResetWithKeys()
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

// The jsii proxy struct for DataSnowflakeHybridTables
type jsiiProxy_DataSnowflakeHybridTables struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataSnowflakeHybridTables) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) HybridTables() DataSnowflakeHybridTablesHybridTablesList {
	var returns DataSnowflakeHybridTablesHybridTablesList
	_jsii_.Get(
		j,
		"hybridTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) In() DataSnowflakeHybridTablesInOutputReference {
	var returns DataSnowflakeHybridTablesInOutputReference
	_jsii_.Get(
		j,
		"in",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) InInput() *DataSnowflakeHybridTablesIn {
	var returns *DataSnowflakeHybridTablesIn
	_jsii_.Get(
		j,
		"inInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) Like() *string {
	var returns *string
	_jsii_.Get(
		j,
		"like",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) LikeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"likeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) Limit() DataSnowflakeHybridTablesLimitOutputReference {
	var returns DataSnowflakeHybridTablesLimitOutputReference
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) LimitInput() *DataSnowflakeHybridTablesLimit {
	var returns *DataSnowflakeHybridTablesLimit
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) StartsWith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) StartsWithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) WithDescribe() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withDescribe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) WithDescribeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withDescribeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) WithIndexes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) WithIndexesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withIndexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) WithKeys() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) WithKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) WithParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeHybridTables) WithParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withParametersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/data-sources/hybrid_tables snowflake_hybrid_tables} Data Source.
func NewDataSnowflakeHybridTables(scope constructs.Construct, id *string, config *DataSnowflakeHybridTablesConfig) DataSnowflakeHybridTables {
	_init_.Initialize()

	if err := validateNewDataSnowflakeHybridTablesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeHybridTables{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeHybridTables.DataSnowflakeHybridTables",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.22.0/docs/data-sources/hybrid_tables snowflake_hybrid_tables} Data Source.
func NewDataSnowflakeHybridTables_Override(d DataSnowflakeHybridTables, scope constructs.Construct, id *string, config *DataSnowflakeHybridTablesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeHybridTables.DataSnowflakeHybridTables",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeHybridTables)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeHybridTables)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeHybridTables)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeHybridTables)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeHybridTables)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeHybridTables)SetLike(val *string) {
	if err := j.validateSetLikeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"like",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeHybridTables)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeHybridTables)SetStartsWith(val *string) {
	if err := j.validateSetStartsWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startsWith",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeHybridTables)SetWithDescribe(val interface{}) {
	if err := j.validateSetWithDescribeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withDescribe",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeHybridTables)SetWithIndexes(val interface{}) {
	if err := j.validateSetWithIndexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withIndexes",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeHybridTables)SetWithKeys(val interface{}) {
	if err := j.validateSetWithKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withKeys",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeHybridTables)SetWithParameters(val interface{}) {
	if err := j.validateSetWithParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withParameters",
		val,
	)
}

// Generates CDKTN code for importing a DataSnowflakeHybridTables resource upon running "cdktn plan <stack-name>".
func DataSnowflakeHybridTables_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataSnowflakeHybridTables_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.dataSnowflakeHybridTables.DataSnowflakeHybridTables",
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
func DataSnowflakeHybridTables_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataSnowflakeHybridTables_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.dataSnowflakeHybridTables.DataSnowflakeHybridTables",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataSnowflakeHybridTables_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataSnowflakeHybridTables_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.dataSnowflakeHybridTables.DataSnowflakeHybridTables",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataSnowflakeHybridTables_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataSnowflakeHybridTables_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.dataSnowflakeHybridTables.DataSnowflakeHybridTables",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataSnowflakeHybridTables_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-snowflake.dataSnowflakeHybridTables.DataSnowflakeHybridTables",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataSnowflakeHybridTables) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeHybridTables) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeHybridTables) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeHybridTables) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeHybridTables) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeHybridTables) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeHybridTables) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeHybridTables) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeHybridTables) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeHybridTables) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeHybridTables) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) PutIn(value *DataSnowflakeHybridTablesIn) {
	if err := d.validatePutInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIn",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) PutLimit(value *DataSnowflakeHybridTablesLimit) {
	if err := d.validatePutLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLimit",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ResetIn() {
	_jsii_.InvokeVoid(
		d,
		"resetIn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ResetLike() {
	_jsii_.InvokeVoid(
		d,
		"resetLike",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ResetLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ResetStartsWith() {
	_jsii_.InvokeVoid(
		d,
		"resetStartsWith",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ResetWithDescribe() {
	_jsii_.InvokeVoid(
		d,
		"resetWithDescribe",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ResetWithIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetWithIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ResetWithKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetWithKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ResetWithParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetWithParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataSnowflakeHybridTables) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeHybridTables) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeHybridTables) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeHybridTables) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

