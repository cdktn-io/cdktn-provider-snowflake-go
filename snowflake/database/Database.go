// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package database

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/database/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/database snowflake_database}.
type Database interface {
	cdktn.TerraformResource
	Catalog() *string
	SetCatalog(val *string)
	CatalogInput() *string
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
	DataRetentionTimeInDays() *float64
	SetDataRetentionTimeInDays(val *float64)
	DataRetentionTimeInDaysInput() *float64
	DefaultDdlCollation() *string
	SetDefaultDdlCollation(val *string)
	DefaultDdlCollationInput() *string
	DefaultNotebookComputePoolCpu() *string
	SetDefaultNotebookComputePoolCpu(val *string)
	DefaultNotebookComputePoolCpuInput() *string
	DefaultNotebookComputePoolGpu() *string
	SetDefaultNotebookComputePoolGpu(val *string)
	DefaultNotebookComputePoolGpuInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DropPublicSchemaOnCreation() interface{}
	SetDropPublicSchemaOnCreation(val interface{})
	DropPublicSchemaOnCreationInput() interface{}
	EnableConsoleOutput() interface{}
	SetEnableConsoleOutput(val interface{})
	EnableConsoleOutputInput() interface{}
	ExternalVolume() *string
	SetExternalVolume(val *string)
	ExternalVolumeInput() *string
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
	IsTransient() interface{}
	SetIsTransient(val interface{})
	IsTransientInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogEventLevel() *string
	SetLogEventLevel(val *string)
	LogEventLevelInput() *string
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	MaxDataExtensionTimeInDays() *float64
	SetMaxDataExtensionTimeInDays(val *float64)
	MaxDataExtensionTimeInDaysInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QuotedIdentifiersIgnoreCase() interface{}
	SetQuotedIdentifiersIgnoreCase(val interface{})
	QuotedIdentifiersIgnoreCaseInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ReplaceInvalidCharacters() interface{}
	SetReplaceInvalidCharacters(val interface{})
	ReplaceInvalidCharactersInput() interface{}
	Replication() DatabaseReplicationOutputReference
	ReplicationInput() *DatabaseReplication
	StorageSerializationPolicy() *string
	SetStorageSerializationPolicy(val *string)
	StorageSerializationPolicyInput() *string
	SuspendTaskAfterNumFailures() *float64
	SetSuspendTaskAfterNumFailures(val *float64)
	SuspendTaskAfterNumFailuresInput() *float64
	TaskAutoRetryAttempts() *float64
	SetTaskAutoRetryAttempts(val *float64)
	TaskAutoRetryAttemptsInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DatabaseTimeoutsOutputReference
	TimeoutsInput() interface{}
	TraceLevel() *string
	SetTraceLevel(val *string)
	TraceLevelInput() *string
	UserTaskManagedInitialWarehouseSize() *string
	SetUserTaskManagedInitialWarehouseSize(val *string)
	UserTaskManagedInitialWarehouseSizeInput() *string
	UserTaskMinimumTriggerIntervalInSeconds() *float64
	SetUserTaskMinimumTriggerIntervalInSeconds(val *float64)
	UserTaskMinimumTriggerIntervalInSecondsInput() *float64
	UserTaskTimeoutMs() *float64
	SetUserTaskTimeoutMs(val *float64)
	UserTaskTimeoutMsInput() *float64
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
	PutReplication(value *DatabaseReplication)
	PutTimeouts(value *DatabaseTimeouts)
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
	ResetCatalog()
	ResetComment()
	ResetDataRetentionTimeInDays()
	ResetDefaultDdlCollation()
	ResetDefaultNotebookComputePoolCpu()
	ResetDefaultNotebookComputePoolGpu()
	ResetDropPublicSchemaOnCreation()
	ResetEnableConsoleOutput()
	ResetExternalVolume()
	ResetId()
	ResetIsTransient()
	ResetLogEventLevel()
	ResetLogLevel()
	ResetMaxDataExtensionTimeInDays()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQuotedIdentifiersIgnoreCase()
	ResetReplaceInvalidCharacters()
	ResetReplication()
	ResetStorageSerializationPolicy()
	ResetSuspendTaskAfterNumFailures()
	ResetTaskAutoRetryAttempts()
	ResetTimeouts()
	ResetTraceLevel()
	ResetUserTaskManagedInitialWarehouseSize()
	ResetUserTaskMinimumTriggerIntervalInSeconds()
	ResetUserTaskTimeoutMs()
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

// The jsii proxy struct for Database
type jsiiProxy_Database struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Database) Catalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) CatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DataRetentionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetentionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DataRetentionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetentionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DefaultDdlCollation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDdlCollation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DefaultDdlCollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDdlCollationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DefaultNotebookComputePoolCpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DefaultNotebookComputePoolCpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolCpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DefaultNotebookComputePoolGpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolGpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DefaultNotebookComputePoolGpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolGpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DropPublicSchemaOnCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropPublicSchemaOnCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) DropPublicSchemaOnCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropPublicSchemaOnCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) EnableConsoleOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConsoleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) EnableConsoleOutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConsoleOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) ExternalVolume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) ExternalVolumeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) IsTransient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTransient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) IsTransientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTransientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) LogEventLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEventLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) LogEventLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEventLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) MaxDataExtensionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataExtensionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) MaxDataExtensionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataExtensionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) QuotedIdentifiersIgnoreCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) QuotedIdentifiersIgnoreCaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) ReplaceInvalidCharacters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) ReplaceInvalidCharactersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceInvalidCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Replication() DatabaseReplicationOutputReference {
	var returns DatabaseReplicationOutputReference
	_jsii_.Get(
		j,
		"replication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) ReplicationInput() *DatabaseReplication {
	var returns *DatabaseReplication
	_jsii_.Get(
		j,
		"replicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) StorageSerializationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSerializationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) StorageSerializationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSerializationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) SuspendTaskAfterNumFailures() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) SuspendTaskAfterNumFailuresInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TaskAutoRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAutoRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TaskAutoRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAutoRetryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) Timeouts() DatabaseTimeoutsOutputReference {
	var returns DatabaseTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TraceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) TraceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"traceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) UserTaskManagedInitialWarehouseSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) UserTaskManagedInitialWarehouseSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) UserTaskMinimumTriggerIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) UserTaskMinimumTriggerIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) UserTaskTimeoutMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Database) UserTaskTimeoutMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userTaskTimeoutMsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/database snowflake_database} Resource.
func NewDatabase(scope constructs.Construct, id *string, config *DatabaseConfig) Database {
	_init_.Initialize()

	if err := validateNewDatabaseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Database{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.database.Database",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/database snowflake_database} Resource.
func NewDatabase_Override(d Database, scope constructs.Construct, id *string, config *DatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.database.Database",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_Database)SetCatalog(val *string) {
	if err := j.validateSetCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalog",
		val,
	)
}

func (j *jsiiProxy_Database)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_Database)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Database)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Database)SetDataRetentionTimeInDays(val *float64) {
	if err := j.validateSetDataRetentionTimeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataRetentionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_Database)SetDefaultDdlCollation(val *string) {
	if err := j.validateSetDefaultDdlCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDdlCollation",
		val,
	)
}

func (j *jsiiProxy_Database)SetDefaultNotebookComputePoolCpu(val *string) {
	if err := j.validateSetDefaultNotebookComputePoolCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNotebookComputePoolCpu",
		val,
	)
}

func (j *jsiiProxy_Database)SetDefaultNotebookComputePoolGpu(val *string) {
	if err := j.validateSetDefaultNotebookComputePoolGpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultNotebookComputePoolGpu",
		val,
	)
}

func (j *jsiiProxy_Database)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Database)SetDropPublicSchemaOnCreation(val interface{}) {
	if err := j.validateSetDropPublicSchemaOnCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dropPublicSchemaOnCreation",
		val,
	)
}

func (j *jsiiProxy_Database)SetEnableConsoleOutput(val interface{}) {
	if err := j.validateSetEnableConsoleOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConsoleOutput",
		val,
	)
}

func (j *jsiiProxy_Database)SetExternalVolume(val *string) {
	if err := j.validateSetExternalVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalVolume",
		val,
	)
}

func (j *jsiiProxy_Database)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Database)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Database)SetIsTransient(val interface{}) {
	if err := j.validateSetIsTransientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isTransient",
		val,
	)
}

func (j *jsiiProxy_Database)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Database)SetLogEventLevel(val *string) {
	if err := j.validateSetLogEventLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logEventLevel",
		val,
	)
}

func (j *jsiiProxy_Database)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_Database)SetMaxDataExtensionTimeInDays(val *float64) {
	if err := j.validateSetMaxDataExtensionTimeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDataExtensionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_Database)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Database)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Database)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Database)SetQuotedIdentifiersIgnoreCase(val interface{}) {
	if err := j.validateSetQuotedIdentifiersIgnoreCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quotedIdentifiersIgnoreCase",
		val,
	)
}

func (j *jsiiProxy_Database)SetReplaceInvalidCharacters(val interface{}) {
	if err := j.validateSetReplaceInvalidCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceInvalidCharacters",
		val,
	)
}

func (j *jsiiProxy_Database)SetStorageSerializationPolicy(val *string) {
	if err := j.validateSetStorageSerializationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSerializationPolicy",
		val,
	)
}

func (j *jsiiProxy_Database)SetSuspendTaskAfterNumFailures(val *float64) {
	if err := j.validateSetSuspendTaskAfterNumFailuresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendTaskAfterNumFailures",
		val,
	)
}

func (j *jsiiProxy_Database)SetTaskAutoRetryAttempts(val *float64) {
	if err := j.validateSetTaskAutoRetryAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskAutoRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_Database)SetTraceLevel(val *string) {
	if err := j.validateSetTraceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceLevel",
		val,
	)
}

func (j *jsiiProxy_Database)SetUserTaskManagedInitialWarehouseSize(val *string) {
	if err := j.validateSetUserTaskManagedInitialWarehouseSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskManagedInitialWarehouseSize",
		val,
	)
}

func (j *jsiiProxy_Database)SetUserTaskMinimumTriggerIntervalInSeconds(val *float64) {
	if err := j.validateSetUserTaskMinimumTriggerIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_Database)SetUserTaskTimeoutMs(val *float64) {
	if err := j.validateSetUserTaskTimeoutMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTaskTimeoutMs",
		val,
	)
}

// Generates CDKTN code for importing a Database resource upon running "cdktn plan <stack-name>".
func Database_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDatabase_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.database.Database",
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
func Database_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.database.Database",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Database_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabase_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.database.Database",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Database_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabase_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.database.Database",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Database_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-snowflake.database.Database",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_Database) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_Database) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_Database) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_Database) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_Database) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_Database) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_Database) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_Database) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_Database) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_Database) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_Database) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_Database) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_Database) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_Database) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := d.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_Database) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_Database) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_Database) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_Database) PutReplication(value *DatabaseReplication) {
	if err := d.validatePutReplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReplication",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_Database) PutTimeouts(value *DatabaseTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_Database) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_Database) ResetCatalog() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalog",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetComment() {
	_jsii_.InvokeVoid(
		d,
		"resetComment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetDataRetentionTimeInDays() {
	_jsii_.InvokeVoid(
		d,
		"resetDataRetentionTimeInDays",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetDefaultDdlCollation() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultDdlCollation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetDefaultNotebookComputePoolCpu() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultNotebookComputePoolCpu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetDefaultNotebookComputePoolGpu() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultNotebookComputePoolGpu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetDropPublicSchemaOnCreation() {
	_jsii_.InvokeVoid(
		d,
		"resetDropPublicSchemaOnCreation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetEnableConsoleOutput() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableConsoleOutput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetExternalVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetIsTransient() {
	_jsii_.InvokeVoid(
		d,
		"resetIsTransient",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetLogEventLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetLogEventLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetLogLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetMaxDataExtensionTimeInDays() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxDataExtensionTimeInDays",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetQuotedIdentifiersIgnoreCase() {
	_jsii_.InvokeVoid(
		d,
		"resetQuotedIdentifiersIgnoreCase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetReplaceInvalidCharacters() {
	_jsii_.InvokeVoid(
		d,
		"resetReplaceInvalidCharacters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetReplication() {
	_jsii_.InvokeVoid(
		d,
		"resetReplication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetStorageSerializationPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageSerializationPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetSuspendTaskAfterNumFailures() {
	_jsii_.InvokeVoid(
		d,
		"resetSuspendTaskAfterNumFailures",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetTaskAutoRetryAttempts() {
	_jsii_.InvokeVoid(
		d,
		"resetTaskAutoRetryAttempts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetTraceLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetTraceLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetUserTaskManagedInitialWarehouseSize() {
	_jsii_.InvokeVoid(
		d,
		"resetUserTaskManagedInitialWarehouseSize",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetUserTaskMinimumTriggerIntervalInSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetUserTaskMinimumTriggerIntervalInSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) ResetUserTaskTimeoutMs() {
	_jsii_.InvokeVoid(
		d,
		"resetUserTaskTimeoutMs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Database) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Database) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

