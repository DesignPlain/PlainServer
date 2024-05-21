package types

type Gkebackup_RestorePlanRestoreConfigSelectedNamespaces struct {
	// A list of Kubernetes Namespaces.
	Namespaces []string `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}
