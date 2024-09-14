package types

type Bigquery_TableExternalDataConfiguration struct {
	/*
	   Additional options if `source_format` is set to
	   "AVRO".  Structure is documented below.
	*/
	AvroOptions Bigquery_TableExternalDataConfigurationAvroOptions `json:"avroOptions,omitempty" yaml:"avroOptions,omitempty"`

	/*
	   Additional options if
	   `source_format` is set to "GOOGLE_SHEETS". Structure is
	   documented below.
	*/
	GoogleSheetsOptions Bigquery_TableExternalDataConfigurationGoogleSheetsOptions `json:"googleSheetsOptions,omitempty" yaml:"googleSheetsOptions,omitempty"`

	/*
	   The data format. Please see sourceFormat under
	   [ExternalDataConfiguration](https://cloud.google.com/bigquery/docs/reference/rest/v2/tables#externaldataconfiguration)
	   in Bigquery's public API documentation for supported formats. To use "GOOGLE_SHEETS"
	   the `scopes` must include "https://www.googleapis.com/auth/drive.readonly".
	*/
	SourceFormat string `json:"sourceFormat,omitempty" yaml:"sourceFormat,omitempty"`

	// Object Metadata is used to create Object Tables. Object Tables contain a listing of objects (with their metadata) found at the sourceUris. If `object_metadata` is set, `source_format` should be omitted.
	ObjectMetadata string `json:"objectMetadata,omitempty" yaml:"objectMetadata,omitempty"`

	/*
	   Additional properties to set if
	   `source_format` is set to "PARQUET". Structure is documented below.
	*/
	ParquetOptions Bigquery_TableExternalDataConfigurationParquetOptions `json:"parquetOptions,omitempty" yaml:"parquetOptions,omitempty"`

	// When creating an external table, the user can provide a reference file with the table schema. This is enabled for the following formats: AVRO, PARQUET, ORC.
	ReferenceFileSchemaUri string `json:"referenceFileSchemaUri,omitempty" yaml:"referenceFileSchemaUri,omitempty"`

	/*
	   Let BigQuery try to autodetect the schema
	   and format of the table.
	*/
	Autodetect bool `json:"autodetect,omitempty" yaml:"autodetect,omitempty"`

	/*
	   Additional properties to set if
	   `source_format` is set to "CSV". Structure is documented below.
	*/
	CsvOptions Bigquery_TableExternalDataConfigurationCsvOptions `json:"csvOptions,omitempty" yaml:"csvOptions,omitempty"`

	/*
	   Additional properties to set if
	   `source_format` is set to "JSON". Structure is documented below.
	*/
	JsonOptions Bigquery_TableExternalDataConfigurationJsonOptions `json:"jsonOptions,omitempty" yaml:"jsonOptions,omitempty"`

	/*
	   The compression type of the data source.
	   Valid values are "NONE" or "GZIP".
	*/
	Compression string `json:"compression,omitempty" yaml:"compression,omitempty"`

	// Metadata Cache Mode for the table. Set this to enable caching of metadata from external data source. Valid values are `AUTOMATIC` and `MANUAL`.
	MetadataCacheMode string `json:"metadataCacheMode,omitempty" yaml:"metadataCacheMode,omitempty"`

	/*
	   A list of the fully-qualified URIs that point to
	   your data in Google Cloud.
	*/
	SourceUris []string `json:"sourceUris,omitempty" yaml:"sourceUris,omitempty"`

	/*
	   Indicates if BigQuery should
	   allow extra values that are not represented in the table schema.
	   If true, the extra values are ignored. If false, records with
	   extra columns are treated as bad records, and if there are too
	   many bad records, an invalid error is returned in the job result.
	   The default value is false.
	*/
	IgnoreUnknownValues bool `json:"ignoreUnknownValues,omitempty" yaml:"ignoreUnknownValues,omitempty"`

	/*
	   The maximum number of bad records that
	   BigQuery can ignore when reading data.
	*/
	MaxBadRecords int `json:"maxBadRecords,omitempty" yaml:"maxBadRecords,omitempty"`

	/*
	   A JSON schema for the external table. Schema is required
	   for CSV and JSON formats if autodetect is not on. Schema is disallowed
	   for Google Cloud Bigtable, Cloud Datastore backups, Avro, Iceberg, ORC and Parquet formats.
	   ~>--NOTE:-- Because this field expects a JSON string, any changes to the
	   string will create a diff, even if the JSON itself hasn't changed.
	   Furthermore drift for this field cannot not be detected because BigQuery
	   only uses this schema to compute the effective schema for the table, therefore
	   any changes on the configured value will force the table to be recreated.
	   This schema is effectively only applied when creating a table from an external
	   datasource, after creation the computed schema will be stored in
	   `google_bigquery_table.schema`

	   ~>--NOTE:-- If you set `external_data_configuration.connection_id`, the
	   table schema must be specified using the top-level `schema` field
	   documented above.
	*/
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`

	/*
	   The connection specifying the credentials to be used to read
	   external storage, such as Azure Blob, Cloud Storage, or S3. The `connection_id` can have
	   the form `{{project}}.{{location}}.{{connection_id}}`
	   or `projects/{{project}}/locations/{{location}}/connections/{{connection_id}}`.

	   ~>--NOTE:-- If you set `external_data_configuration.connection_id`, the
	   table schema must be specified using the top-level `schema` field
	   documented above.
	*/
	ConnectionId string `json:"connectionId,omitempty" yaml:"connectionId,omitempty"`

	/*
	   Specifies how source URIs are interpreted for constructing the file set to load.
	   By default source URIs are expanded against the underlying storage.
	   Other options include specifying manifest files. Only applicable to object storage systems. Docs
	*/
	FileSetSpecType string `json:"fileSetSpecType,omitempty" yaml:"fileSetSpecType,omitempty"`

	/*
	   When set, configures hive partitioning
	   support. Not all storage formats support hive partitioning -- requesting hive
	   partitioning on an unsupported format will lead to an error, as will providing
	   an invalid specification. Structure is documented below.
	*/
	HivePartitioningOptions Bigquery_TableExternalDataConfigurationHivePartitioningOptions `json:"hivePartitioningOptions,omitempty" yaml:"hivePartitioningOptions,omitempty"`
}
