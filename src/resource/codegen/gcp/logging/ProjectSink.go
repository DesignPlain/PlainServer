package logging

import types "libds/gcp/types"

type ProjectSink struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions types.Logging_ProjectSinkBigqueryOptions `json:"bigqueryOptions,omitempty" yaml:"bigqueryOptions,omitempty"`

	/*
	   The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, a BigQuery dataset, a Cloud Logging bucket, or a Google Cloud project. Examples:

	   - `storage.googleapis.com/[GCS_BUCKET]`
	   - `bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]`
	   - `pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]`
	   - `logging.googleapis.com/projects/[PROJECT_ID]/locations/global/buckets/[BUCKET_ID]`
	   - `logging.googleapis.com/projects/[PROJECT_ID]`

	   The writer associated with the sink must have access to write to the above resource.
	*/
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both `filter` and one of `exclusions.filter`, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []types.Logging_ProjectSinkExclusion `json:"exclusions,omitempty" yaml:"exclusions,omitempty"`

	/*
	   The filter to apply when exporting logs. Only log entries that match the filter are exported.
	   See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	   write a filter.
	*/
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`

	// The name of the logging sink. Logging automatically creates two sinks: `_Required` and `_Default`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A user managed service account that will be used to write
	   the log entries. The format must be `serviceAccount:some@email`. This field can only be specified if you are
	   routing logs to a destination outside this sink's project. If not specified, a Logging service account
	   will automatically be generated.
	*/
	CustomWriterIdentity string `json:"customWriterIdentity,omitempty" yaml:"customWriterIdentity,omitempty"`

	// A description of this sink. The maximum length of the description is 8000 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	/*
	   The ID of the project to create the sink in. If omitted, the project associated with the provider is
	   used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Whether or not to create a unique identity associated with this sink. If `false`, then the `writer_identity` used is `serviceAccount:cloud-logs@system.gserviceaccount.com`. If `true` (the default),
	   then a unique service account is created and used for this sink. If you wish to publish logs across projects or utilize
	   `bigquery_options`, you must set `unique_writer_identity` to true.
	*/
	UniqueWriterIdentity bool `json:"uniqueWriterIdentity,omitempty" yaml:"uniqueWriterIdentity,omitempty"`
}
