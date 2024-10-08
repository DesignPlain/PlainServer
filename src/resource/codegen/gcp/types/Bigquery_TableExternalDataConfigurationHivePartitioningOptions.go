package types

type Bigquery_TableExternalDataConfigurationHivePartitioningOptions struct {
	/*
	   When set, what mode of hive partitioning to use when
	   reading data. The following modes are supported.
	   - AUTO: automatically infer partition key name(s) and type(s).
	   - STRINGS: automatically infer partition key name(s). All types are
	   Not all storage formats support hive partitioning. Requesting hive
	   partitioning on an unsupported format will lead to an error.
	   Currently supported formats are: JSON, CSV, ORC, Avro and Parquet.
	   - CUSTOM: when set to `CUSTOM`, you must encode the partition key schema within the `source_uri_prefix` by setting `source_uri_prefix` to `gs://bucket/path_to_table/{key1:TYPE1}/{key2:TYPE2}/{key3:TYPE3}`.
	*/
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	/*
	   If set to true, queries over this table
	   require a partition filter that can be used for partition elimination to be
	   specified.
	*/
	RequirePartitionFilter bool `json:"requirePartitionFilter,omitempty" yaml:"requirePartitionFilter,omitempty"`

	/*
	   When hive partition detection is requested,
	   a common for all source uris must be required. The prefix must end immediately
	   before the partition key encoding begins. For example, consider files following
	   this data layout. `gs://bucket/path_to_table/dt=2019-06-01/country=USA/id=7/file.avro`
	   `gs://bucket/path_to_table/dt=2019-05-31/country=CA/id=3/file.avro` When hive
	   partitioning is requested with either AUTO or STRINGS detection, the common prefix
	   can be either of `gs://bucket/path_to_table` or `gs://bucket/path_to_table/`.
	   Note that when `mode` is set to `CUSTOM`, you must encode the partition key schema within the `source_uri_prefix` by setting `source_uri_prefix` to `gs://bucket/path_to_table/{key1:TYPE1}/{key2:TYPE2}/{key3:TYPE3}`.
	*/
	SourceUriPrefix string `json:"sourceUriPrefix,omitempty" yaml:"sourceUriPrefix,omitempty"`
}
