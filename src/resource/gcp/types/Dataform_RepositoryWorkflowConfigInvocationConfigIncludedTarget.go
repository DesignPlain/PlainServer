package types

type Dataform_RepositoryWorkflowConfigInvocationConfigIncludedTarget struct {
	// The action's database (Google Cloud project ID).
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	// The action's name, within database and schema.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The action's schema (BigQuery dataset ID), within database.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
