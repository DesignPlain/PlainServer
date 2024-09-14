package types

type Dms_EndpointS3Settings struct {
	// ARN of the IAM Role with permissions to read from or write to the S3 Bucket.
	ServiceAccessRoleArn string `json:"serviceAccessRoleArn,omitempty" yaml:"serviceAccessRoleArn,omitempty"`

	// Whether to use `csv_no_sup_value` for columns not included in the supplemental log.
	UseCsvNoSupValue bool `json:"useCsvNoSupValue,omitempty" yaml:"useCsvNoSupValue,omitempty"`

	// Whether to write insert operations to .csv or .parquet output files. Default is `false`.
	CdcInsertsOnly bool `json:"cdcInsertsOnly,omitempty" yaml:"cdcInsertsOnly,omitempty"`

	// Folder path of CDC files. For an S3 source, this setting is required if a task captures change data; otherwise, it's optional. If `cdc_path` is set, AWS DMS reads CDC files from this path and replicates the data changes to the target endpoint. Supported in AWS DMS versions 3.4.2 and later.
	CdcPath string `json:"cdcPath,omitempty" yaml:"cdcPath,omitempty"`

	// Output format for the files that AWS DMS uses to create S3 objects. Valid values are `csv` and `parquet`. Default is `csv`.
	DataFormat string `json:"dataFormat,omitempty" yaml:"dataFormat,omitempty"`

	// Size of one data page in bytes. Default is `1048576` (1 MiB).
	DataPageSize int `json:"dataPageSize,omitempty" yaml:"dataPageSize,omitempty"`

	// Date separating delimiter to use during folder partitioning. Valid values are `SLASH`, `UNDERSCORE`, `DASH`, and `NONE`. Default is `SLASH`.
	DatePartitionDelimiter string `json:"datePartitionDelimiter,omitempty" yaml:"datePartitionDelimiter,omitempty"`

	// Whether DMS saves the transaction order for a CDC load on the S3 target specified by `cdc_path`. Default is `false`.
	PreserveTransactions bool `json:"preserveTransactions,omitempty" yaml:"preserveTransactions,omitempty"`

	// Whether to add column name information to the .csv output file. Default is `false`.
	AddColumnName bool `json:"addColumnName,omitempty" yaml:"addColumnName,omitempty"`

	// Delimiter used to separate columns in the source files. Default is `,`.
	CsvDelimiter string `json:"csvDelimiter,omitempty" yaml:"csvDelimiter,omitempty"`

	// Type of encoding to use. Value values are `rle_dictionary`, `plain`, and `plain_dictionary`. Default is `rle_dictionary`.
	EncodingType string `json:"encodingType,omitempty" yaml:"encodingType,omitempty"`

	// Whether to enable a full load to write INSERT operations to the .csv output files only to indicate how the rows were added to the source database. Default is `false`.
	IncludeOpForFullLoad bool `json:"includeOpForFullLoad,omitempty" yaml:"includeOpForFullLoad,omitempty"`

	// Set to compress target files. Default is `NONE`. Valid values are `GZIP` and `NONE`.
	CompressionType string `json:"compressionType,omitempty" yaml:"compressionType,omitempty"`

	// Partition S3 bucket folders based on transaction commit dates. Default is `false`.
	DatePartitionEnabled bool `json:"datePartitionEnabled,omitempty" yaml:"datePartitionEnabled,omitempty"`

	// Server-side encryption mode that you want to encrypt your .csv or .parquet object files copied to S3. Valid values are `SSE_S3` and `SSE_KMS`. Default is `SSE_S3`.
	EncryptionMode string `json:"encryptionMode,omitempty" yaml:"encryptionMode,omitempty"`

	// For an S3 source, whether each leading double quotation mark has to be followed by an ending double quotation mark. Default is `true`.
	Rfc4180 bool `json:"rfc4180,omitempty" yaml:"rfc4180,omitempty"`

	// Minimum file size condition as defined in kilobytes to output a file to Amazon S3. Default is `32000`. --NOTE:-- Previously, this setting was measured in megabytes but now represents kilobytes. Update configurations accordingly.
	CdcMinFileSize int `json:"cdcMinFileSize,omitempty" yaml:"cdcMinFileSize,omitempty"`

	// Date format to use during folder partitioning. Use this parameter when `date_partition_enabled` is set to true. Valid values are `YYYYMMDD`, `YYYYMMDDHH`, `YYYYMM`, `MMYYYYDD`, and `DDMMYYYY`. Default is `YYYYMMDD`.
	DatePartitionSequence string `json:"datePartitionSequence,omitempty" yaml:"datePartitionSequence,omitempty"`

	// Maximum size in bytes of an encoded dictionary page of a column. Default is `1048576` (1 MiB).
	DictPageSizeLimit int `json:"dictPageSizeLimit,omitempty" yaml:"dictPageSizeLimit,omitempty"`

	// Whether to integrate AWS Glue Data Catalog with an Amazon S3 target. See [Using AWS Glue Data Catalog with an Amazon S3 target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.GlueCatalog) for more information. Default is `false`.
	GlueCatalogGeneration bool `json:"glueCatalogGeneration,omitempty" yaml:"glueCatalogGeneration,omitempty"`

	// Maximum size (in KB) of any .csv file to be created while migrating to an S3 target during full load. Valid values are from `1` to `1048576`. Default is `1048576` (1 GB).
	MaxFileSize int `json:"maxFileSize,omitempty" yaml:"maxFileSize,omitempty"`

	// When set to true, uses the task start time as the timestamp column value instead of the time data is written to target. For full load, when set to true, each row of the timestamp column contains the task start time. For CDC loads, each row of the timestamp column contains the transaction commit time. When set to false, the full load timestamp in the timestamp column increments with the time data arrives at the target. Default is `false`.
	UseTaskStartTimeForFullLoadTimestamp bool `json:"useTaskStartTimeForFullLoadTimestamp,omitempty" yaml:"useTaskStartTimeForFullLoadTimestamp,omitempty"`

	// Maximum length of the interval, defined in seconds, after which to output a file to Amazon S3. Default is `60`.
	CdcMaxBatchInterval int `json:"cdcMaxBatchInterval,omitempty" yaml:"cdcMaxBatchInterval,omitempty"`

	// Column to add with timestamp information to the endpoint data for an Amazon S3 target.
	TimestampColumnName string `json:"timestampColumnName,omitempty" yaml:"timestampColumnName,omitempty"`

	// ARN or Id of KMS Key to use when `encryption_mode` is `SSE_KMS`.
	ServerSideEncryptionKmsKeyId string `json:"serverSideEncryptionKmsKeyId,omitempty" yaml:"serverSideEncryptionKmsKeyId,omitempty"`

	// Predefined (canned) access control list for objects created in an S3 bucket. Valid values include `none`, `private`, `public-read`, `public-read-write`, `authenticated-read`, `aws-exec-read`, `bucket-owner-read`, and `bucket-owner-full-control`. Default is `none`.
	CannedAclForObjects string `json:"cannedAclForObjects,omitempty" yaml:"cannedAclForObjects,omitempty"`

	// Whether to write insert and update operations to .csv or .parquet output files. Default is `false`.
	CdcInsertsAndUpdates bool `json:"cdcInsertsAndUpdates,omitempty" yaml:"cdcInsertsAndUpdates,omitempty"`

	// Whether to enable statistics for Parquet pages and row groups. Default is `true`.
	EnableStatistics bool `json:"enableStatistics,omitempty" yaml:"enableStatistics,omitempty"`

	// JSON document that describes how AWS DMS should interpret the data.
	ExternalTableDefinition string `json:"externalTableDefinition,omitempty" yaml:"externalTableDefinition,omitempty"`

	// When this value is set to `1`, DMS ignores the first row header in a .csv file. Default is `0`.
	IgnoreHeaderRows int `json:"ignoreHeaderRows,omitempty" yaml:"ignoreHeaderRows,omitempty"`

	// Number of rows in a row group. Default is `10000`.
	RowGroupLength int `json:"rowGroupLength,omitempty" yaml:"rowGroupLength,omitempty"`

	// S3 bucket name.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// String to use for all columns not included in the supplemental log.
	CsvNoSupValue string `json:"csvNoSupValue,omitempty" yaml:"csvNoSupValue,omitempty"`

	// Delimiter used to separate rows in the source files. Default is `\n`.
	CsvRowDelimiter string `json:"csvRowDelimiter,omitempty" yaml:"csvRowDelimiter,omitempty"`

	// S3 object prefix.
	BucketFolder string `json:"bucketFolder,omitempty" yaml:"bucketFolder,omitempty"`

	// String to as null when writing to the target.
	CsvNullValue string `json:"csvNullValue,omitempty" yaml:"csvNullValue,omitempty"`

	// Specifies the precision of any TIMESTAMP column values written to an S3 object file in .parquet format. Default is `false`.
	ParquetTimestampInMillisecond bool `json:"parquetTimestampInMillisecond,omitempty" yaml:"parquetTimestampInMillisecond,omitempty"`

	// Version of the .parquet file format. Default is `parquet-1-0`. Valid values are `parquet-1-0` and `parquet-2-0`.
	ParquetVersion string `json:"parquetVersion,omitempty" yaml:"parquetVersion,omitempty"`
}
