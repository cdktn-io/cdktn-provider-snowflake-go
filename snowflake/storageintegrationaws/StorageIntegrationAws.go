// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storageintegrationaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/storageintegrationaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_aws snowflake_storage_integration_aws}.
type StorageIntegrationAws interface {
	cdktn.TerraformResource
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
	DescribeOutput() StorageIntegrationAwsDescribeOutputList
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
	ShowOutput() StorageIntegrationAwsShowOutputList
	StorageAllowedLocations() *[]*string
	SetStorageAllowedLocations(val *[]*string)
	StorageAllowedLocationsInput() *[]*string
	StorageAwsExternalId() *string
	SetStorageAwsExternalId(val *string)
	StorageAwsExternalIdInput() *string
	StorageAwsObjectAcl() *string
	SetStorageAwsObjectAcl(val *string)
	StorageAwsObjectAclInput() *string
	StorageAwsRoleArn() *string
	SetStorageAwsRoleArn(val *string)
	StorageAwsRoleArnInput() *string
	StorageBlockedLocations() *[]*string
	SetStorageBlockedLocations(val *[]*string)
	StorageBlockedLocationsInput() *[]*string
	StorageProvider() *string
	SetStorageProvider(val *string)
	StorageProviderInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() StorageIntegrationAwsTimeoutsOutputReference
	TimeoutsInput() interface{}
	UsePrivatelinkEndpoint() *string
	SetUsePrivatelinkEndpoint(val *string)
	UsePrivatelinkEndpointInput() *string
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
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
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
	PutTimeouts(value *StorageIntegrationAwsTimeouts)
	ResetComment()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStorageAwsExternalId()
	ResetStorageAwsObjectAcl()
	ResetStorageBlockedLocations()
	ResetTimeouts()
	ResetUsePrivatelinkEndpoint()
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
}

// The jsii proxy struct for StorageIntegrationAws
type jsiiProxy_StorageIntegrationAws struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_StorageIntegrationAws) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) DescribeOutput() StorageIntegrationAwsDescribeOutputList {
	var returns StorageIntegrationAwsDescribeOutputList
	_jsii_.Get(
		j,
		"describeOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) ShowOutput() StorageIntegrationAwsShowOutputList {
	var returns StorageIntegrationAwsShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) StorageAllowedLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storageAllowedLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) StorageAllowedLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storageAllowedLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) StorageAwsExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAwsExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) StorageAwsExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAwsExternalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) StorageAwsObjectAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAwsObjectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) StorageAwsObjectAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAwsObjectAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) StorageAwsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAwsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) StorageAwsRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAwsRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) StorageBlockedLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storageBlockedLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) StorageBlockedLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"storageBlockedLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) StorageProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) StorageProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) Timeouts() StorageIntegrationAwsTimeoutsOutputReference {
	var returns StorageIntegrationAwsTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) UsePrivatelinkEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePrivatelinkEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageIntegrationAws) UsePrivatelinkEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePrivatelinkEndpointInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_aws snowflake_storage_integration_aws} Resource.
func NewStorageIntegrationAws(scope constructs.Construct, id *string, config *StorageIntegrationAwsConfig) StorageIntegrationAws {
	_init_.Initialize()

	if err := validateNewStorageIntegrationAwsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageIntegrationAws{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.storageIntegrationAws.StorageIntegrationAws",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.13.0/docs/resources/storage_integration_aws snowflake_storage_integration_aws} Resource.
func NewStorageIntegrationAws_Override(s StorageIntegrationAws, scope constructs.Construct, id *string, config *StorageIntegrationAwsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.storageIntegrationAws.StorageIntegrationAws",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetStorageAllowedLocations(val *[]*string) {
	if err := j.validateSetStorageAllowedLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAllowedLocations",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetStorageAwsExternalId(val *string) {
	if err := j.validateSetStorageAwsExternalIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAwsExternalId",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetStorageAwsObjectAcl(val *string) {
	if err := j.validateSetStorageAwsObjectAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAwsObjectAcl",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetStorageAwsRoleArn(val *string) {
	if err := j.validateSetStorageAwsRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAwsRoleArn",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetStorageBlockedLocations(val *[]*string) {
	if err := j.validateSetStorageBlockedLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageBlockedLocations",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetStorageProvider(val *string) {
	if err := j.validateSetStorageProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageProvider",
		val,
	)
}

func (j *jsiiProxy_StorageIntegrationAws)SetUsePrivatelinkEndpoint(val *string) {
	if err := j.validateSetUsePrivatelinkEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePrivatelinkEndpoint",
		val,
	)
}

// Generates CDKTN code for importing a StorageIntegrationAws resource upon running "cdktn plan <stack-name>".
func StorageIntegrationAws_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateStorageIntegrationAws_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.storageIntegrationAws.StorageIntegrationAws",
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
func StorageIntegrationAws_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageIntegrationAws_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.storageIntegrationAws.StorageIntegrationAws",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageIntegrationAws_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageIntegrationAws_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.storageIntegrationAws.StorageIntegrationAws",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageIntegrationAws_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageIntegrationAws_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.storageIntegrationAws.StorageIntegrationAws",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StorageIntegrationAws_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-snowflake.storageIntegrationAws.StorageIntegrationAws",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StorageIntegrationAws) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StorageIntegrationAws) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StorageIntegrationAws) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StorageIntegrationAws) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StorageIntegrationAws) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StorageIntegrationAws) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StorageIntegrationAws) PutTimeouts(value *StorageIntegrationAwsTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageIntegrationAws) ResetComment() {
	_jsii_.InvokeVoid(
		s,
		"resetComment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageIntegrationAws) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageIntegrationAws) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageIntegrationAws) ResetStorageAwsExternalId() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageAwsExternalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageIntegrationAws) ResetStorageAwsObjectAcl() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageAwsObjectAcl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageIntegrationAws) ResetStorageBlockedLocations() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageBlockedLocations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageIntegrationAws) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageIntegrationAws) ResetUsePrivatelinkEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetUsePrivatelinkEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageIntegrationAws) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageIntegrationAws) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

