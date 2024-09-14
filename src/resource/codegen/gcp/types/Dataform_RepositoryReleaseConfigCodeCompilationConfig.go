package types

type Dataform_RepositoryReleaseConfigCodeCompilationConfig struct {
	// Optional. The default schema (BigQuery dataset ID).
	DefaultSchema string `json:"defaultSchema,omitempty" yaml:"defaultSchema,omitempty"`

	// Optional. The suffix that should be appended to all schema (BigQuery dataset ID) names.
	SchemaSuffix string `json:"schemaSuffix,omitempty" yaml:"schemaSuffix,omitempty"`

	// Optional. The prefix that should be prepended to all table names.
	TablePrefix string `json:"tablePrefix,omitempty" yaml:"tablePrefix,omitempty"`

	/*
	   Optional. User-defined variables that are made available to project code during compilation.
	   An object containing a list of "key": value pairs.
	   Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	*/
	Vars map[string]string `json:"vars,omitempty" yaml:"vars,omitempty"`

	// Optional. The default schema (BigQuery dataset ID) for assertions.
	AssertionSchema string `json:"assertionSchema,omitempty" yaml:"assertionSchema,omitempty"`

	// Optional. The suffix that should be appended to all database (Google Cloud project ID) names.
	DatabaseSuffix string `json:"databaseSuffix,omitempty" yaml:"databaseSuffix,omitempty"`

	// Optional. The default database (Google Cloud project ID).
	DefaultDatabase string `json:"defaultDatabase,omitempty" yaml:"defaultDatabase,omitempty"`

	/*
	   Optional. The default BigQuery location to use. Defaults to "US".
	   See the BigQuery docs for a full list of locations: https://cloud.google.com/bigquery/docs/locations.
	*/
	DefaultLocation string `json:"defaultLocation,omitempty" yaml:"defaultLocation,omitempty"`
}
