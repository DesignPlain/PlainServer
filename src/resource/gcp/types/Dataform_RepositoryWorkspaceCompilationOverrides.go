package types

type Dataform_RepositoryWorkspaceCompilationOverrides struct {
	// The default database (Google Cloud project ID).
	DefaultDatabase string `json:"defaultDatabase,omitempty" yaml:"defaultDatabase,omitempty"`

	// The suffix that should be appended to all schema (BigQuery dataset ID) names.
	SchemaSuffix string `json:"schemaSuffix,omitempty" yaml:"schemaSuffix,omitempty"`

	// The prefix that should be prepended to all table names.
	TablePrefix string `json:"tablePrefix,omitempty" yaml:"tablePrefix,omitempty"`
}
