// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package stageexternals3

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3",
		reflect.TypeOf((*StageExternalS3)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccessPointArn", GoGetter: "AwsAccessPointArn"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccessPointArnInput", GoGetter: "AwsAccessPointArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cloud", GoGetter: "Cloud"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsInput", GoGetter: "CredentialsInput"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInput", GoGetter: "DatabaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "describeOutput", GoGetter: "DescribeOutput"},
			_jsii_.MemberProperty{JsiiProperty: "directory", GoGetter: "Directory"},
			_jsii_.MemberProperty{JsiiProperty: "directoryInput", GoGetter: "DirectoryInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryption", GoGetter: "Encryption"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionInput", GoGetter: "EncryptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileFormat", GoGetter: "FileFormat"},
			_jsii_.MemberProperty{JsiiProperty: "fileFormatInput", GoGetter: "FileFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "fullyQualifiedName", GoGetter: "FullyQualifiedName"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putCredentials", GoMethod: "PutCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "putDirectory", GoMethod: "PutDirectory"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryption", GoMethod: "PutEncryption"},
			_jsii_.MemberMethod{JsiiMethod: "putFileFormat", GoMethod: "PutFileFormat"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsAccessPointArn", GoMethod: "ResetAwsAccessPointArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetCredentials", GoMethod: "ResetCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "resetDirectory", GoMethod: "ResetDirectory"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryption", GoMethod: "ResetEncryption"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileFormat", GoMethod: "ResetFileFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageIntegration", GoMethod: "ResetStorageIntegration"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsePrivatelinkEndpoint", GoMethod: "ResetUsePrivatelinkEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "showOutput", GoGetter: "ShowOutput"},
			_jsii_.MemberProperty{JsiiProperty: "stageType", GoGetter: "StageType"},
			_jsii_.MemberProperty{JsiiProperty: "storageIntegration", GoGetter: "StorageIntegration"},
			_jsii_.MemberProperty{JsiiProperty: "storageIntegrationInput", GoGetter: "StorageIntegrationInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "usePrivatelinkEndpoint", GoGetter: "UsePrivatelinkEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "usePrivatelinkEndpointInput", GoGetter: "UsePrivatelinkEndpointInput"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3Config",
		reflect.TypeOf((*StageExternalS3Config)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3Credentials",
		reflect.TypeOf((*StageExternalS3Credentials)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3CredentialsOutputReference",
		reflect.TypeOf((*StageExternalS3CredentialsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awsKeyId", GoGetter: "AwsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "awsKeyIdInput", GoGetter: "AwsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsRole", GoGetter: "AwsRole"},
			_jsii_.MemberProperty{JsiiProperty: "awsRoleInput", GoGetter: "AwsRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSecretKey", GoGetter: "AwsSecretKey"},
			_jsii_.MemberProperty{JsiiProperty: "awsSecretKeyInput", GoGetter: "AwsSecretKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsToken", GoGetter: "AwsToken"},
			_jsii_.MemberProperty{JsiiProperty: "awsTokenInput", GoGetter: "AwsTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAwsKeyId", GoMethod: "ResetAwsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsRole", GoMethod: "ResetAwsRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSecretKey", GoMethod: "ResetAwsSecretKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsToken", GoMethod: "ResetAwsToken"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3CredentialsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutput",
		reflect.TypeOf((*StageExternalS3DescribeOutput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputDirectoryTable",
		reflect.TypeOf((*StageExternalS3DescribeOutputDirectoryTable)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputDirectoryTableList",
		reflect.TypeOf((*StageExternalS3DescribeOutputDirectoryTableList)(nil)).Elem(),
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
			j := jsiiProxy_StageExternalS3DescribeOutputDirectoryTableList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputDirectoryTableOutputReference",
		reflect.TypeOf((*StageExternalS3DescribeOutputDirectoryTableOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoRefresh", GoGetter: "AutoRefresh"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enable", GoGetter: "Enable"},
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
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3DescribeOutputDirectoryTableOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormat",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormat)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatAvro",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatAvro)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatAvroList",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatAvroList)(nil)).Elem(),
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
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatAvroList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatAvroOutputReference",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatAvroOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "nullIf", GoGetter: "NullIf"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharacters", GoGetter: "ReplaceInvalidCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpace", GoGetter: "TrimSpace"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatAvroOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatCsv",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatCsv)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatCsvList",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatCsvList)(nil)).Elem(),
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
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatCsvList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatCsvOutputReference",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatCsvOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "binaryFormat", GoGetter: "BinaryFormat"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dateFormat", GoGetter: "DateFormat"},
			_jsii_.MemberProperty{JsiiProperty: "emptyFieldAsNull", GoGetter: "EmptyFieldAsNull"},
			_jsii_.MemberProperty{JsiiProperty: "encoding", GoGetter: "Encoding"},
			_jsii_.MemberProperty{JsiiProperty: "errorOnColumnCountMismatch", GoGetter: "ErrorOnColumnCountMismatch"},
			_jsii_.MemberProperty{JsiiProperty: "escape", GoGetter: "Escape"},
			_jsii_.MemberProperty{JsiiProperty: "escapeUnenclosedField", GoGetter: "EscapeUnenclosedField"},
			_jsii_.MemberProperty{JsiiProperty: "fieldDelimiter", GoGetter: "FieldDelimiter"},
			_jsii_.MemberProperty{JsiiProperty: "fieldOptionallyEnclosedBy", GoGetter: "FieldOptionallyEnclosedBy"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtension", GoGetter: "FileExtension"},
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
			_jsii_.MemberProperty{JsiiProperty: "multiLine", GoGetter: "MultiLine"},
			_jsii_.MemberProperty{JsiiProperty: "nullIf", GoGetter: "NullIf"},
			_jsii_.MemberProperty{JsiiProperty: "parseHeader", GoGetter: "ParseHeader"},
			_jsii_.MemberProperty{JsiiProperty: "recordDelimiter", GoGetter: "RecordDelimiter"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharacters", GoGetter: "ReplaceInvalidCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "skipBlankLines", GoGetter: "SkipBlankLines"},
			_jsii_.MemberProperty{JsiiProperty: "skipByteOrderMark", GoGetter: "SkipByteOrderMark"},
			_jsii_.MemberProperty{JsiiProperty: "skipHeader", GoGetter: "SkipHeader"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeFormat", GoGetter: "TimeFormat"},
			_jsii_.MemberProperty{JsiiProperty: "timestampFormat", GoGetter: "TimestampFormat"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpace", GoGetter: "TrimSpace"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "validateUtf8", GoGetter: "ValidateUtf8"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatCsvOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatJson",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatJson)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatJsonList",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatJsonList)(nil)).Elem(),
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
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatJsonList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatJsonOutputReference",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatJsonOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowDuplicate", GoGetter: "AllowDuplicate"},
			_jsii_.MemberProperty{JsiiProperty: "binaryFormat", GoGetter: "BinaryFormat"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dateFormat", GoGetter: "DateFormat"},
			_jsii_.MemberProperty{JsiiProperty: "enableOctal", GoGetter: "EnableOctal"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtension", GoGetter: "FileExtension"},
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
			_jsii_.MemberProperty{JsiiProperty: "ignoreUtf8Errors", GoGetter: "IgnoreUtf8Errors"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "multiLine", GoGetter: "MultiLine"},
			_jsii_.MemberProperty{JsiiProperty: "nullIf", GoGetter: "NullIf"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharacters", GoGetter: "ReplaceInvalidCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "skipByteOrderMark", GoGetter: "SkipByteOrderMark"},
			_jsii_.MemberProperty{JsiiProperty: "stripNullValues", GoGetter: "StripNullValues"},
			_jsii_.MemberProperty{JsiiProperty: "stripOuterArray", GoGetter: "StripOuterArray"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeFormat", GoGetter: "TimeFormat"},
			_jsii_.MemberProperty{JsiiProperty: "timestampFormat", GoGetter: "TimestampFormat"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpace", GoGetter: "TrimSpace"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatJsonOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatList",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatList)(nil)).Elem(),
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
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatOrc",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatOrc)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatOrcList",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatOrcList)(nil)).Elem(),
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
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatOrcList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatOrcOutputReference",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatOrcOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "nullIf", GoGetter: "NullIf"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharacters", GoGetter: "ReplaceInvalidCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpace", GoGetter: "TrimSpace"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatOrcOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatOutputReference",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "avro", GoGetter: "Avro"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "csv", GoGetter: "Csv"},
			_jsii_.MemberProperty{JsiiProperty: "formatName", GoGetter: "FormatName"},
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
			_jsii_.MemberProperty{JsiiProperty: "json", GoGetter: "Json"},
			_jsii_.MemberProperty{JsiiProperty: "orc", GoGetter: "Orc"},
			_jsii_.MemberProperty{JsiiProperty: "parquet", GoGetter: "Parquet"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "xml", GoGetter: "Xml"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatParquet",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatParquet)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatParquetList",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatParquetList)(nil)).Elem(),
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
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatParquetList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatParquetOutputReference",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatParquetOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "binaryAsText", GoGetter: "BinaryAsText"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "nullIf", GoGetter: "NullIf"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharacters", GoGetter: "ReplaceInvalidCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpace", GoGetter: "TrimSpace"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "useLogicalType", GoGetter: "UseLogicalType"},
			_jsii_.MemberProperty{JsiiProperty: "useVectorizedScanner", GoGetter: "UseVectorizedScanner"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatParquetOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatXml",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatXml)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatXmlList",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatXmlList)(nil)).Elem(),
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
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatXmlList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputFileFormatXmlOutputReference",
		reflect.TypeOf((*StageExternalS3DescribeOutputFileFormatXmlOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disableAutoConvert", GoGetter: "DisableAutoConvert"},
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
			_jsii_.MemberProperty{JsiiProperty: "ignoreUtf8Errors", GoGetter: "IgnoreUtf8Errors"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "preserveSpace", GoGetter: "PreserveSpace"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharacters", GoGetter: "ReplaceInvalidCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "skipByteOrderMark", GoGetter: "SkipByteOrderMark"},
			_jsii_.MemberProperty{JsiiProperty: "stripOuterElement", GoGetter: "StripOuterElement"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3DescribeOutputFileFormatXmlOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputList",
		reflect.TypeOf((*StageExternalS3DescribeOutputList)(nil)).Elem(),
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
			j := jsiiProxy_StageExternalS3DescribeOutputList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputLocation",
		reflect.TypeOf((*StageExternalS3DescribeOutputLocation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputLocationList",
		reflect.TypeOf((*StageExternalS3DescribeOutputLocationList)(nil)).Elem(),
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
			j := jsiiProxy_StageExternalS3DescribeOutputLocationList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputLocationOutputReference",
		reflect.TypeOf((*StageExternalS3DescribeOutputLocationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awsAccessPointArn", GoGetter: "AwsAccessPointArn"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3DescribeOutputLocationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputOutputReference",
		reflect.TypeOf((*StageExternalS3DescribeOutputOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "directoryTable", GoGetter: "DirectoryTable"},
			_jsii_.MemberProperty{JsiiProperty: "fileFormat", GoGetter: "FileFormat"},
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
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "privatelink", GoGetter: "Privatelink"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3DescribeOutputOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputPrivatelink",
		reflect.TypeOf((*StageExternalS3DescribeOutputPrivatelink)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputPrivatelinkList",
		reflect.TypeOf((*StageExternalS3DescribeOutputPrivatelinkList)(nil)).Elem(),
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
			j := jsiiProxy_StageExternalS3DescribeOutputPrivatelinkList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DescribeOutputPrivatelinkOutputReference",
		reflect.TypeOf((*StageExternalS3DescribeOutputPrivatelinkOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "usePrivatelinkEndpoint", GoGetter: "UsePrivatelinkEndpoint"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3DescribeOutputPrivatelinkOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3Directory",
		reflect.TypeOf((*StageExternalS3Directory)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3DirectoryOutputReference",
		reflect.TypeOf((*StageExternalS3DirectoryOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoRefresh", GoGetter: "AutoRefresh"},
			_jsii_.MemberProperty{JsiiProperty: "autoRefreshInput", GoGetter: "AutoRefreshInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enable", GoGetter: "Enable"},
			_jsii_.MemberProperty{JsiiProperty: "enableInput", GoGetter: "EnableInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "refreshOnCreate", GoGetter: "RefreshOnCreate"},
			_jsii_.MemberProperty{JsiiProperty: "refreshOnCreateInput", GoGetter: "RefreshOnCreateInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRefresh", GoMethod: "ResetAutoRefresh"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshOnCreate", GoMethod: "ResetRefreshOnCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3DirectoryOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3Encryption",
		reflect.TypeOf((*StageExternalS3Encryption)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3EncryptionAwsCse",
		reflect.TypeOf((*StageExternalS3EncryptionAwsCse)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3EncryptionAwsCseOutputReference",
		reflect.TypeOf((*StageExternalS3EncryptionAwsCseOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "masterKey", GoGetter: "MasterKey"},
			_jsii_.MemberProperty{JsiiProperty: "masterKeyInput", GoGetter: "MasterKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3EncryptionAwsCseOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3EncryptionAwsSseKms",
		reflect.TypeOf((*StageExternalS3EncryptionAwsSseKms)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3EncryptionAwsSseKmsOutputReference",
		reflect.TypeOf((*StageExternalS3EncryptionAwsSseKmsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3EncryptionAwsSseKmsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3EncryptionAwsSseS3",
		reflect.TypeOf((*StageExternalS3EncryptionAwsSseS3)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3EncryptionAwsSseS3OutputReference",
		reflect.TypeOf((*StageExternalS3EncryptionAwsSseS3OutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3EncryptionAwsSseS3OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3EncryptionNone",
		reflect.TypeOf((*StageExternalS3EncryptionNone)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3EncryptionNoneOutputReference",
		reflect.TypeOf((*StageExternalS3EncryptionNoneOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3EncryptionNoneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3EncryptionOutputReference",
		reflect.TypeOf((*StageExternalS3EncryptionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awsCse", GoGetter: "AwsCse"},
			_jsii_.MemberProperty{JsiiProperty: "awsCseInput", GoGetter: "AwsCseInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSseKms", GoGetter: "AwsSseKms"},
			_jsii_.MemberProperty{JsiiProperty: "awsSseKmsInput", GoGetter: "AwsSseKmsInput"},
			_jsii_.MemberProperty{JsiiProperty: "awsSseS3", GoGetter: "AwsSseS3"},
			_jsii_.MemberProperty{JsiiProperty: "awsSseS3Input", GoGetter: "AwsSseS3Input"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "none", GoGetter: "None"},
			_jsii_.MemberProperty{JsiiProperty: "noneInput", GoGetter: "NoneInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAwsCse", GoMethod: "PutAwsCse"},
			_jsii_.MemberMethod{JsiiMethod: "putAwsSseKms", GoMethod: "PutAwsSseKms"},
			_jsii_.MemberMethod{JsiiMethod: "putAwsSseS3", GoMethod: "PutAwsSseS3"},
			_jsii_.MemberMethod{JsiiMethod: "putNone", GoMethod: "PutNone"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsCse", GoMethod: "ResetAwsCse"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSseKms", GoMethod: "ResetAwsSseKms"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsSseS3", GoMethod: "ResetAwsSseS3"},
			_jsii_.MemberMethod{JsiiMethod: "resetNone", GoMethod: "ResetNone"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3EncryptionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormat",
		reflect.TypeOf((*StageExternalS3FileFormat)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatAvro",
		reflect.TypeOf((*StageExternalS3FileFormatAvro)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatAvroOutputReference",
		reflect.TypeOf((*StageExternalS3FileFormatAvroOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberProperty{JsiiProperty: "compressionInput", GoGetter: "CompressionInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "nullIf", GoGetter: "NullIf"},
			_jsii_.MemberProperty{JsiiProperty: "nullIfInput", GoGetter: "NullIfInput"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharacters", GoGetter: "ReplaceInvalidCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharactersInput", GoGetter: "ReplaceInvalidCharactersInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompression", GoMethod: "ResetCompression"},
			_jsii_.MemberMethod{JsiiMethod: "resetNullIf", GoMethod: "ResetNullIf"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplaceInvalidCharacters", GoMethod: "ResetReplaceInvalidCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrimSpace", GoMethod: "ResetTrimSpace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpace", GoGetter: "TrimSpace"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpaceInput", GoGetter: "TrimSpaceInput"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3FileFormatAvroOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatCsv",
		reflect.TypeOf((*StageExternalS3FileFormatCsv)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatCsvOutputReference",
		reflect.TypeOf((*StageExternalS3FileFormatCsvOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "binaryFormat", GoGetter: "BinaryFormat"},
			_jsii_.MemberProperty{JsiiProperty: "binaryFormatInput", GoGetter: "BinaryFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberProperty{JsiiProperty: "compressionInput", GoGetter: "CompressionInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dateFormat", GoGetter: "DateFormat"},
			_jsii_.MemberProperty{JsiiProperty: "dateFormatInput", GoGetter: "DateFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "emptyFieldAsNull", GoGetter: "EmptyFieldAsNull"},
			_jsii_.MemberProperty{JsiiProperty: "emptyFieldAsNullInput", GoGetter: "EmptyFieldAsNullInput"},
			_jsii_.MemberProperty{JsiiProperty: "encoding", GoGetter: "Encoding"},
			_jsii_.MemberProperty{JsiiProperty: "encodingInput", GoGetter: "EncodingInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorOnColumnCountMismatch", GoGetter: "ErrorOnColumnCountMismatch"},
			_jsii_.MemberProperty{JsiiProperty: "errorOnColumnCountMismatchInput", GoGetter: "ErrorOnColumnCountMismatchInput"},
			_jsii_.MemberProperty{JsiiProperty: "escape", GoGetter: "Escape"},
			_jsii_.MemberProperty{JsiiProperty: "escapeInput", GoGetter: "EscapeInput"},
			_jsii_.MemberProperty{JsiiProperty: "escapeUnenclosedField", GoGetter: "EscapeUnenclosedField"},
			_jsii_.MemberProperty{JsiiProperty: "escapeUnenclosedFieldInput", GoGetter: "EscapeUnenclosedFieldInput"},
			_jsii_.MemberProperty{JsiiProperty: "fieldDelimiter", GoGetter: "FieldDelimiter"},
			_jsii_.MemberProperty{JsiiProperty: "fieldDelimiterInput", GoGetter: "FieldDelimiterInput"},
			_jsii_.MemberProperty{JsiiProperty: "fieldOptionallyEnclosedBy", GoGetter: "FieldOptionallyEnclosedBy"},
			_jsii_.MemberProperty{JsiiProperty: "fieldOptionallyEnclosedByInput", GoGetter: "FieldOptionallyEnclosedByInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtension", GoGetter: "FileExtension"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionInput", GoGetter: "FileExtensionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "multiLine", GoGetter: "MultiLine"},
			_jsii_.MemberProperty{JsiiProperty: "multiLineInput", GoGetter: "MultiLineInput"},
			_jsii_.MemberProperty{JsiiProperty: "nullIf", GoGetter: "NullIf"},
			_jsii_.MemberProperty{JsiiProperty: "nullIfInput", GoGetter: "NullIfInput"},
			_jsii_.MemberProperty{JsiiProperty: "parseHeader", GoGetter: "ParseHeader"},
			_jsii_.MemberProperty{JsiiProperty: "parseHeaderInput", GoGetter: "ParseHeaderInput"},
			_jsii_.MemberProperty{JsiiProperty: "recordDelimiter", GoGetter: "RecordDelimiter"},
			_jsii_.MemberProperty{JsiiProperty: "recordDelimiterInput", GoGetter: "RecordDelimiterInput"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharacters", GoGetter: "ReplaceInvalidCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharactersInput", GoGetter: "ReplaceInvalidCharactersInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBinaryFormat", GoMethod: "ResetBinaryFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompression", GoMethod: "ResetCompression"},
			_jsii_.MemberMethod{JsiiMethod: "resetDateFormat", GoMethod: "ResetDateFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmptyFieldAsNull", GoMethod: "ResetEmptyFieldAsNull"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncoding", GoMethod: "ResetEncoding"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorOnColumnCountMismatch", GoMethod: "ResetErrorOnColumnCountMismatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetEscape", GoMethod: "ResetEscape"},
			_jsii_.MemberMethod{JsiiMethod: "resetEscapeUnenclosedField", GoMethod: "ResetEscapeUnenclosedField"},
			_jsii_.MemberMethod{JsiiMethod: "resetFieldDelimiter", GoMethod: "ResetFieldDelimiter"},
			_jsii_.MemberMethod{JsiiMethod: "resetFieldOptionallyEnclosedBy", GoMethod: "ResetFieldOptionallyEnclosedBy"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileExtension", GoMethod: "ResetFileExtension"},
			_jsii_.MemberMethod{JsiiMethod: "resetMultiLine", GoMethod: "ResetMultiLine"},
			_jsii_.MemberMethod{JsiiMethod: "resetNullIf", GoMethod: "ResetNullIf"},
			_jsii_.MemberMethod{JsiiMethod: "resetParseHeader", GoMethod: "ResetParseHeader"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecordDelimiter", GoMethod: "ResetRecordDelimiter"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplaceInvalidCharacters", GoMethod: "ResetReplaceInvalidCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipBlankLines", GoMethod: "ResetSkipBlankLines"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipByteOrderMark", GoMethod: "ResetSkipByteOrderMark"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipHeader", GoMethod: "ResetSkipHeader"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeFormat", GoMethod: "ResetTimeFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimestampFormat", GoMethod: "ResetTimestampFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrimSpace", GoMethod: "ResetTrimSpace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "skipBlankLines", GoGetter: "SkipBlankLines"},
			_jsii_.MemberProperty{JsiiProperty: "skipBlankLinesInput", GoGetter: "SkipBlankLinesInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipByteOrderMark", GoGetter: "SkipByteOrderMark"},
			_jsii_.MemberProperty{JsiiProperty: "skipByteOrderMarkInput", GoGetter: "SkipByteOrderMarkInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipHeader", GoGetter: "SkipHeader"},
			_jsii_.MemberProperty{JsiiProperty: "skipHeaderInput", GoGetter: "SkipHeaderInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeFormat", GoGetter: "TimeFormat"},
			_jsii_.MemberProperty{JsiiProperty: "timeFormatInput", GoGetter: "TimeFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "timestampFormat", GoGetter: "TimestampFormat"},
			_jsii_.MemberProperty{JsiiProperty: "timestampFormatInput", GoGetter: "TimestampFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpace", GoGetter: "TrimSpace"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpaceInput", GoGetter: "TrimSpaceInput"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3FileFormatCsvOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatJson",
		reflect.TypeOf((*StageExternalS3FileFormatJson)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatJsonOutputReference",
		reflect.TypeOf((*StageExternalS3FileFormatJsonOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowDuplicate", GoGetter: "AllowDuplicate"},
			_jsii_.MemberProperty{JsiiProperty: "allowDuplicateInput", GoGetter: "AllowDuplicateInput"},
			_jsii_.MemberProperty{JsiiProperty: "binaryFormat", GoGetter: "BinaryFormat"},
			_jsii_.MemberProperty{JsiiProperty: "binaryFormatInput", GoGetter: "BinaryFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberProperty{JsiiProperty: "compressionInput", GoGetter: "CompressionInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dateFormat", GoGetter: "DateFormat"},
			_jsii_.MemberProperty{JsiiProperty: "dateFormatInput", GoGetter: "DateFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableOctal", GoGetter: "EnableOctal"},
			_jsii_.MemberProperty{JsiiProperty: "enableOctalInput", GoGetter: "EnableOctalInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtension", GoGetter: "FileExtension"},
			_jsii_.MemberProperty{JsiiProperty: "fileExtensionInput", GoGetter: "FileExtensionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "ignoreUtf8Errors", GoGetter: "IgnoreUtf8Errors"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreUtf8ErrorsInput", GoGetter: "IgnoreUtf8ErrorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "multiLine", GoGetter: "MultiLine"},
			_jsii_.MemberProperty{JsiiProperty: "multiLineInput", GoGetter: "MultiLineInput"},
			_jsii_.MemberProperty{JsiiProperty: "nullIf", GoGetter: "NullIf"},
			_jsii_.MemberProperty{JsiiProperty: "nullIfInput", GoGetter: "NullIfInput"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharacters", GoGetter: "ReplaceInvalidCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharactersInput", GoGetter: "ReplaceInvalidCharactersInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowDuplicate", GoMethod: "ResetAllowDuplicate"},
			_jsii_.MemberMethod{JsiiMethod: "resetBinaryFormat", GoMethod: "ResetBinaryFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompression", GoMethod: "ResetCompression"},
			_jsii_.MemberMethod{JsiiMethod: "resetDateFormat", GoMethod: "ResetDateFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableOctal", GoMethod: "ResetEnableOctal"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileExtension", GoMethod: "ResetFileExtension"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreUtf8Errors", GoMethod: "ResetIgnoreUtf8Errors"},
			_jsii_.MemberMethod{JsiiMethod: "resetMultiLine", GoMethod: "ResetMultiLine"},
			_jsii_.MemberMethod{JsiiMethod: "resetNullIf", GoMethod: "ResetNullIf"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplaceInvalidCharacters", GoMethod: "ResetReplaceInvalidCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipByteOrderMark", GoMethod: "ResetSkipByteOrderMark"},
			_jsii_.MemberMethod{JsiiMethod: "resetStripNullValues", GoMethod: "ResetStripNullValues"},
			_jsii_.MemberMethod{JsiiMethod: "resetStripOuterArray", GoMethod: "ResetStripOuterArray"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeFormat", GoMethod: "ResetTimeFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimestampFormat", GoMethod: "ResetTimestampFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrimSpace", GoMethod: "ResetTrimSpace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "skipByteOrderMark", GoGetter: "SkipByteOrderMark"},
			_jsii_.MemberProperty{JsiiProperty: "skipByteOrderMarkInput", GoGetter: "SkipByteOrderMarkInput"},
			_jsii_.MemberProperty{JsiiProperty: "stripNullValues", GoGetter: "StripNullValues"},
			_jsii_.MemberProperty{JsiiProperty: "stripNullValuesInput", GoGetter: "StripNullValuesInput"},
			_jsii_.MemberProperty{JsiiProperty: "stripOuterArray", GoGetter: "StripOuterArray"},
			_jsii_.MemberProperty{JsiiProperty: "stripOuterArrayInput", GoGetter: "StripOuterArrayInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeFormat", GoGetter: "TimeFormat"},
			_jsii_.MemberProperty{JsiiProperty: "timeFormatInput", GoGetter: "TimeFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "timestampFormat", GoGetter: "TimestampFormat"},
			_jsii_.MemberProperty{JsiiProperty: "timestampFormatInput", GoGetter: "TimestampFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpace", GoGetter: "TrimSpace"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpaceInput", GoGetter: "TrimSpaceInput"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3FileFormatJsonOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatOrc",
		reflect.TypeOf((*StageExternalS3FileFormatOrc)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatOrcOutputReference",
		reflect.TypeOf((*StageExternalS3FileFormatOrcOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "nullIf", GoGetter: "NullIf"},
			_jsii_.MemberProperty{JsiiProperty: "nullIfInput", GoGetter: "NullIfInput"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharacters", GoGetter: "ReplaceInvalidCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharactersInput", GoGetter: "ReplaceInvalidCharactersInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNullIf", GoMethod: "ResetNullIf"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplaceInvalidCharacters", GoMethod: "ResetReplaceInvalidCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrimSpace", GoMethod: "ResetTrimSpace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpace", GoGetter: "TrimSpace"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpaceInput", GoGetter: "TrimSpaceInput"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3FileFormatOrcOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatOutputReference",
		reflect.TypeOf((*StageExternalS3FileFormatOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "avro", GoGetter: "Avro"},
			_jsii_.MemberProperty{JsiiProperty: "avroInput", GoGetter: "AvroInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "csv", GoGetter: "Csv"},
			_jsii_.MemberProperty{JsiiProperty: "csvInput", GoGetter: "CsvInput"},
			_jsii_.MemberProperty{JsiiProperty: "formatName", GoGetter: "FormatName"},
			_jsii_.MemberProperty{JsiiProperty: "formatNameInput", GoGetter: "FormatNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "json", GoGetter: "Json"},
			_jsii_.MemberProperty{JsiiProperty: "jsonInput", GoGetter: "JsonInput"},
			_jsii_.MemberProperty{JsiiProperty: "orc", GoGetter: "Orc"},
			_jsii_.MemberProperty{JsiiProperty: "orcInput", GoGetter: "OrcInput"},
			_jsii_.MemberProperty{JsiiProperty: "parquet", GoGetter: "Parquet"},
			_jsii_.MemberProperty{JsiiProperty: "parquetInput", GoGetter: "ParquetInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAvro", GoMethod: "PutAvro"},
			_jsii_.MemberMethod{JsiiMethod: "putCsv", GoMethod: "PutCsv"},
			_jsii_.MemberMethod{JsiiMethod: "putJson", GoMethod: "PutJson"},
			_jsii_.MemberMethod{JsiiMethod: "putOrc", GoMethod: "PutOrc"},
			_jsii_.MemberMethod{JsiiMethod: "putParquet", GoMethod: "PutParquet"},
			_jsii_.MemberMethod{JsiiMethod: "putXml", GoMethod: "PutXml"},
			_jsii_.MemberMethod{JsiiMethod: "resetAvro", GoMethod: "ResetAvro"},
			_jsii_.MemberMethod{JsiiMethod: "resetCsv", GoMethod: "ResetCsv"},
			_jsii_.MemberMethod{JsiiMethod: "resetFormatName", GoMethod: "ResetFormatName"},
			_jsii_.MemberMethod{JsiiMethod: "resetJson", GoMethod: "ResetJson"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrc", GoMethod: "ResetOrc"},
			_jsii_.MemberMethod{JsiiMethod: "resetParquet", GoMethod: "ResetParquet"},
			_jsii_.MemberMethod{JsiiMethod: "resetXml", GoMethod: "ResetXml"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "xml", GoGetter: "Xml"},
			_jsii_.MemberProperty{JsiiProperty: "xmlInput", GoGetter: "XmlInput"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3FileFormatOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatParquet",
		reflect.TypeOf((*StageExternalS3FileFormatParquet)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatParquetOutputReference",
		reflect.TypeOf((*StageExternalS3FileFormatParquetOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "binaryAsText", GoGetter: "BinaryAsText"},
			_jsii_.MemberProperty{JsiiProperty: "binaryAsTextInput", GoGetter: "BinaryAsTextInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberProperty{JsiiProperty: "compressionInput", GoGetter: "CompressionInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "nullIf", GoGetter: "NullIf"},
			_jsii_.MemberProperty{JsiiProperty: "nullIfInput", GoGetter: "NullIfInput"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharacters", GoGetter: "ReplaceInvalidCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharactersInput", GoGetter: "ReplaceInvalidCharactersInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBinaryAsText", GoMethod: "ResetBinaryAsText"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompression", GoMethod: "ResetCompression"},
			_jsii_.MemberMethod{JsiiMethod: "resetNullIf", GoMethod: "ResetNullIf"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplaceInvalidCharacters", GoMethod: "ResetReplaceInvalidCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrimSpace", GoMethod: "ResetTrimSpace"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseLogicalType", GoMethod: "ResetUseLogicalType"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseVectorizedScanner", GoMethod: "ResetUseVectorizedScanner"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpace", GoGetter: "TrimSpace"},
			_jsii_.MemberProperty{JsiiProperty: "trimSpaceInput", GoGetter: "TrimSpaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "useLogicalType", GoGetter: "UseLogicalType"},
			_jsii_.MemberProperty{JsiiProperty: "useLogicalTypeInput", GoGetter: "UseLogicalTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "useVectorizedScanner", GoGetter: "UseVectorizedScanner"},
			_jsii_.MemberProperty{JsiiProperty: "useVectorizedScannerInput", GoGetter: "UseVectorizedScannerInput"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3FileFormatParquetOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatXml",
		reflect.TypeOf((*StageExternalS3FileFormatXml)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3FileFormatXmlOutputReference",
		reflect.TypeOf((*StageExternalS3FileFormatXmlOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberProperty{JsiiProperty: "compressionInput", GoGetter: "CompressionInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disableAutoConvert", GoGetter: "DisableAutoConvert"},
			_jsii_.MemberProperty{JsiiProperty: "disableAutoConvertInput", GoGetter: "DisableAutoConvertInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "ignoreUtf8Errors", GoGetter: "IgnoreUtf8Errors"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreUtf8ErrorsInput", GoGetter: "IgnoreUtf8ErrorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "preserveSpace", GoGetter: "PreserveSpace"},
			_jsii_.MemberProperty{JsiiProperty: "preserveSpaceInput", GoGetter: "PreserveSpaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharacters", GoGetter: "ReplaceInvalidCharacters"},
			_jsii_.MemberProperty{JsiiProperty: "replaceInvalidCharactersInput", GoGetter: "ReplaceInvalidCharactersInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompression", GoMethod: "ResetCompression"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableAutoConvert", GoMethod: "ResetDisableAutoConvert"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreUtf8Errors", GoMethod: "ResetIgnoreUtf8Errors"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreserveSpace", GoMethod: "ResetPreserveSpace"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplaceInvalidCharacters", GoMethod: "ResetReplaceInvalidCharacters"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipByteOrderMark", GoMethod: "ResetSkipByteOrderMark"},
			_jsii_.MemberMethod{JsiiMethod: "resetStripOuterElement", GoMethod: "ResetStripOuterElement"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "skipByteOrderMark", GoGetter: "SkipByteOrderMark"},
			_jsii_.MemberProperty{JsiiProperty: "skipByteOrderMarkInput", GoGetter: "SkipByteOrderMarkInput"},
			_jsii_.MemberProperty{JsiiProperty: "stripOuterElement", GoGetter: "StripOuterElement"},
			_jsii_.MemberProperty{JsiiProperty: "stripOuterElementInput", GoGetter: "StripOuterElementInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3FileFormatXmlOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3ShowOutput",
		reflect.TypeOf((*StageExternalS3ShowOutput)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3ShowOutputList",
		reflect.TypeOf((*StageExternalS3ShowOutputList)(nil)).Elem(),
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
			j := jsiiProxy_StageExternalS3ShowOutputList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3ShowOutputOutputReference",
		reflect.TypeOf((*StageExternalS3ShowOutputOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloud", GoGetter: "Cloud"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "createdOn", GoGetter: "CreatedOn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "directoryEnabled", GoGetter: "DirectoryEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
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
			_jsii_.MemberProperty{JsiiProperty: "hasCredentials", GoGetter: "HasCredentials"},
			_jsii_.MemberProperty{JsiiProperty: "hasEncryptionKey", GoGetter: "HasEncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "ownerRoleType", GoGetter: "OwnerRoleType"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schemaName", GoGetter: "SchemaName"},
			_jsii_.MemberProperty{JsiiProperty: "storageIntegration", GoGetter: "StorageIntegration"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3ShowOutputOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3Timeouts",
		reflect.TypeOf((*StageExternalS3Timeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-snowflake.stageExternalS3.StageExternalS3TimeoutsOutputReference",
		reflect.TypeOf((*StageExternalS3TimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_StageExternalS3TimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
