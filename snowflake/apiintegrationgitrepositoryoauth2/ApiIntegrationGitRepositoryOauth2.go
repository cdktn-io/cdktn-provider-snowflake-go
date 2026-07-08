// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apiintegrationgitrepositoryoauth2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v17/apiintegrationgitrepositoryoauth2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_git_repository_oauth2 snowflake_api_integration_git_repository_oauth2}.
type ApiIntegrationGitRepositoryOauth2 interface {
	cdktn.TerraformResource
	ApiAllowedPrefixes() *[]*string
	SetApiAllowedPrefixes(val *[]*string)
	ApiAllowedPrefixesInput() *[]*string
	ApiBlockedPrefixes() *[]*string
	SetApiBlockedPrefixes(val *[]*string)
	ApiBlockedPrefixesInput() *[]*string
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
	DescribeOutput() ApiIntegrationGitRepositoryOauth2DescribeOutputList
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
	OauthAccessTokenValidity() *float64
	SetOauthAccessTokenValidity(val *float64)
	OauthAccessTokenValidityInput() *float64
	OauthAllowedScopes() *[]*string
	SetOauthAllowedScopes(val *[]*string)
	OauthAllowedScopesInput() *[]*string
	OauthAuthorizationEndpoint() *string
	SetOauthAuthorizationEndpoint(val *string)
	OauthAuthorizationEndpointInput() *string
	OauthClientId() *string
	SetOauthClientId(val *string)
	OauthClientIdInput() *string
	OauthClientSecret() *string
	SetOauthClientSecret(val *string)
	OauthClientSecretInput() *string
	OauthRefreshTokenValidity() *float64
	SetOauthRefreshTokenValidity(val *float64)
	OauthRefreshTokenValidityInput() *float64
	OauthTokenEndpoint() *string
	SetOauthTokenEndpoint(val *string)
	OauthTokenEndpointInput() *string
	OauthUsername() *string
	SetOauthUsername(val *string)
	OauthUsernameInput() *string
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
	ShowOutput() ApiIntegrationGitRepositoryOauth2ShowOutputList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ApiIntegrationGitRepositoryOauth2TimeoutsOutputReference
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
	PutTimeouts(value *ApiIntegrationGitRepositoryOauth2Timeouts)
	ResetApiBlockedPrefixes()
	ResetComment()
	ResetId()
	ResetOauthAccessTokenValidity()
	ResetOauthAllowedScopes()
	ResetOauthRefreshTokenValidity()
	ResetOauthUsername()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for ApiIntegrationGitRepositoryOauth2
type jsiiProxy_ApiIntegrationGitRepositoryOauth2 struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ApiAllowedPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiAllowedPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ApiAllowedPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiAllowedPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ApiBlockedPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiBlockedPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ApiBlockedPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiBlockedPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) DescribeOutput() ApiIntegrationGitRepositoryOauth2DescribeOutputList {
	var returns ApiIntegrationGitRepositoryOauth2DescribeOutputList
	_jsii_.Get(
		j,
		"describeOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) FullyQualifiedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthAccessTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthAccessTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthAccessTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthAccessTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthAllowedScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthAllowedScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthAllowedScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthAllowedScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthAuthorizationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthAuthorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthAuthorizationEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthAuthorizationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthRefreshTokenValidity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthRefreshTokenValidityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"oauthRefreshTokenValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthTokenEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthTokenEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OauthUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ShowOutput() ApiIntegrationGitRepositoryOauth2ShowOutputList {
	var returns ApiIntegrationGitRepositoryOauth2ShowOutputList
	_jsii_.Get(
		j,
		"showOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) Timeouts() ApiIntegrationGitRepositoryOauth2TimeoutsOutputReference {
	var returns ApiIntegrationGitRepositoryOauth2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_git_repository_oauth2 snowflake_api_integration_git_repository_oauth2} Resource.
func NewApiIntegrationGitRepositoryOauth2(scope constructs.Construct, id *string, config *ApiIntegrationGitRepositoryOauth2Config) ApiIntegrationGitRepositoryOauth2 {
	_init_.Initialize()

	if err := validateNewApiIntegrationGitRepositoryOauth2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiIntegrationGitRepositoryOauth2{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.apiIntegrationGitRepositoryOauth2.ApiIntegrationGitRepositoryOauth2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/snowflakedb/snowflake/2.18.0/docs/resources/api_integration_git_repository_oauth2 snowflake_api_integration_git_repository_oauth2} Resource.
func NewApiIntegrationGitRepositoryOauth2_Override(a ApiIntegrationGitRepositoryOauth2, scope constructs.Construct, id *string, config *ApiIntegrationGitRepositoryOauth2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.apiIntegrationGitRepositoryOauth2.ApiIntegrationGitRepositoryOauth2",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetApiAllowedPrefixes(val *[]*string) {
	if err := j.validateSetApiAllowedPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiAllowedPrefixes",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetApiBlockedPrefixes(val *[]*string) {
	if err := j.validateSetApiBlockedPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiBlockedPrefixes",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetOauthAccessTokenValidity(val *float64) {
	if err := j.validateSetOauthAccessTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthAccessTokenValidity",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetOauthAllowedScopes(val *[]*string) {
	if err := j.validateSetOauthAllowedScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthAllowedScopes",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetOauthAuthorizationEndpoint(val *string) {
	if err := j.validateSetOauthAuthorizationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthAuthorizationEndpoint",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetOauthClientId(val *string) {
	if err := j.validateSetOauthClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthClientId",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetOauthClientSecret(val *string) {
	if err := j.validateSetOauthClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthClientSecret",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetOauthRefreshTokenValidity(val *float64) {
	if err := j.validateSetOauthRefreshTokenValidityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthRefreshTokenValidity",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetOauthTokenEndpoint(val *string) {
	if err := j.validateSetOauthTokenEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthTokenEndpoint",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetOauthUsername(val *string) {
	if err := j.validateSetOauthUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthUsername",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApiIntegrationGitRepositoryOauth2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a ApiIntegrationGitRepositoryOauth2 resource upon running "cdktn plan <stack-name>".
func ApiIntegrationGitRepositoryOauth2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateApiIntegrationGitRepositoryOauth2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.apiIntegrationGitRepositoryOauth2.ApiIntegrationGitRepositoryOauth2",
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
func ApiIntegrationGitRepositoryOauth2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiIntegrationGitRepositoryOauth2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.apiIntegrationGitRepositoryOauth2.ApiIntegrationGitRepositoryOauth2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiIntegrationGitRepositoryOauth2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiIntegrationGitRepositoryOauth2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.apiIntegrationGitRepositoryOauth2.ApiIntegrationGitRepositoryOauth2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApiIntegrationGitRepositoryOauth2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApiIntegrationGitRepositoryOauth2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-snowflake.apiIntegrationGitRepositoryOauth2.ApiIntegrationGitRepositoryOauth2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApiIntegrationGitRepositoryOauth2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-snowflake.apiIntegrationGitRepositoryOauth2.ApiIntegrationGitRepositoryOauth2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) PutTimeouts(value *ApiIntegrationGitRepositoryOauth2Timeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ResetApiBlockedPrefixes() {
	_jsii_.InvokeVoid(
		a,
		"resetApiBlockedPrefixes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ResetComment() {
	_jsii_.InvokeVoid(
		a,
		"resetComment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ResetOauthAccessTokenValidity() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthAccessTokenValidity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ResetOauthAllowedScopes() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthAllowedScopes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ResetOauthRefreshTokenValidity() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthRefreshTokenValidity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ResetOauthUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiIntegrationGitRepositoryOauth2) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

