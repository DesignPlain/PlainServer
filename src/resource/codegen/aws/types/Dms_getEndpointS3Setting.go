package types

type Dms_getEndpointS3Setting struct {
	//
	DataPageSize int `json:"dataPageSize,omitempty" yaml:"dataPageSize,omitempty"`

	//
	DatePartitionDelimiter string `json:"datePartitionDelimiter,omitempty" yaml:"datePartitionDelimiter,omitempty"`

	//
	DatePartitionSequence string `json:"datePartitionSequence,omitempty" yaml:"datePartitionSequence,omitempty"`

	//
	EnableStatistics bool `json:"enableStatistics,omitempty" yaml:"enableStatistics,omitempty"`

	//
	ParquetVersion string `json:"parquetVersion,omitempty" yaml:"parquetVersion,omitempty"`

	//
	TimestampColumnName string `json:"timestampColumnName,omitempty" yaml:"timestampColumnName,omitempty"`

	//
	CdcInsertsOnly bool `json:"cdcInsertsOnly,omitempty" yaml:"cdcInsertsOnly,omitempty"`

	//
	CdcPath string `json:"cdcPath,omitempty" yaml:"cdcPath,omitempty"`

	//
	CsvNoSupValue string `json:"csvNoSupValue,omitempty" yaml:"csvNoSupValue,omitempty"`

	//
	ExternalTableDefinition string `json:"externalTableDefinition,omitempty" yaml:"externalTableDefinition,omitempty"`

	//
	IgnoreHeadersRow int `json:"ignoreHeadersRow,omitempty" yaml:"ignoreHeadersRow,omitempty"`

	//
	AddColumnName bool `json:"addColumnName,omitempty" yaml:"addColumnName,omitempty"`

	//
	CompressionType string `json:"compressionType,omitempty" yaml:"compressionType,omitempty"`

	//
	CsvRowDelimiter string `json:"csvRowDelimiter,omitempty" yaml:"csvRowDelimiter,omitempty"`

	//
	DatePartitionEnabled bool `json:"datePartitionEnabled,omitempty" yaml:"datePartitionEnabled,omitempty"`

	//
	DictPageSizeLimit int `json:"dictPageSizeLimit,omitempty" yaml:"dictPageSizeLimit,omitempty"`

	//
	IncludeOpForFullLoad bool `json:"includeOpForFullLoad,omitempty" yaml:"includeOpForFullLoad,omitempty"`

	//
	UseTaskStartTimeForFullLoadTimestamp bool `json:"useTaskStartTimeForFullLoadTimestamp,omitempty" yaml:"useTaskStartTimeForFullLoadTimestamp,omitempty"`

	//
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	//
	CdcInsertsAndUpdates bool `json:"cdcInsertsAndUpdates,omitempty" yaml:"cdcInsertsAndUpdates,omitempty"`

	//
	CdcMaxBatchInterval int `json:"cdcMaxBatchInterval,omitempty" yaml:"cdcMaxBatchInterval,omitempty"`

	//
	CsvNullValue string `json:"csvNullValue,omitempty" yaml:"csvNullValue,omitempty"`

	//
	IgnoreHeaderRows int `json:"ignoreHeaderRows,omitempty" yaml:"ignoreHeaderRows,omitempty"`

	//
	GlueCatalogGeneration bool `json:"glueCatalogGeneration,omitempty" yaml:"glueCatalogGeneration,omitempty"`

	//
	Rfc4180 bool `json:"rfc4180,omitempty" yaml:"rfc4180,omitempty"`

	//
	RowGroupLength int `json:"rowGroupLength,omitempty" yaml:"rowGroupLength,omitempty"`

	//
	BucketFolder string `json:"bucketFolder,omitempty" yaml:"bucketFolder,omitempty"`

	//
	EncryptionMode string `json:"encryptionMode,omitempty" yaml:"encryptionMode,omitempty"`

	//
	PreserveTransactions bool `json:"preserveTransactions,omitempty" yaml:"preserveTransactions,omitempty"`

	//
	UseCsvNoSupValue bool `json:"useCsvNoSupValue,omitempty" yaml:"useCsvNoSupValue,omitempty"`

	//
	CannedAclForObjects string `json:"cannedAclForObjects,omitempty" yaml:"cannedAclForObjects,omitempty"`

	//
	CdcMinFileSize int `json:"cdcMinFileSize,omitempty" yaml:"cdcMinFileSize,omitempty"`

	//
	CsvDelimiter string `json:"csvDelimiter,omitempty" yaml:"csvDelimiter,omitempty"`

	//
	EncodingType string `json:"encodingType,omitempty" yaml:"encodingType,omitempty"`

	//
	ServiceAccessRoleArn string `json:"serviceAccessRoleArn,omitempty" yaml:"serviceAccessRoleArn,omitempty"`

	//
	DataFormat string `json:"dataFormat,omitempty" yaml:"dataFormat,omitempty"`

	//
	MaxFileSize int `json:"maxFileSize,omitempty" yaml:"maxFileSize,omitempty"`

	//
	ParquetTimestampInMillisecond bool `json:"parquetTimestampInMillisecond,omitempty" yaml:"parquetTimestampInMillisecond,omitempty"`

	//
	ServerSideEncryptionKmsKeyId string `json:"serverSideEncryptionKmsKeyId,omitempty" yaml:"serverSideEncryptionKmsKeyId,omitempty"`
}
