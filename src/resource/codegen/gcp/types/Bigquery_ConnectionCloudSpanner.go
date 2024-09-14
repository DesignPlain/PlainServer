package types

type Bigquery_ConnectionCloudSpanner struct {
	// Cloud Spanner database role for fine-grained access control. The Cloud Spanner admin should have provisioned the database role with appropriate permissions, such as `SELECT` and `INSERT`. Other users should only use roles provided by their Cloud Spanner admins. The database role name must start with a letter, and can only contain letters, numbers, and underscores. For more details, see https://cloud.google.com/spanner/docs/fgac-about.
	DatabaseRole string `json:"databaseRole,omitempty" yaml:"databaseRole,omitempty"`

	// Allows setting max parallelism per query when executing on Spanner independent compute resources. If unspecified, default values of parallelism are chosen that are dependent on the Cloud Spanner instance configuration. `useParallelism` and `useDataBoost` must be set when setting max parallelism.
	MaxParallelism int `json:"maxParallelism,omitempty" yaml:"maxParallelism,omitempty"`

	// If set, the request will be executed via Spanner independent compute resources. `use_parallelism` must be set when using data boost.
	UseDataBoost bool `json:"useDataBoost,omitempty" yaml:"useDataBoost,omitempty"`

	// If parallelism should be used when reading from Cloud Spanner.
	UseParallelism bool `json:"useParallelism,omitempty" yaml:"useParallelism,omitempty"`

	/*
	   (Optional, Deprecated)
	   If the serverless analytics service should be used to read data from Cloud Spanner. `useParallelism` must be set when using serverless analytics.

	   > --Warning:-- `useServerlessAnalytics` is deprecated and will be removed in a future major release. Use `useDataBoost` instead.
	*/
	UseServerlessAnalytics bool `json:"useServerlessAnalytics,omitempty" yaml:"useServerlessAnalytics,omitempty"`

	// Cloud Spanner database in the form `project/instance/database'.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`
}
