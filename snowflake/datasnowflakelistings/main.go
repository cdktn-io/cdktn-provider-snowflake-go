// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakelistings

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListings",
		reflect.TypeOf((*DataSnowflakeListings)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "like", GoGetter: "Like"},
			_jsii_.MemberProperty{JsiiProperty: "likeInput", GoGetter: "LikeInput"},
			_jsii_.MemberProperty{JsiiProperty: "limit", GoGetter: "Limit"},
			_jsii_.MemberProperty{JsiiProperty: "limitInput", GoGetter: "LimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "listings", GoGetter: "Listings"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putLimit", GoMethod: "PutLimit"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLike", GoMethod: "ResetLike"},
			_jsii_.MemberMethod{JsiiMethod: "resetLimit", GoMethod: "ResetLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartsWith", GoMethod: "ResetStartsWith"},
			_jsii_.MemberMethod{JsiiMethod: "resetWithDescribe", GoMethod: "ResetWithDescribe"},
			_jsii_.MemberProperty{JsiiProperty: "startsWith", GoGetter: "StartsWith"},
			_jsii_.MemberProperty{JsiiProperty: "startsWithInput", GoGetter: "StartsWithInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "withDescribe", GoGetter: "WithDescribe"},
			_jsii_.MemberProperty{JsiiProperty: "withDescribeInput", GoGetter: "WithDescribeInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataSnowflakeListings{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsConfig",
		reflect.TypeOf((*DataSnowflakeListingsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsLimit",
		reflect.TypeOf((*DataSnowflakeListingsLimit)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsLimitOutputReference",
		reflect.TypeOf((*DataSnowflakeListingsLimitOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "from", GoGetter: "From"},
			_jsii_.MemberProperty{JsiiProperty: "fromInput", GoGetter: "FromInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetFrom", GoMethod: "ResetFrom"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rows", GoGetter: "Rows"},
			_jsii_.MemberProperty{JsiiProperty: "rowsInput", GoGetter: "RowsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataSnowflakeListingsLimitOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsListings",
		reflect.TypeOf((*DataSnowflakeListingsListings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsListingsDescribeOutput",
		reflect.TypeOf((*DataSnowflakeListingsListingsDescribeOutput)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsListingsDescribeOutputList",
		reflect.TypeOf((*DataSnowflakeListingsListingsDescribeOutputList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataSnowflakeListingsListingsDescribeOutputList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsListingsDescribeOutputOutputReference",
		reflect.TypeOf((*DataSnowflakeListingsListingsDescribeOutputOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationPackage", GoGetter: "ApplicationPackage"},
			_jsii_.MemberProperty{JsiiProperty: "approverContact", GoGetter: "ApproverContact"},
			_jsii_.MemberProperty{JsiiProperty: "businessNeeds", GoGetter: "BusinessNeeds"},
			_jsii_.MemberProperty{JsiiProperty: "categories", GoGetter: "Categories"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "createdOn", GoGetter: "CreatedOn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customizedContactInfo", GoGetter: "CustomizedContactInfo"},
			_jsii_.MemberProperty{JsiiProperty: "dataAttributes", GoGetter: "DataAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "dataDictionary", GoGetter: "DataDictionary"},
			_jsii_.MemberProperty{JsiiProperty: "dataPreview", GoGetter: "DataPreview"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "distribution", GoGetter: "Distribution"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "globalName", GoGetter: "GlobalName"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isApplication", GoGetter: "IsApplication"},
			_jsii_.MemberProperty{JsiiProperty: "isByRequest", GoGetter: "IsByRequest"},
			_jsii_.MemberProperty{JsiiProperty: "isLimitedTrial", GoGetter: "IsLimitedTrial"},
			_jsii_.MemberProperty{JsiiProperty: "isMonetized", GoGetter: "IsMonetized"},
			_jsii_.MemberProperty{JsiiProperty: "isMountlessQueryable", GoGetter: "IsMountlessQueryable"},
			_jsii_.MemberProperty{JsiiProperty: "isShare", GoGetter: "IsShare"},
			_jsii_.MemberProperty{JsiiProperty: "isTargeted", GoGetter: "IsTargeted"},
			_jsii_.MemberProperty{JsiiProperty: "lastCommittedVersionAlias", GoGetter: "LastCommittedVersionAlias"},
			_jsii_.MemberProperty{JsiiProperty: "lastCommittedVersionName", GoGetter: "LastCommittedVersionName"},
			_jsii_.MemberProperty{JsiiProperty: "lastCommittedVersionUri", GoGetter: "LastCommittedVersionUri"},
			_jsii_.MemberProperty{JsiiProperty: "legacyUniformListingLocators", GoGetter: "LegacyUniformListingLocators"},
			_jsii_.MemberProperty{JsiiProperty: "limitedTrialPlan", GoGetter: "LimitedTrialPlan"},
			_jsii_.MemberProperty{JsiiProperty: "listingTerms", GoGetter: "ListingTerms"},
			_jsii_.MemberProperty{JsiiProperty: "liveVersionUri", GoGetter: "LiveVersionUri"},
			_jsii_.MemberProperty{JsiiProperty: "manifestYaml", GoGetter: "ManifestYaml"},
			_jsii_.MemberProperty{JsiiProperty: "monetizationDisplayOrder", GoGetter: "MonetizationDisplayOrder"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "organizationProfileName", GoGetter: "OrganizationProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "ownerRoleType", GoGetter: "OwnerRoleType"},
			_jsii_.MemberProperty{JsiiProperty: "profile", GoGetter: "Profile"},
			_jsii_.MemberProperty{JsiiProperty: "publishedOn", GoGetter: "PublishedOn"},
			_jsii_.MemberProperty{JsiiProperty: "publishedVersionAlias", GoGetter: "PublishedVersionAlias"},
			_jsii_.MemberProperty{JsiiProperty: "publishedVersionName", GoGetter: "PublishedVersionName"},
			_jsii_.MemberProperty{JsiiProperty: "publishedVersionUri", GoGetter: "PublishedVersionUri"},
			_jsii_.MemberProperty{JsiiProperty: "refreshSchedule", GoGetter: "RefreshSchedule"},
			_jsii_.MemberProperty{JsiiProperty: "refreshType", GoGetter: "RefreshType"},
			_jsii_.MemberProperty{JsiiProperty: "regions", GoGetter: "Regions"},
			_jsii_.MemberProperty{JsiiProperty: "rejectionReason", GoGetter: "RejectionReason"},
			_jsii_.MemberProperty{JsiiProperty: "requestApprovalType", GoGetter: "RequestApprovalType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resources", GoGetter: "Resources"},
			_jsii_.MemberProperty{JsiiProperty: "retriedOn", GoGetter: "RetriedOn"},
			_jsii_.MemberProperty{JsiiProperty: "reviewState", GoGetter: "ReviewState"},
			_jsii_.MemberProperty{JsiiProperty: "revisions", GoGetter: "Revisions"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledDropTime", GoGetter: "ScheduledDropTime"},
			_jsii_.MemberProperty{JsiiProperty: "share", GoGetter: "Share"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "subtitle", GoGetter: "Subtitle"},
			_jsii_.MemberProperty{JsiiProperty: "supportContact", GoGetter: "SupportContact"},
			_jsii_.MemberProperty{JsiiProperty: "targetAccounts", GoGetter: "TargetAccounts"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trialDetails", GoGetter: "TrialDetails"},
			_jsii_.MemberProperty{JsiiProperty: "uniformListingLocator", GoGetter: "UniformListingLocator"},
			_jsii_.MemberProperty{JsiiProperty: "unpublishedByAdminReasons", GoGetter: "UnpublishedByAdminReasons"},
			_jsii_.MemberProperty{JsiiProperty: "updatedOn", GoGetter: "UpdatedOn"},
			_jsii_.MemberProperty{JsiiProperty: "usageExamples", GoGetter: "UsageExamples"},
		},
		func() interface{} {
			j := jsiiProxy_DataSnowflakeListingsListingsDescribeOutputOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsListingsList",
		reflect.TypeOf((*DataSnowflakeListingsListingsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataSnowflakeListingsListingsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsListingsOutputReference",
		reflect.TypeOf((*DataSnowflakeListingsListingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "describeOutput", GoGetter: "DescribeOutput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "showOutput", GoGetter: "ShowOutput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataSnowflakeListingsListingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsListingsShowOutput",
		reflect.TypeOf((*DataSnowflakeListingsListingsShowOutput)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsListingsShowOutputList",
		reflect.TypeOf((*DataSnowflakeListingsListingsShowOutputList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataSnowflakeListingsListingsShowOutputList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.dataSnowflakeListings.DataSnowflakeListingsListingsShowOutputOutputReference",
		reflect.TypeOf((*DataSnowflakeListingsListingsShowOutputOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "createdOn", GoGetter: "CreatedOn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "detailedTargetAccounts", GoGetter: "DetailedTargetAccounts"},
			_jsii_.MemberProperty{JsiiProperty: "distribution", GoGetter: "Distribution"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "globalName", GoGetter: "GlobalName"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isApplication", GoGetter: "IsApplication"},
			_jsii_.MemberProperty{JsiiProperty: "isByRequest", GoGetter: "IsByRequest"},
			_jsii_.MemberProperty{JsiiProperty: "isLimitedTrial", GoGetter: "IsLimitedTrial"},
			_jsii_.MemberProperty{JsiiProperty: "isMonetized", GoGetter: "IsMonetized"},
			_jsii_.MemberProperty{JsiiProperty: "isMountlessQueryable", GoGetter: "IsMountlessQueryable"},
			_jsii_.MemberProperty{JsiiProperty: "isTargeted", GoGetter: "IsTargeted"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "organizationProfileName", GoGetter: "OrganizationProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "ownerRoleType", GoGetter: "OwnerRoleType"},
			_jsii_.MemberProperty{JsiiProperty: "profile", GoGetter: "Profile"},
			_jsii_.MemberProperty{JsiiProperty: "publishedOn", GoGetter: "PublishedOn"},
			_jsii_.MemberProperty{JsiiProperty: "regions", GoGetter: "Regions"},
			_jsii_.MemberProperty{JsiiProperty: "rejectedOn", GoGetter: "RejectedOn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "reviewState", GoGetter: "ReviewState"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "subtitle", GoGetter: "Subtitle"},
			_jsii_.MemberProperty{JsiiProperty: "targetAccounts", GoGetter: "TargetAccounts"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "uniformListingLocator", GoGetter: "UniformListingLocator"},
			_jsii_.MemberProperty{JsiiProperty: "updatedOn", GoGetter: "UpdatedOn"},
		},
		func() interface{} {
			j := jsiiProxy_DataSnowflakeListingsListingsShowOutputOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
