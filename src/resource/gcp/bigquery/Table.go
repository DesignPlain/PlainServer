package bigquery

import types "DesignSphere_Server/src/resource/gcp/types"

type Table struct {
	/*
	   Specifies column names to use for data clustering.
	   Up to four top-level columns are allowed, and should be specified in
	   descending priority order.
	*/
	Clusterings []string `json:"clusterings,omitempty" yaml:"clusterings,omitempty"`

	/*
	   A unique ID for the resource.
	   Changing this forces a new resource to be created.
	*/
	TableId string `json:"tableId,omitempty" yaml:"tableId,omitempty"`

	/*
	   If specified, configures time-based
	   partitioning for this table. Structure is documented below.
	*/
	TimePartitioning types.Bigquery_TableTimePartitioning `json:"timePartitioning,omitempty" yaml:"timePartitioning,omitempty"`

	/*
	   Whether or not to allow the provider to destroy the instance. Unless this field is set to false
	   in state, a `=destroy` or `=update` that would delete the instance will fail.
	*/
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	// The field description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The time when this table expires, in
	   milliseconds since the epoch. If not present, the table will persist
	   indefinitely. Expired tables will be deleted and their storage
	   reclaimed.
	*/
	ExpirationTime int `json:"expirationTime,omitempty" yaml:"expirationTime,omitempty"`

	/*
	   A mapping of labels to assign to the resource.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field 'effective_labels' for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Defines the primary key and foreign keys.
	   Structure is documented below.
	*/
	TableConstraints types.Bigquery_TableTableConstraints `json:"tableConstraints,omitempty" yaml:"tableConstraints,omitempty"`

	// Replication info of a table created using "AS REPLICA" DDL like: "CREATE MATERIALIZED VIEW mv1 AS REPLICA OF src_mv".
	TableReplicationInfo types.Bigquery_TableTableReplicationInfo `json:"tableReplicationInfo,omitempty" yaml:"tableReplicationInfo,omitempty"`

	/*
	   Describes the data format,
	   location, and other properties of a table stored outside of BigQuery.
	   By defining these properties, the data source can then be queried as
	   if it were a standard BigQuery table. Structure is documented below.
	*/
	ExternalDataConfiguration types.Bigquery_TableExternalDataConfiguration `json:"externalDataConfiguration,omitempty" yaml:"externalDataConfiguration,omitempty"`

	// A descriptive name for the table.
	FriendlyName string `json:"friendlyName,omitempty" yaml:"friendlyName,omitempty"`

	/*
	   The maximum staleness of data that could be
	   returned when the table (or stale MV) is queried. Staleness encoded as a
	   string encoding of [SQL IntervalValue
	   type](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#interval_type).
	*/
	MaxStaleness string `json:"maxStaleness,omitempty" yaml:"maxStaleness,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   If specified, configures range-based
	   partitioning for this table. Structure is documented below.
	*/
	RangePartitioning types.Bigquery_TableRangePartitioning `json:"rangePartitioning,omitempty" yaml:"rangePartitioning,omitempty"`

	/*
	   If set to true, queries over this table
	   require a partition filter that can be used for partition elimination to be
	   specified.
	*/
	RequirePartitionFilter bool `json:"requirePartitionFilter,omitempty" yaml:"requirePartitionFilter,omitempty"`

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
	   The dataset ID to create the table in.
	   Changing this forces a new resource to be created.
	*/
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`

	/*
	   Specifies how the table should be encrypted.
	   If left blank, the table will be encrypted with a Google-managed key; that process
	   is transparent to the user.  Structure is documented below.
	*/
	EncryptionConfiguration types.Bigquery_TableEncryptionConfiguration `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`

	/*
	   If specified, configures this table as a materialized view.
	   Structure is documented below.
	*/
	MaterializedView types.Bigquery_TableMaterializedView `json:"materializedView,omitempty" yaml:"materializedView,omitempty"`

	/*
	   If specified, configures this table as a view.
	   Structure is documented below.
	*/
	View types.Bigquery_TableView `json:"view,omitempty" yaml:"view,omitempty"`
}
