package dms

type S3Endpoint struct {
	// Date separating delimiter to use during folder partitioning. Valid values are `SLASH`, `UNDERSCORE`, `DASH`, and `NONE`. (AWS default is `SLASH`.) (Ignored for source endpoints.)
	DatePartitionDelimiter string `json:"datePartitionDelimiter,omitempty" yaml:"datePartitionDelimiter,omitempty"`

	// Whether to enable statistics for Parquet pages and row groups. Default is `true`.
	EnableStatistics bool `json:"enableStatistics,omitempty" yaml:"enableStatistics,omitempty"`

	// Bucket owner to prevent sniping. Value is an AWS account ID.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`

	// Delimiter used to separate rows in the source files. Default is newline (_i.e._, `\n`).
	CsvRowDelimiter string `json:"csvRowDelimiter,omitempty" yaml:"csvRowDelimiter,omitempty"`

	// Date format to use during folder partitioning. Use this parameter when `date_partition_enabled` is set to true. Valid values are `YYYYMMDD`, `YYYYMMDDHH`, `YYYYMM`, `MMYYYYDD`, and `DDMMYYYY`. (AWS default is `YYYYMMDD`.) (Ignored for source endpoints.)
	DatePartitionSequence string `json:"datePartitionSequence,omitempty" yaml:"datePartitionSequence,omitempty"`

	// Type of endpoint. Valid values are `source`, `target`.
	EndpointType string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`

	// Version of the .parquet file format. Valid values are `parquet-1-0` and `parquet-2-0`. (AWS default is `parquet-1-0`.) (Ignored for source endpoints.)
	ParquetVersion string `json:"parquetVersion,omitempty" yaml:"parquetVersion,omitempty"`

	// Whether to use `csv_no_sup_value` for columns not included in the supplemental log. (Ignored for source endpoints.)
	UseCsvNoSupValue bool `json:"useCsvNoSupValue,omitempty" yaml:"useCsvNoSupValue,omitempty"`

	// Minimum file size condition as defined in kilobytes to output a file to Amazon S3. (AWS default is 32000 KB.)
	CdcMinFileSize int `json:"cdcMinFileSize,omitempty" yaml:"cdcMinFileSize,omitempty"`

	// ARN for the certificate.
	CertificateArn string `json:"certificateArn,omitempty" yaml:"certificateArn,omitempty"`

	// Set to compress target files. Valid values are `GZIP` and `NONE`. Default is `NONE`. (Ignored for source endpoints.)
	CompressionType string `json:"compressionType,omitempty" yaml:"compressionType,omitempty"`

	// When `encryption_mode` is `SSE_KMS`, ARN for the AWS KMS key. (Ignored for source endpoints -- only `SSE_S3` `encryption_mode` is valid.)
	ServerSideEncryptionKmsKeyId string `json:"serverSideEncryptionKmsKeyId,omitempty" yaml:"serverSideEncryptionKmsKeyId,omitempty"`

	// Whether to write insert operations to .csv or .parquet output files. Default is `false`.
	CdcInsertsOnly bool `json:"cdcInsertsOnly,omitempty" yaml:"cdcInsertsOnly,omitempty"`

	// Size of one data page in bytes. (AWS default is 1 MiB, _i.e._, `1048576`.)
	DataPageSize int `json:"dataPageSize,omitempty" yaml:"dataPageSize,omitempty"`

	// Whether to integrate AWS Glue Data Catalog with an Amazon S3 target. See [Using AWS Glue Data Catalog with an Amazon S3 target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.GlueCatalog) for more information. Default is `false`.
	GlueCatalogGeneration bool `json:"glueCatalogGeneration,omitempty" yaml:"glueCatalogGeneration,omitempty"`

	// Server-side encryption mode that you want to encrypt your .csv or .parquet object files copied to S3. Valid values are `SSE_S3` and `SSE_KMS`. (AWS default is `SSE_S3`.) (Ignored for source endpoints -- only `SSE_S3` is valid.)
	EncryptionMode string `json:"encryptionMode,omitempty" yaml:"encryptionMode,omitempty"`

	// JSON document that describes how AWS DMS should interpret the data.
	ExternalTableDefinition string `json:"externalTableDefinition,omitempty" yaml:"externalTableDefinition,omitempty"`

	// When this value is set to `1`, DMS ignores the first row header in a .csv file. (AWS default is `0`.)
	IgnoreHeaderRows int `json:"ignoreHeaderRows,omitempty" yaml:"ignoreHeaderRows,omitempty"`

	// Whether to enable a full load to write INSERT operations to the .csv output files only to indicate how the rows were added to the source database. Default is `false`.
	IncludeOpForFullLoad bool `json:"includeOpForFullLoad,omitempty" yaml:"includeOpForFullLoad,omitempty"`

	// Number of rows in a row group. (AWS default is `10000`.)
	RowGroupLength int `json:"rowGroupLength,omitempty" yaml:"rowGroupLength,omitempty"`

	// Delimiter used to separate columns in the source files. Default is `,`.
	CsvDelimiter string `json:"csvDelimiter,omitempty" yaml:"csvDelimiter,omitempty"`

	// Maximum size in bytes of an encoded dictionary page of a column. (AWS default is 1 MiB, _i.e._, `1048576`.)
	DictPageSizeLimit int `json:"dictPageSizeLimit,omitempty" yaml:"dictPageSizeLimit,omitempty"`

	// Type of encoding to use. Value values are `rle_dictionary`, `plain`, and `plain_dictionary`. (AWS default is `rle_dictionary`.)
	EncodingType string `json:"encodingType,omitempty" yaml:"encodingType,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Undocumented argument for use as directed by AWS Support.
	DetachTargetOnLobLookupFailureParquet bool `json:"detachTargetOnLobLookupFailureParquet,omitempty" yaml:"detachTargetOnLobLookupFailureParquet,omitempty"`

	// Database endpoint identifier. Identifiers must contain from 1 to 255 alphanumeric characters or hyphens, begin with a letter, contain only ASCII letters, digits, and hyphens, not end with a hyphen, and not contain two consecutive hyphens.
	EndpointId string `json:"endpointId,omitempty" yaml:"endpointId,omitempty"`

	// Maximum size (in KB) of any .csv file to be created while migrating to an S3 target during full load. Valid values are from `1` to `1048576`. (AWS default is 1 GB, _i.e._, `1048576`.)
	MaxFileSize int `json:"maxFileSize,omitempty" yaml:"maxFileSize,omitempty"`

	// Whether DMS saves the transaction order for a CDC load on the S3 target specified by `cdc_path`. Default is `false`. (Ignored for source endpoints.)
	PreserveTransactions bool `json:"preserveTransactions,omitempty" yaml:"preserveTransactions,omitempty"`

	// For an S3 source, whether each leading double quotation mark has to be followed by an ending double quotation mark. Default is `true`.
	Rfc4180 bool `json:"rfc4180,omitempty" yaml:"rfc4180,omitempty"`

	// S3 object prefix.
	BucketFolder string `json:"bucketFolder,omitempty" yaml:"bucketFolder,omitempty"`

	// String to as null when writing to the target. (AWS default is `NULL`.)
	CsvNullValue string `json:"csvNullValue,omitempty" yaml:"csvNullValue,omitempty"`

	// Partition S3 bucket folders based on transaction commit dates. Default is `false`. (Ignored for source endpoints.)
	DatePartitionEnabled bool `json:"datePartitionEnabled,omitempty" yaml:"datePartitionEnabled,omitempty"`

	// SSL mode to use for the connection. Valid values are `none`, `require`, `verify-ca`, `verify-full`. (AWS default is `none`.)
	SslMode string `json:"sslMode,omitempty" yaml:"sslMode,omitempty"`

	// Maximum length of the interval, defined in seconds, after which to output a file to Amazon S3. (AWS default is `60`.)
	CdcMaxBatchInterval int `json:"cdcMaxBatchInterval,omitempty" yaml:"cdcMaxBatchInterval,omitempty"`

	// Column to add with timestamp information to the endpoint data for an Amazon S3 target.
	TimestampColumnName string `json:"timestampColumnName,omitempty" yaml:"timestampColumnName,omitempty"`

	// Whether to add padding. Default is `false`. (Ignored for source endpoints.)
	AddTrailingPaddingCharacter bool `json:"addTrailingPaddingCharacter,omitempty" yaml:"addTrailingPaddingCharacter,omitempty"`

	// S3 bucket name.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// Predefined (canned) access control list for objects created in an S3 bucket. Valid values include `none`, `private`, `public-read`, `public-read-write`, `authenticated-read`, `aws-exec-read`, `bucket-owner-read`, and `bucket-owner-full-control`. Default is `none`.
	CannedAclForObjects string `json:"cannedAclForObjects,omitempty" yaml:"cannedAclForObjects,omitempty"`

	// Whether to write insert and update operations to .csv or .parquet output files. Default is `false`.
	CdcInsertsAndUpdates bool `json:"cdcInsertsAndUpdates,omitempty" yaml:"cdcInsertsAndUpdates,omitempty"`

	// Convert the current UTC time to a timezone. The conversion occurs when a date partition folder is created and a CDC filename is generated. The timezone format is Area/Location (_e.g._, `Europe/Paris`). Use this when `date_partition_enabled` is `true`. (Ignored for source endpoints.)
	DatePartitionTimezone string `json:"datePartitionTimezone,omitempty" yaml:"datePartitionTimezone,omitempty"`

	// Output format for the files that AWS DMS uses to create S3 objects. Valid values are `csv` and `parquet`.  (Ignored for source endpoints -- only `csv` is valid.)
	DataFormat string `json:"dataFormat,omitempty" yaml:"dataFormat,omitempty"`

	// ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// Specifies the precision of any TIMESTAMP column values written to an S3 object file in .parquet format. Default is `false`. (Ignored for source endpoints.)
	ParquetTimestampInMillisecond bool `json:"parquetTimestampInMillisecond,omitempty" yaml:"parquetTimestampInMillisecond,omitempty"`

	/*
	   ARN of the IAM role with permissions to the S3 Bucket.

	   The following arguments are optional:
	*/
	ServiceAccessRoleArn string `json:"serviceAccessRoleArn,omitempty" yaml:"serviceAccessRoleArn,omitempty"`

	// When set to `true`, uses the task start time as the timestamp column value instead of the time data is written to target. For full load, when set to `true`, each row of the timestamp column contains the task start time. For CDC loads, each row of the timestamp column contains the transaction commit time.When set to false, the full load timestamp in the timestamp column increments with the time data arrives at the target. Default is `false`.
	UseTaskStartTimeForFullLoadTimestamp bool `json:"useTaskStartTimeForFullLoadTimestamp,omitempty" yaml:"useTaskStartTimeForFullLoadTimestamp,omitempty"`

	// Whether to add column name information to the .csv output file. Default is `false`.
	AddColumnName bool `json:"addColumnName,omitempty" yaml:"addColumnName,omitempty"`

	// Folder path of CDC files. If `cdc_path` is set, AWS DMS reads CDC files from this path and replicates the data changes to the target endpoint. Supported in AWS DMS versions 3.4.2 and later.
	CdcPath string `json:"cdcPath,omitempty" yaml:"cdcPath,omitempty"`

	// Only applies if output files for a CDC load are written in .csv format. If `use_csv_no_sup_value` is set to `true`, string to use for all columns not included in the supplemental log. If you do not specify a string value, DMS uses the null value for these columns regardless of `use_csv_no_sup_value`. (Ignored for source endpoints.)
	CsvNoSupValue string `json:"csvNoSupValue,omitempty" yaml:"csvNoSupValue,omitempty"`
}
