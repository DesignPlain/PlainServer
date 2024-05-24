package types

type Dms_getEndpointS3Setting struct {
	//
	CannedAclForObjects string `json:"cannedAclForObjects,omitempty" yaml:"cannedAclForObjects,omitempty"`

	//
	CdcInsertsAndUpdates bool `json:"cdcInsertsAndUpdates,omitempty" yaml:"cdcInsertsAndUpdates,omitempty"`

	//
	CsvDelimiter string `json:"csvDelimiter,omitempty" yaml:"csvDelimiter,omitempty"`

	//
	ExternalTableDefinition string `json:"externalTableDefinition,omitempty" yaml:"externalTableDefinition,omitempty"`

	//
	MaxFileSize int `json:"maxFileSize,omitempty" yaml:"maxFileSize,omitempty"`

	//
	PreserveTransactions bool `json:"preserveTransactions,omitempty" yaml:"preserveTransactions,omitempty"`

	//
	CsvRowDelimiter string `json:"csvRowDelimiter,omitempty" yaml:"csvRowDelimiter,omitempty"`

	//
	ParquetTimestampInMillisecond bool `json:"parquetTimestampInMillisecond,omitempty" yaml:"parquetTimestampInMillisecond,omitempty"`

	//
	ServiceAccessRoleArn string `json:"serviceAccessRoleArn,omitempty" yaml:"serviceAccessRoleArn,omitempty"`

	//
	AddColumnName bool `json:"addColumnName,omitempty" yaml:"addColumnName,omitempty"`

	//
	CdcInsertsOnly bool `json:"cdcInsertsOnly,omitempty" yaml:"cdcInsertsOnly,omitempty"`

	//
	CdcMaxBatchInterval int `json:"cdcMaxBatchInterval,omitempty" yaml:"cdcMaxBatchInterval,omitempty"`

	//
	Rfc4180 bool `json:"rfc4180,omitempty" yaml:"rfc4180,omitempty"`

	//
	BucketFolder string `json:"bucketFolder,omitempty" yaml:"bucketFolder,omitempty"`

	//
	CdcMinFileSize int `json:"cdcMinFileSize,omitempty" yaml:"cdcMinFileSize,omitempty"`

	//
	DataFormat string `json:"dataFormat,omitempty" yaml:"dataFormat,omitempty"`

	//
	DictPageSizeLimit int `json:"dictPageSizeLimit,omitempty" yaml:"dictPageSizeLimit,omitempty"`

	//
	EncryptionMode string `json:"encryptionMode,omitempty" yaml:"encryptionMode,omitempty"`

	//
	IgnoreHeadersRow int `json:"ignoreHeadersRow,omitempty" yaml:"ignoreHeadersRow,omitempty"`

	//
	CsvNullValue string `json:"csvNullValue,omitempty" yaml:"csvNullValue,omitempty"`

	//
	DataPageSize int `json:"dataPageSize,omitempty" yaml:"dataPageSize,omitempty"`

	//
	DatePartitionSequence string `json:"datePartitionSequence,omitempty" yaml:"datePartitionSequence,omitempty"`

	//
	ParquetVersion string `json:"parquetVersion,omitempty" yaml:"parquetVersion,omitempty"`

	//
	RowGroupLength int `json:"rowGroupLength,omitempty" yaml:"rowGroupLength,omitempty"`

	//
	UseCsvNoSupValue bool `json:"useCsvNoSupValue,omitempty" yaml:"useCsvNoSupValue,omitempty"`

	//
	UseTaskStartTimeForFullLoadTimestamp bool `json:"useTaskStartTimeForFullLoadTimestamp,omitempty" yaml:"useTaskStartTimeForFullLoadTimestamp,omitempty"`

	//
	CdcPath string `json:"cdcPath,omitempty" yaml:"cdcPath,omitempty"`

	//
	CompressionType string `json:"compressionType,omitempty" yaml:"compressionType,omitempty"`

	//
	DatePartitionDelimiter string `json:"datePartitionDelimiter,omitempty" yaml:"datePartitionDelimiter,omitempty"`

	//
	DatePartitionEnabled bool `json:"datePartitionEnabled,omitempty" yaml:"datePartitionEnabled,omitempty"`

	//
	GlueCatalogGeneration bool `json:"glueCatalogGeneration,omitempty" yaml:"glueCatalogGeneration,omitempty"`

	//
	ServerSideEncryptionKmsKeyId string `json:"serverSideEncryptionKmsKeyId,omitempty" yaml:"serverSideEncryptionKmsKeyId,omitempty"`

	//
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	//
	EnableStatistics bool `json:"enableStatistics,omitempty" yaml:"enableStatistics,omitempty"`

	//
	EncodingType string `json:"encodingType,omitempty" yaml:"encodingType,omitempty"`

	//
	TimestampColumnName string `json:"timestampColumnName,omitempty" yaml:"timestampColumnName,omitempty"`

	//
	CsvNoSupValue string `json:"csvNoSupValue,omitempty" yaml:"csvNoSupValue,omitempty"`

	//
	IgnoreHeaderRows int `json:"ignoreHeaderRows,omitempty" yaml:"ignoreHeaderRows,omitempty"`

	//
	IncludeOpForFullLoad bool `json:"includeOpForFullLoad,omitempty" yaml:"includeOpForFullLoad,omitempty"`
}
