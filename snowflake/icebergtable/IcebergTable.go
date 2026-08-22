// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package icebergtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/icebergtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table snowflake_iceberg_table}.
type IcebergTable interface {
	cdktn.TerraformResource
	AggregationPolicy() IcebergTableAggregationPolicyOutputReference
	AggregationPolicyInput() *IcebergTableAggregationPolicy
	BaseLocation() *string
	SetBaseLocation(val *string)
	BaseLocationInput() *string
	Catalog() *string
	SetCatalog(val *string)
	CatalogInput() *string
	CatalogSync() *string
	SetCatalogSync(val *string)
	CatalogSyncInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ChangeTracking() *string
	SetChangeTracking(val *string)
	ChangeTrackingInput() *string
	CheckConstraint() IcebergTableCheckConstraintList
	CheckConstraintInput() interface{}
	ClusterBy() *[]*string
	SetClusterBy(val *[]*string)
	ClusterByInput() *[]*string
	Column() IcebergTableColumnList
	ColumnInput() interface{}
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
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	DataRetentionTimeInDays() *float64
	SetDataRetentionTimeInDays(val *float64)
	DataRetentionTimeInDaysInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DescribeOutput() IcebergTableDescribeOutputList
	EnableDataCompaction() interface{}
	SetEnableDataCompaction(val interface{})
	EnableDataCompactionInput() interface{}
	EnableIcebergMergeOnRead() interface{}
	SetEnableIcebergMergeOnRead(val interface{})
	EnableIcebergMergeOnReadInput() interface{}
	ErrorLogging() *string
	SetErrorLogging(val *string)
	ErrorLoggingInput() *string
	ExternalVolume() *string
	SetExternalVolume(val *string)
	ExternalVolumeInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	ForeignKeyConstraint() IcebergTableForeignKeyConstraintList
	ForeignKeyConstraintInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullyQualifiedName() *string
	IcebergVersion() *float64
	SetIcebergVersion(val *float64)
	IcebergVersionInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaxDataExtensionTimeInDays() *float64
	SetMaxDataExtensionTimeInDays(val *float64)
	MaxDataExtensionTimeInDaysInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Parameters() IcebergTableParametersList
	PartitionBy() IcebergTablePartitionByList
	PartitionByInput() interface{}
	PathLayout() *string
	SetPathLayout(val *string)
	PathLayoutInput() *string
	PrimaryKeyConstraint() IcebergTablePrimaryKeyConstraintOutputReference
	PrimaryKeyConstraintInput() *IcebergTablePrimaryKeyConstraint
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
	RowAccessPolicy() IcebergTableRowAccessPolicyOutputReference
	RowAccessPolicyInput() *IcebergTableRowAccessPolicy
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	ShowOutput() IcebergTableShowOutputList
	StorageSerializationPolicy() *string
	SetStorageSerializationPolicy(val *string)
	StorageSerializationPolicyInput() *string
	TargetFileSize() *string
	SetTargetFileSize(val *string)
	TargetFileSizeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() IcebergTableTimeoutsOutputReference
	TimeoutsInput() interface{}
	UniqueConstraint() IcebergTableUniqueConstraintList
	UniqueConstraintInput() interface{}
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
	PutAggregationPolicy(value *IcebergTableAggregationPolicy)
	PutCheckConstraint(value interface{})
	PutColumn(value interface{})
	PutForeignKeyConstraint(value interface{})
	PutPartitionBy(value interface{})
	PutPrimaryKeyConstraint(value *IcebergTablePrimaryKeyConstraint)
	PutRowAccessPolicy(value *IcebergTableRowAccessPolicy)
	PutTimeouts(value *IcebergTableTimeouts)
	PutUniqueConstraint(value interface{})
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
	ResetAggregationPolicy()
	ResetBaseLocation()
	ResetCatalog()
	ResetCatalogSync()
	ResetChangeTracking()
	ResetCheckConstraint()
	ResetClusterBy()
	ResetComment()
	ResetDataRetentionTimeInDays()
	ResetEnableDataCompaction()
	ResetEnableIcebergMergeOnRead()
	ResetErrorLogging()
	ResetExternalVolume()
	ResetForeignKeyConstraint()
	ResetIcebergVersion()
	ResetId()
	ResetMaxDataExtensionTimeInDays()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartitionBy()
	ResetPathLayout()
	ResetPrimaryKeyConstraint()
	ResetRowAccessPolicy()
	ResetStorageSerializationPolicy()
	ResetTargetFileSize()
	ResetTimeouts()
	ResetUniqueConstraint()
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

// The jsii proxy struct for IcebergTable
type jsiiProxy_IcebergTable struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_IcebergTable) AggregationPolicy() IcebergTableAggregationPolicyOutputReference {
	var returns IcebergTableAggregationPolicyOutputReference
	_jsii_.Get(
		j,
		"aggregationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) AggregationPolicyInput() *IcebergTableAggregationPolicy {
	var returns *IcebergTableAggregationPolicy
	_jsii_.Get(
		j,
		"aggregationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) BaseLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) BaseLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Catalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) CatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) CatalogSync() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) CatalogSyncInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogSyncInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ChangeTracking() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeTracking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ChangeTrackingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"changeTrackingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) CheckConstraint() IcebergTableCheckConstraintList {
	var returns IcebergTableCheckConstraintList
	_jsii_.Get(
		j,
		"checkConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) CheckConstraintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ClusterBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ClusterByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Column() IcebergTableColumnList {
	var returns IcebergTableColumnList
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) DataRetentionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetentionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) DataRetentionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetentionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) DescribeOutput() IcebergTableDescribeOutputList {
	var returns IcebergTableDescribeOutputList
	_jsii_.Get(
		j,
		"describeOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) EnableDataCompaction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDataCompaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) EnableDataCompactionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDataCompactionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) EnableIcebergMergeOnRead() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIcebergMergeOnRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) EnableIcebergMergeOnReadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIcebergMergeOnReadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ErrorLogging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ErrorLoggingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ExternalVolume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ExternalVolumeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ForeignKeyConstraint() IcebergTableForeignKeyConstraintList {
	var returns IcebergTableForeignKeyConstraintList
	_jsii_.Get(
		j,
		"foreignKeyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ForeignKeyConstraintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"foreignKeyConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) IcebergVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icebergVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) IcebergVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"icebergVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) MaxDataExtensionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataExtensionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) MaxDataExtensionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataExtensionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Parameters() IcebergTableParametersList {
	var returns IcebergTableParametersList
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) PartitionBy() IcebergTablePartitionByList {
	var returns IcebergTablePartitionByList
	_jsii_.Get(
		j,
		"partitionBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) PartitionByInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) PathLayout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) PathLayoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathLayoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) PrimaryKeyConstraint() IcebergTablePrimaryKeyConstraintOutputReference {
	var returns IcebergTablePrimaryKeyConstraintOutputReference
	_jsii_.Get(
		j,
		"primaryKeyConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) PrimaryKeyConstraintInput() *IcebergTablePrimaryKeyConstraint {
	var returns *IcebergTablePrimaryKeyConstraint
	_jsii_.Get(
		j,
		"primaryKeyConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) RowAccessPolicy() IcebergTableRowAccessPolicyOutputReference {
	var returns IcebergTableRowAccessPolicyOutputReference
	_jsii_.Get(
		j,
		"rowAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) RowAccessPolicyInput() *IcebergTableRowAccessPolicy {
	var returns *IcebergTableRowAccessPolicy
	_jsii_.Get(
		j,
		"rowAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) ShowOutput() IcebergTableShowOutputList {
	var returns IcebergTableShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) StorageSerializationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSerializationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) StorageSerializationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageSerializationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) TargetFileSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) TargetFileSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) Timeouts() IcebergTableTimeoutsOutputReference {
	var returns IcebergTableTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) UniqueConstraint() IcebergTableUniqueConstraintList {
	var returns IcebergTableUniqueConstraintList
	_jsii_.Get(
		j,
		"uniqueConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IcebergTable) UniqueConstraintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uniqueConstraintInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table snowflake_iceberg_table} Resource.
func NewIcebergTable(scope constructs.Construct, id *string, config *IcebergTableConfig) IcebergTable {
	_init_.Initialize()

	if err := validateNewIcebergTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IcebergTable{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.icebergTable.IcebergTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/iceberg_table snowflake_iceberg_table} Resource.
func NewIcebergTable_Override(i IcebergTable, scope constructs.Construct, id *string, config *IcebergTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.icebergTable.IcebergTable",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IcebergTable)SetBaseLocation(val *string) {
	if err := j.validateSetBaseLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseLocation",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetCatalog(val *string) {
	if err := j.validateSetCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalog",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetCatalogSync(val *string) {
	if err := j.validateSetCatalogSyncParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogSync",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetChangeTracking(val *string) {
	if err := j.validateSetChangeTrackingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeTracking",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetClusterBy(val *[]*string) {
	if err := j.validateSetClusterByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterBy",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetDataRetentionTimeInDays(val *float64) {
	if err := j.validateSetDataRetentionTimeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataRetentionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetEnableDataCompaction(val interface{}) {
	if err := j.validateSetEnableDataCompactionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDataCompaction",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetEnableIcebergMergeOnRead(val interface{}) {
	if err := j.validateSetEnableIcebergMergeOnReadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIcebergMergeOnRead",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetErrorLogging(val *string) {
	if err := j.validateSetErrorLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorLogging",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetExternalVolume(val *string) {
	if err := j.validateSetExternalVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalVolume",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetIcebergVersion(val *float64) {
	if err := j.validateSetIcebergVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icebergVersion",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetMaxDataExtensionTimeInDays(val *float64) {
	if err := j.validateSetMaxDataExtensionTimeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDataExtensionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetPathLayout(val *string) {
	if err := j.validateSetPathLayoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathLayout",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetStorageSerializationPolicy(val *string) {
	if err := j.validateSetStorageSerializationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSerializationPolicy",
		val,
	)
}

func (j *jsiiProxy_IcebergTable)SetTargetFileSize(val *string) {
	if err := j.validateSetTargetFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetFileSize",
		val,
	)
}

// Generates CDKTN code for importing a IcebergTable resource upon running "cdktn plan <stack-name>".
func IcebergTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateIcebergTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.icebergTable.IcebergTable",
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
func IcebergTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIcebergTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.icebergTable.IcebergTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IcebergTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIcebergTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.icebergTable.IcebergTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IcebergTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIcebergTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.icebergTable.IcebergTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IcebergTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-snowflake.icebergTable.IcebergTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IcebergTable) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IcebergTable) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IcebergTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IcebergTable) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := i.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IcebergTable) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IcebergTable) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IcebergTable) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IcebergTable) PutAggregationPolicy(value *IcebergTableAggregationPolicy) {
	if err := i.validatePutAggregationPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAggregationPolicy",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IcebergTable) PutCheckConstraint(value interface{}) {
	if err := i.validatePutCheckConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCheckConstraint",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IcebergTable) PutColumn(value interface{}) {
	if err := i.validatePutColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putColumn",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IcebergTable) PutForeignKeyConstraint(value interface{}) {
	if err := i.validatePutForeignKeyConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putForeignKeyConstraint",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IcebergTable) PutPartitionBy(value interface{}) {
	if err := i.validatePutPartitionByParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPartitionBy",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IcebergTable) PutPrimaryKeyConstraint(value *IcebergTablePrimaryKeyConstraint) {
	if err := i.validatePutPrimaryKeyConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPrimaryKeyConstraint",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IcebergTable) PutRowAccessPolicy(value *IcebergTableRowAccessPolicy) {
	if err := i.validatePutRowAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRowAccessPolicy",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IcebergTable) PutTimeouts(value *IcebergTableTimeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IcebergTable) PutUniqueConstraint(value interface{}) {
	if err := i.validatePutUniqueConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putUniqueConstraint",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IcebergTable) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := i.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (i *jsiiProxy_IcebergTable) ResetAggregationPolicy() {
	_jsii_.InvokeVoid(
		i,
		"resetAggregationPolicy",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetBaseLocation() {
	_jsii_.InvokeVoid(
		i,
		"resetBaseLocation",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetCatalog() {
	_jsii_.InvokeVoid(
		i,
		"resetCatalog",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetCatalogSync() {
	_jsii_.InvokeVoid(
		i,
		"resetCatalogSync",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetChangeTracking() {
	_jsii_.InvokeVoid(
		i,
		"resetChangeTracking",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetCheckConstraint() {
	_jsii_.InvokeVoid(
		i,
		"resetCheckConstraint",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetClusterBy() {
	_jsii_.InvokeVoid(
		i,
		"resetClusterBy",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetComment() {
	_jsii_.InvokeVoid(
		i,
		"resetComment",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetDataRetentionTimeInDays() {
	_jsii_.InvokeVoid(
		i,
		"resetDataRetentionTimeInDays",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetEnableDataCompaction() {
	_jsii_.InvokeVoid(
		i,
		"resetEnableDataCompaction",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetEnableIcebergMergeOnRead() {
	_jsii_.InvokeVoid(
		i,
		"resetEnableIcebergMergeOnRead",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetErrorLogging() {
	_jsii_.InvokeVoid(
		i,
		"resetErrorLogging",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetExternalVolume() {
	_jsii_.InvokeVoid(
		i,
		"resetExternalVolume",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetForeignKeyConstraint() {
	_jsii_.InvokeVoid(
		i,
		"resetForeignKeyConstraint",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetIcebergVersion() {
	_jsii_.InvokeVoid(
		i,
		"resetIcebergVersion",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetMaxDataExtensionTimeInDays() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxDataExtensionTimeInDays",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetPartitionBy() {
	_jsii_.InvokeVoid(
		i,
		"resetPartitionBy",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetPathLayout() {
	_jsii_.InvokeVoid(
		i,
		"resetPathLayout",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetPrimaryKeyConstraint() {
	_jsii_.InvokeVoid(
		i,
		"resetPrimaryKeyConstraint",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetRowAccessPolicy() {
	_jsii_.InvokeVoid(
		i,
		"resetRowAccessPolicy",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetStorageSerializationPolicy() {
	_jsii_.InvokeVoid(
		i,
		"resetStorageSerializationPolicy",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetTargetFileSize() {
	_jsii_.InvokeVoid(
		i,
		"resetTargetFileSize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) ResetUniqueConstraint() {
	_jsii_.InvokeVoid(
		i,
		"resetUniqueConstraint",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IcebergTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IcebergTable) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

