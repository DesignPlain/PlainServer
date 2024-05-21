package types

type Dataform_RepositoryWorkflowConfigInvocationConfig struct {
	// Optional. When set to true, any incremental tables will be fully refreshed.
	FullyRefreshIncrementalTablesEnabled bool `json:"fullyRefreshIncrementalTablesEnabled,omitempty" yaml:"fullyRefreshIncrementalTablesEnabled,omitempty"`

	// Optional. The set of tags to include.
	IncludedTags []string `json:"includedTags,omitempty" yaml:"includedTags,omitempty"`

	/*
	   Optional. The set of action identifiers to include.
	   Structure is documented below.
	*/
	IncludedTargets []Dataform_RepositoryWorkflowConfigInvocationConfigIncludedTarget `json:"includedTargets,omitempty" yaml:"includedTargets,omitempty"`

	// Optional. The service account to run workflow invocations under.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// Optional. When set to true, transitive dependencies of included actions will be executed.
	TransitiveDependenciesIncluded bool `json:"transitiveDependenciesIncluded,omitempty" yaml:"transitiveDependenciesIncluded,omitempty"`

	// Optional. When set to true, transitive dependents of included actions will be executed.
	TransitiveDependentsIncluded bool `json:"transitiveDependentsIncluded,omitempty" yaml:"transitiveDependentsIncluded,omitempty"`
}
