// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package catalogintegrationicebergrest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/catalogintegrationicebergrest/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/catalog_integration_iceberg_rest snowflake_catalog_integration_iceberg_rest}.
type CatalogIntegrationIcebergRest interface {
	cdktn.TerraformResource
	BearerRestAuthentication() CatalogIntegrationIcebergRestBearerRestAuthenticationOutputReference
	BearerRestAuthenticationInput() *CatalogIntegrationIcebergRestBearerRestAuthentication
	CatalogNamespace() *string
	SetCatalogNamespace(val *string)
	CatalogNamespaceInput() *string
	CatalogSource() *string
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
	DescribeOutput() CatalogIntegrationIcebergRestDescribeOutputList
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OauthRestAuthentication() CatalogIntegrationIcebergRestOauthRestAuthenticationOutputReference
	OauthRestAuthenticationInput() *CatalogIntegrationIcebergRestOauthRestAuthentication
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
	RefreshIntervalSeconds() *float64
	SetRefreshIntervalSeconds(val *float64)
	RefreshIntervalSecondsInput() *float64
	RestConfig() CatalogIntegrationIcebergRestRestConfigOutputReference
	RestConfigInput() *CatalogIntegrationIcebergRestRestConfig
	ShowOutput() CatalogIntegrationIcebergRestShowOutputList
	Sigv4RestAuthentication() CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference
	Sigv4RestAuthenticationInput() *CatalogIntegrationIcebergRestSigv4RestAuthentication
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CatalogIntegrationIcebergRestTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutBearerRestAuthentication(value *CatalogIntegrationIcebergRestBearerRestAuthentication)
	PutOauthRestAuthentication(value *CatalogIntegrationIcebergRestOauthRestAuthentication)
	PutRestConfig(value *CatalogIntegrationIcebergRestRestConfig)
	PutSigv4RestAuthentication(value *CatalogIntegrationIcebergRestSigv4RestAuthentication)
	PutTimeouts(value *CatalogIntegrationIcebergRestTimeouts)
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
	ResetBearerRestAuthentication()
	ResetCatalogNamespace()
	ResetComment()
	ResetId()
	ResetOauthRestAuthentication()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRefreshIntervalSeconds()
	ResetSigv4RestAuthentication()
	ResetTimeouts()
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

// The jsii proxy struct for CatalogIntegrationIcebergRest
type jsiiProxy_CatalogIntegrationIcebergRest struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) BearerRestAuthentication() CatalogIntegrationIcebergRestBearerRestAuthenticationOutputReference {
	var returns CatalogIntegrationIcebergRestBearerRestAuthenticationOutputReference
	_jsii_.Get(
		j,
		"bearerRestAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) BearerRestAuthenticationInput() *CatalogIntegrationIcebergRestBearerRestAuthentication {
	var returns *CatalogIntegrationIcebergRestBearerRestAuthentication
	_jsii_.Get(
		j,
		"bearerRestAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) CatalogNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) CatalogNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) CatalogSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) DescribeOutput() CatalogIntegrationIcebergRestDescribeOutputList {
	var returns CatalogIntegrationIcebergRestDescribeOutputList
	_jsii_.Get(
		j,
		"describeOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) OauthRestAuthentication() CatalogIntegrationIcebergRestOauthRestAuthenticationOutputReference {
	var returns CatalogIntegrationIcebergRestOauthRestAuthenticationOutputReference
	_jsii_.Get(
		j,
		"oauthRestAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) OauthRestAuthenticationInput() *CatalogIntegrationIcebergRestOauthRestAuthentication {
	var returns *CatalogIntegrationIcebergRestOauthRestAuthentication
	_jsii_.Get(
		j,
		"oauthRestAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) RefreshIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) RefreshIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) RestConfig() CatalogIntegrationIcebergRestRestConfigOutputReference {
	var returns CatalogIntegrationIcebergRestRestConfigOutputReference
	_jsii_.Get(
		j,
		"restConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) RestConfigInput() *CatalogIntegrationIcebergRestRestConfig {
	var returns *CatalogIntegrationIcebergRestRestConfig
	_jsii_.Get(
		j,
		"restConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) ShowOutput() CatalogIntegrationIcebergRestShowOutputList {
	var returns CatalogIntegrationIcebergRestShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Sigv4RestAuthentication() CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference {
	var returns CatalogIntegrationIcebergRestSigv4RestAuthenticationOutputReference
	_jsii_.Get(
		j,
		"sigv4RestAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Sigv4RestAuthenticationInput() *CatalogIntegrationIcebergRestSigv4RestAuthentication {
	var returns *CatalogIntegrationIcebergRestSigv4RestAuthentication
	_jsii_.Get(
		j,
		"sigv4RestAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) Timeouts() CatalogIntegrationIcebergRestTimeoutsOutputReference {
	var returns CatalogIntegrationIcebergRestTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/catalog_integration_iceberg_rest snowflake_catalog_integration_iceberg_rest} Resource.
func NewCatalogIntegrationIcebergRest(scope constructs.Construct, id *string, config *CatalogIntegrationIcebergRestConfig) CatalogIntegrationIcebergRest {
	_init_.Initialize()

	if err := validateNewCatalogIntegrationIcebergRestParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CatalogIntegrationIcebergRest{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.catalogIntegrationIcebergRest.CatalogIntegrationIcebergRest",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.20.0/docs/resources/catalog_integration_iceberg_rest snowflake_catalog_integration_iceberg_rest} Resource.
func NewCatalogIntegrationIcebergRest_Override(c CatalogIntegrationIcebergRest, scope constructs.Construct, id *string, config *CatalogIntegrationIcebergRestConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.catalogIntegrationIcebergRest.CatalogIntegrationIcebergRest",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest)SetCatalogNamespace(val *string) {
	if err := j.validateSetCatalogNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogNamespace",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CatalogIntegrationIcebergRest)SetRefreshIntervalSeconds(val *float64) {
	if err := j.validateSetRefreshIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshIntervalSeconds",
		val,
	)
}

// Generates CDKTN code for importing a CatalogIntegrationIcebergRest resource upon running "cdktn plan <stack-name>".
func CatalogIntegrationIcebergRest_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateCatalogIntegrationIcebergRest_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.catalogIntegrationIcebergRest.CatalogIntegrationIcebergRest",
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
func CatalogIntegrationIcebergRest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCatalogIntegrationIcebergRest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.catalogIntegrationIcebergRest.CatalogIntegrationIcebergRest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CatalogIntegrationIcebergRest_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCatalogIntegrationIcebergRest_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.catalogIntegrationIcebergRest.CatalogIntegrationIcebergRest",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CatalogIntegrationIcebergRest_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCatalogIntegrationIcebergRest_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.catalogIntegrationIcebergRest.CatalogIntegrationIcebergRest",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CatalogIntegrationIcebergRest_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-snowflake.catalogIntegrationIcebergRest.CatalogIntegrationIcebergRest",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRest) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRest) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRest) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRest) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRest) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRest) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRest) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRest) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRest) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CatalogIntegrationIcebergRest) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := c.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) PutBearerRestAuthentication(value *CatalogIntegrationIcebergRestBearerRestAuthentication) {
	if err := c.validatePutBearerRestAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBearerRestAuthentication",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) PutOauthRestAuthentication(value *CatalogIntegrationIcebergRestOauthRestAuthentication) {
	if err := c.validatePutOauthRestAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOauthRestAuthentication",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) PutRestConfig(value *CatalogIntegrationIcebergRestRestConfig) {
	if err := c.validatePutRestConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRestConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) PutSigv4RestAuthentication(value *CatalogIntegrationIcebergRestSigv4RestAuthentication) {
	if err := c.validatePutSigv4RestAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSigv4RestAuthentication",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) PutTimeouts(value *CatalogIntegrationIcebergRestTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := c.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ResetBearerRestAuthentication() {
	_jsii_.InvokeVoid(
		c,
		"resetBearerRestAuthentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ResetCatalogNamespace() {
	_jsii_.InvokeVoid(
		c,
		"resetCatalogNamespace",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ResetOauthRestAuthentication() {
	_jsii_.InvokeVoid(
		c,
		"resetOauthRestAuthentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ResetRefreshIntervalSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetRefreshIntervalSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ResetSigv4RestAuthentication() {
	_jsii_.InvokeVoid(
		c,
		"resetSigv4RestAuthentication",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CatalogIntegrationIcebergRest) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

