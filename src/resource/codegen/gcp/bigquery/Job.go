package bigquery

import types "libds/gcp/types"

type Job struct {
	/*
	   Configures an extract job.
	   Structure is documented below.
	*/
	Extract types.Bigquery_JobExtract `json:"extract,omitempty" yaml:"extract,omitempty"`

	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs string `json:"jobTimeoutMs,omitempty" yaml:"jobTimeoutMs,omitempty"`

	// The geographic location of the job. The default value is US.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Copies a table.
	   Structure is documented below.
	*/
	Copy types.Bigquery_JobCopy `json:"copy,omitempty" yaml:"copy,omitempty"`

	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId string `json:"jobId,omitempty" yaml:"jobId,omitempty"`

	/*
	   The labels associated with this job. You can use these to organize and group your jobs.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Configures a load job.
	   Structure is documented below.
	*/
	Load types.Bigquery_JobLoad `json:"load,omitempty" yaml:"load,omitempty"`

	/*
	   SQL query text to execute. The useLegacySql field can be used to indicate whether the query uses legacy SQL or standard SQL.
	   -NOTE-: queries containing [DML language](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
	   (`DELETE`, `UPDATE`, `MERGE`, `INSERT`) must specify `create_disposition = ""` and `write_disposition = ""`.
	*/
	Query types.Bigquery_JobQuery `json:"query,omitempty" yaml:"query,omitempty"`
}
