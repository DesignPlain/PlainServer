package logging

import types "libds/gcp/types"

type FolderSink struct {
	/*
	   The filter to apply when exporting logs. Only log entries that match the filter are exported.
	   See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	   write a filter.
	*/
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`

	/*
	   The folder to be exported to the sink. Note that either `[FOLDER_ID]` or `folders/[FOLDER_ID]` is
	   accepted.
	*/
	Folder string `json:"folder,omitempty" yaml:"folder,omitempty"`

	/*
	   Whether or not to include children folders in the sink export. If true, logs
	   associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	*/
	IncludeChildren bool `json:"includeChildren,omitempty" yaml:"includeChildren,omitempty"`

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions types.Logging_FolderSinkBigqueryOptions `json:"bigqueryOptions,omitempty" yaml:"bigqueryOptions,omitempty"`

	// A description of this sink. The maximum length of the description is 8000 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

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

	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both `filter` and one of `exclusions.filter`, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []types.Logging_FolderSinkExclusion `json:"exclusions,omitempty" yaml:"exclusions,omitempty"`

	// The name of the logging sink.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
