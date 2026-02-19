// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakelistings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v16/datasnowflakelistings/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataSnowflakeListingsListingsDescribeOutputOutputReference interface {
	cdktn.ComplexObject
	ApplicationPackage() *string
	ApproverContact() *string
	BusinessNeeds() *string
	Categories() *string
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
	CreatedOn() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomizedContactInfo() *string
	DataAttributes() *string
	DataDictionary() *string
	DataPreview() *string
	Description() *string
	Distribution() *string
	// Experimental.
	Fqn() *string
	GlobalName() *string
	InternalValue() *DataSnowflakeListingsListingsDescribeOutput
	SetInternalValue(val *DataSnowflakeListingsListingsDescribeOutput)
	IsApplication() cdktn.IResolvable
	IsByRequest() cdktn.IResolvable
	IsLimitedTrial() cdktn.IResolvable
	IsMonetized() cdktn.IResolvable
	IsMountlessQueryable() cdktn.IResolvable
	IsShare() cdktn.IResolvable
	IsTargeted() cdktn.IResolvable
	LastCommittedVersionAlias() *string
	LastCommittedVersionName() *string
	LastCommittedVersionUri() *string
	LegacyUniformListingLocators() *string
	LimitedTrialPlan() *string
	ListingTerms() *string
	LiveVersionUri() *string
	ManifestYaml() *string
	MonetizationDisplayOrder() *string
	Name() *string
	OrganizationProfileName() *string
	Owner() *string
	OwnerRoleType() *string
	Profile() *string
	PublishedOn() *string
	PublishedVersionAlias() *string
	PublishedVersionName() *string
	PublishedVersionUri() *string
	RefreshSchedule() *string
	RefreshType() *string
	Regions() *string
	RejectionReason() *string
	RequestApprovalType() *string
	Resources() *string
	RetriedOn() *string
	ReviewState() *string
	Revisions() *string
	ScheduledDropTime() *string
	Share() *string
	State() *string
	Subtitle() *string
	SupportContact() *string
	TargetAccounts() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Title() *string
	TrialDetails() *string
	UniformListingLocator() *string
	UnpublishedByAdminReasons() *string
	UpdatedOn() *string
	UsageExamples() *string
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

// The jsii proxy struct for DataSnowflakeListingsListingsDescribeOutputOutputReference
type jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) ApplicationPackage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationPackage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) ApproverContact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approverContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) BusinessNeeds() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessNeeds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Categories() *string {
	var returns *string
	_jsii_.Get(
		j,
		"categories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) CustomizedContactInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customizedContactInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) DataAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) DataDictionary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDictionary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) DataPreview() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataPreview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Distribution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) GlobalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) InternalValue() *DataSnowflakeListingsListingsDescribeOutput {
	var returns *DataSnowflakeListingsListingsDescribeOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) IsApplication() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) IsByRequest() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isByRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) IsLimitedTrial() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isLimitedTrial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) IsMonetized() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isMonetized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) IsMountlessQueryable() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isMountlessQueryable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) IsShare() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isShare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) IsTargeted() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isTargeted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) LastCommittedVersionAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastCommittedVersionAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) LastCommittedVersionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastCommittedVersionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) LastCommittedVersionUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastCommittedVersionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) LegacyUniformListingLocators() *string {
	var returns *string
	_jsii_.Get(
		j,
		"legacyUniformListingLocators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) LimitedTrialPlan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitedTrialPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) ListingTerms() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listingTerms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) LiveVersionUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"liveVersionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) ManifestYaml() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestYaml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) MonetizationDisplayOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monetizationDisplayOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) OrganizationProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) OwnerRoleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerRoleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) PublishedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) PublishedVersionAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishedVersionAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) PublishedVersionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishedVersionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) PublishedVersionUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishedVersionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) RefreshSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) RefreshType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Regions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) RejectionReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rejectionReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) RequestApprovalType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestApprovalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Resources() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) RetriedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retriedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) ReviewState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reviewState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Revisions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) ScheduledDropTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledDropTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Share() *string {
	var returns *string
	_jsii_.Get(
		j,
		"share",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Subtitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) SupportContact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportContact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) TargetAccounts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) TrialDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trialDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) UniformListingLocator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniformListingLocator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) UnpublishedByAdminReasons() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unpublishedByAdminReasons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) UpdatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) UsageExamples() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageExamples",
		&returns,
	)
	return returns
}


func NewDataSnowflakeListingsListingsDescribeOutputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeListingsListingsDescribeOutputOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeListingsListingsDescribeOutputOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsListingsDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeListingsListingsDescribeOutputOutputReference_Override(d DataSnowflakeListingsListingsDescribeOutputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsListingsDescribeOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference)SetInternalValue(val *DataSnowflakeListingsListingsDescribeOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

