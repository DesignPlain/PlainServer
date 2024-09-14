package types

type Gkebackup_RestorePlanRestoreConfigExcludedNamespaces struct {
	// A list of Kubernetes Namespaces.
	Namespaces []string `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}
