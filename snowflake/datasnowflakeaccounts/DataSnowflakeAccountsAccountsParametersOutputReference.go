// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasnowflakeaccounts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/jsii"

	"github.com/cdktn-io/cdktn-provider-snowflake-go/snowflake/v18/datasnowflakeaccounts/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataSnowflakeAccountsAccountsParametersOutputReference interface {
	cdktn.ComplexObject
	AbortDetachedQuery() DataSnowflakeAccountsAccountsParametersAbortDetachedQueryList
	ActivePythonProfiler() DataSnowflakeAccountsAccountsParametersActivePythonProfilerList
	AllowBindValuesAccess() DataSnowflakeAccountsAccountsParametersAllowBindValuesAccessList
	AllowClientMfaCaching() DataSnowflakeAccountsAccountsParametersAllowClientMfaCachingList
	AllowedSpcsWorkloadTypes() DataSnowflakeAccountsAccountsParametersAllowedSpcsWorkloadTypesList
	AllowIdToken() DataSnowflakeAccountsAccountsParametersAllowIdTokenList
	Autocommit() DataSnowflakeAccountsAccountsParametersAutocommitList
	BaseLocationPrefix() DataSnowflakeAccountsAccountsParametersBaseLocationPrefixList
	BinaryInputFormat() DataSnowflakeAccountsAccountsParametersBinaryInputFormatList
	BinaryOutputFormat() DataSnowflakeAccountsAccountsParametersBinaryOutputFormatList
	Catalog() DataSnowflakeAccountsAccountsParametersCatalogList
	CatalogSync() DataSnowflakeAccountsAccountsParametersCatalogSyncList
	ClientEnableLogInfoStatementParameters() DataSnowflakeAccountsAccountsParametersClientEnableLogInfoStatementParametersList
	ClientEncryptionKeySize() DataSnowflakeAccountsAccountsParametersClientEncryptionKeySizeList
	ClientMemoryLimit() DataSnowflakeAccountsAccountsParametersClientMemoryLimitList
	ClientMetadataRequestUseConnectionCtx() DataSnowflakeAccountsAccountsParametersClientMetadataRequestUseConnectionCtxList
	ClientMetadataUseSessionDatabase() DataSnowflakeAccountsAccountsParametersClientMetadataUseSessionDatabaseList
	ClientPrefetchThreads() DataSnowflakeAccountsAccountsParametersClientPrefetchThreadsList
	ClientResultChunkSize() DataSnowflakeAccountsAccountsParametersClientResultChunkSizeList
	ClientResultColumnCaseInsensitive() DataSnowflakeAccountsAccountsParametersClientResultColumnCaseInsensitiveList
	ClientSessionKeepAlive() DataSnowflakeAccountsAccountsParametersClientSessionKeepAliveList
	ClientSessionKeepAliveHeartbeatFrequency() DataSnowflakeAccountsAccountsParametersClientSessionKeepAliveHeartbeatFrequencyList
	ClientTimestampTypeMapping() DataSnowflakeAccountsAccountsParametersClientTimestampTypeMappingList
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
	CortexCodeCliDailyEstCreditLimitPerUser() DataSnowflakeAccountsAccountsParametersCortexCodeCliDailyEstCreditLimitPerUserList
	CortexCodeDesktopDailyEstCreditLimitPerUser() DataSnowflakeAccountsAccountsParametersCortexCodeDesktopDailyEstCreditLimitPerUserList
	CortexCodeSnowsightDailyEstCreditLimitPerUser() DataSnowflakeAccountsAccountsParametersCortexCodeSnowsightDailyEstCreditLimitPerUserList
	CortexEnabledCrossRegion() DataSnowflakeAccountsAccountsParametersCortexEnabledCrossRegionList
	CortexModelsAllowlist() DataSnowflakeAccountsAccountsParametersCortexModelsAllowlistStructList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CsvTimestampFormat() DataSnowflakeAccountsAccountsParametersCsvTimestampFormatList
	DataMetricSchedule() DataSnowflakeAccountsAccountsParametersDataMetricScheduleList
	DataRetentionTimeInDays() DataSnowflakeAccountsAccountsParametersDataRetentionTimeInDaysList
	DateInputFormat() DataSnowflakeAccountsAccountsParametersDateInputFormatList
	DateOutputFormat() DataSnowflakeAccountsAccountsParametersDateOutputFormatList
	DefaultDbtVersion() DataSnowflakeAccountsAccountsParametersDefaultDbtVersionList
	DefaultDdlCollation() DataSnowflakeAccountsAccountsParametersDefaultDdlCollationList
	DefaultNotebookComputePoolCpu() DataSnowflakeAccountsAccountsParametersDefaultNotebookComputePoolCpuList
	DefaultNotebookComputePoolGpu() DataSnowflakeAccountsAccountsParametersDefaultNotebookComputePoolGpuList
	DefaultNullOrdering() DataSnowflakeAccountsAccountsParametersDefaultNullOrderingList
	DefaultStreamlitComputePool() DataSnowflakeAccountsAccountsParametersDefaultStreamlitComputePoolList
	DefaultStreamlitNotebookWarehouse() DataSnowflakeAccountsAccountsParametersDefaultStreamlitNotebookWarehouseList
	DisableUiDownloadButton() DataSnowflakeAccountsAccountsParametersDisableUiDownloadButtonList
	DisableUserPrivilegeGrants() DataSnowflakeAccountsAccountsParametersDisableUserPrivilegeGrantsList
	DisallowedSpcsWorkloadTypes() DataSnowflakeAccountsAccountsParametersDisallowedSpcsWorkloadTypesList
	EnableAutomaticSensitiveDataClassificationLog() DataSnowflakeAccountsAccountsParametersEnableAutomaticSensitiveDataClassificationLogList
	EnableBudgetEventLogging() DataSnowflakeAccountsAccountsParametersEnableBudgetEventLoggingList
	EnableConsoleOutput() DataSnowflakeAccountsAccountsParametersEnableConsoleOutputList
	EnableCortexAnalyst() DataSnowflakeAccountsAccountsParametersEnableCortexAnalystList
	EnableDataCompaction() DataSnowflakeAccountsAccountsParametersEnableDataCompactionList
	EnableEgressCostOptimizer() DataSnowflakeAccountsAccountsParametersEnableEgressCostOptimizerList
	EnableGetDdlUseDataTypeAlias() DataSnowflakeAccountsAccountsParametersEnableGetDdlUseDataTypeAliasList
	EnableIcebergMergeOnRead() DataSnowflakeAccountsAccountsParametersEnableIcebergMergeOnReadList
	EnableIdentifierFirstLogin() DataSnowflakeAccountsAccountsParametersEnableIdentifierFirstLoginList
	EnableInternalStagesPrivatelink() DataSnowflakeAccountsAccountsParametersEnableInternalStagesPrivatelinkList
	EnableNotebookCreationInPersonalDb() DataSnowflakeAccountsAccountsParametersEnableNotebookCreationInPersonalDbList
	EnablePerAccountAppServicePrivatelinkUrl() DataSnowflakeAccountsAccountsParametersEnablePerAccountAppServicePrivatelinkUrlList
	EnablePersonalDatabase() DataSnowflakeAccountsAccountsParametersEnablePersonalDatabaseList
	EnableSpcsBlockStorageSnowflakeFullEncryptionEnforcement() DataSnowflakeAccountsAccountsParametersEnableSpcsBlockStorageSnowflakeFullEncryptionEnforcementList
	EnableTagPropagationEventLogging() DataSnowflakeAccountsAccountsParametersEnableTagPropagationEventLoggingList
	EnableTriSecretAndRekeyOptOutForImageRepository() DataSnowflakeAccountsAccountsParametersEnableTriSecretAndRekeyOptOutForImageRepositoryList
	EnableTriSecretAndRekeyOptOutForSpcsBlockStorage() DataSnowflakeAccountsAccountsParametersEnableTriSecretAndRekeyOptOutForSpcsBlockStorageList
	EnableUnhandledExceptionsReporting() DataSnowflakeAccountsAccountsParametersEnableUnhandledExceptionsReportingList
	EnableUnloadPhysicalTypeOptimization() DataSnowflakeAccountsAccountsParametersEnableUnloadPhysicalTypeOptimizationList
	EnableUnredactedQuerySyntaxError() DataSnowflakeAccountsAccountsParametersEnableUnredactedQuerySyntaxErrorList
	EnableUnredactedSecureObjectError() DataSnowflakeAccountsAccountsParametersEnableUnredactedSecureObjectErrorList
	EnforceNetworkRulesForInternalStages() DataSnowflakeAccountsAccountsParametersEnforceNetworkRulesForInternalStagesList
	ErrorOnNondeterministicMerge() DataSnowflakeAccountsAccountsParametersErrorOnNondeterministicMergeList
	ErrorOnNondeterministicUpdate() DataSnowflakeAccountsAccountsParametersErrorOnNondeterministicUpdateList
	EventTable() DataSnowflakeAccountsAccountsParametersEventTableList
	ExternalOauthAddPrivilegedRolesToBlockedList() DataSnowflakeAccountsAccountsParametersExternalOauthAddPrivilegedRolesToBlockedListStructList
	ExternalVolume() DataSnowflakeAccountsAccountsParametersExternalVolumeList
	// Experimental.
	Fqn() *string
	GeographyOutputFormat() DataSnowflakeAccountsAccountsParametersGeographyOutputFormatList
	GeometryOutputFormat() DataSnowflakeAccountsAccountsParametersGeometryOutputFormatList
	HybridTableLockTimeout() DataSnowflakeAccountsAccountsParametersHybridTableLockTimeoutList
	IcebergVersionDefault() DataSnowflakeAccountsAccountsParametersIcebergVersionDefaultList
	InitialReplicationSizeLimitInTb() DataSnowflakeAccountsAccountsParametersInitialReplicationSizeLimitInTbList
	InternalValue() *DataSnowflakeAccountsAccountsParameters
	SetInternalValue(val *DataSnowflakeAccountsAccountsParameters)
	JdbcTreatDecimalAsInt() DataSnowflakeAccountsAccountsParametersJdbcTreatDecimalAsIntList
	JdbcTreatTimestampNtzAsUtc() DataSnowflakeAccountsAccountsParametersJdbcTreatTimestampNtzAsUtcList
	JdbcUseSessionTimezone() DataSnowflakeAccountsAccountsParametersJdbcUseSessionTimezoneList
	JsonIndent() DataSnowflakeAccountsAccountsParametersJsonIndentList
	JsTreatIntegerAsBigint() DataSnowflakeAccountsAccountsParametersJsTreatIntegerAsBigintList
	ListingAutoFulfillmentReplicationRefreshSchedule() DataSnowflakeAccountsAccountsParametersListingAutoFulfillmentReplicationRefreshScheduleList
	LockTimeout() DataSnowflakeAccountsAccountsParametersLockTimeoutList
	LogEventLevel() DataSnowflakeAccountsAccountsParametersLogEventLevelList
	LogLevel() DataSnowflakeAccountsAccountsParametersLogLevelList
	MaxConcurrencyLevel() DataSnowflakeAccountsAccountsParametersMaxConcurrencyLevelList
	MaxDataExtensionTimeInDays() DataSnowflakeAccountsAccountsParametersMaxDataExtensionTimeInDaysList
	MetricLevel() DataSnowflakeAccountsAccountsParametersMetricLevelList
	MinDataRetentionTimeInDays() DataSnowflakeAccountsAccountsParametersMinDataRetentionTimeInDaysList
	MultiStatementCount() DataSnowflakeAccountsAccountsParametersMultiStatementCountList
	NetworkPolicy() DataSnowflakeAccountsAccountsParametersNetworkPolicyList
	NoorderSequenceAsDefault() DataSnowflakeAccountsAccountsParametersNoorderSequenceAsDefaultList
	OauthAddPrivilegedRolesToBlockedList() DataSnowflakeAccountsAccountsParametersOauthAddPrivilegedRolesToBlockedListStructList
	OdbcTreatDecimalAsInt() DataSnowflakeAccountsAccountsParametersOdbcTreatDecimalAsIntList
	PeriodicDataRekeying() DataSnowflakeAccountsAccountsParametersPeriodicDataRekeyingList
	PipeExecutionPaused() DataSnowflakeAccountsAccountsParametersPipeExecutionPausedList
	PreventLoadFromInlineUrl() DataSnowflakeAccountsAccountsParametersPreventLoadFromInlineUrlList
	PreventUnloadToInlineUrl() DataSnowflakeAccountsAccountsParametersPreventUnloadToInlineUrlList
	PreventUnloadToInternalStages() DataSnowflakeAccountsAccountsParametersPreventUnloadToInternalStagesList
	PythonProfilerModules() DataSnowflakeAccountsAccountsParametersPythonProfilerModulesList
	PythonProfilerTargetStage() DataSnowflakeAccountsAccountsParametersPythonProfilerTargetStageList
	QueryTag() DataSnowflakeAccountsAccountsParametersQueryTagList
	QuotedIdentifiersIgnoreCase() DataSnowflakeAccountsAccountsParametersQuotedIdentifiersIgnoreCaseList
	ReadConsistencyMode() DataSnowflakeAccountsAccountsParametersReadConsistencyModeList
	ReplaceInvalidCharacters() DataSnowflakeAccountsAccountsParametersReplaceInvalidCharactersList
	RequireStorageIntegrationForStageCreation() DataSnowflakeAccountsAccountsParametersRequireStorageIntegrationForStageCreationList
	RequireStorageIntegrationForStageOperation() DataSnowflakeAccountsAccountsParametersRequireStorageIntegrationForStageOperationList
	RowsPerResultset() DataSnowflakeAccountsAccountsParametersRowsPerResultsetList
	RowTimestampDefault() DataSnowflakeAccountsAccountsParametersRowTimestampDefaultList
	S3StageVpceDnsName() DataSnowflakeAccountsAccountsParametersS3StageVpceDnsNameList
	SearchPath() DataSnowflakeAccountsAccountsParametersSearchPathList
	ServerlessTaskMaxStatementSize() DataSnowflakeAccountsAccountsParametersServerlessTaskMaxStatementSizeList
	ServerlessTaskMinStatementSize() DataSnowflakeAccountsAccountsParametersServerlessTaskMinStatementSizeList
	ShareRestrictions() DataSnowflakeAccountsAccountsParametersShareRestrictionsList
	SimulatedDataSharingConsumer() DataSnowflakeAccountsAccountsParametersSimulatedDataSharingConsumerList
	SqlTraceQueryText() DataSnowflakeAccountsAccountsParametersSqlTraceQueryTextList
	SsoLoginPage() DataSnowflakeAccountsAccountsParametersSsoLoginPageList
	StatementQueuedTimeoutInSeconds() DataSnowflakeAccountsAccountsParametersStatementQueuedTimeoutInSecondsList
	StatementTimeoutInSeconds() DataSnowflakeAccountsAccountsParametersStatementTimeoutInSecondsList
	StorageSerializationPolicy() DataSnowflakeAccountsAccountsParametersStorageSerializationPolicyList
	StrictJsonOutput() DataSnowflakeAccountsAccountsParametersStrictJsonOutputList
	SuspendTaskAfterNumFailures() DataSnowflakeAccountsAccountsParametersSuspendTaskAfterNumFailuresList
	TaskAutoRetryAttempts() DataSnowflakeAccountsAccountsParametersTaskAutoRetryAttemptsList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeInputFormat() DataSnowflakeAccountsAccountsParametersTimeInputFormatList
	TimeOutputFormat() DataSnowflakeAccountsAccountsParametersTimeOutputFormatList
	TimestampDayIsAlways24H() DataSnowflakeAccountsAccountsParametersTimestampDayIsAlways24HList
	TimestampInputFormat() DataSnowflakeAccountsAccountsParametersTimestampInputFormatList
	TimestampLtzOutputFormat() DataSnowflakeAccountsAccountsParametersTimestampLtzOutputFormatList
	TimestampNtzOutputFormat() DataSnowflakeAccountsAccountsParametersTimestampNtzOutputFormatList
	TimestampOutputFormat() DataSnowflakeAccountsAccountsParametersTimestampOutputFormatList
	TimestampTypeMapping() DataSnowflakeAccountsAccountsParametersTimestampTypeMappingList
	TimestampTzOutputFormat() DataSnowflakeAccountsAccountsParametersTimestampTzOutputFormatList
	Timezone() DataSnowflakeAccountsAccountsParametersTimezoneList
	TraceLevel() DataSnowflakeAccountsAccountsParametersTraceLevelList
	TransactionAbortOnError() DataSnowflakeAccountsAccountsParametersTransactionAbortOnErrorList
	TransactionDefaultIsolationLevel() DataSnowflakeAccountsAccountsParametersTransactionDefaultIsolationLevelList
	TwoDigitCenturyStart() DataSnowflakeAccountsAccountsParametersTwoDigitCenturyStartList
	UnsupportedDdlAction() DataSnowflakeAccountsAccountsParametersUnsupportedDdlActionList
	UseCachedResult() DataSnowflakeAccountsAccountsParametersUseCachedResultList
	UserTaskManagedInitialWarehouseSize() DataSnowflakeAccountsAccountsParametersUserTaskManagedInitialWarehouseSizeList
	UserTaskMinimumTriggerIntervalInSeconds() DataSnowflakeAccountsAccountsParametersUserTaskMinimumTriggerIntervalInSecondsList
	UserTaskTimeoutMs() DataSnowflakeAccountsAccountsParametersUserTaskTimeoutMsList
	UseWorkspacesForSql() DataSnowflakeAccountsAccountsParametersUseWorkspacesForSqlList
	WeekOfYearPolicy() DataSnowflakeAccountsAccountsParametersWeekOfYearPolicyList
	WeekStart() DataSnowflakeAccountsAccountsParametersWeekStartList
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

// The jsii proxy struct for DataSnowflakeAccountsAccountsParametersOutputReference
type jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) AbortDetachedQuery() DataSnowflakeAccountsAccountsParametersAbortDetachedQueryList {
	var returns DataSnowflakeAccountsAccountsParametersAbortDetachedQueryList
	_jsii_.Get(
		j,
		"abortDetachedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ActivePythonProfiler() DataSnowflakeAccountsAccountsParametersActivePythonProfilerList {
	var returns DataSnowflakeAccountsAccountsParametersActivePythonProfilerList
	_jsii_.Get(
		j,
		"activePythonProfiler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) AllowBindValuesAccess() DataSnowflakeAccountsAccountsParametersAllowBindValuesAccessList {
	var returns DataSnowflakeAccountsAccountsParametersAllowBindValuesAccessList
	_jsii_.Get(
		j,
		"allowBindValuesAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) AllowClientMfaCaching() DataSnowflakeAccountsAccountsParametersAllowClientMfaCachingList {
	var returns DataSnowflakeAccountsAccountsParametersAllowClientMfaCachingList
	_jsii_.Get(
		j,
		"allowClientMfaCaching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) AllowedSpcsWorkloadTypes() DataSnowflakeAccountsAccountsParametersAllowedSpcsWorkloadTypesList {
	var returns DataSnowflakeAccountsAccountsParametersAllowedSpcsWorkloadTypesList
	_jsii_.Get(
		j,
		"allowedSpcsWorkloadTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) AllowIdToken() DataSnowflakeAccountsAccountsParametersAllowIdTokenList {
	var returns DataSnowflakeAccountsAccountsParametersAllowIdTokenList
	_jsii_.Get(
		j,
		"allowIdToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) Autocommit() DataSnowflakeAccountsAccountsParametersAutocommitList {
	var returns DataSnowflakeAccountsAccountsParametersAutocommitList
	_jsii_.Get(
		j,
		"autocommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) BaseLocationPrefix() DataSnowflakeAccountsAccountsParametersBaseLocationPrefixList {
	var returns DataSnowflakeAccountsAccountsParametersBaseLocationPrefixList
	_jsii_.Get(
		j,
		"baseLocationPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) BinaryInputFormat() DataSnowflakeAccountsAccountsParametersBinaryInputFormatList {
	var returns DataSnowflakeAccountsAccountsParametersBinaryInputFormatList
	_jsii_.Get(
		j,
		"binaryInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) BinaryOutputFormat() DataSnowflakeAccountsAccountsParametersBinaryOutputFormatList {
	var returns DataSnowflakeAccountsAccountsParametersBinaryOutputFormatList
	_jsii_.Get(
		j,
		"binaryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) Catalog() DataSnowflakeAccountsAccountsParametersCatalogList {
	var returns DataSnowflakeAccountsAccountsParametersCatalogList
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) CatalogSync() DataSnowflakeAccountsAccountsParametersCatalogSyncList {
	var returns DataSnowflakeAccountsAccountsParametersCatalogSyncList
	_jsii_.Get(
		j,
		"catalogSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ClientEnableLogInfoStatementParameters() DataSnowflakeAccountsAccountsParametersClientEnableLogInfoStatementParametersList {
	var returns DataSnowflakeAccountsAccountsParametersClientEnableLogInfoStatementParametersList
	_jsii_.Get(
		j,
		"clientEnableLogInfoStatementParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ClientEncryptionKeySize() DataSnowflakeAccountsAccountsParametersClientEncryptionKeySizeList {
	var returns DataSnowflakeAccountsAccountsParametersClientEncryptionKeySizeList
	_jsii_.Get(
		j,
		"clientEncryptionKeySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ClientMemoryLimit() DataSnowflakeAccountsAccountsParametersClientMemoryLimitList {
	var returns DataSnowflakeAccountsAccountsParametersClientMemoryLimitList
	_jsii_.Get(
		j,
		"clientMemoryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ClientMetadataRequestUseConnectionCtx() DataSnowflakeAccountsAccountsParametersClientMetadataRequestUseConnectionCtxList {
	var returns DataSnowflakeAccountsAccountsParametersClientMetadataRequestUseConnectionCtxList
	_jsii_.Get(
		j,
		"clientMetadataRequestUseConnectionCtx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ClientMetadataUseSessionDatabase() DataSnowflakeAccountsAccountsParametersClientMetadataUseSessionDatabaseList {
	var returns DataSnowflakeAccountsAccountsParametersClientMetadataUseSessionDatabaseList
	_jsii_.Get(
		j,
		"clientMetadataUseSessionDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ClientPrefetchThreads() DataSnowflakeAccountsAccountsParametersClientPrefetchThreadsList {
	var returns DataSnowflakeAccountsAccountsParametersClientPrefetchThreadsList
	_jsii_.Get(
		j,
		"clientPrefetchThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ClientResultChunkSize() DataSnowflakeAccountsAccountsParametersClientResultChunkSizeList {
	var returns DataSnowflakeAccountsAccountsParametersClientResultChunkSizeList
	_jsii_.Get(
		j,
		"clientResultChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ClientResultColumnCaseInsensitive() DataSnowflakeAccountsAccountsParametersClientResultColumnCaseInsensitiveList {
	var returns DataSnowflakeAccountsAccountsParametersClientResultColumnCaseInsensitiveList
	_jsii_.Get(
		j,
		"clientResultColumnCaseInsensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ClientSessionKeepAlive() DataSnowflakeAccountsAccountsParametersClientSessionKeepAliveList {
	var returns DataSnowflakeAccountsAccountsParametersClientSessionKeepAliveList
	_jsii_.Get(
		j,
		"clientSessionKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ClientSessionKeepAliveHeartbeatFrequency() DataSnowflakeAccountsAccountsParametersClientSessionKeepAliveHeartbeatFrequencyList {
	var returns DataSnowflakeAccountsAccountsParametersClientSessionKeepAliveHeartbeatFrequencyList
	_jsii_.Get(
		j,
		"clientSessionKeepAliveHeartbeatFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ClientTimestampTypeMapping() DataSnowflakeAccountsAccountsParametersClientTimestampTypeMappingList {
	var returns DataSnowflakeAccountsAccountsParametersClientTimestampTypeMappingList
	_jsii_.Get(
		j,
		"clientTimestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) CortexCodeCliDailyEstCreditLimitPerUser() DataSnowflakeAccountsAccountsParametersCortexCodeCliDailyEstCreditLimitPerUserList {
	var returns DataSnowflakeAccountsAccountsParametersCortexCodeCliDailyEstCreditLimitPerUserList
	_jsii_.Get(
		j,
		"cortexCodeCliDailyEstCreditLimitPerUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) CortexCodeDesktopDailyEstCreditLimitPerUser() DataSnowflakeAccountsAccountsParametersCortexCodeDesktopDailyEstCreditLimitPerUserList {
	var returns DataSnowflakeAccountsAccountsParametersCortexCodeDesktopDailyEstCreditLimitPerUserList
	_jsii_.Get(
		j,
		"cortexCodeDesktopDailyEstCreditLimitPerUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) CortexCodeSnowsightDailyEstCreditLimitPerUser() DataSnowflakeAccountsAccountsParametersCortexCodeSnowsightDailyEstCreditLimitPerUserList {
	var returns DataSnowflakeAccountsAccountsParametersCortexCodeSnowsightDailyEstCreditLimitPerUserList
	_jsii_.Get(
		j,
		"cortexCodeSnowsightDailyEstCreditLimitPerUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) CortexEnabledCrossRegion() DataSnowflakeAccountsAccountsParametersCortexEnabledCrossRegionList {
	var returns DataSnowflakeAccountsAccountsParametersCortexEnabledCrossRegionList
	_jsii_.Get(
		j,
		"cortexEnabledCrossRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) CortexModelsAllowlist() DataSnowflakeAccountsAccountsParametersCortexModelsAllowlistStructList {
	var returns DataSnowflakeAccountsAccountsParametersCortexModelsAllowlistStructList
	_jsii_.Get(
		j,
		"cortexModelsAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) CsvTimestampFormat() DataSnowflakeAccountsAccountsParametersCsvTimestampFormatList {
	var returns DataSnowflakeAccountsAccountsParametersCsvTimestampFormatList
	_jsii_.Get(
		j,
		"csvTimestampFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DataMetricSchedule() DataSnowflakeAccountsAccountsParametersDataMetricScheduleList {
	var returns DataSnowflakeAccountsAccountsParametersDataMetricScheduleList
	_jsii_.Get(
		j,
		"dataMetricSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DataRetentionTimeInDays() DataSnowflakeAccountsAccountsParametersDataRetentionTimeInDaysList {
	var returns DataSnowflakeAccountsAccountsParametersDataRetentionTimeInDaysList
	_jsii_.Get(
		j,
		"dataRetentionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DateInputFormat() DataSnowflakeAccountsAccountsParametersDateInputFormatList {
	var returns DataSnowflakeAccountsAccountsParametersDateInputFormatList
	_jsii_.Get(
		j,
		"dateInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DateOutputFormat() DataSnowflakeAccountsAccountsParametersDateOutputFormatList {
	var returns DataSnowflakeAccountsAccountsParametersDateOutputFormatList
	_jsii_.Get(
		j,
		"dateOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DefaultDbtVersion() DataSnowflakeAccountsAccountsParametersDefaultDbtVersionList {
	var returns DataSnowflakeAccountsAccountsParametersDefaultDbtVersionList
	_jsii_.Get(
		j,
		"defaultDbtVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DefaultDdlCollation() DataSnowflakeAccountsAccountsParametersDefaultDdlCollationList {
	var returns DataSnowflakeAccountsAccountsParametersDefaultDdlCollationList
	_jsii_.Get(
		j,
		"defaultDdlCollation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DefaultNotebookComputePoolCpu() DataSnowflakeAccountsAccountsParametersDefaultNotebookComputePoolCpuList {
	var returns DataSnowflakeAccountsAccountsParametersDefaultNotebookComputePoolCpuList
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DefaultNotebookComputePoolGpu() DataSnowflakeAccountsAccountsParametersDefaultNotebookComputePoolGpuList {
	var returns DataSnowflakeAccountsAccountsParametersDefaultNotebookComputePoolGpuList
	_jsii_.Get(
		j,
		"defaultNotebookComputePoolGpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DefaultNullOrdering() DataSnowflakeAccountsAccountsParametersDefaultNullOrderingList {
	var returns DataSnowflakeAccountsAccountsParametersDefaultNullOrderingList
	_jsii_.Get(
		j,
		"defaultNullOrdering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DefaultStreamlitComputePool() DataSnowflakeAccountsAccountsParametersDefaultStreamlitComputePoolList {
	var returns DataSnowflakeAccountsAccountsParametersDefaultStreamlitComputePoolList
	_jsii_.Get(
		j,
		"defaultStreamlitComputePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DefaultStreamlitNotebookWarehouse() DataSnowflakeAccountsAccountsParametersDefaultStreamlitNotebookWarehouseList {
	var returns DataSnowflakeAccountsAccountsParametersDefaultStreamlitNotebookWarehouseList
	_jsii_.Get(
		j,
		"defaultStreamlitNotebookWarehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DisableUiDownloadButton() DataSnowflakeAccountsAccountsParametersDisableUiDownloadButtonList {
	var returns DataSnowflakeAccountsAccountsParametersDisableUiDownloadButtonList
	_jsii_.Get(
		j,
		"disableUiDownloadButton",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DisableUserPrivilegeGrants() DataSnowflakeAccountsAccountsParametersDisableUserPrivilegeGrantsList {
	var returns DataSnowflakeAccountsAccountsParametersDisableUserPrivilegeGrantsList
	_jsii_.Get(
		j,
		"disableUserPrivilegeGrants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) DisallowedSpcsWorkloadTypes() DataSnowflakeAccountsAccountsParametersDisallowedSpcsWorkloadTypesList {
	var returns DataSnowflakeAccountsAccountsParametersDisallowedSpcsWorkloadTypesList
	_jsii_.Get(
		j,
		"disallowedSpcsWorkloadTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableAutomaticSensitiveDataClassificationLog() DataSnowflakeAccountsAccountsParametersEnableAutomaticSensitiveDataClassificationLogList {
	var returns DataSnowflakeAccountsAccountsParametersEnableAutomaticSensitiveDataClassificationLogList
	_jsii_.Get(
		j,
		"enableAutomaticSensitiveDataClassificationLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableBudgetEventLogging() DataSnowflakeAccountsAccountsParametersEnableBudgetEventLoggingList {
	var returns DataSnowflakeAccountsAccountsParametersEnableBudgetEventLoggingList
	_jsii_.Get(
		j,
		"enableBudgetEventLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableConsoleOutput() DataSnowflakeAccountsAccountsParametersEnableConsoleOutputList {
	var returns DataSnowflakeAccountsAccountsParametersEnableConsoleOutputList
	_jsii_.Get(
		j,
		"enableConsoleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableCortexAnalyst() DataSnowflakeAccountsAccountsParametersEnableCortexAnalystList {
	var returns DataSnowflakeAccountsAccountsParametersEnableCortexAnalystList
	_jsii_.Get(
		j,
		"enableCortexAnalyst",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableDataCompaction() DataSnowflakeAccountsAccountsParametersEnableDataCompactionList {
	var returns DataSnowflakeAccountsAccountsParametersEnableDataCompactionList
	_jsii_.Get(
		j,
		"enableDataCompaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableEgressCostOptimizer() DataSnowflakeAccountsAccountsParametersEnableEgressCostOptimizerList {
	var returns DataSnowflakeAccountsAccountsParametersEnableEgressCostOptimizerList
	_jsii_.Get(
		j,
		"enableEgressCostOptimizer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableGetDdlUseDataTypeAlias() DataSnowflakeAccountsAccountsParametersEnableGetDdlUseDataTypeAliasList {
	var returns DataSnowflakeAccountsAccountsParametersEnableGetDdlUseDataTypeAliasList
	_jsii_.Get(
		j,
		"enableGetDdlUseDataTypeAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableIcebergMergeOnRead() DataSnowflakeAccountsAccountsParametersEnableIcebergMergeOnReadList {
	var returns DataSnowflakeAccountsAccountsParametersEnableIcebergMergeOnReadList
	_jsii_.Get(
		j,
		"enableIcebergMergeOnRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableIdentifierFirstLogin() DataSnowflakeAccountsAccountsParametersEnableIdentifierFirstLoginList {
	var returns DataSnowflakeAccountsAccountsParametersEnableIdentifierFirstLoginList
	_jsii_.Get(
		j,
		"enableIdentifierFirstLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableInternalStagesPrivatelink() DataSnowflakeAccountsAccountsParametersEnableInternalStagesPrivatelinkList {
	var returns DataSnowflakeAccountsAccountsParametersEnableInternalStagesPrivatelinkList
	_jsii_.Get(
		j,
		"enableInternalStagesPrivatelink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableNotebookCreationInPersonalDb() DataSnowflakeAccountsAccountsParametersEnableNotebookCreationInPersonalDbList {
	var returns DataSnowflakeAccountsAccountsParametersEnableNotebookCreationInPersonalDbList
	_jsii_.Get(
		j,
		"enableNotebookCreationInPersonalDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnablePerAccountAppServicePrivatelinkUrl() DataSnowflakeAccountsAccountsParametersEnablePerAccountAppServicePrivatelinkUrlList {
	var returns DataSnowflakeAccountsAccountsParametersEnablePerAccountAppServicePrivatelinkUrlList
	_jsii_.Get(
		j,
		"enablePerAccountAppServicePrivatelinkUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnablePersonalDatabase() DataSnowflakeAccountsAccountsParametersEnablePersonalDatabaseList {
	var returns DataSnowflakeAccountsAccountsParametersEnablePersonalDatabaseList
	_jsii_.Get(
		j,
		"enablePersonalDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableSpcsBlockStorageSnowflakeFullEncryptionEnforcement() DataSnowflakeAccountsAccountsParametersEnableSpcsBlockStorageSnowflakeFullEncryptionEnforcementList {
	var returns DataSnowflakeAccountsAccountsParametersEnableSpcsBlockStorageSnowflakeFullEncryptionEnforcementList
	_jsii_.Get(
		j,
		"enableSpcsBlockStorageSnowflakeFullEncryptionEnforcement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableTagPropagationEventLogging() DataSnowflakeAccountsAccountsParametersEnableTagPropagationEventLoggingList {
	var returns DataSnowflakeAccountsAccountsParametersEnableTagPropagationEventLoggingList
	_jsii_.Get(
		j,
		"enableTagPropagationEventLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableTriSecretAndRekeyOptOutForImageRepository() DataSnowflakeAccountsAccountsParametersEnableTriSecretAndRekeyOptOutForImageRepositoryList {
	var returns DataSnowflakeAccountsAccountsParametersEnableTriSecretAndRekeyOptOutForImageRepositoryList
	_jsii_.Get(
		j,
		"enableTriSecretAndRekeyOptOutForImageRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableTriSecretAndRekeyOptOutForSpcsBlockStorage() DataSnowflakeAccountsAccountsParametersEnableTriSecretAndRekeyOptOutForSpcsBlockStorageList {
	var returns DataSnowflakeAccountsAccountsParametersEnableTriSecretAndRekeyOptOutForSpcsBlockStorageList
	_jsii_.Get(
		j,
		"enableTriSecretAndRekeyOptOutForSpcsBlockStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableUnhandledExceptionsReporting() DataSnowflakeAccountsAccountsParametersEnableUnhandledExceptionsReportingList {
	var returns DataSnowflakeAccountsAccountsParametersEnableUnhandledExceptionsReportingList
	_jsii_.Get(
		j,
		"enableUnhandledExceptionsReporting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableUnloadPhysicalTypeOptimization() DataSnowflakeAccountsAccountsParametersEnableUnloadPhysicalTypeOptimizationList {
	var returns DataSnowflakeAccountsAccountsParametersEnableUnloadPhysicalTypeOptimizationList
	_jsii_.Get(
		j,
		"enableUnloadPhysicalTypeOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableUnredactedQuerySyntaxError() DataSnowflakeAccountsAccountsParametersEnableUnredactedQuerySyntaxErrorList {
	var returns DataSnowflakeAccountsAccountsParametersEnableUnredactedQuerySyntaxErrorList
	_jsii_.Get(
		j,
		"enableUnredactedQuerySyntaxError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnableUnredactedSecureObjectError() DataSnowflakeAccountsAccountsParametersEnableUnredactedSecureObjectErrorList {
	var returns DataSnowflakeAccountsAccountsParametersEnableUnredactedSecureObjectErrorList
	_jsii_.Get(
		j,
		"enableUnredactedSecureObjectError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EnforceNetworkRulesForInternalStages() DataSnowflakeAccountsAccountsParametersEnforceNetworkRulesForInternalStagesList {
	var returns DataSnowflakeAccountsAccountsParametersEnforceNetworkRulesForInternalStagesList
	_jsii_.Get(
		j,
		"enforceNetworkRulesForInternalStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ErrorOnNondeterministicMerge() DataSnowflakeAccountsAccountsParametersErrorOnNondeterministicMergeList {
	var returns DataSnowflakeAccountsAccountsParametersErrorOnNondeterministicMergeList
	_jsii_.Get(
		j,
		"errorOnNondeterministicMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ErrorOnNondeterministicUpdate() DataSnowflakeAccountsAccountsParametersErrorOnNondeterministicUpdateList {
	var returns DataSnowflakeAccountsAccountsParametersErrorOnNondeterministicUpdateList
	_jsii_.Get(
		j,
		"errorOnNondeterministicUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) EventTable() DataSnowflakeAccountsAccountsParametersEventTableList {
	var returns DataSnowflakeAccountsAccountsParametersEventTableList
	_jsii_.Get(
		j,
		"eventTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ExternalOauthAddPrivilegedRolesToBlockedList() DataSnowflakeAccountsAccountsParametersExternalOauthAddPrivilegedRolesToBlockedListStructList {
	var returns DataSnowflakeAccountsAccountsParametersExternalOauthAddPrivilegedRolesToBlockedListStructList
	_jsii_.Get(
		j,
		"externalOauthAddPrivilegedRolesToBlockedList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ExternalVolume() DataSnowflakeAccountsAccountsParametersExternalVolumeList {
	var returns DataSnowflakeAccountsAccountsParametersExternalVolumeList
	_jsii_.Get(
		j,
		"externalVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) GeographyOutputFormat() DataSnowflakeAccountsAccountsParametersGeographyOutputFormatList {
	var returns DataSnowflakeAccountsAccountsParametersGeographyOutputFormatList
	_jsii_.Get(
		j,
		"geographyOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) GeometryOutputFormat() DataSnowflakeAccountsAccountsParametersGeometryOutputFormatList {
	var returns DataSnowflakeAccountsAccountsParametersGeometryOutputFormatList
	_jsii_.Get(
		j,
		"geometryOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) HybridTableLockTimeout() DataSnowflakeAccountsAccountsParametersHybridTableLockTimeoutList {
	var returns DataSnowflakeAccountsAccountsParametersHybridTableLockTimeoutList
	_jsii_.Get(
		j,
		"hybridTableLockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) IcebergVersionDefault() DataSnowflakeAccountsAccountsParametersIcebergVersionDefaultList {
	var returns DataSnowflakeAccountsAccountsParametersIcebergVersionDefaultList
	_jsii_.Get(
		j,
		"icebergVersionDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) InitialReplicationSizeLimitInTb() DataSnowflakeAccountsAccountsParametersInitialReplicationSizeLimitInTbList {
	var returns DataSnowflakeAccountsAccountsParametersInitialReplicationSizeLimitInTbList
	_jsii_.Get(
		j,
		"initialReplicationSizeLimitInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) InternalValue() *DataSnowflakeAccountsAccountsParameters {
	var returns *DataSnowflakeAccountsAccountsParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) JdbcTreatDecimalAsInt() DataSnowflakeAccountsAccountsParametersJdbcTreatDecimalAsIntList {
	var returns DataSnowflakeAccountsAccountsParametersJdbcTreatDecimalAsIntList
	_jsii_.Get(
		j,
		"jdbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) JdbcTreatTimestampNtzAsUtc() DataSnowflakeAccountsAccountsParametersJdbcTreatTimestampNtzAsUtcList {
	var returns DataSnowflakeAccountsAccountsParametersJdbcTreatTimestampNtzAsUtcList
	_jsii_.Get(
		j,
		"jdbcTreatTimestampNtzAsUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) JdbcUseSessionTimezone() DataSnowflakeAccountsAccountsParametersJdbcUseSessionTimezoneList {
	var returns DataSnowflakeAccountsAccountsParametersJdbcUseSessionTimezoneList
	_jsii_.Get(
		j,
		"jdbcUseSessionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) JsonIndent() DataSnowflakeAccountsAccountsParametersJsonIndentList {
	var returns DataSnowflakeAccountsAccountsParametersJsonIndentList
	_jsii_.Get(
		j,
		"jsonIndent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) JsTreatIntegerAsBigint() DataSnowflakeAccountsAccountsParametersJsTreatIntegerAsBigintList {
	var returns DataSnowflakeAccountsAccountsParametersJsTreatIntegerAsBigintList
	_jsii_.Get(
		j,
		"jsTreatIntegerAsBigint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ListingAutoFulfillmentReplicationRefreshSchedule() DataSnowflakeAccountsAccountsParametersListingAutoFulfillmentReplicationRefreshScheduleList {
	var returns DataSnowflakeAccountsAccountsParametersListingAutoFulfillmentReplicationRefreshScheduleList
	_jsii_.Get(
		j,
		"listingAutoFulfillmentReplicationRefreshSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) LockTimeout() DataSnowflakeAccountsAccountsParametersLockTimeoutList {
	var returns DataSnowflakeAccountsAccountsParametersLockTimeoutList
	_jsii_.Get(
		j,
		"lockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) LogEventLevel() DataSnowflakeAccountsAccountsParametersLogEventLevelList {
	var returns DataSnowflakeAccountsAccountsParametersLogEventLevelList
	_jsii_.Get(
		j,
		"logEventLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) LogLevel() DataSnowflakeAccountsAccountsParametersLogLevelList {
	var returns DataSnowflakeAccountsAccountsParametersLogLevelList
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) MaxConcurrencyLevel() DataSnowflakeAccountsAccountsParametersMaxConcurrencyLevelList {
	var returns DataSnowflakeAccountsAccountsParametersMaxConcurrencyLevelList
	_jsii_.Get(
		j,
		"maxConcurrencyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) MaxDataExtensionTimeInDays() DataSnowflakeAccountsAccountsParametersMaxDataExtensionTimeInDaysList {
	var returns DataSnowflakeAccountsAccountsParametersMaxDataExtensionTimeInDaysList
	_jsii_.Get(
		j,
		"maxDataExtensionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) MetricLevel() DataSnowflakeAccountsAccountsParametersMetricLevelList {
	var returns DataSnowflakeAccountsAccountsParametersMetricLevelList
	_jsii_.Get(
		j,
		"metricLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) MinDataRetentionTimeInDays() DataSnowflakeAccountsAccountsParametersMinDataRetentionTimeInDaysList {
	var returns DataSnowflakeAccountsAccountsParametersMinDataRetentionTimeInDaysList
	_jsii_.Get(
		j,
		"minDataRetentionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) MultiStatementCount() DataSnowflakeAccountsAccountsParametersMultiStatementCountList {
	var returns DataSnowflakeAccountsAccountsParametersMultiStatementCountList
	_jsii_.Get(
		j,
		"multiStatementCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) NetworkPolicy() DataSnowflakeAccountsAccountsParametersNetworkPolicyList {
	var returns DataSnowflakeAccountsAccountsParametersNetworkPolicyList
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) NoorderSequenceAsDefault() DataSnowflakeAccountsAccountsParametersNoorderSequenceAsDefaultList {
	var returns DataSnowflakeAccountsAccountsParametersNoorderSequenceAsDefaultList
	_jsii_.Get(
		j,
		"noorderSequenceAsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) OauthAddPrivilegedRolesToBlockedList() DataSnowflakeAccountsAccountsParametersOauthAddPrivilegedRolesToBlockedListStructList {
	var returns DataSnowflakeAccountsAccountsParametersOauthAddPrivilegedRolesToBlockedListStructList
	_jsii_.Get(
		j,
		"oauthAddPrivilegedRolesToBlockedList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) OdbcTreatDecimalAsInt() DataSnowflakeAccountsAccountsParametersOdbcTreatDecimalAsIntList {
	var returns DataSnowflakeAccountsAccountsParametersOdbcTreatDecimalAsIntList
	_jsii_.Get(
		j,
		"odbcTreatDecimalAsInt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) PeriodicDataRekeying() DataSnowflakeAccountsAccountsParametersPeriodicDataRekeyingList {
	var returns DataSnowflakeAccountsAccountsParametersPeriodicDataRekeyingList
	_jsii_.Get(
		j,
		"periodicDataRekeying",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) PipeExecutionPaused() DataSnowflakeAccountsAccountsParametersPipeExecutionPausedList {
	var returns DataSnowflakeAccountsAccountsParametersPipeExecutionPausedList
	_jsii_.Get(
		j,
		"pipeExecutionPaused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) PreventLoadFromInlineUrl() DataSnowflakeAccountsAccountsParametersPreventLoadFromInlineUrlList {
	var returns DataSnowflakeAccountsAccountsParametersPreventLoadFromInlineUrlList
	_jsii_.Get(
		j,
		"preventLoadFromInlineUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) PreventUnloadToInlineUrl() DataSnowflakeAccountsAccountsParametersPreventUnloadToInlineUrlList {
	var returns DataSnowflakeAccountsAccountsParametersPreventUnloadToInlineUrlList
	_jsii_.Get(
		j,
		"preventUnloadToInlineUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) PreventUnloadToInternalStages() DataSnowflakeAccountsAccountsParametersPreventUnloadToInternalStagesList {
	var returns DataSnowflakeAccountsAccountsParametersPreventUnloadToInternalStagesList
	_jsii_.Get(
		j,
		"preventUnloadToInternalStages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) PythonProfilerModules() DataSnowflakeAccountsAccountsParametersPythonProfilerModulesList {
	var returns DataSnowflakeAccountsAccountsParametersPythonProfilerModulesList
	_jsii_.Get(
		j,
		"pythonProfilerModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) PythonProfilerTargetStage() DataSnowflakeAccountsAccountsParametersPythonProfilerTargetStageList {
	var returns DataSnowflakeAccountsAccountsParametersPythonProfilerTargetStageList
	_jsii_.Get(
		j,
		"pythonProfilerTargetStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) QueryTag() DataSnowflakeAccountsAccountsParametersQueryTagList {
	var returns DataSnowflakeAccountsAccountsParametersQueryTagList
	_jsii_.Get(
		j,
		"queryTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) QuotedIdentifiersIgnoreCase() DataSnowflakeAccountsAccountsParametersQuotedIdentifiersIgnoreCaseList {
	var returns DataSnowflakeAccountsAccountsParametersQuotedIdentifiersIgnoreCaseList
	_jsii_.Get(
		j,
		"quotedIdentifiersIgnoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ReadConsistencyMode() DataSnowflakeAccountsAccountsParametersReadConsistencyModeList {
	var returns DataSnowflakeAccountsAccountsParametersReadConsistencyModeList
	_jsii_.Get(
		j,
		"readConsistencyMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ReplaceInvalidCharacters() DataSnowflakeAccountsAccountsParametersReplaceInvalidCharactersList {
	var returns DataSnowflakeAccountsAccountsParametersReplaceInvalidCharactersList
	_jsii_.Get(
		j,
		"replaceInvalidCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) RequireStorageIntegrationForStageCreation() DataSnowflakeAccountsAccountsParametersRequireStorageIntegrationForStageCreationList {
	var returns DataSnowflakeAccountsAccountsParametersRequireStorageIntegrationForStageCreationList
	_jsii_.Get(
		j,
		"requireStorageIntegrationForStageCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) RequireStorageIntegrationForStageOperation() DataSnowflakeAccountsAccountsParametersRequireStorageIntegrationForStageOperationList {
	var returns DataSnowflakeAccountsAccountsParametersRequireStorageIntegrationForStageOperationList
	_jsii_.Get(
		j,
		"requireStorageIntegrationForStageOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) RowsPerResultset() DataSnowflakeAccountsAccountsParametersRowsPerResultsetList {
	var returns DataSnowflakeAccountsAccountsParametersRowsPerResultsetList
	_jsii_.Get(
		j,
		"rowsPerResultset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) RowTimestampDefault() DataSnowflakeAccountsAccountsParametersRowTimestampDefaultList {
	var returns DataSnowflakeAccountsAccountsParametersRowTimestampDefaultList
	_jsii_.Get(
		j,
		"rowTimestampDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) S3StageVpceDnsName() DataSnowflakeAccountsAccountsParametersS3StageVpceDnsNameList {
	var returns DataSnowflakeAccountsAccountsParametersS3StageVpceDnsNameList
	_jsii_.Get(
		j,
		"s3StageVpceDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) SearchPath() DataSnowflakeAccountsAccountsParametersSearchPathList {
	var returns DataSnowflakeAccountsAccountsParametersSearchPathList
	_jsii_.Get(
		j,
		"searchPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ServerlessTaskMaxStatementSize() DataSnowflakeAccountsAccountsParametersServerlessTaskMaxStatementSizeList {
	var returns DataSnowflakeAccountsAccountsParametersServerlessTaskMaxStatementSizeList
	_jsii_.Get(
		j,
		"serverlessTaskMaxStatementSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ServerlessTaskMinStatementSize() DataSnowflakeAccountsAccountsParametersServerlessTaskMinStatementSizeList {
	var returns DataSnowflakeAccountsAccountsParametersServerlessTaskMinStatementSizeList
	_jsii_.Get(
		j,
		"serverlessTaskMinStatementSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ShareRestrictions() DataSnowflakeAccountsAccountsParametersShareRestrictionsList {
	var returns DataSnowflakeAccountsAccountsParametersShareRestrictionsList
	_jsii_.Get(
		j,
		"shareRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) SimulatedDataSharingConsumer() DataSnowflakeAccountsAccountsParametersSimulatedDataSharingConsumerList {
	var returns DataSnowflakeAccountsAccountsParametersSimulatedDataSharingConsumerList
	_jsii_.Get(
		j,
		"simulatedDataSharingConsumer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) SqlTraceQueryText() DataSnowflakeAccountsAccountsParametersSqlTraceQueryTextList {
	var returns DataSnowflakeAccountsAccountsParametersSqlTraceQueryTextList
	_jsii_.Get(
		j,
		"sqlTraceQueryText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) SsoLoginPage() DataSnowflakeAccountsAccountsParametersSsoLoginPageList {
	var returns DataSnowflakeAccountsAccountsParametersSsoLoginPageList
	_jsii_.Get(
		j,
		"ssoLoginPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) StatementQueuedTimeoutInSeconds() DataSnowflakeAccountsAccountsParametersStatementQueuedTimeoutInSecondsList {
	var returns DataSnowflakeAccountsAccountsParametersStatementQueuedTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementQueuedTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) StatementTimeoutInSeconds() DataSnowflakeAccountsAccountsParametersStatementTimeoutInSecondsList {
	var returns DataSnowflakeAccountsAccountsParametersStatementTimeoutInSecondsList
	_jsii_.Get(
		j,
		"statementTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) StorageSerializationPolicy() DataSnowflakeAccountsAccountsParametersStorageSerializationPolicyList {
	var returns DataSnowflakeAccountsAccountsParametersStorageSerializationPolicyList
	_jsii_.Get(
		j,
		"storageSerializationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) StrictJsonOutput() DataSnowflakeAccountsAccountsParametersStrictJsonOutputList {
	var returns DataSnowflakeAccountsAccountsParametersStrictJsonOutputList
	_jsii_.Get(
		j,
		"strictJsonOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) SuspendTaskAfterNumFailures() DataSnowflakeAccountsAccountsParametersSuspendTaskAfterNumFailuresList {
	var returns DataSnowflakeAccountsAccountsParametersSuspendTaskAfterNumFailuresList
	_jsii_.Get(
		j,
		"suspendTaskAfterNumFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TaskAutoRetryAttempts() DataSnowflakeAccountsAccountsParametersTaskAutoRetryAttemptsList {
	var returns DataSnowflakeAccountsAccountsParametersTaskAutoRetryAttemptsList
	_jsii_.Get(
		j,
		"taskAutoRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TimeInputFormat() DataSnowflakeAccountsAccountsParametersTimeInputFormatList {
	var returns DataSnowflakeAccountsAccountsParametersTimeInputFormatList
	_jsii_.Get(
		j,
		"timeInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TimeOutputFormat() DataSnowflakeAccountsAccountsParametersTimeOutputFormatList {
	var returns DataSnowflakeAccountsAccountsParametersTimeOutputFormatList
	_jsii_.Get(
		j,
		"timeOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TimestampDayIsAlways24H() DataSnowflakeAccountsAccountsParametersTimestampDayIsAlways24HList {
	var returns DataSnowflakeAccountsAccountsParametersTimestampDayIsAlways24HList
	_jsii_.Get(
		j,
		"timestampDayIsAlways24H",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TimestampInputFormat() DataSnowflakeAccountsAccountsParametersTimestampInputFormatList {
	var returns DataSnowflakeAccountsAccountsParametersTimestampInputFormatList
	_jsii_.Get(
		j,
		"timestampInputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TimestampLtzOutputFormat() DataSnowflakeAccountsAccountsParametersTimestampLtzOutputFormatList {
	var returns DataSnowflakeAccountsAccountsParametersTimestampLtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampLtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TimestampNtzOutputFormat() DataSnowflakeAccountsAccountsParametersTimestampNtzOutputFormatList {
	var returns DataSnowflakeAccountsAccountsParametersTimestampNtzOutputFormatList
	_jsii_.Get(
		j,
		"timestampNtzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TimestampOutputFormat() DataSnowflakeAccountsAccountsParametersTimestampOutputFormatList {
	var returns DataSnowflakeAccountsAccountsParametersTimestampOutputFormatList
	_jsii_.Get(
		j,
		"timestampOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TimestampTypeMapping() DataSnowflakeAccountsAccountsParametersTimestampTypeMappingList {
	var returns DataSnowflakeAccountsAccountsParametersTimestampTypeMappingList
	_jsii_.Get(
		j,
		"timestampTypeMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TimestampTzOutputFormat() DataSnowflakeAccountsAccountsParametersTimestampTzOutputFormatList {
	var returns DataSnowflakeAccountsAccountsParametersTimestampTzOutputFormatList
	_jsii_.Get(
		j,
		"timestampTzOutputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) Timezone() DataSnowflakeAccountsAccountsParametersTimezoneList {
	var returns DataSnowflakeAccountsAccountsParametersTimezoneList
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TraceLevel() DataSnowflakeAccountsAccountsParametersTraceLevelList {
	var returns DataSnowflakeAccountsAccountsParametersTraceLevelList
	_jsii_.Get(
		j,
		"traceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TransactionAbortOnError() DataSnowflakeAccountsAccountsParametersTransactionAbortOnErrorList {
	var returns DataSnowflakeAccountsAccountsParametersTransactionAbortOnErrorList
	_jsii_.Get(
		j,
		"transactionAbortOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TransactionDefaultIsolationLevel() DataSnowflakeAccountsAccountsParametersTransactionDefaultIsolationLevelList {
	var returns DataSnowflakeAccountsAccountsParametersTransactionDefaultIsolationLevelList
	_jsii_.Get(
		j,
		"transactionDefaultIsolationLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) TwoDigitCenturyStart() DataSnowflakeAccountsAccountsParametersTwoDigitCenturyStartList {
	var returns DataSnowflakeAccountsAccountsParametersTwoDigitCenturyStartList
	_jsii_.Get(
		j,
		"twoDigitCenturyStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) UnsupportedDdlAction() DataSnowflakeAccountsAccountsParametersUnsupportedDdlActionList {
	var returns DataSnowflakeAccountsAccountsParametersUnsupportedDdlActionList
	_jsii_.Get(
		j,
		"unsupportedDdlAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) UseCachedResult() DataSnowflakeAccountsAccountsParametersUseCachedResultList {
	var returns DataSnowflakeAccountsAccountsParametersUseCachedResultList
	_jsii_.Get(
		j,
		"useCachedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) UserTaskManagedInitialWarehouseSize() DataSnowflakeAccountsAccountsParametersUserTaskManagedInitialWarehouseSizeList {
	var returns DataSnowflakeAccountsAccountsParametersUserTaskManagedInitialWarehouseSizeList
	_jsii_.Get(
		j,
		"userTaskManagedInitialWarehouseSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) UserTaskMinimumTriggerIntervalInSeconds() DataSnowflakeAccountsAccountsParametersUserTaskMinimumTriggerIntervalInSecondsList {
	var returns DataSnowflakeAccountsAccountsParametersUserTaskMinimumTriggerIntervalInSecondsList
	_jsii_.Get(
		j,
		"userTaskMinimumTriggerIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) UserTaskTimeoutMs() DataSnowflakeAccountsAccountsParametersUserTaskTimeoutMsList {
	var returns DataSnowflakeAccountsAccountsParametersUserTaskTimeoutMsList
	_jsii_.Get(
		j,
		"userTaskTimeoutMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) UseWorkspacesForSql() DataSnowflakeAccountsAccountsParametersUseWorkspacesForSqlList {
	var returns DataSnowflakeAccountsAccountsParametersUseWorkspacesForSqlList
	_jsii_.Get(
		j,
		"useWorkspacesForSql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) WeekOfYearPolicy() DataSnowflakeAccountsAccountsParametersWeekOfYearPolicyList {
	var returns DataSnowflakeAccountsAccountsParametersWeekOfYearPolicyList
	_jsii_.Get(
		j,
		"weekOfYearPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) WeekStart() DataSnowflakeAccountsAccountsParametersWeekStartList {
	var returns DataSnowflakeAccountsAccountsParametersWeekStartList
	_jsii_.Get(
		j,
		"weekStart",
		&returns,
	)
	return returns
}


func NewDataSnowflakeAccountsAccountsParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataSnowflakeAccountsAccountsParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataSnowflakeAccountsAccountsParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeAccounts.DataSnowflakeAccountsAccountsParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataSnowflakeAccountsAccountsParametersOutputReference_Override(d DataSnowflakeAccountsAccountsParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-snowflake.dataSnowflakeAccounts.DataSnowflakeAccountsAccountsParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference)SetInternalValue(val *DataSnowflakeAccountsAccountsParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataSnowflakeAccountsAccountsParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

