package types

type Opensearch_DomainSoftwareUpdateOptions struct {
	// Whether automatic service software updates are enabled for the domain. Defaults to `false`.
	AutoSoftwareUpdateEnabled bool `json:"autoSoftwareUpdateEnabled,omitempty" yaml:"autoSoftwareUpdateEnabled,omitempty"`
}
