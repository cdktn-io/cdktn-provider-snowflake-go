// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package warehouseinteractive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/warehouseinteractive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/warehouse_interactive snowflake_warehouse_interactive}.
type WarehouseInteractive interface {
	cdktn.TerraformResource
	AutoResume() *string
	SetAutoResume(val *string)
	AutoResumeInput() *string
	AutoSuspend() *float64
	SetAutoSuspend(val *float64)
	AutoSuspendInput() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
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
	FallbackWarehouse() *string
	SetFallbackWarehouse(val *string)
	FallbackWarehouseInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullyQualifiedName() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InitiallySuspended() interface{}
	SetInitiallySuspended(val interface{})
	InitiallySuspendedInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaxClusterCount() *float64
	SetMaxClusterCount(val *float64)
	MaxClusterCountInput() *float64
	MaxConcurrencyLevel() *float64
	SetMaxConcurrencyLevel(val *float64)
	MaxConcurrencyLevelInput() *float64
	MinClusterCount() *float64
	SetMinClusterCount(val *float64)
	MinClusterCountInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Parameters() WarehouseInteractiveParametersList
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResourceMonitor() *string
	SetResourceMonitor(val *string)
	ResourceMonitorInput() *string
	ShowOutput() WarehouseInteractiveShowOutputList
	StatementQueuedTimeoutInSeconds() *float64
	SetStatementQueuedTimeoutInSeconds(val *float64)
	StatementQueuedTimeoutInSecondsInput() *float64
	StatementTimeoutInSeconds() *float64
	SetStatementTimeoutInSeconds(val *float64)
	StatementTimeoutInSecondsInput() *float64
	Tables() *[]*string
	SetTables(val *[]*string)
	TablesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() WarehouseInteractiveTimeoutsOutputReference
	TimeoutsInput() interface{}
	WarehouseSize() *string
	SetWarehouseSize(val *string)
	WarehouseSizeInput() *string
	WarehouseType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *WarehouseInteractiveTimeouts)
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
	ResetAutoResume()
	ResetAutoSuspend()
	ResetComment()
	ResetFallbackWarehouse()
	ResetId()
	ResetInitiallySuspended()
	ResetMaxClusterCount()
	ResetMaxConcurrencyLevel()
	ResetMinClusterCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceMonitor()
	ResetStatementQueuedTimeoutInSeconds()
	ResetStatementTimeoutInSeconds()
	ResetTables()
	ResetTimeouts()
	ResetWarehouseSize()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for WarehouseInteractive
type jsiiProxy_WarehouseInteractive struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_WarehouseInteractive) AutoResume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoResume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) AutoResumeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoResumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) AutoSuspend() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoSuspend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) AutoSuspendInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoSuspendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) FallbackWarehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fallbackWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) FallbackWarehouseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fallbackWarehouseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) InitiallySuspended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initiallySuspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) InitiallySuspendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initiallySuspendedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) MaxClusterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClusterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) MaxClusterCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClusterCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) MaxConcurrencyLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) MaxConcurrencyLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) MinClusterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minClusterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) MinClusterCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minClusterCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) Parameters() WarehouseInteractiveParametersList {
	var returns WarehouseInteractiveParametersList
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) ResourceMonitor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) ResourceMonitorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) ShowOutput() WarehouseInteractiveShowOutputList {
	var returns WarehouseInteractiveShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) StatementQueuedTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) StatementQueuedTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) StatementTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) StatementTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) Tables() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) TablesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) Timeouts() WarehouseInteractiveTimeoutsOutputReference {
	var returns WarehouseInteractiveTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) WarehouseSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) WarehouseSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarehouseInteractive) WarehouseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/warehouse_interactive snowflake_warehouse_interactive} Resource.
func NewWarehouseInteractive(scope constructs.Construct, id *string, config *WarehouseInteractiveConfig) WarehouseInteractive {
	_init_.Initialize()

	if err := validateNewWarehouseInteractiveParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WarehouseInteractive{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.warehouseInteractive.WarehouseInteractive",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/warehouse_interactive snowflake_warehouse_interactive} Resource.
func NewWarehouseInteractive_Override(w WarehouseInteractive, scope constructs.Construct, id *string, config *WarehouseInteractiveConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.warehouseInteractive.WarehouseInteractive",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetAutoResume(val *string) {
	if err := j.validateSetAutoResumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoResume",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetAutoSuspend(val *float64) {
	if err := j.validateSetAutoSuspendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSuspend",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetFallbackWarehouse(val *string) {
	if err := j.validateSetFallbackWarehouseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fallbackWarehouse",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetInitiallySuspended(val interface{}) {
	if err := j.validateSetInitiallySuspendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initiallySuspended",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetMaxClusterCount(val *float64) {
	if err := j.validateSetMaxClusterCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxClusterCount",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetMaxConcurrencyLevel(val *float64) {
	if err := j.validateSetMaxConcurrencyLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrencyLevel",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetMinClusterCount(val *float64) {
	if err := j.validateSetMinClusterCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minClusterCount",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetResourceMonitor(val *string) {
	if err := j.validateSetResourceMonitorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceMonitor",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetStatementQueuedTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementQueuedTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementQueuedTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetStatementTimeoutInSeconds(val *float64) {
	if err := j.validateSetStatementTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetTables(val *[]*string) {
	if err := j.validateSetTablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tables",
		val,
	)
}

func (j *jsiiProxy_WarehouseInteractive)SetWarehouseSize(val *string) {
	if err := j.validateSetWarehouseSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warehouseSize",
		val,
	)
}

// Generates CDKTN code for importing a WarehouseInteractive resource upon running "cdktn plan <stack-name>".
func WarehouseInteractive_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateWarehouseInteractive_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.warehouseInteractive.WarehouseInteractive",
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
func WarehouseInteractive_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWarehouseInteractive_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.warehouseInteractive.WarehouseInteractive",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WarehouseInteractive_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWarehouseInteractive_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.warehouseInteractive.WarehouseInteractive",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WarehouseInteractive_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWarehouseInteractive_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.warehouseInteractive.WarehouseInteractive",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WarehouseInteractive_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-snowflake.warehouseInteractive.WarehouseInteractive",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WarehouseInteractive) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WarehouseInteractive) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WarehouseInteractive) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WarehouseInteractive) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := w.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WarehouseInteractive) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WarehouseInteractive) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WarehouseInteractive) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WarehouseInteractive) PutTimeouts(value *WarehouseInteractiveTimeouts) {
	if err := w.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WarehouseInteractive) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := w.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetAutoResume() {
	_jsii_.InvokeVoid(
		w,
		"resetAutoResume",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetAutoSuspend() {
	_jsii_.InvokeVoid(
		w,
		"resetAutoSuspend",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetComment() {
	_jsii_.InvokeVoid(
		w,
		"resetComment",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetFallbackWarehouse() {
	_jsii_.InvokeVoid(
		w,
		"resetFallbackWarehouse",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetInitiallySuspended() {
	_jsii_.InvokeVoid(
		w,
		"resetInitiallySuspended",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetMaxClusterCount() {
	_jsii_.InvokeVoid(
		w,
		"resetMaxClusterCount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetMaxConcurrencyLevel() {
	_jsii_.InvokeVoid(
		w,
		"resetMaxConcurrencyLevel",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetMinClusterCount() {
	_jsii_.InvokeVoid(
		w,
		"resetMinClusterCount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetResourceMonitor() {
	_jsii_.InvokeVoid(
		w,
		"resetResourceMonitor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetStatementQueuedTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		w,
		"resetStatementQueuedTimeoutInSeconds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetStatementTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		w,
		"resetStatementTimeoutInSeconds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetTables() {
	_jsii_.InvokeVoid(
		w,
		"resetTables",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) ResetWarehouseSize() {
	_jsii_.InvokeVoid(
		w,
		"resetWarehouseSize",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarehouseInteractive) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarehouseInteractive) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		w,
		"with",
		args,
		&returns,
	)

	return returns
}

